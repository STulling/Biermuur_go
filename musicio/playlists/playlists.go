package playlists

import (
	"path"

	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/musicio/musicutil"
)

var (
	playlistFolder = path.Join(musicutil.MusicFolder, "playlists")
)

func listPlaylists() []string {
	return musicutil.ListFiles(musicutil.MusicFolder)
}

func playPlaylist(name string) {
	audio.MusicQueue.PlayList = musicutil.ReadLines(name)
}
