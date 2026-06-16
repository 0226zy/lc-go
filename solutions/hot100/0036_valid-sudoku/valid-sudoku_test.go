package validsudoku

import "testing"

func TestValidSudoku(t *testing.T) {
	// TODO: 根据 LeetCode 官方示例补充测试用例
	tests := []struct {
		name string
		// TODO: 补充输入字段
		want interface{}
	}{
		// {
		//     name: "示例1",
		//     want: nil,
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: 调用 ValidSudoku 并比较结果
		})
	}
}

func BenchmarkValidSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// TODO: 调用 ValidSudoku
	}
}
