package main

import (
    "github.com/SMAWK/smawk-bot"
    "net/http"
    "os"
    "os/exec"
    "bytes"
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
        generateCertificate()
        os.Exit(1)
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

func generateCertificate() {
    // FILL OUT THESE FIELDS
    // These will be used to generate the certificate
    country := "US"
    state   := "South Carolina"
    city    := "Lexington"
    org     := "My Simple Things"
    domain  := "mysimplethings.xyz"

    // Generate our string for the certificate
    certstring := "\"/C="+country+"/ST="+state+"/L="+city+"/O="+org+"/CN="+domain+"\""

    cmdname := "openssl"
    cmdargs := []string{"req","-newkey","rsa:2048","-sha256","-nodes","-keyout",bot_key,"-x509","-days","365","-out",bot_cert,"-subj","/"+certstring}

    cmd := exec.Command(cmdname,cmdargs...)
    var stderr bytes.Buffer
    cmd.Stderr = &stderr
    err := cmd.Run()
    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
    }
}
