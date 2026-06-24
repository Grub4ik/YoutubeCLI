package browser

import (
	"context"
	"fmt"

	"github.com/ytget/ytdlp"
)

func Download(url string) error {
	d := ytdlp.New()

	d = d.WithProgress(func(p ytdlp.Progress) {
		fmt.Printf("\rDownloading: %.1f%%", p.Percent)
	})

	info, err := d.Download(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Printf("\nSaved: %s\n", info.Title)
	return nil
}
