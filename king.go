package main

type King struct {
	color color
}

var _ Piece = King{}

func (k King) String() string {
	if k.color == white {
		return "♔"
	}

	return "♚"
}

func (k King) Value() int {
	return 1000 * int(k.color)
}

func (q King) GetMoves(coord Coord, board [8][8]Piece) []Move {
	moves := []Move{}

	for y := -1; y <= 1; y++ {
		for x := -1; x <= 1; x++ {
			if x == 0 && y == 0 {
				continue
			}

			if coord.y+y < 0 || coord.y+y >= 8 || coord.x+x < 0 || coord.x+x >= 8 {
				continue
			}

			if board[coord.y+y][coord.x+x] != nil {
				continue
			}

			moves = append(moves, Move{
				From:  coord,
				To:    Coord{x: coord.x + x, y: coord.y + y},
				Piece: q,
			})
		}
	}

	return moves
}

func (b King) Color() color {
	return b.color
}
