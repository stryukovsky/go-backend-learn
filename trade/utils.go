package trade

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"

	"github.com/samber/lo"
	"github.com/samber/lo/mutable"
	"golang.org/x/sync/errgroup"
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
	ValuesCh       chan ParsedEvent
}

func ParseEVMEvents[RawEvent, ParsedEvent any](
	parallelFactor int,
	workerName string,
	chainId string,
	events []RawEvent,
	parseFn func(ParallelEVMParserTask[ParsedEvent], RawEvent) error,
) ([]ParsedEvent, error) {
	if len(events) == 0 {
		return []ParsedEvent{}, nil
	}
	eventChunks := Chunks(events, parallelFactor)
	valuesCh := make(chan ParsedEvent, len(events))
	wg, ctx := errgroup.WithContext(context.Background())
	task := ParallelEVMParserTask[ParsedEvent]{
		ParallelFactor: parallelFactor,
		WorkerName:     workerName,
		ChainId:        chainId,
		ValuesCh:       valuesCh,
	}
	for i, chunk := range eventChunks {
		wg.Go(func() error {
			slog.Debug(fmt.Sprintf("[%s] Parsing %d-th chunk of Supply Events", workerName, i+1))
			for _, generalEvent := range chunk {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					if err := parseFn(task, generalEvent); err != nil {
						return err
					}
				}
			}
			return nil
		})
	}
	err := wg.Wait()
	if err != nil {
		close(valuesCh)
		// drain channel
		for range valuesCh {
		}
		return nil, err
	}
	close(valuesCh)
	result := make([]ParsedEvent, 0, len(events))
	for item := range valuesCh {
		result = append(result, item)
	}
	return result, nil
}

type WithURL interface {
	URL() string
}

func RetryEthCall[CallerPtr WithURL, R any](listCallers func() []CallerPtr, call func(CallerPtr) (R, error)) (R, error) {
	originalCallers := listCallers()
	shuffledCallers := make([]CallerPtr, len(originalCallers))
	copy(shuffledCallers, originalCallers)
	mutable.Shuffle(shuffledCallers)
	var zero R
	if len(shuffledCallers) == 0 {
		return zero, fmt.Errorf("No callers provided")
	}
	// firstly attempt random client

	firstAttemptCaller := RandomChoice(shuffledCallers)
	result, err := call(firstAttemptCaller)
	if err == nil {
		// skip any next attempts of calling
		return result, err
	}
	for _, client := range shuffledCallers {
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
