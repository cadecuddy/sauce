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
	// red := color.New(color.FgRed)

	color.New(color.FgGreen).Add(color.Bold).Printf("✅ sauce found : [%f similarity]\n", res.Similarity)

	// regulate size & print top flower border
	size := getBorderSize(res.Anilist.Title.Romaji, res.Anilist.Title.English)
	if size >= 35 {
		size = 28
	}
	println(strings.Repeat("🌸", size))
	println()

	formatTitle(res.Anilist.Title.Native, res.Anilist.Title.English, size)
	formatType(res, malData)
	formatScore(malData.Score)
	b.Print("🏆 Ranking: ")
	color.New(color.FgHiMagenta).Printf("#%s\n", humanize.Comma(int64(malData.Rank)))
	b.Print("📕 Source: ")
	// color.New(color.FgRed).Add(color.FgYellow).Println("IS THIS ORANGE?")
	color.Red(" %s", malData.Source)
	// Movies won't have their year load from MAL Data
	if malData.Year != 0 {
		b.Print("📅 Year: ")
		color.Cyan("   %s %d", strings.Title(malData.Season), malData.Year)
	}
	formatGenre(malData.Genres)
	b.Print("🎬 Studio:  ")
	color.New(color.FgGreen).Println(malData.Studios[0].Name)

	println()
	println(strings.Repeat("🌸", size))
}

// Helper function to calculate the total flower border size
// and determine if the title needs to be translated.
func getBorderSize(nativeTitle string, englishTitle string) int {
	var borderLength = PREFIX_LENGTH

	if nativeTitle == englishTitle {
		borderLength += len(nativeTitle)
	} else {
		borderLength += len(nativeTitle) + len(englishTitle) + 3
	}

	return int(float32(borderLength))
}

func formatTitle(nativeTitle string, englishTitle string, borderSize int) {
	color.New(color.Bold).Print("✨ Anime:   ")

	title := fmt.Sprintf("%s (%s)", nativeTitle, englishTitle)
	color.New(color.Bold).Printf("%s\n", title)
}

func formatType(res types.Result, malData jikan.AnimeBase) {
	b := color.New(color.Bold)
	b.Print("❓ Type:    ")

	if malData.Type == "Movie" {
		fmt.Println("Movie 🎥")
		b.Print("🕐 Scene:  ")
		formatTimestamp(res.From, res.To)
		return
	} else {
		fmt.Println("TV Show 📺")
		formatEpisodes(res.Episode, malData.Episodes, res.From, res.To)
		return
	}

}

// Helper for formatting Episode information to output
func formatEpisodes(episode int, totalEpisodes int, timestampFrom float64, timestampTo float64) {
	color.New(color.Bold).Print("🕐 Episode: ")

	if totalEpisodes != 0 {
		fmt.Printf("%d/%d @", episode, totalEpisodes)
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
	r := color.New(color.FgBlue)

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
		color.New(color.FgHiGreen).Add(color.Bold).Println(score)
		return
	}
	if score >= 7 {
		color.New(color.FgCyan).Add(color.Bold).Println(score)
		return
	}

	color.New(color.FgRed).Add(color.Bold).Println(score)
}
