package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
)

func main() {
	disableInputBuffering()
	disableCharacterEcho()
	defer enableCharacterEcho()

	input := listenKeyboardInput()
	output := make(chan Board)

	game := NewGame(input, output)

	board := NewBoard(10, 20)
	game.play(board)
}

func disableInputBuffering() {
	// exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	cmd := exec.Command("stty", "cbreak", "min", "1")
	cmd.Stdin = os.Stdin
	cmd.Run()
}

func disableCharacterEcho() {
	// do not display entered characters on the screen
	// exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	cmd2 := exec.Command("stty", "-echo")
	cmd2.Stdin = os.Stdin
	cmd2.Run()
}

func enableCharacterEcho() {
	cmd := exec.Command("stty", "echo")
	cmd.Stdin = os.Stdin
	cmd.Run()
}

type Key int

const (
	Up Key = iota
	Down
	Left
	Right
	Space
)

func listenKeyboardInput() chan Key {
	input := make(chan byte)

	go func(ch chan byte) {
		reader := bufio.NewReader(os.Stdin)
		for {
			b, err := reader.ReadByte()
			if err != nil {
				close(ch)
				return
			}
			ch <- b
		}
	}(input)

	output := make(chan Key)
	go func(in chan byte, out chan Key) {
		for {
			select {
			case b, ok := <-in:
				if !ok {
					log.Fatal("Cannot read from channel ch")
				} else {
					switch b {
					case 'a':
						out <- Left
					case 'd':
						out <- Right
					case 'w':
						out <- Up
					case 's':
						out <- Down
					case ' ':
						out <- Space
					case 27:
						b2, ok := <-in
						if !ok {
							log.Fatal("Cannot read from channel ch")
						}
						if b2 != 91 {
							continue
						}

						if b3, ok := <-in; !ok {
							log.Fatal("Cannot read from channel ch")
						} else {
							switch b3 {
							case 65:
								out <- Up
							case 66:
								out <- Down
							case 67:
								out <- Right
							case 68:
								out <- Left
							}
						}
					}
				}
			}
		}
	}(input, output)

	return output
}
