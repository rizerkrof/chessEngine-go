package main

type Knight struct {
	color color
}

var _ Piece = Knight{}

func (k Knight) String() string {
	if k.color == white {
		return "♘"
	}

	return "♞"
}

func (k Knight) Value() int {
	return 3 * int(k.color)
}

func (q Knight) GetMoves(coord Coord, board [8][8]Piece) []Move {
	moves := []Move{}

	knightOffsets := []int{-2, -1, 1, 2}

	for _, y := range knightOffsets {
		for _, x := range knightOffsets {
			if x == y || x == -y {
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

func (b Knight) Color() color {
	return b.color
}
