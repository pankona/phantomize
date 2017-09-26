// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene"
)

func main() {
	sim := simra.NewSimra()
	sim.Start(&scene.Title{})
}
