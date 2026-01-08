package trade

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"sync"

	"github.com/samber/lo"
)

func RandomChoice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}

func Chunks[T any](slice []T, chunksCount int) [][]T {
	chunkSize := len(slice) / chunksCount
	if chunkSize == 0 {
		// this condition also covers case when len(slice) < chunksCount so this is case when chunk split is not needed
		return [][]T{slice}
	}
	return lo.Chunk(slice, chunkSize)
}

type ParallelEVMParserTask[ParsedEvent any] struct {
	ParallelFactor int
	WorkerName     string
	ChainId        string
	Wg             *sync.WaitGroup
	ValuesCh       chan ParsedEvent
	Cancel         func()
}

func ParseEVMEvents[RawEvent, ParsedEvent any](
	parallelFactor int,
	workerName string,
	chainId string,
	events []RawEvent,
	parseFn func(*ParallelEVMParserTask[ParsedEvent], RawEvent),
) ([]ParsedEvent, error) {
	if len(events) == 0 {
		return []ParsedEvent{}, nil
	}
	eventChunks := Chunks(events, parallelFactor)
	var wg sync.WaitGroup
	wg.Add(len(eventChunks))
	valuesCh := make(chan ParsedEvent, parallelFactor)
	ctx, cancel := context.WithCancel(context.Background())
	defer ctx.Done()
	task := ParallelEVMParserTask[ParsedEvent]{
		ParallelFactor: parallelFactor,
		WorkerName:     workerName,
		ChainId:        chainId,
		Wg:             &wg,
		ValuesCh:       valuesCh,
		Cancel:         cancel,
	}
	for i, chunk := range eventChunks {
		go func() {
			defer wg.Done()
			slog.Debug(fmt.Sprintf("[%s] Parsing %d-th chunk of Supply Events", workerName, i+1))
			for _, generalEvent := range chunk {
				select {
				case <-ctx.Done():
					return
				default:
					parseFn(&task, generalEvent)
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(valuesCh)
	}()
	result := make([]ParsedEvent, 0, len(events))
	for item := range valuesCh {
		result = append(result, item)
	}
	cancel()
	return result, nil
}

type WithURL interface {
	URL() string
}

func RetryEthCall[CallerPtr WithURL, R any](listCallers func() []CallerPtr, call func(CallerPtr) (R, error)) (R, error) {
	callers := listCallers()
	 var zero R
	if len(callers) == 0 {
		return zero, fmt.Errorf("No callers provided")
	}
	// firstly attempt random client

	firstAttemptCaller := RandomChoice(callers)
	result, err := call(firstAttemptCaller)
	if err == nil {
		// skip any next attempts of calling
		return result, err
	}
	for _, client := range callers {
		r, e := call(client)
		result = r
		err = e
		if e != nil {
			slog.Warn(fmt.Sprintf("Client with url %s failed to get  %s", client.URL(), err.Error()))
		} else {
			err = nil
			// finally get value; leave the cycle
			break
		}
	}
	if err != nil {
		return zero, fmt.Errorf("All clients could not perform call")
	}
	return result, nil
}
