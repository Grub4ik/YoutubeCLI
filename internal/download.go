package browser

import (
	"context"
	"fmt"

	"github.com/lrstanley/go-ytdlp"
)

func Download(url string) error {
	ytdlp.MustInstall(context.Background(), nil)

	dl := ytdlp.New().
		Output("~/Downloads/%(title)s.%(ext)s")

	_, err := dl.Run(context.Background(), url)
	if err != nil {
		return fmt.Errorf("ошибка скачивания: %w", err)
	}

	fmt.Printf("Saved: ~/Downloads")
	return nil
}
