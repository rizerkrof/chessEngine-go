package main

type Pawn struct {
	color color
}

var _ Piece = Pawn{}

func (p Pawn) String() string {
	if p.color == white {
		return "♙"
	}

	return "♟"
}

func (p Pawn) Value() int {
	return 1 * int(p.color)
}

func (p Pawn) GetMoves(coord Coord, board [8][8]Piece) []Move {
	var moves []Move

	if coord.x < 0 || coord.x >= 8 || coord.y < 0 || coord.y >= 8 || coord.y == 0 || coord.y == 7 {
		return moves
	}

	color := int(p.color)

	if board[coord.y+color][coord.x] == nil {
		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: coord.x, y: coord.y + 1*color},
			Piece: p,
		})
	}

	pawnFirstRow := int(3.5 - float32(color)*2.5) // returns 6 for black and 1 for white
	if coord.y == pawnFirstRow && board[coord.y+color][coord.x] == nil && board[coord.y+2*color][coord.x] == nil {
		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: coord.x, y: coord.y + 2*color},
			Piece: p,
		})
	}

	return moves
}

func (p Pawn) Color() color {
	return p.color
}
