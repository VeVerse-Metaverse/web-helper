// Package cmd
// Author: Egor Pristavka <e@veverse.com>
// Copyright © 2023 LE7EL AS
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"web-helper/internal"
)

// ytCmd represents the yt command
var ytCmd = &cobra.Command{
	Use:   "yt",
	Short: "Get the YT video details",
	Long:  `Request the YT video details from the YT API and return the details in JSON format.`,
	Run: func(cmd *cobra.Command, args []string) {
		videoId, _ := cmd.Flags().GetString("videoId")
		if videoId == "" {
			cmd.Help()
			return
		}

		response, err := internal.GetPlayerResponse(videoId)
		if err != nil {
			return
		}

		serializedResponse, err := json.Marshal(response)
		if err != nil {
			return
		}

		cmd.Println(string(serializedResponse))
	},
}

func init() {
	rootCmd.AddCommand(ytCmd)

	ytCmd.Flags().StringP("videoId", "v", "", "The video ID")
}
