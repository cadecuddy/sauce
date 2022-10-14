package cmd

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cadecuddy/sauce/utils"
	"github.com/spf13/cobra"
)

const URL_SEARCH string = "https://api.trace.moe/search?anilistInfo&url="

// urlCmd represents the url command
var urlCmd = &cobra.Command{
	Use:   "url",
	Short: "Search for anime source using the url to the media.",
	Long:  `Search for anime source using the url to the media.`,
	Run: func(cmd *cobra.Command, args []string) {
		urlSearch(args[0])
	},
}

// Makes a search via URL to the unidentified anime media.
//
// Checks the URL and once verififed, queries the trace.moe API,
// getting the core information of the show.
func urlSearch(linkToMedia string) {
	// validate URL
	_, err := url.ParseRequestURI(linkToMedia)
	if err != nil {
		fmt.Println("❌ Invalid URL")
		return
	}
	s := utils.ConfigSpinner()
	s.Start()

	// make GET request to trace.moe API
	resp, err := http.Get(URL_SEARCH + linkToMedia)
	identifiedAnime, malData := utils.HandleResponse(resp, err, s)

	s.Stop()

	// use highest similarity accuracy
	utils.PrintSauce(identifiedAnime, malData.Data)
}
