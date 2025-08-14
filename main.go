package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return nil
	}

	rand.NewSource(time.Now().UnixNano())

	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(size * 10)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, value := range data[1:] {
		if value > max {
			max = value
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}

	size := len(data) / CHUNKS
	maxSlice := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		go func(index int) {
			defer wg.Done()
			start := index * size
			end := start + size

			// Для последнего чанка берем все оставшиеся элементы
			if index == CHUNKS-1 {
				end = len(data)
			}

			chunk := data[start:end]
			maxSlice[index] = maximum(chunk)
		}(i)
	}

	wg.Wait()
	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE) //1
	// ваш код здесь

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)
	if len(data) == 0 {
		fmt.Println("Не удалось сгенерировать данные")
		return
	}

	fmt.Println("Ищем максимальное значение в один поток") //2
	// ваш код здесь

	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed) //3

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS) //4
	// ваш код здесь

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed) //4
}
