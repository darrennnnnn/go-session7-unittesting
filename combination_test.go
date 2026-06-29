package main

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		name     string
		n, k     int
		expected [][]int
	}{
		{
			name:     "n = 4, k = 2",
			n:        4,
			k:        2,
			expected: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name:     "k = 1",
			n:        3,
			k:        1,
			expected: [][]int{{1}, {2}, {3}},
		},
		{
			name:     "n = k",
			n:        2,
			k:        2,
			expected: [][]int{{1, 2}},
		},
		{
			name:     "k = 0",
			n:        5,
			k:        0,
			expected: [][]int{},
		},
		{
			name: "n = 5,k = 3",
			n:    5,
			k:    3,
			expected: [][]int{
				{1, 2, 3}, {1, 2, 4}, {1, 2, 5},
				{1, 3, 4}, {1, 3, 5}, {1, 4, 5},
				{2, 3, 4}, {2, 3, 5}, {2, 4, 5},
				{3, 4, 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combination(tt.n, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("combine(%d, %d) = %v, want %v",
					tt.n, tt.k, got, tt.expected)
			}
		})
	}
}

func TestCombineEdgeCases(t *testing.T) {
	t.Run("n = 0 should return empty", func(t *testing.T) {
		result := combination(0, 1)
		if len(result) != 0 {
			t.Errorf("Expected empty result for n = 0, got %v", result)
		}
	})

	t.Run("k > n should return empty", func(t *testing.T) {
		result := combination(3, 5)
		if len(result) != 0 {
			t.Errorf("Expected empty result when k > n, got %v", result)
		}
	})
}

func BenchmarkCombine(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
		k    int
	}{
		{"n = 5,k = 2", 5, 2},
		{"n = 10,k = 3", 10, 3},
		{"n = 15,k = 4", 15, 4},
	}

	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				combination(bm.n, bm.k)
			}
		})
	}
}
