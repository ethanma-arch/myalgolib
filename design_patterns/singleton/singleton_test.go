package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance_ReturnsSameInstance(t *testing.T) {
	a := GetInstance()
	b := GetInstance()
	if a != b {
		t.Fatal("GetInstance() should return the same instance")
	}
}

func TestGetInstance_ConcurrentSafe(t *testing.T) {
	const n = 100
	var instances [n]*Singleton
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			instances[i] = GetInstance()
		}(i)
	}
	wg.Wait()
	first := instances[0]
	for i := 1; i < n; i++ {
		if instances[i] != first {
			t.Fatalf("instance %d != instance 0", i)
		}
	}
}
