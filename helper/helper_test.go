package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// sebelum
	fmt.Println("before unit test")

	m.Run()

	// sesudah
	fmt.Println("after unit test")
}

// untuk mengecek os
func TestSkip(t *testing.T) {
	// jika runtime berilai darwin akan skip
	if runtime.GOOS == "darwin" {
		t.Skip("Skipping test on Mac OS X")
	}
	result := HelloWorld("wahyu")
	t.Error("result must be hello wahyu", result)
}

func TestHelloWorldFadly(t *testing.T) {
	result := HelloWorld("fadly")
	assert.Equal(t, "hello fadly", result)
}
