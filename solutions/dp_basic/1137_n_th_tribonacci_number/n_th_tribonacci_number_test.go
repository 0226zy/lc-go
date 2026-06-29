package nthtribonumber

import "testing"

func TestTribonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"T(0)=0", 0, 0},
		{"T(1)=1", 1, 1},
		{"T(2)=1", 2, 1},
		{"T(3)=2", 3, 2},
		{"T(4)=4", 4, 4},
		{"T(5)=7", 5, 7},
		{"T(10)=149", 10, 149},
		{"T(25)=1389537", 25, 1389537},
		{"T(37)=2082876103", 37, 2082876103},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Tribonacci(tt.n)
			if got != tt.want {
				t.Errorf("Tribonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func BenchmarkTribonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Tribonacci(37)
	}
}
