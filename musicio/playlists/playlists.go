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

func PlayPlaylist(name string) {
	audio.MusicQueue.PlayList = musicutil.ReadLines(path.Join(playlistFolder, name))
}
