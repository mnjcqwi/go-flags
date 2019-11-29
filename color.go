package main

import (
	"flag"
	"fmt"
	"image/color"
	"strconv"
)

func main() {
	fg := flag.String("fg", "000000", "background color")
	flag.Parse()
	v, err := strconv.ParseInt(*fg, 16, 64)
	if err != nil {
		return
	}
	b := uint8(v % 256)
	g := uint8(v / 256 % 256)
	r := uint8(v / (256 * 256) % 256)
	c := color.RGBA{R: r, G: b, B: g, A: 255}
	draw(c)
}

func draw(c color.Color) {
	r, g, b, a := c.RGBA()
	r, g, b, a = r/256, g/256, b/256, a/256
	fmt.Printf("drawing background with RGBA(%v,%v,%v,%v)", r, g, b, a)

}
