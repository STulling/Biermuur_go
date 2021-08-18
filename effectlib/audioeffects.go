package effectlib

import (
	"math"

	"github.com/STulling/Biermuur_go/display"
)

var (
	t       float64     = 0
	x_array []int       = make([]int, display.Width)
	snake   [][2]uint32 = make([][2]uint32, display.Width)
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
		display.SetPixelColor(x, x_array[x], display.Primary)
		display.SetPixelColor(x, x_array[x]-1, display.Primary)
	}
	display.Render()
}

func DebugWave(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	dt := 0.3 //* (1 + 3*pitch)
	t += dt
	for x := 0; x < display.Width; x++ {
		x_val := 3. * math.Pi * float64(x) / (display.Width - 1)
		x_array[x] = int(rms*display.Height/2*math.Sin(x_val+t) + display.Height/2)
	}
	for x := 0; x < display.Width; x++ {
		display.SetPixelColor(x, x_array[x], display.RGBToColor(0, 255, 0))
		display.SetPixelColor(x, x_array[x]-1, display.RGBToColor(0, 255, 0))
	}
	display.Render()
}

func Snake(rms float64, pitch float64) {
	display.SetStrip(display.Secondary)
	color := display.Primary
	height := uint32(pitch * display.Height)
	snake = append(snake[1:], [2]uint32{height, color})
	for i, data := range snake {
		display.SetPixelColor(i, int(data[0]), data[1])
		display.SetPixelColor(i, int(data[0])+1, data[1])
	}
	display.Render()
}

func Clear(rms float64, pitch float64) {
	display.Clear()
	display.Render()
}
