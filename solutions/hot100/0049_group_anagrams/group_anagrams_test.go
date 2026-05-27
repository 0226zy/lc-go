package groupanagrams

import (
	"testing"

	"github.com/0226zy/lc-go/pkg/utils"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want [][]string
	}{
		// TODO: 补充测试用例
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: 调用 GroupAnagrams 并验证结果
			_ = tt.want
			_ = utils.Equal2DIntSlice
		})
	}
}
