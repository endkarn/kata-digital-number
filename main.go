package main

import (
	"fmt"
	"kata-digital-number/number"
)

func main() {
	render := number.RenderLcdNumber(213)
	fmt.Println(render)
}
