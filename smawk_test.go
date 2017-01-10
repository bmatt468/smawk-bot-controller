package main_test

import (
    "github.com/SMAWK/smawk-bot"
    "log"
    "testing"
)

// Create our constants for use throughout the testing functions
const (
	SMAWKToken              = "249930361:AAHz1Gksb-eT0SQG47lDb7WbJxujr7kGCkU"
	ChatID                 	= 55997207
)

/* ================================================ */
/*                 Helper functions                 */
/* ================================================ */

// GetBot is a helper function to return an instance of
// a SmawkBot struct to each of the testing methods
func GetBot(t *testing.T) (*smawk.SmawkBot, error) {
	return smawk.Connect(SMAWKToken, true)
}

/* ================================================ */
/*                Testing functions                 */
/* ================================================ */

// TestLoadBot tests to see if the bot is loading and authenticated properly
func TestLoadBot(t *testing.T) {
	// Fetch our bot using the helper function
	_, err := GetBot(t)

	// Check to see if something bad happened and break if need be
	if err != nil {
		log.Fatal(err)
		t.FailNow()
	}
}

// TestParseHello emulates the /hello{@smawk_bot} command and responds as it does in production
func TestParseHello(t *testing.T) {
	// Fetch our bot using the helper function
	bot, _ := GetBot(t)

	// Generate our update using the helper function
	update, _  := bot.GenerateUpdate("/id")
	bot.ParseAndExecuteUpdate(update)
}
