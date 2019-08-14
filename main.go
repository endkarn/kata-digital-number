package main

import (
	"fmt"
	"kata-digital-number/number"
)

func main() {
	render := number.RenderLcdNumber(1234567890)
	fmt.Println(render)
}
