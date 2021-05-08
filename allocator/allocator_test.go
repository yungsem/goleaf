package allocator

import (
	"fmt"
	"testing"
)

func BenchmarkAllocateId(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 要测试的操作
		AllocateId("mes")
	}
}

func TestAllocateId(t *testing.T) {
	id, err := AllocateId("mes")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(id)
}
