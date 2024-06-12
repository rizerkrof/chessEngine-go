package main

import "fmt"

type color int

const black color = -1
const white color = 1

type Piece interface {
	Value() int
	String() string
	GetMoves(Coord, [8][8]Piece) []Move
	Color() color
}

type Coord struct {
	x, y int
}

type Move struct {
	From, To Coord
	Piece    Piece
}

type Engine struct {
	board [8][8]Piece
	turn  color
}

func (e Engine) eval() float64 {
	eval := 0.0

	for _, row := range e.board {
		for _, piece := range row {
			if piece != nil {
				eval += float64(piece.Value())
			}
		}
	}

	return eval
}

func (e Engine) String() string {
	var s string

	for i := len(e.board) - 1; i >= 0; i-- {
		row := e.board[i]
		s += fmt.Sprintf("--------------------------------\n")
		for _, piece := range row {
			if piece != nil {
				s += " " + piece.String() + " |"
			} else {
				s += "   |"
			}
		}
		s += "\n"
	}

	return s
}

func (e Engine) GetMoves() []Move {
	moves := []Move{}

	for y, row := range e.board {
		for x, piece := range row {
			if piece != nil {
				moves = append(moves, piece.GetMoves(Coord{x, y}, e.board)...)
			}
		}
	}

	return moves
}

func (e Engine) MakeMove(m Move) Engine {
	board := e.board
	board[m.To.y][m.To.x] = m.Piece
	board[m.From.y][m.From.x] = nil

	return Engine{
		board: board,
		turn:  e.turn * -1,
	}
}

func new() Engine {
	/*return Engine{
		board: [8][8]Piece{
			{Rook{white}, Knight{white}, Bishop{white}, Queen{white}, King{white}, Bishop{white}, Knight{white}, Rook{white}},
			{Pawn{white}, Pawn{white}, Pawn{white}, Pawn{white}, Pawn{white}, Pawn{white}, Pawn{white}, Pawn{white}},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, King{white}, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, Pawn{black}, Pawn{black}, Pawn{black}, Pawn{black}, Pawn{black}, Pawn{black}, nil},
			{Rook{black}, nil, nil, Queen{black}, King{black}, nil, nil, Rook{black}},
		},
		turn: white,
	} */
	return Engine{
		board: [8][8]Piece{
			{Rook{white}, Knight{white}, Bishop{white}, Queen{white}, King{white}, Bishop{white}, Knight{white}, Rook{white}},
			{Pawn{white}, nil, nil, nil, nil, nil, nil, Pawn{white}},
			{nil, nil, nil, nil, nil, nil, nil, Pawn{white}},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil, nil, nil, nil},
			{Rook{black}, nil, nil, Queen{black}, King{black}, nil, nil, Rook{black}},
		},
		turn: white,
	}
}

func main() {
	e := new()
	// fmt.Println(e)
	// fmt.Println(e.GetMoves())
	// for _, move := range e.GetMoves() {
	// 	fmt.Println(e.MakeMove(move).eval())
	// 	fmt.Println(e.MakeMove(move).String())
	// }

}

type moveEval struct {
	move Move
	eval float64
}

func minimax(e Engine, color color) Move {
	myMoves := e.GetMoves()
	for _, myMove := range myMoves {
		engineAfterMove := e.MakeMove(myMove)
		opponentMoves := engineAfterMove.GetMoves()
		opponentMovesEval := []moveEval{}
		for _, opponentMove := range opponentMoves {
			engineAfterOpponentMove := engineAfterMove.MakeMove(opponentMove)
			eval := engineAfterOpponentMove.eval()
			opponentMovesEval = append(opponentMovesEval, moveEval{opponentMove, eval})
		}
		minEval := 100000.0
		for _, moveEval := range opponentMovesEval {
			minEval = min(minEval, moveEval.eval)
		}
	}
}
