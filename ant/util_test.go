package ant

import (
	"strconv"
	"testing"
)

func TestHash(t *testing.T) {
	for i := 1; i < 100000; i++ {
		r := StrToInt32Hash("sdds" + strconv.Itoa(i))
		//fmt.Println(r)
		if r < 0 {
			t.Error("r is negative")
		}
	}
}

func TestHashEqual1Miliion(t *testing.T) {
	N := 20000
	mp := make(map[uint32]bool, N)
	for i := 1; i < N; i++ {
		r := StrToInt32Hash("method" + strconv.Itoa(i))
		if b := mp[r]; b {
			t.Error("r is already exists: ", r)
		}
		mp[r] = true
	}
}

func TestMethodUnique1Miliion(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	N := 1000000
	for i := 1; i < N; i++ {
		uniqueMethodHash("method" + strconv.Itoa(i))
	}
}

func TestMethodUniqueFew(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	hashMp = make(map[uint32]string)
	N := 1000
	for i := 1; i < N; i++ {
		uniqueMethodHash("method" + strconv.Itoa(i))
	}

	method := "AddGroup"
	uniqueMethodHash(method)
	uniqueMethodHash(method)

}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
