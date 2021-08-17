package display

import (
	ws281x "github.com/rpi-ws281x/rpi-ws281x-go"
)

const (
	brightness = 255
	ledCounts  = 360
	gpioPin    = 21
	freq       = 800000
	width      = 20
	height     = 18
)

var (
	strip ws
)

type ws struct {
	ws2811 *ws281x.WS2811
}

func (ws *ws) init() error {
	err := ws.ws2811.Init()
	if err != nil {
		return err
	}

	return nil
}

func (ws *ws) close() {
	ws.ws2811.Fini()
}

func (ws *ws) leds() []uint32 {
	return strip.ws2811.Leds(0)
}

func SetPixelColor(x int, y int, color uint32) {
	if x < 0 || y < 0 {
		return
	}
	if x >= width || y >= height {
		return
	}
	if y%2 == 1 {
		x = width - 1 - x
	}
	strip.leds()[x+y*width] = color
}

func Clear() {
	for i := 0; i < len(strip.leds()); i++ {
		strip.leds()[i] = 0
	}
}

func RGBToColor(r uint8, g uint8, b uint8) uint32 {
	return uint32(uint32(r)<<16 | uint32(g)<<8 | uint32(b))
}

func Init() {
	opt := ws281x.DefaultOptions
	opt.Channels[0].Brightness = brightness
	opt.Channels[0].LedCount = ledCounts
	opt.Channels[0].GpioPin = gpioPin
	opt.Frequency = freq

	ws2811, err := ws281x.MakeWS2811(&opt)
	if err != nil {
		panic(err)
	}

	strip := ws{
		ws2811: ws2811,
	}

	err = strip.init()
	if err != nil {
		panic(err)
	}
	defer strip.close()
}
