package main

import (
  "fmt"
  "math"
  "sync"
)

// Функция для проверки простого числа
func isPrime(n int) bool {
  if n < 2 {
    return false
  }
  for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
    if n%i == 0 {
      return false
    }
  }
  return true
}

// Функция задачи, проверяющая простоту числа
func primeTask(n int, wg *sync.WaitGroup, ch chan int) {
  defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении задачи
  if isPrime(n) {
    ch <- n // Отправляем простое число в канал
  }
}

func main() {
  // Массив чисел для проверки
  numbers := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

  // Канал для хранения простых чисел
  ch := make(chan int, len(numbers))

  // WaitGroup для синхронизации задач
  var wg sync.WaitGroup

  // Для каждого числа создаем задачу
  for _, num := range numbers {
    wg.Add(1)
    go primeTask(num, &wg, ch)
  }

  // Ждем завершения всех задач
  wg.Wait()

  // Закрываем канал, так как все горутины завершены
  close(ch)

  // Вывод всех простых чисел
  primes := []int{}
  for prime := range ch {
    primes = append(primes, prime)
  }

  // Выводим результат
  fmt.Printf("Количество простых чисел: %d\n", len(primes))
  fmt.Printf("Простые числа: %v\n", primes)
}
