package main

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Piece [][][]int

var OPiece = Piece{
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

var IPiece = Piece{
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

var TPiece = Piece{
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

var LPiece = Piece{
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

var JPiece = Piece{
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

var SPiece = Piece{
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

var ZPiece = Piece{
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

var Pieces = []Piece{
	OPiece,
	IPiece,
	TPiece,
	LPiece,
	JPiece,
	SPiece,
	ZPiece,
}

func getRandomPiece() Piece {
	return Pieces[rand.Intn(len(Pieces))]
}
