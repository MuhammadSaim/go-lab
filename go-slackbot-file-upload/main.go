package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/MuhammadSaim/go-lab/go-slackbot-file-upload/config"
	"github.com/slack-go/slack"
)

func main() {
	// load the config file
	cfg := config.LoadConfig()

	// make a slack api instance
	slackAPI := slack.New(cfg.SLACK_BOT_TOKEN)

	// set the important data
	filesSlice := []string{"./files/martin_joo_proper_api_design_with_laravel.pdf"}

	// loop through the files
	for _, filePath := range filesSlice {

		// get the file name
		fileName := filepath.Base(filePath)

		// get the size of the file
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		// set params for file upload
		params := slack.UploadFileV2Parameters{
			Channel:  cfg.SLACK_CHANNEL_ID,
			File:     filePath,
			Filename: fileName,
			FileSize: int(fileInfo.Size()),
		}

		// upload the file
		file, err := slackAPI.UploadFileV2(params)
		if err != nil {
			fmt.Printf("Error! %v\n", err)
		}

		// Explicitly share the file link in the channel
		_, _, err = slackAPI.PostMessage(cfg.SLACK_CHANNEL_ID, slack.MsgOptionText(
			fmt.Sprintf("File uploaded successfully: <%s|%s>", file.ID, file.Title), false))
		if err != nil {
			fmt.Printf("Error sharing file link: %v\n", err)
		}

		fmt.Printf("Name: %v", file)

	}
}
