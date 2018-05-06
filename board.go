package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var clearFunc map[string]func() error

func init() {
	clearFunc = map[string]func() error{
		"linux": func() error {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			return cmd.Run()
		},
		"darwin": func() error {
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			return cmd.Run()
		},
		"windows": func() error {
			cmd := exec.Command("cmd", "/c", "cls")
			cmd.Stdout = os.Stdout
			return cmd.Run()
		},
	}
}

type Block int

const (
	EmptyBlock Block = iota
	PieceBlock
)

type Line []Block
type Board struct {
	width, height int
	lines         []Line
}

func NewBoard(width, height int) Board {
	board := Board{
		width:  width,
		height: height,
	}

	for j := 0; j < height; j++ {
		line := Line{}
		for i := 0; i < width; i++ {
			line = append(line, EmptyBlock)
		}
		board.lines = append(board.lines, line)
	}
	return board
}

func (b *Board) PutPiece(p Piece, rot, x, y int) {
	for j := y; j < y+4; j++ {
		for i := x; i < x+4; i++ {
			block := p[rot][j-y][i-x]
			switch block {
			case 0:
			case 1:
				b.lines[j][i] = PieceBlock
			}
		}
	}
}

func (b *Board) PieceAllowed(p Piece, rot, x, y int) bool {
	for j := y; j < y+4; j++ {
		for i := x; i < x+4; i++ {
			pieceBlock := p[rot][j-y][i-x]
			if pieceBlock == 0 {
				continue
			}

			if i < 0 || i > b.width-1 {
				return false
			}
			if j > b.height-1 {
				return false
			}

			belowBlock := b.lines[j][i]
			if belowBlock == PieceBlock {
				return false
			}
		}
	}
	return true
}

func (b *Board) RemoveFullLines() int {
	linesRemoved := 0
	for j := b.height - 1; j > 0; j-- {
		if isFull(b.lines[j]) {
			b.moveLinesDownOn(j)
			linesRemoved++
			j++ // Clean this line again
		}
	}
	return linesRemoved
}

func (b *Board) moveLinesDownOn(j int) {
	for j > 0 {
		copy(b.lines[j], b.lines[j-1])
		j--
	}
	b.lines[0] = make([]Block, b.width)
}

func (b *Board) HasFullLine() bool {
	for _, line := range b.lines {
		if isFull(line) {
			return true
		}
	}
	return false
}

func isFull(line Line) bool {
	for _, block := range line {
		if block != PieceBlock {
			return false
		}
	}
	return true
}

func (b *Board) Copy() Board {
	new := Board{
		width:  b.width,
		height: b.height,
	}
	for j := 0; j < b.height; j++ {
		line := Line{}
		for i := 0; i < b.width; i++ {
			line = append(line, b.lines[j][i])
		}
		new.lines = append(new.lines, line)
	}
	return new
}

func printBoard(board Board) {
	fmt.Println("|--------------------|")
	for _, line := range board.lines {
		fmt.Print("|")
		for _, block := range line {
			switch block {
			case EmptyBlock:
				fmt.Print("  ")
			case PieceBlock:
				// fmt.Print(string("\u001B[41m\u25a2\u25a2\u001B[0m")) // u25a0, 1, 2, 3, 6, (7), 8 are good blocks characters
				fmt.Print(string("\u001B[41m  \u001B[0m")) // u25a0, 1, 2, 3, 6, (7), 8 are good blocks characters
			default:
				fmt.Print("??")
			}
		}
		fmt.Println("|")
	}
	fmt.Println("|--------------------|")
}

func clearScreen() {
	value, ok := clearFunc[runtime.GOOS]
	if ok {
		if err := value(); err != nil {
			log.Fatalf("cannot run clear command: %v", err)
		}
	} else {
		log.Fatalf("cannot run clear command, unsupported OS")
	}
}
