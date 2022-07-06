package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	m          [][]int
	rowsNumber int
	colsNumber int
}

func New(s string) (*Matrix, error) {
	rows := strings.Split(s, "\n")
	result := Matrix{
		rowsNumber: len(rows),
		m:          make([][]int, len(rows)),
	}

	for rowIndex, row := range rows {
		values := strings.Split(strings.TrimSpace(row), " ")
		if len(values) == 0 {
			return nil, errors.New("empty row")
		}
		if rowIndex == 0 {
			result.colsNumber = len(values)
		} else {
			if result.colsNumber != len(values) {
				return nil, errors.New("different column number")
			}
		}

		result.m[rowIndex] = make([]int, result.colsNumber)

		for colIndex, value := range values {
			intValue, err := strconv.Atoi(value)

			if err != nil {
				return nil, errors.New("no valid int value")
			}

			result.m[rowIndex][colIndex] = intValue
		}
	}

	return &result, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Cols() [][]int {
	cols := make([][]int, m.colsNumber)

	for col := 0; col < m.colsNumber; col++ {
		cols[col] = make([]int, m.rowsNumber)
		for row := 0; row < m.rowsNumber; row++ {
			cols[col][row] = m.m[row][col]
		}
	}

	return cols
}

func (m *Matrix) Rows() [][]int {
	rows := make([][]int, m.rowsNumber)

	for row := 0; row < m.rowsNumber; row++ {
		rows[row] = make([]int, m.colsNumber)
		for col := 0; col < m.colsNumber; col++ {
			rows[row][col] = m.m[row][col]
		}
	}

	return rows
}

func (m *Matrix) Set(row, col, val int) bool {
	if row < 0 || col < 0 {
		return false
	}

	if row >= m.rowsNumber || col >= m.colsNumber {
		return false
	}

	m.m[row][col] = val

	return true
}
