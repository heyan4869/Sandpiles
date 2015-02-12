package main

import (
	"fmt"
	"os"
	"strconv"
)

// Build a new struct Board
type Board struct {
	row int
	column int
	value [][]int
}

// Create a new board
func CreateBoard(size, pile int) Board {
	var rownum int = size
	var columnnum int = size
	var newvalue [][]int = make([][]int, size)
	for row := range newvalue {
		newvalue[row] = make([]int, size)
	}
	newvalue[size/2][size/2] = pile
	return Board{rownum, columnnum, newvalue}

}

// Topples squares until the board converged to a stable configuration
func (b *Board) ComputeSteadyState() {
	for i := 0; i < b.row; i++ {
		for j := 0; j < b.column; j++ {
			b.Topple(i, j)
		}
	}
}

// Draw the board
func (b *Board) DrawBoard() {
	c := CreateNewCanvas(b.row, b.column)
	c.SetLineWidth(1)
	for i := 0; i < b.NumRows(); i++ {
		for j := 0; j < b.NumCols(); j++ {
			if b.Cell(i, j) == 0 {
				c.SetFillColor(MakeColor(0, 0, 0))
			} else if b.Cell(i, j) == 1 {
				c.SetFillColor(MakeColor(85, 85, 85))
			} else if b.Cell(i, j) == 2 {
				c.SetFillColor(MakeColor(170, 170, 170))
			} else if b.Cell(i, j) == 3 {
				c.SetFillColor(MakeColor(255, 255, 255))
			}
			c.MoveTo(float64(i), float64(j))
			c.LineTo(float64(i+1), float64(j))
			c.LineTo(float64(i+1), float64(j+1))
			c.LineTo(float64(i), float64(j+1))
			c.LineTo(float64(i), float64(j))
			c.Fill()
		}
	}
	c.SaveToPNG("board.png")
}

// Topple (r, c) until it can鈥檛 be toppled any more
func (b *Board) Topple(r, c int) {
	if b.value[r][c] >= 4 {
		var count int
		if b.Contains(r-1, c) == true {
			var currvalue int = b.Cell(r-1, c)
			b.Set(r-1, c, currvalue+1)
			count++
		}
		if b.Contains(r, c-1) == true {
			var currvalue int = b.Cell(r, c-1)
			b.Set(r, c-1, currvalue+1)
			count++
		}
		if b.Contains(r, c+1) == true {
			var currvalue int = b.Cell(r, c+1)
			b.Set(r, c+1, currvalue+1)
			count++
		}
		if b.Contains(r+1, c) == true {
			var currvalue int = b.Cell(r+1, c)
			b.Set(r+1, c, currvalue+1)
			count++
		}
		var selfvalue = b.Cell(r, c)
		b.Set(r, c, selfvalue-count)
	}
}

// Return true if (r, c) is within the field
func (b *Board) Contains(r, c int) bool {
	if r >=0 && r < b.NumRows() && c >=0 && c < b.NumCols() {
		return true
	} else {
		return false
	}

}

// Set the value of cell (r, c)
func (b *Board) Set(r,c, value int) {
	b.value[r][c] = value
}

// Return the value of the cell (r, c)
func (b *Board) Cell(r, c int) int {
	return b.value[r][c]
}

// Return true if there are no cells with 鈮� 4 coins on them
func (b *Board) IsConverged() bool {
	for i := 0; i < b.NumRows(); i++ {
		for j := 0; j <b.NumCols(); j++ {
			if b.Cell(i, j) >= 4 {
				return false
			}
		}
	}
	return true
}

// Return the number of rows on the board
func (b *Board) NumRows() int {
	return b.row
}

// Return the number of columns on the board
func (b *Board) NumCols() int {
	return b.column
}

// Main function of this sandpile program
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error: Should input the value of size and pile")
		return
	}

	size, err := strconv.Atoi(os.Args[1])
	if err != nil || size <= 0 {
		fmt.Println("Error: Input size should be a positive number")
		return
	}

	pile, err := strconv.Atoi(os.Args[2])
	if err != nil || pile <= 0 {
		fmt.Println("Error: Input pile should be a positive number")
		return
	}

	// Create a new Board
	b := CreateBoard(size, pile)
	// Topple the cell until the board converged to a stable configuration
	for b.IsConverged() == false {
		b.ComputeSteadyState()
	}
	// Draw the stable configuration of the board
	b.DrawBoard()
}