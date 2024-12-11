package assignment1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Reactangle(t *testing.T) {
	got := Reactangle(2, 3)
	assert.Equal(t, 10, got)
}

func Test_Squareperimeter(t *testing.T) {
	got := Squareperimeter(2)
	assert.Equal(t, 8, got)
}
func Test_Cube(t *testing.T) {
	got := Cube(2, 1, 4)
	assert.Equal(t, 8, got)
}
func Test_Sphere(t *testing.T) {
	got := Sphere(2)
	assert.Equal(t, 25.12, got)
}

// Benchmarks
func BenchmarkReactangle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reactangle(2, 3)
	}
}

func BenchmarkSquareperimeter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Squareperimeter(2)
	}
}

func BenchmarkCube(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cube(2, 1, 4)
	}
}

func BenchmarkSphere(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sphere(2)
	}
}
