package main

import (
	"fmt"
	"github.com/STulling/Biermuur_go/audio"
	"github.com/STulling/Biermuur_go/displaycontroller"
	"github.com/STulling/Biermuur_go/mathprocessor"
	"github.com/faiface/beep"
	"syscall"
	"testing"
	"time"
)

func Test(t *testing.T) {
	go displaycontroller.RunDisplayPipe()
	go mathprocessor.RunCalculationPipe(44100)
	audio.Init(beep.SampleRate(44100), 16)
	audio.Play()
	audio.MusicQueue.AddSong("BENEE")
	time.Sleep(time.Second * 10)
}

/*

	Ok, deze shit werkt wellicht op linux,
	windows timer is cringe. Max 15.5 ofzo timer resolution
	Niet goed genoeg, oftewel tijd voor linux debugging....

	zodra ik terug ben in delft

	Oftewel Simon kom van je luie reet en fix deze shit,
	niet meer kloten met audio libraries maar hele gore time based shit


	NEVERMIND BITCH
 */

var (
	modwinmm    = syscall.NewLazyDLL("winmm.dll")
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	proctimeBeginPeriod = modwinmm.NewProc("timeBeginPeriod")
	proctimeEndPeriod   = modwinmm.NewProc("timeEndPeriod")

	procCreateEvent = modkernel32.NewProc("CreateEventW")
	procSetEvent    = modkernel32.NewProc("SetEvent")
)

func timeBeginPeriod(period uint32) {
	syscall.Syscall(proctimeBeginPeriod.Addr(), 1, uintptr(period), 0, 0)
}

func Test2(t *testing.T) {
	timeBeginPeriod(1)
	ticker := time.NewTicker(time.Second/100)
	timey := time.Now()
	for {
		select {
		case t := <-ticker.C:
			fmt.Println(t.Sub(timey))
			timey = time.Now()
		}
	}
}