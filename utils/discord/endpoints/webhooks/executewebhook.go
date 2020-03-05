package webhooks

import (
	"github.com/TicketsBot/GoPanel/utils/discord"
	"github.com/TicketsBot/GoPanel/utils/discord/objects"
)

type ExecuteWebhookBody struct {
	Content         string                 `json:"content"`
	Username        string                 `json:"username"`
	AvatarUrl       string                 `json:"avatar_url"`
	AllowedMentions objects.AllowedMention `json:"allowed_mentions"`
}

func ExecuteWebhook(url string) discord.Endpoint {
	return discord.Endpoint{
		RequestType:       discord.POST,
		AuthorizationType: discord.NONE,
		Endpoint:          url,
	}
}