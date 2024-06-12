package main

import "fmt"

type Rook struct {
	color color
}

var _ Piece = Rook{}

func (r Rook) String() string {
	if r.color == white {
		return "♖"
	}

	return "♜"
}

func (r Rook) Value() int {
	return 5 * int(r.color)
}

func (p Rook) GetMoves(coord Coord, board [8][8]Piece) []Move {
	moves := []Move{}

	newCoordX := coord.x - 1
	for newCoordX >= 0 {
		if board[coord.y][newCoordX] == nil || board[coord.y][newCoordX].Color() != p.color {
			moves = append(moves, Move{
				From:  coord,
				To:    Coord{x: newCoordX, y: coord.y},
				Piece: p,
			})
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: coord.y},
			Piece: p,
		})

		newCoordX--
	}

	newCoordY := coord.y - 1
	for newCoordY >= 0 {
		if board[newCoordY][coord.x] == nil || board[newCoordY][coord.x].Color() != p.color {
			moves = append(moves, Move{
				From:  coord,
				To:    Coord{x: coord.x, y: newCoordY},
				Piece: p,
			})
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: coord.x, y: newCoordY},
			Piece: p,
		})

		newCoordY--
	}

	newCoordX = coord.x + 1
	for newCoordX < 8 {
		if board[coord.y][newCoordX] == nil || board[coord.y][newCoordX].Color() != p.color {
			moves = append(moves, Move{
				From:  coord,
				To:    Coord{x: newCoordX, y: coord.y},
				Piece: p,
			})
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: coord.y},
			Piece: p,
		})

		newCoordX++
	}

	newCoordY = coord.y + 1
	for newCoordY < 8 {
		if board[newCoordY][coord.x] == nil || board[newCoordY][coord.x].Color() != p.color {
			moves = append(moves, Move{
				From:  coord,
				To:    Coord{x: coord.x, y: newCoordY},
				Piece: p,
			})
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: coord.x, y: newCoordY},
			Piece: p,
		})

		newCoordY++
	}
	fmt.Println("okkk")

	return moves
}

func (b Rook) Color() color {
	return b.color
}
