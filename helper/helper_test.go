package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// membuat func BenchmarkHelloworld
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Println("Hello world")
	}
}

// membuat table test
func TestHelloWorldTable(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "HelloWorld(fadly)",
			request:  "fadly",
			expected: "Hello Fadly",
		},
	}
	// iterasi menggunakan subtest
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}

}

// membuat subtest
func TestSubTest(t *testing.T) {
	t.Run("fadly", func(t *testing.T) {
		result := HelloWorld("fadly")
		require.Equal(t, "fello fadly", result)
	})
	t.Run("pangestu", func(t *testing.T) {
		result := HelloWorld("pangestu")
		require.Equal(t, "fello pangestu", result)
	})
}

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
