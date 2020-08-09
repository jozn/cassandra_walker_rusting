package ant

import (
	"fmt"
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	for i := 1; i < 100000; i++ {
		r := Hash("sdds" + strconv.Itoa(i))
		fmt.Println(r)
		if r < 0 {
			t.Error("r is negative")
		}
	}
}

func TestHashEqual1Miliion(t *testing.T) {
	N := 20000
	mp := make(map[int]bool, N)
	for i := 1; i < N; i++ {
		r := Hash("method" + strconv.Itoa(i))
		if b := mp[r]; b {
			t.Error("r is already exists: ", r)
		}
		mp[r] = true
	}
}
