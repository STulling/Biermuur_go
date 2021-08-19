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
		print(len(data))
		//rms, tone := float64(data[0]), 0. //processing.ProcessBlock(data)
		displaycontroller.ToDisplay <- [2]float64{0.8, 0}
	}
}
