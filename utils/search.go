package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/briandowns/spinner"
	"github.com/cadecuddy/sauce/types"
	"github.com/darenliang/jikan-go"
)

// Analyzes initial trace.moe API response and handles errors.
// A 3rd party MAL (MyAnimeList) API is queried to get more
// detailed anime data to supplement the trace.moe data.
func HandleResponse(res *http.Response, err error, s *spinner.Spinner) (types.Result, jikan.AnimeById) {
	if err != nil {
		fmt.Println("❌ Error with request")
		s.Stop()
		os.Exit(1)
	}
	defer res.Body.Close()

	// read body to buffer
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		os.Exit(1)
	}

	// Unmarshall JSON to custom trace.moe response type
	var traceMoeResponse types.MoeResponse
	json.Unmarshal(body, &traceMoeResponse)
	if traceMoeResponse.Error != "" {
		s.Stop()
		fmt.Println("❌ Invalid Media")
		os.Exit(1)
	}

	// Query jikan API for MAL data
	identifiedAnime := traceMoeResponse.Result[0]
	malData, err := jikan.GetAnimeById(identifiedAnime.Anilist.MalID)
	if err != nil {
		s.Stop()
		fmt.Println("❌ Error getting MAL data")
		os.Exit(1)
	}

	return identifiedAnime, *malData
}
