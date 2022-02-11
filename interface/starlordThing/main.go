package main

import (
	"github.com/codecamp-prem/interface/starlordThing/gadget"
)

// Player : interface
type Player interface {
	Play(string)
	Stop()
}

// Tryout : When you have a value of a concrete type assigned to a variable with an interface type, a type assertion lets you get the concrete type back. It’s kind of like a type conversion. Its syntax even looks like a cross between a method call and a type conversion. After an interface value, you type a dot,followed by a pair of parentheses with the concrete type. */
// type assertion to get a value of a concrete type back,
func Tryout(player Player) {
	player.Play("Sexy Back")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	songsList := []string{"Jai Ho", "Rangeela marooo", "Ringadridham"}
	var player Player = gadget.TapePlayer{} // {} for struct defined type
	playList(player, songsList)
	player = gadget.TapeRecorder{}
	playList(player, songsList)

	Tryout(gadget.TapePlayer{})
	Tryout(gadget.TapeRecorder{})
}
