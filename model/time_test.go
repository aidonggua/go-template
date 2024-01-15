package model

import (
	"testing"
	"time"
)

func TestTime(t *testing.T) {
	uft := time.Now()
	t.Logf("uft: %v", uft)
	ft := Now()
	t.Logf("ft: %v", ft)
	ft.SetPattern(DateFormat)
	t.Logf("ft: %v", ft)
	ft.SetPattern(TimeFormat)
	t.Logf("ft: %v", ft)
}
