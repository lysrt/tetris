package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Shape [][][]int

type Piece struct {
	Shape
	Color Block
}

var OShape = Shape{
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
	},
}

var IShape = Shape{
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 1, 0},
		[]int{0, 0, 1, 0},
		[]int{0, 0, 1, 0},
		[]int{0, 0, 1, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{1, 1, 1, 1},
		[]int{0, 0, 0, 0},
	},
}

var TShape = Shape{
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
}

var LShape = Shape{
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{1, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 1, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
}

var JShape = Shape{
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{1, 1, 1, 0},
		[]int{0, 0, 1, 0},
		[]int{0, 0, 0, 0},
	},
}

var SShape = Shape{
	[][]int{
		[]int{0, 1, 1, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 1, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{1, 0, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
}

var ZShape = Shape{
	[][]int{
		[]int{1, 1, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 1, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 1, 0, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 0, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{0, 1, 1, 0},
		[]int{0, 0, 0, 0},
	},
	[][]int{
		[]int{0, 1, 0, 0},
		[]int{1, 1, 0, 0},
		[]int{1, 0, 0, 0},
		[]int{0, 0, 0, 0},
	},
}

var Pieces = []*Piece{
	&Piece{Shape: OShape, Color: ColorOPiece},
	&Piece{Shape: IShape, Color: ColorIPiece},
	&Piece{Shape: TShape, Color: ColorTPiece},
	&Piece{Shape: LShape, Color: ColorLPiece},
	&Piece{Shape: JShape, Color: ColorJPiece},
	&Piece{Shape: SShape, Color: ColorSPiece},
	&Piece{Shape: ZShape, Color: ColorZPiece},
}

func getRandomPiece() *Piece {
	return Pieces[rand.Intn(len(Pieces))]
}
