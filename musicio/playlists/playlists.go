package playlists

import (
	"os"
	"path"

	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/musicio/musicutil"
)

var (
	playlistFolder = path.Join(musicutil.MusicFolder, "playlists")
)

func chk(err error)  {
	if err != nil {
		panic(err)
	}
}

func ListPlaylists() []string {
	return musicutil.ListFiles(playlistFolder)
}

func RemovePlaylist(name string) {
	chk(os.Remove(path.Join(playlistFolder, name)))
}

func AddPlaylist(name string, song string) {
	f, err := os.OpenFile(path.Join(playlistFolder, name),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	chk(err)
	defer f.Close()
	if _, err := f.WriteString(song + "\n"); err != nil {
		panic(err)
	}
}

func NewPlaylist(name string) {
	f, err := os.Create(path.Join(playlistFolder, name))
	chk(err)
	f.Close()
}

func PlayPlaylist(name string) {
	audio.MusicQueue.PlayList = musicutil.ReadLines(path.Join(playlistFolder, name))
}
