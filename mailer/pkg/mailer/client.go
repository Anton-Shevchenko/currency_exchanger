package mailer

import (
	"github.com/mailjet/mailjet-apiv3-go"
)

type Config struct {
	PubKey      string
	PrivKey     string
	SenderEmail string
}

type Client struct {
	config Config
}

func NewClient(config Config) *Client {
	return &Client{
		config: config,
	}
}

func (c Client) SendEmail(to, msg string) error {
	mailjetClient := mailjet.NewMailjetClient(c.config.PubKey, c.config.PrivKey)
	messagesInfo := []mailjet.InfoMessagesV31{
		{
			From: &mailjet.RecipientV31{
				Email: c.config.SenderEmail,
				Name:  "Currency Rate",
			},
			To: &mailjet.RecipientsV31{
				mailjet.RecipientV31{
					Email: to,
					Name:  "Currency Rate",
				},
			},
			Subject:  "Daily currency rate",
			TextPart: "Daily currency rate",
			HTMLPart: msg,
			CustomID: "AppGettingStartedTest",
		},
	}
	messages := mailjet.MessagesV31{Info: messagesInfo}
	_, err := mailjetClient.SendMailV31(&messages)

	if err != nil {
		return err
	}

	return nil
}
