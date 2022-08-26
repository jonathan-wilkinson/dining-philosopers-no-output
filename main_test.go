package main

import (
	"testing"
)

func Test_main(t *testing.T) {

	for i := 0; i < 1_000_000; i++ {
		main()
		if len(orderFinished) != 5 {
			t.Error("incorrect length of slice orderFinished")
		}
		orderFinished = []string{}
	}
}
