package dailytemperatures

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name         string
		temperatures []int
		want         []int
	}{
		{
			name:         "官方示例1",
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name:         "官方示例2",
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "官方示例3",
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
		{
			name:         "单元素",
			temperatures: []int{50},
			want:         []int{0},
		},
		{
			name:         "全部递减",
			temperatures: []int{90, 80, 70, 60},
			want:         []int{0, 0, 0, 0},
		},
		{
			name:         "全部递增",
			temperatures: []int{60, 70, 80, 90},
			want:         []int{1, 1, 1, 0},
		},
		{
			name:         "全部相同",
			temperatures: []int{50, 50, 50, 50},
			want:         []int{0, 0, 0, 0},
		},
		{
			name:         "空数组",
			temperatures: []int{},
			want:         []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := DailyTemperatures(tt.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DailyTemperatures() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkDailyTemperatures(b *testing.B) {
	temperatures := []int{73, 74, 75, 71, 69, 72, 76, 73}
	for i := 0; i < b.N; i++ {
		DailyTemperatures(temperatures)
	}
}
