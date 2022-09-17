package bigcache

//
//import (
//	"fmt"
//	"github.com/cespare/xxhash"
//	"math/rand"
//	"strconv"
//	"testing"
//	"time"
//)
//
//type tempHaser struct {
//}
//
//func (t *tempHaser) Sum64(key string) uint64 {
//	return xxhash.Sum64([]byte(key))
//}
//func BenchmarkReadFromCache2(b *testing.B) {
//	for _, shards := range []int{1, 512, 1024, 8192} {
//		b.Run(fmt.Sprintf("%d-shards", shards), func(b *testing.B) {
//			readFromCache(b, shards, false)
//		})
//	}
//}
//func BenchmarkWriteToCache2(b *testing.B) {
//	for _, shards := range []int{1, 512, 1024, 8192} {
//		b.Run(fmt.Sprintf("%d-shards", shards), func(b *testing.B) {
//			writeToCache(b, shards, 100*time.Second, b.N)
//		})
//	}
//}
//func BenchmarkXXWriteToCache(b *testing.B) {
//	for _, shards := range []int{1, 512, 1024, 8192} {
//		b.Run(fmt.Sprintf("%d-shards", shards), func(b *testing.B) {
//			xxWriteToCache(b, shards, 100*time.Second, b.N)
//		})
//	}
//}
//
//func BenchmarkXXReadFromCache(b *testing.B) {
//	for _, shards := range []int{1, 512, 1024, 8192} {
//		b.Run(fmt.Sprintf("%d-shards", shards), func(b *testing.B) {
//			xxReadFromCache(b, shards, false)
//		})
//	}
//}
//
//func xxWriteToCache(b *testing.B, shards int, lifeWindow time.Duration, requestsInLifeWindow int) {
//	cache, _ := NewBigCache(Config{
//		Shards:             shards,
//		LifeWindow:         lifeWindow,
//		MaxEntriesInWindow: max(requestsInLifeWindow, 100),
//		MaxEntrySize:       500,
//		Hasher:             &tempHaser{},
//	})
//	rand.Seed(time.Now().Unix())
//
//	b.RunParallel(func(pb *testing.PB) {
//		id := rand.Int()
//		counter := 0
//
//		b.ReportAllocs()
//		for pb.Next() {
//			cache.Set(fmt.Sprintf("key-%d-%d", id, counter), message)
//			counter = counter + 1
//		}
//	})
//}
//
////func appendToCache(b *testing.B, shards int, lifeWindow time.Duration, requestsInLifeWindow int) {
////	cache, _ := NewBigCache(Config{
////		Shards:             shards,
////		LifeWindow:         lifeWindow,
////		MaxEntriesInWindow: max(requestsInLifeWindow, 100),
////		MaxEntrySize:       2000,
////	})
////	rand.Seed(time.Now().Unix())
////
////	b.RunParallel(func(pb *testing.PB) {
////		id := rand.Int()
////		counter := 0
////
////		b.ReportAllocs()
////		for pb.Next() {
////			key := fmt.Sprintf("key-%d-%d", id, counter)
////			for j := 0; j < 7; j++ {
////				cache.Append(key, message)
////			}
////			counter = counter + 1
////		}
////	})
////}
//
//func xxReadFromCache(b *testing.B, shards int, info bool) {
//	cache, _ := NewBigCache(Config{
//		Shards:             shards,
//		LifeWindow:         1000 * time.Second,
//		MaxEntriesInWindow: max(b.N, 100),
//		MaxEntrySize:       500,
//		Hasher:             &tempHaser{},
//	})
//	for i := 0; i < b.N; i++ {
//		cache.Set(strconv.Itoa(i), message)
//	}
//	b.ResetTimer()
//
//	b.RunParallel(func(pb *testing.PB) {
//		b.ReportAllocs()
//
//		for pb.Next() {
//			if info {
//				cache.GetWithInfo(strconv.Itoa(rand.Intn(b.N)))
//			} else {
//				cache.Get(strconv.Itoa(rand.Intn(b.N)))
//			}
//		}
//	})
//}
//
////func readFromCacheNonExistentKeys(b *testing.B, shards int) {
////	cache, _ := NewBigCache(Config{
////		Shards:             shards,
////		LifeWindow:         1000 * time.Second,
////		MaxEntriesInWindow: max(b.N, 100),
////		MaxEntrySize:       500,
////	})
////	b.ResetTimer()
////
////	b.RunParallel(func(pb *testing.PB) {
////		b.ReportAllocs()
////
////		for pb.Next() {
////			cache.Get(strconv.Itoa(rand.Intn(b.N)))
////		}
////	})
////}
