package gadget

import "fmt"

// TapePlayer struct
type TapePlayer struct {
	Batteries string
}

// Play :TapePlayer.Play(string)
func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

// Stop :TapePlayer.Stop()
func (t TapePlayer) Stop() {
	fmt.Println("Stop")
}

// TapeRecorder struct
type TapeRecorder struct {
	Microphones int
}

// Play :TapeRecorder.Play()
func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

// Record : TapeRecorder.Record()
func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

// Stop : TapeRecorder.Stop()
func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}
