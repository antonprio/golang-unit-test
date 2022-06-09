package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "Anton",
			request: "Anton",
		},
		{
			name:    "Prio",
			request: "Prio",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("Anton", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Anton")
		}
	})

	b.Run("Prio", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Prio")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Anton")
	}
}

func TestMain(m *testing.M) {
	// do something before running all unit test
	fmt.Println("RUNNING PHASE BEFORE UNIT TEST...")

	m.Run()

	// do something after running all unit test
	fmt.Println("RUNNING PHASE AFTER UNIT TEST...")
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Anton",
			request:  "Anton",
			expected: "Hello Anton",
		},
		{
			name:     "Prio",
			request:  "Prio",
			expected: "Hello Prio",
		},
		{
			name:     "Hutomo",
			request:  "Hutomo",
			expected: "Hello Hutomo",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Anton", func(t *testing.T) {
		result := HelloWorld("Anton")
		assert.Equal(t, "Hello Anton", result, "Result must be 'Hello Anton'")
	})

	t.Run("Prio", func(t *testing.T) {
		result := HelloWorld("Prio")
		assert.Equal(t, "Hello Prio", result, "Result must be 'Hello Prio '")
	})
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test skipped on Mac OS")
	}

	result := HelloWorld("Anton")
	assert.Equal(t, "Hello Anton", result, "Result must be 'Hello Anton'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Anton")
	assert.Equal(t, "Hello Anton", result, "Result must be 'Hello Anton'")

	fmt.Println("TestHelloWorld with assert done ...")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Anton")
	require.Equal(t, "Hello Anton", result, "Result must be 'Hello Anton'")

	fmt.Println("TestHelloWorld with require done ...")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Anton")

	if result != "Hello Anton" {
		// error
		t.Error("Result must be Hello Anton")
	}
}

func TestHelloWorldNew(t *testing.T) {
	result := HelloWorld("Anton new")

	if result != "Hello Anton new" {
		t.Fatal("Result must be Hello Anton new")
	}
}
