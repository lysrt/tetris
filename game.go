package main

import (
	"fmt"
	"log"
	"time"
)

type Game struct {
	score int

	piece          Piece
	rotation       int
	pieceX, pieceY int

	input  chan byte
	output chan Board
}

func NewGame(input chan byte, output chan Board) *Game {
	g := &Game{
		score:    0,
		piece:    nil,
		rotation: 0,
		input:    input,
		output:   output,
	}

	return g
}

func (g *Game) play(board Board) {
	for !g.lost(board) {

		if g.piece == nil {
			g.piece = getRandomPiece()
			g.pieceX = 4
			g.pieceY = 0
		}

		// What is the last key pressed?
		select {
		case stdin, ok := <-g.input:
			if !ok {
				log.Fatal("Cannot read from channel ch")
			} else {
				// fmt.Println("Read input from stdin:", stdin)
				switch stdin {
				case 'a':
					if board.PieceAllowed(g.piece, g.rotation, g.pieceX-1, g.pieceY) {
						g.pieceX--
					}
				case 'd':
					if board.PieceAllowed(g.piece, g.rotation, g.pieceX+1, g.pieceY) {
						g.pieceX++
					}
				case 'w':
					if board.PieceAllowed(g.piece, (g.rotation+1)%4, g.pieceX, g.pieceY) {
						g.rotation++
						g.rotation %= 4
					}
				case 's':
					if board.PieceAllowed(g.piece, g.rotation, g.pieceX, g.pieceY+1) {
						g.pieceY++
					} else {
						board.PutPiece(g.piece, g.rotation, g.pieceX, g.pieceY)
						g.piece = nil
						if board.HasFullLine() {
							removedLines := board.RemoveFullLines()
							g.score += removedLines
							time.Sleep(250 * time.Millisecond)
						}
					}
				}
			}
		case <-time.After(250 * time.Millisecond):
			// Do something when there is nothing to read from stdin
			if board.PieceAllowed(g.piece, g.rotation, g.pieceX, g.pieceY+1) {
				g.pieceY++
			} else {
				board.PutPiece(g.piece, g.rotation, g.pieceX, g.pieceY)
				g.piece = nil
				if board.HasFullLine() {
					removedLines := board.RemoveFullLines()
					g.score += removedLines
					time.Sleep(250 * time.Millisecond)
				}
			}
		}

		screen := board.Copy()

		// Apply the moving piece on the screen
		if g.piece != nil {
			screen.PutPiece(g.piece, g.rotation, g.pieceX, g.pieceY)
		}

		// Refresh screen
		clearScreen()
		printBoard(screen)
		g.printScore()
	}
	close(g.input)
	close(g.output)
}

func (g *Game) PieceCanGoDown(board Board) bool {
	if g.pieceY > 5 {
		return false
	}
	return true
}

func (g *Game) printScore() {
	fmt.Printf("Score: %d\n", g.score)
}

func (g *Game) lost(board Board) bool {
	for _, block := range board.lines[0] {
		if block == PieceBlock {
			return true
		}
	}
	return false
}
