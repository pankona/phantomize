// +build darwin linux

package main

import (
	"github.com/pankona/gomo-simra/simra"
	"github.com/pankona/phantomize/scene"
)

func onStart() {
	simra.LogDebug("receive onStart")
	engine := simra.GetInstance()
	// TODO: this will be called on rotation.
	// to keep state on rotation, SetScene must not call
	// every onStart.
	engine.SetScene(&scene.Title{})

}

func onStop() {
	simra.LogDebug("receive onStop")
}

func main() {
	simra.LogDebug("[IN]")
	engine := simra.GetInstance()
	engine.Start(onStart, onStop)
	simra.LogDebug("[OUT]")
}
