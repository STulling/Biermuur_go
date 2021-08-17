package effectlib

import (
	"math/rand"
	"time"

	"github.com/STulling/Biermuur_go/display"
)

func RandomRGB() uint32 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return display.RGBToColor(uint8(r1.Int()), uint8(r1.Int()), uint8(r1.Int()))
}
