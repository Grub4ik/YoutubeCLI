# YoutubeCLI

YoutubeCLI is a command-line tool for Linux that lets you search, open, and download YouTube videos directly from the terminal.
# Requirements

Go 1.21 or higher, yt-dlp, ffmpeg.
Installation

Clone the repository, build the binary, and install it system-wide:

git clone https://github.com/Grub4ik/youtubecli.git
cd youtubecli
go build -o yt
sudo cp yt /usr/local/bin/

After this, the yt command is available from anywhere in your terminal.
# Usage

Open YouTube homepage:
yt

Search for videos:
yt "never gonna give you up"

Download a video or a playlist:
yt -d "https://www.youtube.com/watch?v=dQw4w9WgXcQ"

Download a video or a playlist as mp3:
yt -d "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -mp3

Download a video or a playlist choosing save directory:
yt -d "https://www.youtube.com/watch?v=dQw4w9WgXcQ" -a ~/Downloads

Show help:
yt -h
# How it works

The project is structured into three main components:
main.go handles command-line arguments and routes calls.
internal/browser.go opens the browser using xdg-open on Linux, open on macOS, or start on Windows.
internal/download.go downloads videos via yt-dlp with format and output path parameters.

Videos are saved to the ~/Downloads folder in the best available quality.
# Dependencies

Install required system utilities:

Ubuntu / Debian:
sudo apt install yt-dlp ffmpeg

Arch Linux:
sudo pacman -S yt-dlp ffmpeg

Fedora:
sudo dnf install yt-dlp ffmpeg

The Go wrapper for yt-dlp installs automatically on first run.
# Troubleshooting

Download errors: make sure yt-dlp and ffmpeg are installed:
yt-dlp --version
ffmpeg -version

Browser doesn't open: check that xdg-open is installed:
which xdg-open

Install it if missing:
sudo apt install xdg-utils
# Development

Build and run locally:
go build -o yt
./yt "test query"

Format and lint:
go fmt ./...
go vet ./...
# Roadmap

Playlist support, quality and format selection, direct streaming to media players.
License

MIT License
# Author

Grub4ik
# Contributing

Issues and pull requests are welcome. Open an issue on GitHub or contact the author directly.

Thanks for using YoutubeCLI!
