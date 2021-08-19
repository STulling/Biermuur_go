package mathprocessor

import (
	"github.com/STulling/Biermuur_go/displaycontroller"
)

var (
	ToCalculate = make(chan []byte, 0)
)

func RunCalculationPipe() {
	for {
		<-ToCalculate
		//rms, tone := float64(data[0]), 0. //processing.ProcessBlock(data)
		displaycontroller.ToDisplay <- [2]float64{0.8, 0}
	}
}
