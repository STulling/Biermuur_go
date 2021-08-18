package musicutil

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var (
	MusicFolder = initMusicFolder()
)

func initMusicFolder() string {
	envVar := os.Getenv("FLASK_MEDIA_DIR")
	if envVar == "" {
		return "."
	}
	return envVar
}

func ReadLines(name string) []string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := make([]string, 0)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		result = append(result, strings.TrimSuffix(scanner.Text(), "\n"))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}

func ListFiles(path string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() {
			result = append(result, f.Name())
		}
	}
	return result
}

func ListFilesExtension(path string, ext string) []string {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	result := make([]string, 0)
	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ext) {
			result = append(result, f.Name())
		}
	}
	return result
}
