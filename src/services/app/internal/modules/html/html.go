package html

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
)

const maxRetries = 5

type RedirectError struct{}

func (r RedirectError) Error() string {
	return "redirect occured"
}

type HTML struct {
	finalUrl string
	url      string
	content  string
	html     *html.Node
	loaded   bool
}

func NewHTMLFromURL(url string) *HTML {
	return &HTML{url: url}
}

func (h *HTML) lazyLoad() error {
	if h.loaded {
		return nil
	}

	c := http.Client{}

	req, err := http.NewRequest(http.MethodGet, h.url, nil)
	if err != nil {
		return errors.Wrapf(err, "failed creating request for lazy loading html")
	}

	// HACK: news.google.com proxy will only return 302 for cURL, otherwise it will do a javascript redirect.
	req.Header.Set("User-Agent", "curl/8.1.2")
	
	// Run with retries
	resp, err := h.doWithRetries(req, c)
	if err != nil {
		return errors.Wrap(err, "failed loading html with retries")
	}
	defer resp.Body.Close()

	h.finalUrl = resp.Request.URL.String()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "failed reading bodu")
	}
	h.content = string(content)

	// Parse the HTML content
	doc, err := html.Parse(bytes.NewBuffer(content))
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}
	h.html = doc

	h.loaded = true
	return nil
}

func (h *HTML) doRequest(req *http.Request, c http.Client) (*http.Response, error) {
	resp, err := c.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "unable to retrieve item HTML")
	}

	if resp.StatusCode != 200 {
		_ = resp.Body.Close()
		return nil, errors.Errorf("got bad status code %d retrieving item HTML", resp.StatusCode)
	}

	return resp, nil
}

func (h *HTML) doWithRetries(req *http.Request, c http.Client) (resp *http.Response, err error) {
	for i := 0; i < maxRetries; i++ {
		resp, err = h.doRequest(req, c)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}

	return
}

func (h *HTML) GetHTML() (string, error) {
	if err := h.lazyLoad(); err != nil {
		return "", err
	}

	return h.content, nil
}

// GetFinalURL gets the URL of the request, and the last URL that was redirected to
func (h *HTML) GetFinalURL() (string, error) {
	if err := h.lazyLoad(); err != nil {
		return "", err
	}

	return h.finalUrl, nil
}

func (h *HTML) GetOGImage() string {
	if err := h.lazyLoad(); err != nil {
		logrus.WithError(err).Errorf("failed loading html, can't get OG image")
		return ""
	}

	// Find the meta tag with property="og:image"
	ogImage := findMetaProperty(h.html, "og:image")
	if ogImage != "" {
		return ogImage
	}

	return ""
}

// Helper function to find the content of a meta tag with a given property value
func findMetaProperty(n *html.Node, property string) string {
	if n.Type == html.ElementNode && n.Data == "meta" {
		for _, attr := range n.Attr {
			if attr.Key == "property" && attr.Val == property {
				// If property matches, return the content attribute
				for _, attr := range n.Attr {
					if attr.Key == "content" {
						return attr.Val
					}
				}
			}
		}
	}

	// If the property is not found, recursively search in the children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		result := findMetaProperty(c, property)
		if result != "" {
			return result
		}
	}

	return ""
}
