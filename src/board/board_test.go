package board

import (
	"reflect"
	"testing"
)

func TestNewBoardShouldReturnNewBoard(t *testing.T) {

	board, _ := New(12, 4)
	expectedBoard := Board{4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}

	if !reflect.DeepEqual(board, expectedBoard) {
		t.Error("New Board doesn't return expected Board")
	}
}

func TestNewBoardShouldReturnSizeErrors(t *testing.T) {

	_, errors := New(13, 4)

	if errors == nil {
		t.Error("New Board should return an error with an invalid board size")
	}
}
