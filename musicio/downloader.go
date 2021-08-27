package musicio

import (
	"fmt"
	"github.com/STulling/Biermuur_go/musicio/musicutil"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"strings"
)

func AddSong(name string) {
	tmpFolder := path.Join(musicutil.MusicFolder, "tmp")
	tmpName := randomString(8)
	command := fmt.Sprintf("youtube-dl -x -f bestaudio -x --audio-format mp3 --postprocessor-args \"-ar 44100 -ac 2\" -o \"%s/%s.%%(ext)s\" \"ytsearch1:%s\"",
		tmpFolder, tmpName, name)
	exec.Command(command).Run()
	volume := findVolume(path.Join(tmpFolder, tmpName))
	normalizeFile(path.Join(tmpFolder, tmpName), path.Join(musicutil.MusicFolder, name + ".mp3"), volume)
	e := os.Remove(path.Join(tmpFolder, tmpName))
	if e != nil {
		panic(e)
	}
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func removeHeader(s string) string {
	if idx := strings.Index(s, "max_volume: "); idx != -1 {
		return s[idx+12:]
	}
	return s
}

func findVolume(inputFile string) string {
	// Get the peak of the audio file using 'volumedetect'
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-af", "volumedetect", "-f", "null", "-y", "nul")

	// Output of the command
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	// Start command
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// Read the output of the command
	ffmpegCLIOutput, _ := ioutil.ReadAll(stderr)

	// Convert the command output from bytes to a string
	convertCLIOutput := string(ffmpegCLIOutput[:])

	// Index the command output and output the 'max_volume" value
	currentVol := removeHeader(convertCLIOutput)

	// Remove the 'dB' extension
	var removedDB string

	if strings.Contains(currentVol[:], "0.0 dB") {
		removedDB = currentVol[0:4]
	} else {
		removedDB = currentVol[1:5]
	}
	normalizedVol := "volume=" + removedDB

	return normalizedVol
}

func normalizeFile(inputFile, outputFile, normalizedVol string) {
	cmd2 := exec.Command("ffmpeg", "-i", inputFile, "-filter:a", normalizedVol, outputFile)
	if err := cmd2.Start(); err != nil {
		panic(err)
	}
}

