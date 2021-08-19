package mathprocessor

import (
	"fmt"
	"time"

	"github.com/STulling/Biermuur_go/displaycontroller"
)

var (
	ToCalculate = make(chan []byte, 0)
	prevTime    = time.Now()
)

func RunCalculationPipe() {
	for {
		data := <-ToCalculate
		fmt.Println(time.Since(prevTime))
		prevTime = time.Now()
		fmt.Println(data)
		//rms, tone := processing.ProcessBlock(data)
		displaycontroller.ToDisplay <- [2]float64{0.8, 0}
	}
}
