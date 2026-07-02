package download

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lrstanley/go-ytdlp"
)

func Video(url, directory string) error {
	_, err := ytdlp.Install(context.Background(), nil)
	if err != nil {
		return fmt.Errorf("failed to install yt-dlp: %w", err)
	}

	var outputPath string
	var saveDir string

	if directory != "" {
		saveDir = directory
		outputPath = filepath.Join(directory, "%(title)s.%(ext)s")
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("cannot get home dir: %w", err)
		}
		saveDir = filepath.Join(home, "Downloads")
		outputPath = filepath.Join(saveDir, "%(title)s.%(ext)s")
	}

	if saveDir != "" {
		if err := os.MkdirAll(saveDir, 0755); err != nil {
			return fmt.Errorf("cannot create directory: %w", err)
		}
	}

	dl := ytdlp.New().Output(outputPath)
	result, err := dl.Run(context.Background(), url)
	if err != nil {
		return fmt.Errorf("download error: %w", err)
	}

	fmt.Printf("Download complete!\n")
	fmt.Printf("Saved to: %s\n", outputPath)
	fmt.Printf("Title: %s\n", result)

	return nil
}
