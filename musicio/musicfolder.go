package musicio

import (
	"strings"

	"github.com/STulling/Biermuur_go/musicio/musicutil"
)

func ListSongs() []string {
	songs := musicutil.ListFiles(musicutil.MusicFolder)
	for i, song := range songs {
		songs[i] = strings.TrimSuffix(song, ".mp3")
	}
	return songs
}
