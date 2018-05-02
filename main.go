package main

import (
	"bufio"
	"os"
	"os/exec"
)

func main() {
	// disable input buffering
	// exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	cmd := exec.Command("stty", "cbreak", "min", "1")
	cmd.Stdin = os.Stdin
	cmd.Run()
	// do not display entered characters on the screen
	// exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
	cmd2 := exec.Command("stty", "-echo")
	cmd2.Stdin = os.Stdin
	cmd2.Run()

	// restore the echoing state when exiting
	defer func() {
		cmd := exec.Command("stty", "echo")
		cmd.Stdin = os.Stdin
		cmd.Run()
	}()

	input := make(chan byte)
	output := make(chan Board)

	go func(ch chan byte) {
		reader := bufio.NewReader(os.Stdin)
		for {
			// buf := make([]byte, 1)
			// _, err := os.Stdin.Read(buf)
			// if err != nil {
			// 	log.Fatalf("Cannot read stdin: %v", err)
			// }
			b, err := reader.ReadByte()
			if err != nil { // Maybe log non io.EOF errors, if you want
				close(ch)
				return
			}
			ch <- b
		}
	}(input)

	game := NewGame(input, output)

	board := NewBoard(10, 12)
	game.play(board)
}
