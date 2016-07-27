#SMÄWKBot Controller
This repo contains the controller that implements the SMÄWKBot API

# Creating a controller
A controller is required to implement the SMÄWKBot [SB] API. This controller is a standalone program that Go will compile; once compiled, this program can be run from most anywhere on a file system.

To specify that the controller will be compiled (and not implemented like a library) you **must** specify `package main` at the top of your main file. If this is not done, the program will not be built.

Inside the controller, we have six main tasks:

- Fetching the SB API
- Generating a certificate
- Generating a SB instance
- Opening a webhook
- Listening for updates
- Sending the updates to the SB API

These tasks will be explained below.

## Fetching the SB API
The SB API is on a public repository; therefore, Go makes it very easy to grab this dependency when we build the controller.
The code to fetch the API is as follows:
```Go
import (
    "github.com/SMAWK/smawk-bot"
)
```
With this one simple command, Go will take care of fetching the SB API from its repo, and adding it to the proper location inside of the `$GOPATH/src` directory.

## Generating a certificate
Telegram requires that each webhook be on an https:// URL. This task is easily remedied by using a self-signed certificate. The SB API has a method that will take care of the generation of the certificate for you. The certificate and key will then be located in the working directory of the binary so that it can be easily accessed by the program. To generate the certificate properly, you will need to fill out a couple fields before calling the `GenerateCertificate()` Method.

The following code block does two tasks: it first checks to see if a certificate exists, then it creates a certificate if needed:
```Go
if _, err := os.Stat(bot_cert); os.IsNotExist(err) {
    country := "Your Country"
    state   := "Your State (no abbreviations)"
    city    := "Your City"
    org     := "Your Organization"
    domain  := "Your Domain"
    smawk.GenerateCertificate(country,state,city,org,domain,bot_key,bot_cert)
}
```

## Generating a SB instance
To use the SB API you will need an instance of the bot to work with. To get the bot, call the `Connect` method and provide your personal token that was given to you by the BotFather.
Note that this method is static, and must be prefaced by the package name `smawk`. If you want to enable debug mode (and see info passed to the console), set the second argument of the the `Connect` method to true.

## Opening a webhook
Go and the SB API make it easy to open a webhook on your url. Telegram highly recommends using a personal token to make sure that only telegram requests come to your webhook. All you need to do to get the bot up and running on a webhook is provide the following code:
```Go
bot.OpenWebhookWithCert(bot_url+bot_url_token, bot_cert)
```
This code will pass in your self-signed certificate so that telegram can properly use TLS to encrypt the commands sent to the webhook.

## Listening for updates
The SB API gives you a method to listed on your token. Also, you will need to spin up a Go server to listen for the requests coming in. To spin the webserver up, you will need to add a line to your import block at the top of your code. When you use the `ListenAndServeTLS` method, pass in an IP address of 0.0.0.0:<your_port>. This will make sure that the server only listens locally.
```Go
import (
    "net/http"
)

updates := bot.Listen(bot_url_token)
go http.ListenAndServeTLS("0.0.0.0:8443", bot_cert, bot_key, nil)
```

## Sending the updates to the SB API
Once you have received an update, you will need to pass it to the SB API to be parsed and executed. The API will take care of any response that is needed for each of the commands. To properly send updates to the API, the following code block is needed:
```Go
for update := range updates {
    bot.ParseAndExecuteUpdate(update)
}
```
This block loops through the `updates` channel and sends each of updates through to the API. Telegram will attempt to send an update to your bot until it is accepted, so be careful about adding to many to the queue when the bot is disabled.

# Keeping the SB API up to date
From time to time, the API will be updated with fixes, new commands, etc. When the API is updated, all you have to is run the following command from your terminal to bring your local API up to date:
```Shell
cd $GOPATH
go get -u github.com/<your_github_user_name>/<your_controller_repo>
```
