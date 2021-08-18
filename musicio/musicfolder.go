package musicio

import "github.com/STulling/Biermuur_go/musicio/musicutil"

func ListSongs() []string {
	return musicutil.ListFiles(musicutil.MusicFolder)
}
