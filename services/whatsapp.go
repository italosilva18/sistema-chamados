package services

import (
	"github.com/twilio/twilio-go"
)

func EnviarMensagemWhatsApp(to string, body string) error {
	client := twilio.NewRestClient()

	params := &v2010.CreateMessageParams{}
	params.SetTo(to)
	params.SetFrom("whatsapp:+14155238886") // Número do Twilio
	params.SetBody(body)

	_, err := client.ApiV2010.CreateMessage(params)
	return err
}
