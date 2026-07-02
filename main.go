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
		outputdir = flag.String("a", "", "Output irection")
	)

	flag.Parse()

	args := flag.Args()

	if *download != "" {
		if err := videoDownload.Video(*download, *outputdir); err != nil {
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
