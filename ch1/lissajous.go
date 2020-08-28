package main

import(
    "image"
    "image/color"
    "image/gif"
    "io"
    "math"
    "math/rand"
    "os"
    "time"
)

var palette = []color.Color{
    color.White,
    color.Black
}

const (
    whiteIndex = 0  //first color in palette
    blackIndex = 1  //next color in palette
)

func main() {
    // The sequence of images is deterministic unless we seed
    // the pseudo-random number generator using the current time.
    rand.Seed(time.Now().UTC().UnixNano())
    lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
    const (
        cycles  = 5      // number of complete x oscillator revolutions
        res     = 0.001  // angular resolution
        size    = 100    // image canvas covers [-size..+size]
        nframes = 64     // number of animation frames
        delay   = 8      // delay between frames in 10ms unites
    )

    freq := rand.Float64() * 3.0    //relative frequence of y oscillator
}
