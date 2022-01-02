package singleton

import "sync"

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var (
	instance *singleton
	once     sync.Once
)

func init() {
	instance = &singleton{count: 100}
}

func GetInstance() Singleton {
	// if instance == nil {
	// if there isn't data -> constructor
	//	instance = &singleton{count: 100}
	// }

	/* --- Lazy load
	When multiple processes enter at the same time due to Goroutine, they give different addresses
	Solution: use Do() in the Sync package
	run only once -> same the memory address
	once.Do(func() {
		instance = &singleton{count: 100}
	}) */

	// --- Diligent load in line 18
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
