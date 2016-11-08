package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math/rand"
)

// rNVolume is a basic type to hold the layer's information
type rNVolume struct {
	Fields []mat64.Dense
}

func (vol rNVolume) Dims() (int, int, int) {
	i1, i2 := vol.Fields[0].Dims()
	i3 := len(vol.Fields)
	return i1, i2, i3

}

//newRNVolume generates a rNVolume of fixed size filled with zeros
func newRNVolume(h int, w int, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *mat64.NewDense(h, w, nil))
	}
	return v
}

//newRNVolumeRandom generates a rNVolume of fixed size filled with values between 0 and 1
func newRNVolumeRandom(h int, w int, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	data := make([]float64, w*h)
	for i := range data {
		data[i] = rand.Float64()
	}
	a := mat64.NewDense(w, h, data)

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *a)
	}
	return v
}

func (vol rNVolume) Height() int {
	return len(vol.Fields)
}

func (vol rNVolume) Print() {

	for i := range vol.Fields {
		fa := mat64.Formatted(&vol.Fields[i], mat64.Prefix(" "))
		fmt.Printf("Layer %v:\n\n %v\n\n", i, fa)
	}
}

func (vol *rNVolume) Width() int {
	_, c := vol.Fields[0].Dims()
	return c
}

func (vol rNVolume) Depth() int {
	r, _ := vol.Fields[0].Dims()
	return r
}

//Equal3Dim checks if the size of two volumes are the same
func (vol rNVolume) EqualSize(a rNVolume) bool {
	i1, i2, i3 := vol.Dims()
	e1, e2, e3 := a.Dims()
	return Equal3Dim(i1, i2, i3, e1, e2, e3)
}
