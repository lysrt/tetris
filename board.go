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
	ColorOPiece
	ColorIPiece
	ColorTPiece
	ColorLPiece
	ColorJPiece
	ColorSPiece
	ColorZPiece
	Ghost
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

func (b *Board) PutPiece(p *Piece, rot, x, y int) {
	for j := y; j < y+4; j++ {
		for i := x; i < x+4; i++ {
			block := p.Shape[rot][j-y][i-x]
			switch block {
			case 0:
			case 1:
				b.lines[j][i] = p.Color
			}
		}
	}
}

func (b *Board) PieceAllowed(p *Piece, rot, x, y int) bool {
	for j := y; j < y+4; j++ {
		for i := x; i < x+4; i++ {
			pieceBlock := p.Shape[rot][j-y][i-x]
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
			if belowBlock != EmptyBlock {
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
			j++ // Clean this line again as blocks go down
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
		if block == EmptyBlock {
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
			case ColorOPiece:
				fmt.Print(string("\u001B[43m  \u001B[0m")) // Yellow
			case ColorIPiece:
				fmt.Print(string("\u001B[46m  \u001B[0m")) // Cyan
			case ColorTPiece:
				fmt.Print(string("\u001B[45m  \u001B[0m")) // Purple
			case ColorLPiece:
				fmt.Print(string("\u001B[47m  \u001B[0m")) // White
			case ColorJPiece:
				fmt.Print(string("\u001B[44m  \u001B[0m")) // Blue
			case ColorSPiece:
				fmt.Print(string("\u001B[41m  \u001B[0m")) // Red
			case ColorZPiece:
				fmt.Print(string("\u001B[42m  \u001B[0m")) // Green
			case Ghost:
				fmt.Print(string("\u001B[100m  \u001B[0m")) // Black
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
