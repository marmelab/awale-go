package render

import (
	"strings"
	"testing"
)

func TestCurrentPlayerIndicatorReturnArrow(t *testing.T) {
	indicator := CurrentPlayerIndicator(1, 1)

	if !strings.Contains(indicator, "=>") {
		t.Error("Indicator doesn't return expected Indicator current player")
	}
}

func TestCurrentPlayerIndicatorReturnEmptySelector(t *testing.T) {
	indicator := CurrentPlayerIndicator(1, 0)

	if strings.TrimSpace(indicator) != "" {
		t.Error("Indicator doesn't return arrow")
	}
}

func TestRenderScoreShouldReturnScoreFormat(t *testing.T) {
	score := [2]int{0, 5}
	expectedRender := "Score:\tPlayer (1): 0\tPlayer (2): 5\n"

	render := RenderScore(score)
	if render != expectedRender {
		t.Error("Render score doesn't return expected render")
	}
}
