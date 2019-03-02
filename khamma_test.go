package khamma

import "testing"

func BenchmarkAnalyze(b *testing.B) {
	err := InitializeWithDefault()
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		Analyze("집에 가서 잠을 자고 싶다.", "")
	}
}
