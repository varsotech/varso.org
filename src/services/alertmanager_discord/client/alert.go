package client

import (
	"github.com/luminancetech/varso/src/common/api"
)

type AlertResponse struct {
}

type AlertManAlert struct {
	Annotations struct {
		Description string `json:"description"`
		Summary     string `json:"summary"`
	} `json:"annotations"`
	EndsAt       string            `json:"endsAt"`
	GeneratorURL string            `json:"generatorURL"`
	Labels       map[string]string `json:"labels"`
	StartsAt     string            `json:"startsAt"`
	Status       string            `json:"status"`
}

type AlertRequest struct {
	Alerts            []AlertManAlert `json:"alerts"`
	CommonAnnotations struct {
		Summary string `json:"summary"`
	} `json:"commonAnnotations"`
	CommonLabels struct {
		Alertname string `json:"alertname"`
	} `json:"commonLabels"`
	ExternalURL string `json:"externalURL"`
	GroupKey    string `json:"groupKey"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	Receiver string `json:"receiver"`
	Status   string `json:"status"`
	Version  string `json:"version"`
}

func (c *Client) Alert() (*AlertResponse, *api.Error) {
	return api.SendRequest[AlertRequest, AlertResponse](c.Client.Client, c.AuthHeader, "POST", "alertmanager_discord/alert", &AlertRequest{})
}
