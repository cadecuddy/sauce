package utils

import (
	"fmt"
	"strings"

	"github.com/cadecuddy/sauce/types"
	"github.com/darenliang/jikan-go"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
)

const PREFIX_LENGTH int = 12

// Split printing statements into separate functions
// TODO: include a 'type' category to differentiate between TV and movie
// => separate formatEpisodes and make new getTimestamp function

func PrintSauce(res types.Result, malData jikan.AnimeBase) {
	// look @ https://github.com/fatih/color for color formatting
	b := color.New(color.Bold)
	red := color.New(color.FgRed)

	color.New(color.FgGreen).Add(color.Bold).Printf("✅ sauce found : [%f similarity]\n", res.Similarity)

	size, translate := getBorderSize(res.Anilist.Title.Romaji, res.Anilist.Title.English)
	println(strings.Repeat("🌸", size))
	println("🌸")
	formatTitle(res.Anilist.Title.Romaji, res.Anilist.Title.English, translate, size)
	formatEpisodes(res.Episode, malData.Episodes, res.From, res.To, malData.Type)

	// Format score based on how good it is
	if malData.Year != 0 {
		b.Print("🌸 📅 Year: ")
		color.Red("   %s %d", strings.Title(malData.Season), malData.Year)
	}
	b.Print("🌸 📈 Score:   ")
	red.Printf("%.2f\n", malData.Score)
	b.Print("🌸 🏆 Ranking: ")
	red.Printf("#%s\n", humanize.Comma(int64(malData.Rank)))
	b.Print("🌸 📕 Source: ")
	color.Red(" %s", malData.Source)

	formatGenre(malData.Genres)

	println("🌸")
	println(strings.Repeat("🌸", size))
}

// Helper function to calculate the total flower border size
// and determine if the title needs to be translated.
func getBorderSize(romanjiTitle string, englishTitle string) (int, bool) {
	var translateTitle bool
	var borderLength = PREFIX_LENGTH

	if romanjiTitle == englishTitle {
		borderLength += len(romanjiTitle)
		translateTitle = false
	} else {
		borderLength += len(romanjiTitle) + len(englishTitle) + 3
		translateTitle = true
	}

	// flower emoji is roughly 2 characters wide
	// (print len(prefix) + title) / 1.8 flower emojis
	return int(float32(float32(borderLength) / float32(1.8))), translateTitle
}

func formatTitle(romanjiTitle string, englishTitle string, translate bool, borderSize int) {
	b := color.New(color.Bold)
	var title string
	// borderChars := int(float32(borderSize) * float32(1.8))
	// only print english title if title is in english
	if translate {
		b.Print("🌸 ✨ Anime:   ")
		title = fmt.Sprintf("%s (%s)", romanjiTitle, englishTitle)
		// suffix := strings.Repeat(" ", borderChars-len(title)) + "🌸"
		color.New(color.FgRed).Printf("%s\n", title)
	} else {
		b.Print("🌸 ✨ Anime:   ")
		title = romanjiTitle
		// suffix := strings.Repeat(" ", borderChars-len(title)) + "🌸"
		color.New(color.FgRed).Printf("%s\n", title)
	}

}

func formatType()

// Helper for formatting Episode information to output
func formatEpisodes(episode int, totalEpisodes int, timestampTo float64, timestampFrom float64, mediaType string) {
	b := color.New(color.Bold)
	// If jikan can't get an accurate total episode count (usual of large series i.e Once Piece)
	// only print the episode number
	if mediaType == "Movie" {
		b.Printf("🌸 🎥 Movie:  ")
		color.New(color.FgHiBlue).Add(color.Bold).Printf(" [%s - %s]\n", ConvertTimestamp(timestampTo), ConvertTimestamp(timestampFrom))
		return
	}

	b.Printf("🌸 📺 Episode: ")
	if totalEpisodes != 0 {
		color.New(color.FgRed).Printf("%d/%d @", episode, totalEpisodes)
	} else {
		color.New(color.FgRed).Printf("%d @", episode)
	}
	color.New(color.FgHiBlue).Add(color.Bold).Printf(" [%s - %s]\n", ConvertTimestamp(timestampTo), ConvertTimestamp(timestampFrom))
}

func formatGenre(genres []jikan.MalItem) {
	b := color.New(color.Bold)
	r := color.New(color.FgRed)

	b.Print("🌸 📜 Genres:  ")
	for i, genre := range genres {
		if i != 0 {
			r.Print(", " + genre.Name)
		} else {
			r.Print(genre.Name)
		}
	}
	b.Print("\n")
}
