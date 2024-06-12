package main

import "fmt"

type Queen struct {
	color color
}

var _ Piece = Queen{}

func (q Queen) String() string {
	if q.color == white {
		return "♕"
	}

	return "♛"
}

func (q Queen) Value() int {
	return 9 * int(q.color)
}

func (q Queen) GetMoves(coord Coord, board [8][8]Piece) []Move {
	moves := []Move{}

	newCoordX := coord.x - 1
	for newCoordX >= 0 {
		if board[coord.y][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: coord.y},
			Piece: q,
		})

		newCoordX--
	}

	newCoordY := coord.y - 1
	for newCoordY >= 0 {
		if board[newCoordY][coord.x] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: coord.x, y: newCoordY},
			Piece: q,
		})

		newCoordY--
	}

	newCoordX = coord.x + 1
	for newCoordX < 8 {
		if board[coord.y][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: coord.y},
			Piece: q,
		})

		newCoordX++
	}

	newCoordY = coord.y + 1
	for newCoordY < 8 {
		if board[newCoordY][coord.x] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: coord.x, y: newCoordY},
			Piece: q,
		})

		newCoordY++
	}

	newCoordX = coord.x - 1
	newCoordY = coord.y - 1
	for newCoordX >= 0 && newCoordY >= 0 {
		fmt.Println(newCoordX, newCoordY)
		if board[newCoordY][newCoordX] != nil {
			break
		}

		moves = append(moves, Move{
			From:  coord,
			To:    Coord{x: newCoordX, y: newCoordY},
			Piece: q,
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
			Piece: q,
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
			Piece: q,
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
			Piece: q,
		})

		newCoordX++
		newCoordY++
	}

	return moves
}

func (b Queen) Color() color {
	return b.color
}
