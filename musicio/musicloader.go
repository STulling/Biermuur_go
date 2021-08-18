package musicio

import (
	"log"
	"os"
	"path"

	"github.com/STulling/Biermuur_go/musicio/musicutil"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
)

func Load(file string) beep.Streamer {
	f, err := os.Open(path.Join(musicutil.MusicFolder, file) + ".mp3")
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, _ := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return streamer
}
