package main

import (
	"os"
	"strconv"
	"strings"
)

func Parse(filepath string) (*Puzzle, error) {
	data, err := read(filepath)

	if err != nil {
		return nil, err
	}

	if data != nil {
		puzzle, err := createPuzzle(data)
		if err != nil {
			return nil, err
		}
		return puzzle, err
	}

	return nil, nil
}

func read(filepath string) ([]byte, error) {
	var data []byte
	buf := make([]byte, 8)
	reader, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer reader.Close()

	for {
		length, err := reader.Read(buf)

		if err != nil {
			if err.Error() != "EOF" {
				return nil, err
			}
			break
		}

		data = append(data, buf[:length]...)
	}

	return data, nil
}

func clean(data []byte) ([][]int, []int, error) {
	dimension := make([]int, 2)
	board := make([][]int, 0)
	lines := strings.Split(string(data), "\n")
	for _, v := range lines {
		if v == "" {
			break
		}
		line := strings.TrimSpace(v)
		l := strings.Split(line, "\t")
		newLines := make([]int, 0)
		for _, m := range l {
			n, err := strconv.Atoi(m)
			if err != nil {
				return nil, nil, err
			}
			newLines = append(newLines, n)
		}
		if len(line) == 0 {
			continue
		}
		board = append(board, newLines)
	}

	dimension[0], dimension[1] = len(board), len(board[0])
	return board, dimension, nil
}

func createPuzzle(data []byte) (*Puzzle, error) {
	board, dimension, err := clean(data)

	if err != nil {
		return nil, err
	}

	return &Puzzle{board: board, dimension: dimension}, nil
}

