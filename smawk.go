// This is an executable package; it must be named main
package main

// Here we import the required libraries for the controller
// github.com/bmatt468/smawk-bot: fetch the SMÄWKBot API from its repo
// net/http: used for spooling up a webserver to listen for updates
// os: used for checking if a file exists
import (
    "github.com/bmatt468/smawk-bot"
    "net/http"
    "os"
)

// The following constants are needed for the controller to function properly
// SMAWKToken: the token for your bot provided by the BotFather
// bot_url: the URL for the webhook that the bot will listen on (note that https:// is required)
// bot_url_token: (optional but hightly recommended) a token that will be added onto the end of a url so that only a specific url known only to you will be pung by telegram
// bot_cert: the filename of the self-signed certificate
// bot_key: the filename of the public key for the certificate
const (

    SMAWKToken      = "249930361:AAHz1Gksb-eT0SQG47lDb7WbJxujr7kGCkU"
    bot_url         = "https://www.benjaminrmatthews.com:8443"
    bot_url_token   = "/309LKj2390gklj1LJF2"
    bot_cert        = "smawk_cert.pem"
    bot_key         = "smawk_key.pem"
)

func main() {
    // First, we need to see if the certificate we specify is present or not. This certificate is required by telegram if you want to recieve updates
    if _, err := os.Stat(bot_cert); os.IsNotExist(err) {
        // No cert found. We need to generate one
        // FILL OUT THESE FIELDS
        // These will be used to generate the certificate
        country := "US"
        state   := "South Carolina"
        city    := "Lexington"
        org     := "MattCorp"
        domain  := "www.benjaminrmatthews.com"
        smawk.GenerateCertificate(country,state,city,org,domain,bot_key,bot_cert)
    }

    // Create the bot using the provided access token and tell it that we aren't in debug mode
    bot, _ := smawk.Connect(SMAWKToken, false)

    // Open the webhook with our self-signed certificate
    bot.OpenWebhookWithCert(bot_url+bot_url_token, bot_cert)

    // Start listening on our webhook for the commands
    // Spin off a goroutine to handle listening elsewhere
    updates := bot.Listen(bot_url_token)
    go http.ListenAndServeTLS("0.0.0.0:8443", bot_cert, bot_key, nil)

    // Once we receive an update in the update channel, pass it through to the update handler in the SB API
    for update := range updates {
        bot.ParseAndExecuteUpdate(update)
    }
}
