package displaycontroller

import (
	"fmt"

	"github.com/STulling/Biermuur_go/display"
	"github.com/STulling/Biermuur_go/effectlib"
)

var (
	callbacks = map[string]func(float64, float64){
		"wave":  effectlib.Wave,
		"clear": effectlib.Clear,
		"debug": debug,
	}
)

func debug(arg1 float64, arg2 float64) {
	fmt.Println(fmt.Sprintf("%f", arg1) + " " + fmt.Sprintf("%f", arg2))
}

var (
	callback  func(float64, float64) = callbacks["wave"]
	ToDisplay chan [2]float64        = make(chan [2]float64, 0)
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
