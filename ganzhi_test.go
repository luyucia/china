package china

import (
	"testing"
	"time"
)

func TestGetGanzhi(t *testing.T) {

	t.Log(GetGanzhi("1989"))

	t.Log(GetNowShichen())

	t.Log(GetNowLiuzhu())

	tim, _ := time.Parse("2006-01-02", "2020-01-01")
	t.Log(GetYuandanGanzhi(tim))
}
