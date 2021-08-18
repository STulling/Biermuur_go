package effectlib

import (
	"fmt"
	"math"

	"github.com/STulling/Biermuur_go/display"
)

var (
	t       float64 = 0
	x_array []int   = make([]int, display.Width)
)

func Wave(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	dt := 0.1 * (1 + 3*pitch)
	t += dt
	for x := 0; x < display.Width; x++ {
		x_val := 3. * math.Pi * float64(x) / (display.Width - 1)
		x_array[x] = int(rms*display.Height/2*math.Sin(x_val+t) + display.Height/2)
	}
	for x := 0; x < display.Width; x++ {
		display.SetPixelColor(x, int(x_array[x]), display.Primary)
		display.SetPixelColor(x, int(x_array[x]-1), display.Primary)
	}
	display.Strip.Render()
}

func Clear(rms float64, pitch float64) {
	display.Init()
	fmt.Println("here")
	display.Clear()
	fmt.Println("there")
	display.Strip.Render()
	fmt.Println("done")
}
