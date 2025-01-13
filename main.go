package main

import (
	"calculator/math"
	"calculator/physical"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	println(e)

	sum := math.Sum(1, 2)
	fmt.Println(sum)

	sub := math.Sub(3, 1)
	fmt.Println(sub)

	leis := physical.LeisNewton(5)
	fmt.Println(leis)
}
