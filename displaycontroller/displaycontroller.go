package displaycontroller

import (
	"fmt"

	"github.com/STulling/Biermuur_go/effectlib"
)

var (
	callbacks = map[string]func(float64, float64){
		"wave":  effectlib.Wave,
		"debug": debug,
	}
)

func debug(arg1 float64, arg2 float64) {
	fmt.Println(fmt.Sprintf("%f", arg1) + " " + fmt.Sprintf("%f", arg2))
}

var (
	callback  func(float64, float64)
	ToDisplay chan [2]float64 = make(chan [2]float64, 0)
)

func SetCallback(name string) {
	callback = callbacks[name]
}

func RunDisplayPipe() {
	for {
		data := <-ToDisplay
		callback(data[0], data[1])
	}
}