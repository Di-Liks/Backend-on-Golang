package main

import (
    "fmt"
    "math"
    "sync"
)

type Task func() // Тип для задания

// Структура пула задач с мьютексом для безопасной работы с потоками
type TaskPool struct {
    tasks []Task
    mu    sync.Mutex
}

// Интерфейс для пула задач
type ITaskPool interface {
    NextTask() Task
    Push(Task)
}

// Интерфейс для выполнения задач
type Executor interface {
    ExecNext()
}

// Метод добавления задачи в пул
func (tp *TaskPool) AddTask(t Task) {
    tp.mu.Lock()
    tp.tasks = append(tp.tasks, t)
    tp.mu.Unlock()
}

// Метод получения следующей задачи
func (tp *TaskPool) GetNextTask() Task {
    tp.mu.Lock()
    defer tp.mu.Unlock()

    if len(tp.tasks) == 0 {
        return nil
    }

    nextTask := tp.tasks[0]
    tp.tasks = tp.tasks[1:]
    return nextTask
}

// Метод выполнения следующей задачи
func (tp *TaskPool) ExecuteNext() {
    task := tp.GetNextTask()
    if task != nil {
        task()
    }
}

// Глобальные переменные для хранения простых чисел и их количества
var primeNumbers []int
var totalPrimes int
var primeMu sync.Mutex

// Функция для проверки простого числа
func isNumberPrime(n int) bool {
    if n < 2 {
        return false
    }
    for _, prime := range primeNumbers {
        if prime > int(math.Sqrt(float64(n))) {
            break
        }
        if n%prime == 0 {
            return false
        }
    }
    return true
}

// Создание задачи для проверки числа на простоту
func createPrimeCheckTask(n int) Task {
    return func() {
        if isNumberPrime(n) {
            primeMu.Lock()
            primeNumbers = append(primeNumbers, n)
            totalPrimes++
            primeMu.Unlock()
        }
    }
}

// Генерация последовательности чисел от 1 до n
func generateNumbers(limit int) []int {
    numbers := make([]int, limit)
    for i := 0; i < limit; i++ {
        numbers[i] = i + 1
    }
    return numbers
}

func main() {
    numbers := generateNumbers(10) // Генерация чисел от 1 до 10
    var pool TaskPool               // Инициализация пула задач

    // Добавление задач для каждого числа
    for _, n := range numbers {
        pool.AddTask(createPrimeCheckTask(n))
    }

    // Выполнение задач, пока не закончатся
    for len(pool.tasks) > 0 {
        pool.ExecuteNext()
    }

    // Вывод результата
    fmt.Printf("Общее количество простых чисел: %d\n", totalPrimes)
    fmt.Println("Простые числа:", primeNumbers)
}
