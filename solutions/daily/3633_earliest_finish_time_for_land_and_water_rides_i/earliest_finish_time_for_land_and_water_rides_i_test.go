package earliestfinish

import "testing"

func TestEarliestFinishTime(t *testing.T) {
	tests := []struct {
		name          string
		landStartTime []int
		landDuration  []int
		waterStartTime []int
		waterDuration []int
		want          int
	}{
		{
			name:          "示例1",
			landStartTime: []int{2, 8},
			landDuration:  []int{4, 1},
			waterStartTime: []int{6},
			waterDuration: []int{3},
			want:          9,
		},
		{
			name:          "示例2",
			landStartTime: []int{5},
			landDuration:  []int{3},
			waterStartTime: []int{1},
			waterDuration: []int{10},
			want:          14,
		},
		{
			name:          "单项目且时间相同",
			landStartTime: []int{1},
			landDuration:  []int{1},
			waterStartTime: []int{1},
			waterDuration: []int{1},
			want:          3, // 1+1=2 结束第一个，2+1=3 结束第二个
		},
		{
			name:          "第一类需要等待第二类",
			landStartTime: []int{10},
			landDuration:  []int{1},
			waterStartTime: []int{1},
			waterDuration: []int{1},
			want:          11,
		},
		{
			name:          "多个项目选最优",
			landStartTime: []int{1, 5, 10},
			landDuration:  []int{10, 2, 1},
			waterStartTime: []int{3, 8},
			waterDuration: []int{5, 2},
			want:          10, // land->water: minEnd=7, water最优=8+2=10; water->land: minEnd=8, land最优=8+2=10
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EarliestFinishTime(tt.landStartTime, tt.landDuration, tt.waterStartTime, tt.waterDuration)
			if got != tt.want {
				t.Errorf("EarliestFinishTime() = %d, want %d", got, tt.want)
			}
		})
	}
}

func BenchmarkEarliestFinishTime(b *testing.B) {
	landStartTime := []int{2, 8, 15, 20, 25}
	landDuration := []int{4, 1, 3, 2, 5}
	waterStartTime := []int{6, 10, 18}
	waterDuration := []int{3, 7, 2}
	for i := 0; i < b.N; i++ {
		EarliestFinishTime(landStartTime, landDuration, waterStartTime, waterDuration)
	}
}
