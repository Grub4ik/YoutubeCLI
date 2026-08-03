package download

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Video(url, directory string, asMP3 bool) error {
	ytdlpPath, err := exec.LookPath("yt-dlp")
	if err != nil {
		return fmt.Errorf("yt-dlp is not installed or is not available in PATH")
	}

	if _, err := exec.LookPath("ffmpeg"); err != nil {
		return fmt.Errorf("ffmpeg is not installed or is not available in PATH")
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

	args := []string{"--output", outputPath}

	if asMP3 {
		args = append(args,
			"--format", "ba/b",
			"--extract-audio",
			"--audio-format", "mp3",
			"--audio-quality", "0",
		)
	} else {
		args = append(args,
			"--format", "bv*[ext=mp4]+ba[ext=m4a]/b[ext=mp4]/b",
			"--merge-output-format", "mp4",
		)
	}

	// The separator prevents a URL beginning with '-' from being parsed as an option.
	args = append(args, "--", url)

	cmd := exec.CommandContext(context.Background(), ytdlpPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp failed: %w", err)
	}

	fmt.Println("Download complete!")
	fmt.Printf("Saved to: %s\n", saveDir)

	return nil
}
