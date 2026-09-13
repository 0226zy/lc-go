package findmedianfromdatastream

import (
	"math"
	"testing"
)

const eps = 1e-5

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < eps
}

func TestMedianFinderExample(t *testing.T) {
	// LeetCode 官方示例
	mf := Constructor()
	mf.AddNum(1)
	mf.AddNum(2)
	if got := mf.FindMedian(); !almostEqual(got, 1.5) {
		t.Errorf("FindMedian() = %v, want 1.5", got)
	}
	mf.AddNum(3)
	if got := mf.FindMedian(); !almostEqual(got, 2.0) {
		t.Errorf("FindMedian() = %v, want 2.0", got)
	}
}

func TestMedianFinderCases(t *testing.T) {
	tests := []struct {
		name   string
		ops    []string
		args   []int
		want   []float64 // 仅对 FindMedian 有期望值，其它位置填 NaN 占位
		hasMed []bool
	}{
		{
			name:   "单元素",
			ops:    []string{"AddNum", "FindMedian"},
			args:   []int{5, 0},
			want:   []float64{0, 5.0},
			hasMed: []bool{false, true},
		},
		{
			name:   "偶数个元素",
			ops:    []string{"AddNum", "AddNum", "FindMedian"},
			args:   []int{1, 2, 0},
			want:   []float64{0, 0, 1.5},
			hasMed: []bool{false, false, true},
		},
		{
			name:   "递减插入",
			ops:    []string{"AddNum", "AddNum", "AddNum", "FindMedian"},
			args:   []int{3, 2, 1, 0},
			want:   []float64{0, 0, 0, 2.0},
			hasMed: []bool{false, false, false, true},
		},
		{
			name:   "含负数",
			ops:    []string{"AddNum", "AddNum", "AddNum", "FindMedian"},
			args:   []int{-1, -2, -3, 0},
			want:   []float64{0, 0, 0, -2.0},
			hasMed: []bool{false, false, false, true},
		},
		{
			name:   "重复元素",
			ops:    []string{"AddNum", "AddNum", "AddNum", "AddNum", "FindMedian"},
			args:   []int{2, 2, 2, 2, 0},
			want:   []float64{0, 0, 0, 0, 2.0},
			hasMed: []bool{false, false, false, false, true},
		},
		{
			name:   "交替查询",
			ops:    []string{"AddNum", "FindMedian", "AddNum", "FindMedian", "AddNum", "FindMedian", "AddNum", "FindMedian"},
			args:   []int{6, 0, 10, 0, 2, 0, 6, 0},
			want:   []float64{0, 6.0, 0, 8.0, 0, 6.0, 0, 6.0},
			hasMed: []bool{false, true, false, true, false, true, false, true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mf := Constructor()
			for i, op := range tt.ops {
				switch op {
				case "AddNum":
					mf.AddNum(tt.args[i])
				case "FindMedian":
					got := mf.FindMedian()
					if tt.hasMed[i] && !almostEqual(got, tt.want[i]) {
						t.Errorf("step %d FindMedian() = %v, want %v", i, got, tt.want[i])
					}
				}
			}
		})
	}
}

func TestMedianFinderLarge(t *testing.T) {
	mf := Constructor()
	n := 1000
	for i := 1; i <= n; i++ {
		mf.AddNum(i)
		if i%2 == 1 {
			want := float64((i + 1) / 2)
			if got := mf.FindMedian(); !almostEqual(got, want) {
				t.Fatalf("i=%d FindMedian() = %v, want %v", i, got, want)
			}
		} else {
			want := float64(i/2+i/2+1) / 2.0
			if got := mf.FindMedian(); !almostEqual(got, want) {
				t.Fatalf("i=%d FindMedian() = %v, want %v", i, got, want)
			}
		}
	}
}

func BenchmarkMedianFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := Constructor()
		for x := 0; x < 10000; x++ {
			mf.AddNum(x)
			if x%100 == 0 {
				_ = mf.FindMedian()
			}
		}
	}
}
