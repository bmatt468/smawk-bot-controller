package main

import (
    "github.com/SMAWK/smawk-bot"
)

const (
    SMAWKToken      = "249930361:AAHz1Gksb-eT0SQG47lDb7WbJxujr7kGCkU"
    bot_url         = "https://mysimplethings.xyz:8443"
    bot_url_token   = "/309LKj2390gklj1LJF2"
    bot_cert        = "smawk_cert.pem"
)

func main() {
    // Create the bot using the provided access token
    bot := smawk.Connect(SMAWKToken, false)

    // Open the webhook
    bot.OpenWebhookWithCert(bot_url+bot_url_token, bot_cert)

    // Start listening on our webhook for the commands
    // Spin off a goroutine to handle listening elsewhere
    /*updates := bot.ListenForWebhook("/309LKj2390gklj1LJF2")
    go http.ListenAndServeTLS("0.0.0.0:8443", "smawk_cert.pem", "smawk_key.pem", nil)

    // Parse and execute each of the commands that come through the pipe
    for update := range updates {
        cmd := update.Message.Text
        if (cmd == "/start" || cmd == "/start@smawk_bot") {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Lo, the official SMÄWKBot rises!")
            bot.Send(msg)
        } else if (cmd == "/hello" || cmd == "/hello@smawk_bot") {
            msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, @" + update.Message.From.UserName + "!")
            bot.Send(msg)
        }
    }*/
}
