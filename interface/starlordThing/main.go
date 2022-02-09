package main

import (
	"github.com/codecamp-prem/interface/starlordThing/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	songsList := []string{"Jai Ho", "Rangeela marooo", "Ringadridham"}
	var player Player = gadget.TapePlayer{}
	playList(player, songsList)
	player = gadget.TapeRecorder{}
	playList(player, songsList)
}
