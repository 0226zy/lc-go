package topkfrequentwords

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	t.Skip("未实现")

	tests := []struct {
		name  string
		words []string
		k     int
		want  []string
	}{
		{
			name:  "示例1",
			words: []string{"i","love","leetcode","i","love","coding"},
			k:     2,
			want:  []string{"i","love"},
		},
		{
			name:  "示例2",
			words: []string{"the","day","is","sunny","the","the","the","sunny","is","is"},
			k:     4,
			want:  []string{"the","is","sunny","day"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TopKFrequent(tt.words, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
