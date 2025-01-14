package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/MuhammadSaim/go-lab/go-slackbot-age-calculator/config"
	"github.com/slack-io/slacker"
)

func main() {
	// load the config file
	cfg := config.LoadConfig()

	// make a slack api instance
	bot := slacker.NewClient(cfg.SLACK_BOT_TOKEN, cfg.SLACK_APP_TOKEN)

	// register the bot command
	bot.AddCommand(&slacker.CommandDefinition{
		Description: "Calculate age",
		Command:     "my yob is <year>",
		Examples:    []string{"my yob is 1990"},
		Handler: func(ctx *slacker.CommandContext) {
			// get the param
			year := ctx.Request().Param("year")

			// check params is not empty
			if year == "" {
				ctx.Response().Reply("Invalid year format")
				return
			}

			// convert string to an int
			yob, err := strconv.Atoi(year)
			if err != nil {
				ctx.Response().Reply("Not able to parse the year.")
				return
			}

			// get the current year
			currentYear := time.Now().Year()

			// check age year should not be greater then curret year
			if currentYear <= yob {
				ctx.Response().Reply("Your age year should be less then current year.")
				return
			}

			// Calculate the age
			age := time.Now().Year() - yob

			// return Response back to slack
			ctx.Response().Reply(fmt.Sprintf("Your age is %d years.", age))
		},
	})

	// context to run
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// bot lisntener
	err := bot.Listen(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
