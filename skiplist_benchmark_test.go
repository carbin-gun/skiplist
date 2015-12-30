package skiplist

import (
	"fmt"
	"testing"
)

func BenchmarkInsert_1000(b *testing.B)    { benchInsert(b, 1000) }
func BenchmarkInsert_10000(b *testing.B)   { benchInsert(b, 10000) }
func BenchmarkInsert_100000(b *testing.B)  { benchInsert(b, 100000) }
func BenchmarkInsert_1000000(b *testing.B) { benchInsert(b, 1000000) }

func benchInsert(b *testing.B, total int64) {
	var i int64
	for i = 0; i < int64(b.N); i++ {
		l := NewList()
		var j int64
		for j = 0; j < total; j++ {
			l.Insert(j, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
		}
	}
}

func BenchmarkParallelInsert(b *testing.B) {
	l := NewList()
	var i int64 = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
		}
	})
}

func BenchmarkDelete_1000(b *testing.B)    { benchDelete(b, 1000) }
func BenchmarkDelete_10000(b *testing.B)   { benchDelete(b, 10000) }
func BenchmarkDelete_100000(b *testing.B)  { benchDelete(b, 100000) }
func BenchmarkDelete_1000000(b *testing.B) { benchDelete(b, 1000000) }

func benchDelete(b *testing.B, total int64) {
	l := NewList()
	var i int64
	for i = 0; i < total; i++ {
		l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
	}
	b.ResetTimer()
	for i = 0; i < int64(b.N); i++ {
		var j int64
		for j = 0; j < total; j++ {
			l.Delete(i)
		}
	}
}

func BenchmarkParallelDelete(b *testing.B) {
	l := NewList()
	var i int64
	for i = 0; i < 1000000; i++ {
		l.Insert(i, []byte(fmt.Sprintf("Skiplist node insert %d", i)))
	}
	b.ResetTimer()
	i = 0
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i++
			l.Delete(i)
		}
	})
}
