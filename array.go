package main

import "fmt"

func modifySlice(s []int) ([]int, []int) {
  var a []int
  var b []int
  for i, v := range s {
    if i%2 == 0 {
      a = append(a, v)
    } else {
      b = append(b, v)
    }
  }
  return a, b
}

func revSlice(s []int) []int {
  rev := make([]int, len(s))
  copy(rev, s)
  for i := 0; i < len(rev)/2; i++ {
    rev[i], rev[len(rev)-1-i] = rev[len(rev)-1-i], rev[i]
  }
  return rev
}

func mergeSlice(a []int, b []int) []int {
  return append(a, b...)
}

func main() {
  s := []int{1, 2, 3, 4, 5, 6}
  a, b := modifySlice(s)
  one := mergeSlice(a, revSlice(b))
  two := mergeSlice(a, b)
  three := mergeSlice(revSlice(a), revSlice(b))
  four := mergeSlice(revSlice(a), b)
  fmt.Println("Чётные элементы:", a)
  fmt.Println("Нечётные элементы:", b)
  fmt.Println("Второй массив реверсивен:", one)
  fmt.Println("Нет реверсивных массивов:", two)
  fmt.Println("Все массивы реверсивны:", three)
  fmt.Println("Первый массив реверсивен:", four)
  fmt.Println("Исходный массив:", s)
}
