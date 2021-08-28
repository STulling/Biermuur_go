package musicio

import (
	"github.com/STulling/Biermuur_go/musicio/musicutil"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"time"
)

func chk(err error) {
	if err != nil {
		panic(err)
	}
}

func AddSong(name string) {
	tmpFolder := path.Join(musicutil.MusicFolder, "tmp")
	tmpName := randomString(8)
	youtubePath, err := exec.LookPath("youtube-dl")
	chk(err)
	cmd := &exec.Cmd{
		Path: youtubePath,
		Args: []string{
			youtubePath,
			"-x",
			"-f bestaudio[ext=m4a]",
			"-x",
			"--postprocessor-args",
			"-ar 44100 -ac 2 -acodec libmp3lame -af loudnorm=I=-16:LRA=11:TP=-1.5",
			"-o",
			tmpFolder + "/" + tmpName + ".m4a",
			"ytsearch:" + name,
		},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Stdin: os.Stdin,
	}
	chk(cmd.Run())
	soxPath, err := exec.LookPath("sox")
	chk(err)
	soxCmd := &exec.Cmd{
		Path: soxPath,
		Args: []string{
			soxPath,
			"--norm=0",
			path.Join(tmpFolder, tmpName+".mp3"),
			path.Join(tmpFolder, tmpName+"-norm.mp3"),
		},
		Stderr: os.Stderr,
		Stdout: os.Stdout,
		Stdin: os.Stdin,
	}
	chk(soxCmd.Run())
	chk(os.Rename(path.Join(tmpFolder, tmpName+"-norm.mp3"),
		path.Join(musicutil.MusicFolder, name + ".mp3")))
	chk(os.Remove(path.Join(tmpFolder, tmpName+".mp3")))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func randomString(n int) string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}