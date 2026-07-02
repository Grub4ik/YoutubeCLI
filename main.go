package main

import (
	"flag"
	"fmt"
	browser "youtubecli/internal/browser"
	videoDownload "youtubecli/internal/download"
)

func OpenMainPage() {
	browser.OpenBrowser("https://www.youtube.com/")
}

func main() {
	var (
		download  = flag.String("d", "", "Download video")
		outputdir = flag.String("a", "", "Output direction")
		mp3       = flag.Bool("mp3", false, "Download as MP3 (audio only)")
		help      = flag.Bool("h", false, "Show help")
	)

	flag.Parse()

	args := flag.Args()

	if *help {
		fmt.Println(`
		Usage:
  youtubecli [flags] [search query]

Flags:
  -h          Show this help
  -d <url>    Download video
  -a <dir>    Output directory (default: current dir)

Examples:
  youtubecli                      Open YouTube homepage
  youtubecli "cats"               Search and open first video
  youtubecli -d "https://..."     Download video
  youtubecli -d "url" -a ./videos Download video to ./videos
		`)
		return
	}

	if *download != "" {
		if err := videoDownload.Video(*download, *outputdir, *mp3); err != nil {
			fmt.Printf("Error while downloading: %v\n", err)
		}
		return
	}

	if len(args) == 0 {
		OpenMainPage()
		return
	}

	if len(args) >= 1 {
		browser.OpenBrowser("https://www.youtube.com/results?search_query=" + args[0])
		return
	}

}
