package main

import "fmt"

type Bishop struct {
	color color
}

var _ Piece = Bishop{}

func (b Bishop) String() string {
	if b.color == white {
		return "♗"
	}

	return "♝"
}

func (b Bishop) Value() int {
	return 3 * int(b.color)
}

func (b Bishop) GetMoves(coord Coord, board [8][8]Piece) []Move {
	moves := []Move{}
	newCoordX := coord.x - 1
	newCoordY := coord.y - 1
	for newCoordX >= 0 && newCoordY >= 0 {
		fmt.Println(newCoordX, newCoordY)
		if board[newCoordY][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: newCoordY},
			Piece: b,
		})

		newCoordX--
		newCoordY--
	}

	newCoordX = coord.x + 1
	newCoordY = coord.y - 1
	for newCoordX < 8 && newCoordY >= 0 {
		if board[newCoordY][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: newCoordY},
			Piece: b,
		})

		newCoordX++
		newCoordY--
	}

	newCoordX = coord.x - 1
	newCoordY = coord.y + 1
	for newCoordX >= 0 && newCoordY < 8 {
		if board[newCoordY][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: newCoordY},
			Piece: b,
		})

		newCoordX--
		newCoordY++
	}

	newCoordX = coord.x + 1
	newCoordY = coord.y + 1
	for newCoordX < 8 && newCoordY < 8 {
		if board[newCoordY][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: newCoordY},
			Piece: b,
		})

		newCoordX++
		newCoordY++
	}

	return moves
}

func (b Bishop) Color() color {
	return b.color
}
