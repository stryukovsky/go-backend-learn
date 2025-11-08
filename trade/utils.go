package trade

import "math/rand"

func RandomChoice[T any](slice []T) T {
	return slice[rand.Intn(len(slice))]
}
