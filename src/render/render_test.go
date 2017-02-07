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
