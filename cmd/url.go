package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/cadecuddy/sauce/types"
	"github.com/cadecuddy/sauce/utils"
	"github.com/darenliang/jikan-go"
	"github.com/spf13/cobra"
)

const BASE_URL string = "https://api.trace.moe/search?anilistInfo&url="

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
// getting the core information of the show. Once the response has
// been recieved, a 3rd party MAL (MyAnimeList) API is queried to get more
// detailed show data to supplement the trace.moe data.
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
	resp, err := http.Get(BASE_URL + linkToMedia)
	if err != nil {
		fmt.Println("❌ Error with request")
		s.Stop()
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read body to buffer
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Unmarshall JSON to custom trace.moe response type
	var traceMoeResponse types.MoeResponse
	json.Unmarshal(body, &traceMoeResponse)
	if traceMoeResponse.Error != "" {
		s.Stop()
		fmt.Println("❌ URL yielded no results")
		return
	}

	// Use jikan API for MAL data
	identifiedAnime := traceMoeResponse.Result[0]
	malData, err := jikan.GetAnimeById(identifiedAnime.Anilist.MalID)
	if err != nil {
		s.Stop()
		fmt.Println("❌ Error getting MAL data")
		return
	}

	s.Stop()

	// use highest similarity accuracy
	utils.PrintSauce(identifiedAnime, malData.Data)
}

func init() {
	rootCmd.AddCommand(urlCmd)
}
