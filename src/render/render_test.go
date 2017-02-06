package render

import (
	"strings"
	"testing"
)

func TestIndicatorCurrentPlayerReturnArrow(t *testing.T) {
	indicator := IndicatorCurrentPlayer(1, 1)

	if !strings.Contains(indicator, "=>") {
		t.Error("Indicator doesn't return expected Indicator current player")
	}
}

func TestIndicatorCurrentPlayerReturnEmptySelector(t *testing.T) {
	indicator := IndicatorCurrentPlayer(1, 0)

	if strings.TrimSpace(indicator) != "" {
		t.Error("Indicator doesn't return arrow")
	}
}
