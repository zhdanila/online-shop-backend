package main

import (
	"math/rand"
	"testing"
	"time"
)

func generateTestData(size int, wordLen int) []string {
	rand.Seed(time.Now().UnixNano())
	letters := "abcdefghijklmnopqrstuvwxyz"
	data := make([]string, size)

	for i := 0; i < size; i++ {
		word := make([]byte, wordLen)
		for j := 0; j < wordLen; j++ {
			word[j] = letters[rand.Intn(len(letters))]
		}
		data[i] = string(word)
	}

	return data
}

var (
	smallData  = generateTestData(26, 1)
	mediumData = generateTestData(1000, 5)
	largeData  = generateTestData(100000, 10)
)

func benchmarkConcatFunction(b *testing.B, concatFunc func([]string) string, data []string) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = concatFunc(data)
	}
}

func BenchmarkConcat_Small(b *testing.B)  { benchmarkConcatFunction(b, concat, smallData) }
func BenchmarkConcat_Medium(b *testing.B) { benchmarkConcatFunction(b, concat, mediumData) }
func BenchmarkConcat_Large(b *testing.B)  { benchmarkConcatFunction(b, concat, largeData) }

func BenchmarkConcatOptimizedFirst_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFirst, smallData)
}
func BenchmarkConcatOptimizedFirst_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFirst, mediumData)
}
func BenchmarkConcatOptimizedFirst_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFirst, largeData)
}

func BenchmarkConcatOptimizedSecond_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSecond, smallData)
}
func BenchmarkConcatOptimizedSecond_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSecond, mediumData)
}
func BenchmarkConcatOptimizedSecond_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSecond, largeData)
}

func BenchmarkConcatOptimizedThird_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedThird, smallData)
}
func BenchmarkConcatOptimizedThird_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedThird, mediumData)
}
func BenchmarkConcatOptimizedThird_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedThird, largeData)
}

func BenchmarkConcatOptimizedFourth_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFourth, smallData)
}
func BenchmarkConcatOptimizedFourth_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFourth, mediumData)
}
func BenchmarkConcatOptimizedFourth_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFourth, largeData)
}

func BenchmarkConcatOptimizedFifth_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFivth, smallData)
}
func BenchmarkConcatOptimizedFifth_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFivth, mediumData)
}
func BenchmarkConcatOptimizedFifth_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedFivth, largeData)
}

func BenchmarkConcatOptimizedSixth_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSixth, smallData)
}
func BenchmarkConcatOptimizedSixth_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSixth, mediumData)
}
func BenchmarkConcatOptimizedSixth_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSixth, largeData)
}

func BenchmarkConcatOptimizedSeventh_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSeventh, smallData)
}
func BenchmarkConcatOptimizedSeventh_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSeventh, mediumData)
}
func BenchmarkConcatOptimizedSeventh_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedSeventh, largeData)
}

func BenchmarkConcatOptimizedEighth_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedEights, smallData)
}
func BenchmarkConcatOptimizedEighth_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedEights, mediumData)
}
func BenchmarkConcatOptimizedEighth_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedEights, largeData)
}

func BenchmarkConcatOptimizedNinth_Small(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedNinth, smallData)
}
func BenchmarkConcatOptimizedNinth_Medium(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedNinth, mediumData)
}
func BenchmarkConcatOptimizedNinth_Large(b *testing.B) {
	benchmarkConcatFunction(b, concatOptimizedNinth, largeData)
}
