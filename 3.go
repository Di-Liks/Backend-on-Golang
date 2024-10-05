package main

import "fmt"

type Mutex struct {
	Count  int
	signal chan struct{}
	done   chan struct{}
}

func (m *Mutex) Unlock() {
	m.signal <- struct{}{}
}

func main() {
	m := Mutex{
		Count: 3, signal: make(chan struct{}, 3),
		done: make(chan struct{}, 3)}
	for i := 0; i < m.Count; i++ {
		m.signal <- struct{}{}
	}
	for i := 0; i < 3; i++ {
		go func() {
			defer m.Unlock()
			fmt.Println("Hello, 世界")
			m.done <- struct{}{}
		}()
	}
	m.Wait()
}
