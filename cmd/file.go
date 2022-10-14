package cmd

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/cadecuddy/sauce/utils"
	"github.com/spf13/cobra"
)

const FILE_SEARCH = "https://api.trace.moe/search"

// fileCmd represents the file command
var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "Search for anime source using a path to local file.",
	Long: `Search for anime source using a path to local file.
	
File searches support most visual file mediums under 25MB including:
- png
- jpeg/jpg
- gif
- webp
- many more!
`,
	Run: func(cmd *cobra.Command, args []string) {
		fileSearch(args[0])
	},
	Args: cobra.MinimumNArgs(1),
}

// Searches for anime based on file path input
func fileSearch(filepath string) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("❌ Invalid file")
		return
	}
	defer file.Close()

	s := utils.ConfigSpinner()
	s.Start()

	// form file upload via https://gist.github.com/andrewmilson/19185aab2347f6ad29f5
	buffer := &bytes.Buffer{}
	writer := multipart.NewWriter(buffer)
	part, _ := writer.CreateFormFile("file", filepath)
	io.Copy(part, file)
	writer.Close()

	res, err := http.Post(URL_SEARCH, writer.FormDataContentType(), buffer)
	identifiedAnime, malData := utils.HandleResponse(res, err, s)

	s.Stop()

	// do things
	utils.PrintSauce(identifiedAnime, malData.Data)
}
