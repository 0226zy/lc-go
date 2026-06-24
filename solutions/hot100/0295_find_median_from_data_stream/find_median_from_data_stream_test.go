package findmedianfromdatastream

import (
	"math"
	"testing"
)

func TestMedianFinder(t *testing.T) {
t.Skip("未实现")
	t.Run("基本用例", func(t *testing.T) {
		mf := Constructor()
		mf.AddNum(1)
		mf.AddNum(2)
		got := mf.FindMedian()
		want := 1.5
		if math.Abs(got-want) > 1e-5 {
			t.Errorf("第一次 FindMedian() = %f, want %f", got, want)
		}

		mf.AddNum(3)
		got = mf.FindMedian()
		want = 2.0
		if math.Abs(got-want) > 1e-5 {
			t.Errorf("第二次 FindMedian() = %f, want %f", got, want)
		}
	})

	t.Run("单元素", func(t *testing.T) {
		// TODO: 实现后取消注释
		// mf := Constructor()
		// mf.AddNum(42)
		// if mf.FindMedian() != 42.0 {
		// 	t.Errorf("单元素中位数应为该值本身")
		// }
	})

	t.Run("负数和正数", func(t *testing.T) {
		// TODO: 实现后取消注释
		// mf := Constructor()
		// mf.AddNum(-1)
		// mf.AddNum(-2)
		// mf.AddNum(-3)
		// if mf.FindMedian() != -2.0 {
		// 	t.Errorf("中位数应为 -2.0")
		// }
	})

	t.Run("大量相同值", func(t *testing.T) {
		// TODO: 实现后取消注释
		// mf := Constructor()
		// for i := 0; i < 100; i++ {
		// 	mf.AddNum(5)
		// }
		// if mf.FindMedian() != 5.0 {
		// 	t.Errorf("全部相同值的中位数应为 5.0")
		// }
	})
}

func BenchmarkMedianFinder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		mf := Constructor()
		mf.AddNum(1)
		mf.AddNum(2)
		mf.FindMedian()
	}
}
