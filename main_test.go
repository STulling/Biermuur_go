package main

import (
	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/STulling/Biermuur_go/mathprocessor"
	"github.com/faiface/beep"
	"testing"
	"time"
)

func Test(t *testing.T) {
	go displaycontroller.RunDisplayPipe()
	go mathprocessor.RunCalculationPipe()
	audio.Init(beep.SampleRate(44100), 8*1024)
	audio.Play()
	audio.MusicQueue.AddSong("BENEE")
	time.Sleep(time.Second * 10)
}
