package main

import (
    "github.com/SMAWK/smawk-bot"
    "net/http"
)

const (
    SMAWKToken      = "249930361:AAHz1Gksb-eT0SQG47lDb7WbJxujr7kGCkU"
    bot_url         = "https://mysimplethings.xyz:8443"
    bot_url_token   = "/309LKj2390gklj1LJF2"
    bot_cert        = "smawk_cert.pem"
    bot_key         = "smawk_key.pem"
)

func main() {
    // Create the bot using the provided access token
    bot := smawk.Connect(SMAWKToken, false)

    // Open the webhook
    bot.OpenWebhookWithCert(bot_url+bot_url_token, bot_cert)

    // Start listening on our webhook for the commands
    // Spin off a goroutine to handle listening elsewhere
    updates := bot.Listen(bot_url_token)
    go http.ListenAndServeTLS("0.0.0.0:8443", bot_cert, bot_key, nil)

    // Parse and execute each of the commands that come through the pipe
    for update := range updates {
        bot.ParseAndExecuteUpdate(update)
    }
}
