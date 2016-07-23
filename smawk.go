package main

import (
    "github.com/SMAWK/smawk-bot"
    "net/http"
    "os"
)

const (
    SMAWKToken      = "249930361:AAHz1Gksb-eT0SQG47lDb7WbJxujr7kGCkU"
    bot_url         = "https://mysimplethings.xyz:8443"
    bot_url_token   = "/309LKj2390gklj1LJF2"
    bot_cert        = "smawk_cert.pem"
    bot_key         = "smawk_key.pem"
)

func main() {
    // Check to see if bot_cert exists
    if _, err := os.Stat(bot_cert); os.IsNotExist(err) {
        // No cert found. We need to generate one
        // FILL OUT THESE FIELDS
        // These will be used to generate the certificate
        country := "US"
        state   := "South Carolina"
        city    := "Lexington"
        org     := "My Simple Things"
        domain  := "mysimplethings.xyz"
        smawk.GenerateCertificate(country,state,city,org,domain,bot_key,bot_cert)
    }

    // Create the bot using the provided access token
    bot, _ := smawk.Connect(SMAWKToken, false)

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
