package io

import (
	"log"
	"os"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
)

func Load(file string) beep.Streamer {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	streamer, _, _ := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	return streamer
}
