package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/luminancetech/varso/src/common/api"
	"github.com/luminancetech/varso/src/services/alertmanager_discord/client"
	"github.com/luminancetech/varso/src/services/alertmanager_discord/internal/config"
)

// Discord color values
const (
	ColorRed   = 0x992D22
	ColorGreen = 0x2ECC71
	ColorGrey  = 0x95A5A6
)

func Alert(_ *api.Writer, r *http.Request, _ httprouter.Params, jwt *api.JWT, amo client.AlertRequest) (*client.AlertResponse, *api.Error) {
	groupedAlerts := make(map[string][]client.AlertManAlert)

	for _, alert := range amo.Alerts {
		groupedAlerts[alert.Status] = append(groupedAlerts[alert.Status], alert)
	}

	for status, alerts := range groupedAlerts {
		DO := discordOut{}

		RichEmbed := discordEmbed{
			Title:       fmt.Sprintf("[%s:%d] %s", strings.ToUpper(status), len(alerts), amo.CommonLabels.Alertname),
			Description: amo.CommonAnnotations.Summary,
			Color:       ColorGrey,
			Fields:      []discordEmbedField{},
		}

		if status == "firing" {
			RichEmbed.Color = ColorRed
		} else if status == "resolved" {
			RichEmbed.Color = ColorGreen
		}

		if amo.CommonAnnotations.Summary != "" {
			DO.Content = fmt.Sprintf(" === %s === \n", amo.CommonAnnotations.Summary)
		}

		for _, alert := range alerts {
			realname := alert.Labels["instance"]
			if strings.Contains(realname, "localhost") && alert.Labels["exported_instance"] != "" {
				realname = alert.Labels["exported_instance"]
			}

			RichEmbed.Fields = append(RichEmbed.Fields, discordEmbedField{
				Name:  fmt.Sprintf("[%s]: %s on %s", strings.ToUpper(status), alert.Labels["alertname"], realname),
				Value: alert.Annotations.Description,
			})
		}

		DO.Embeds = []discordEmbed{RichEmbed}

		DOD, err := json.Marshal(DO)
		if err != nil {
			return nil, api.NewInternalError(err, "failed marshalling json")
		}

		_, err = http.Post(config.AlertsDiscordWebhook, "application/json", bytes.NewReader(DOD))
		if err != nil {
			return nil, api.NewInternalError(err, "failed calling alert webhook")
		}
	}
	return &client.AlertResponse{}, nil
}

type discordOut struct {
	Content string         `json:"content"`
	Embeds  []discordEmbed `json:"embeds"`
}

type discordEmbed struct {
	Title       string              `json:"title"`
	Description string              `json:"description"`
	Color       int                 `json:"color"`
	Fields      []discordEmbedField `json:"fields"`
}

type discordEmbedField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
