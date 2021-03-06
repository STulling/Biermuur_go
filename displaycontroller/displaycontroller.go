package displaycontroller

import (
	"fmt"
	"github.com/STulling/Biermuur_go/display"
	"github.com/STulling/Biermuur_go/effectlib"
	"time"
)

const (
	offset = 0.05
)

var (
	callbacks = map[string]func(float64, float64){
		"wave":      effectlib.Wave,
		"debugwave": effectlib.DebugWave,
		"slowwave":  effectlib.SlowWave,
		"sparkle":   effectlib.Sparkle,
		"mond":      effectlib.Mond,
		"fill":      effectlib.Fill,
		"diamond":   effectlib.Ruit,
		"circle":    effectlib.Cirkel,
		"bars":      effectlib.Simple,
		"clear":     effectlib.Clear,
		"snake":     effectlib.Snake,
		"debug":     debug,
	}
	tone = 0.
	prevTime = time.Now()
)

func debug(arg1 float64, arg2 float64) {
	fmt.Println(time.Now().Sub(prevTime), arg1, arg2)
	prevTime = time.Now()
}

var (
	callback  = callbacks["wave"]
	ToDisplay = make(chan [2]float64, 0)
)

func SetCallback(name string) {
	callback = callbacks[name]
}

func RunDisplayPipe() {
	display.Init()
	for {
		data := <-ToDisplay
		display.Primary = effectlib.Wheel(uint8(data[1] * 255))
		callback(data[0], data[1])
	}
}
