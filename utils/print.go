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

func PrintSauce(res types.Result, malData jikan.AnimeBase) {
	// look @ https://github.com/fatih/color for color formatting
	b := color.New(color.Bold)
	red := color.New(color.FgRed)

	color.New(color.FgGreen).Add(color.Bold).Printf("✅ sauce found : [%f similarity]\n", res.Similarity)

	size, translate := getBorderSize(res.Anilist.Title.Romaji, res.Anilist.Title.English)
	println(strings.Repeat("🌸", size))
	println()
	formatTitle(res.Anilist.Title.Native, res.Anilist.Title.English, translate, size)
	formatType(res, malData)

	// Format score based on how good it is
	if malData.Year != 0 {
		b.Print("📅 Year: ")
		color.Red("   %s %d", strings.Title(malData.Season), malData.Year)
	}
	formatScore(malData.Score)
	b.Print("🏆 Ranking: ")
	red.Printf("#%s\n", humanize.Comma(int64(malData.Rank)))
	b.Print("📕 Source: ")
	color.Red(" %s", malData.Source)

	formatGenre(malData.Genres)

	println()
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

	return int(float32(float32(borderLength) / float32(1.8))), translateTitle
}

func formatTitle(romanjiTitle string, englishTitle string, translate bool, borderSize int) {
	b := color.New(color.Bold)
	var title string

	// only print english title if title is in english
	if translate {
		b.Print("✨ Anime:   ")
		title = fmt.Sprintf("%s (%s)", romanjiTitle, englishTitle)
		// suffix := strings.Repeat(" ", borderChars-len(title)) + "🌸"
		color.New(color.FgRed).Printf("%s\n", title)
	} else {
		b.Print("✨ Anime:   ")
		title = romanjiTitle
		// suffix := strings.Repeat(" ", borderChars-len(title)) + "🌸"
		color.New(color.FgRed).Printf("%s\n", title)
	}
}

func formatType(res types.Result, malData jikan.AnimeBase) {
	b := color.New(color.Bold)
	b.Print("❓ Type:    ")

	if malData.Type == "Movie" {
		color.New(color.FgRed).Println("Movie 🎥")
		b.Print("🕐 Scene:  ")
		formatTimestamp(res.From, res.To)
		return
	} else {
		color.New(color.FgRed).Println("TV Show 📺")
		formatEpisodes(res.Episode, malData.Episodes, res.From, res.To)
		return
	}

}

// Helper for formatting Episode information to output
func formatEpisodes(episode int, totalEpisodes int, timestampTo float64, timestampFrom float64) {
	color.New(color.Bold).Print("🕐 Episode: ")

	if totalEpisodes != 0 {
		color.New(color.FgRed).Printf("%d/%d @", episode, totalEpisodes)
	} else {
		color.New(color.FgRed).Printf("%d @", episode)
	}
	formatTimestamp(timestampFrom, timestampTo)
}

func formatTimestamp(from float64, to float64) {
	color.New(color.FgHiBlue).Add(color.Bold).Printf(" [%s - %s]\n", ConvertTimestamp(from), ConvertTimestamp(to))
}

// Prints the anime's genres as found on MAL
func formatGenre(genres []jikan.MalItem) {
	color.New(color.Bold).Print("📜 Genres:  ")
	r := color.New(color.FgRed)

	for i, genre := range genres {
		if i != 0 {
			r.Print(", " + genre.Name)
		} else {
			r.Print(genre.Name)
		}
	}
	print("\n")
}

// Prints the score with a color dependent on how high it is.
func formatScore(score float64) {
	color.New(color.Bold).Print("📈 Score:   ")

	if score >= 8 {
		color.New(color.FgHiGreen).Println(score)
		return
	}
	if score >= 7 {
		color.New(color.FgHiYellow).Println(score)
		return
	}

	color.New(color.FgHiRed).Println(score)
}
