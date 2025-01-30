package main

import (
	"bytes"
	"strings"
	"sync"
	"unsafe"
)

func concat(str []string) string {
	result := ""
	for _, v := range str {
		result += v
	}
	return result
}

func concatOptimizedFirst(str []string) string {
	var builder strings.Builder
	for _, v := range str {
		builder.WriteString(v)
	}
	return builder.String()
}

func concatOptimizedSecond(str []string) string {
	var buffer bytes.Buffer

	for _, v := range str {
		buffer.WriteString(v)
	}
	return buffer.String()
}

func concatOptimizedThird(str []string) string {
	return strings.Join(str, "")
}

func concatOptimizedFourth(str []string) string {
	totalLen := 0
	for _, v := range str {
		totalLen += len(v)
	}

	result := make([]byte, totalLen)
	pos := 0
	for _, v := range str {
		copy(result[pos:], v)
		pos += len(v)
	}

	return string(result)
}

func concatOptimizedFivth(str []string) string {
	if len(str) == 0 {
		return ""
	}

	var wg sync.WaitGroup
	numWorkers := 4
	chunkSize := (len(str) + numWorkers - 1) / numWorkers
	results := make([]string, numWorkers)
	var mu sync.Mutex

	for i := 0; i < numWorkers; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(str) {
			end = len(str)
		}

		wg.Add(1)
		go func(i int, s []string) {
			defer wg.Done()
			result := strings.Join(s, "")
			mu.Lock()
			results[i] = result
			mu.Unlock()
		}(i, str[start:end])
	}

	wg.Wait()
	return strings.Join(results, "")
}

func concatOptimizedSixth(str []string) string {
	totalLen := 0
	for _, v := range str {
		totalLen += len(v)
	}

	result := make([]byte, 0, totalLen)

	for _, v := range str {
		result = append(result, v...)
	}

	return string(result)
}

func concatOptimizedSeventh(str []string) string {
	totalLen := 0
	for _, v := range str {
		totalLen += len(v)
	}

	result := make([]byte, totalLen)
	pos := 0
	for _, v := range str {
		copy(result[pos:], v)
		pos += len(v)
	}

	return *(*string)(unsafe.Pointer(&result))
}

func concatOptimizedEights(str []string) string {
	totalLen := 0
	for _, v := range str {
		totalLen += len(v)
	}

	result := make([]byte, totalLen)
	pos := 0
	for _, v := range str {
		copy(result[pos:], v)
		pos += len(v)
	}

	return unsafe.String(&result[0], len(result))
}

func concatOptimizedNinth(str []string) string {
	byteSlices := make([][]byte, len(str))
	for i, v := range str {
		byteSlices[i] = []byte(v)
	}
	return string(bytes.Join(byteSlices, nil))
}
