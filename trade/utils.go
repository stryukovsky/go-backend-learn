package trade

import (
	"math/rand"

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
