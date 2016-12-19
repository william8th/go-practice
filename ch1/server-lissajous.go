package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

var palette = []color.Color{
	color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
	color.RGBA{0x00, 0xFF, 0x00, 0xFF},
	color.RGBA{0x00, 0x00, 0xFF, 0xFF},
}

const (
	whiteIndex = 0
	greenIndex = 1
	blueIndex  = 2
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		cycles, _ := strconv.Atoi(req.Header.Get("cycles"))
		// We don't care about error as long as we get a 0 by default
		if cycles <= 0 {
			cycles = 5
		}
		lissajous(w, float64(cycles))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func lissajous(out io.Writer, cycles float64) {
	const (
		//cycles = number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [--size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {

		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {

			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			var colourIndex uint8 = greenIndex
			if i%2 == 0 {
				colourIndex = blueIndex
			}

			img.SetColorIndex(
				size+int(x*size+0.5),
				size+int(y*size+0.5),
				colourIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
}
