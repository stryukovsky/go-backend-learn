package parallel

import "sync"

func Split[T any](source []T, n int) []chan T {
	var wg sync.WaitGroup
	wg.Add(n)

	chans := make([]chan T, n)
	for i := range chans {
		chans[i] = make(chan T)
	}
	go func() {
		iter := 0
		for _, elem := range source {
			channel := chans[iter%n]
			channel <- elem
			iter++
		}
		for _, ch := range chans {
			close(ch)
		}
	}()
	return chans
}
