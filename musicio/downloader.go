package musicio

import (
	"fmt"
	"os/exec"

	"github.com/STulling/Biermuur_go/musicio/musicutil"
)

func AddSong(name string) {
	command := fmt.Sprintf("youtube-dl -x -f bestaudio -x --audio-format mp3 --postprocessor-args \"-ar 44100 -ac 2\" -o \"%s/%%(title)s.%%(ext)s\" \"ytsearch1:%s\"", musicutil.MusicFolder, name)
	exec.Command(command)
}
