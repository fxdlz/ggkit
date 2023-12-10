package slice

import "errors"

const LoadFactor float32 = 0.625

func Delete[T any](slice []T, index int) ([]T, T, error) {
	err := checkRange(slice, index)
	if err != nil {
		return nil, nil, err
	}
	target := slice[index]
	for i := index + 1; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]
	slice = shrink(slice)
	return slice, target, nil
}

func checkRange[T any](slice []T, index int) error {
	if index >= len(slice) || index < 0 {
		return errors.New("index out of range")
	}
	return nil
}

func shrink[T any](slice []T) []T {
	c, changed := calCapacity(slice)
	if !changed {
		return slice
	}
	newSlice := make([]T, 0, c)
	newSlice = append(newSlice, slice...)
	return newSlice
}

func calCapacity[T any](slice []T) (int, bool) {
	l, c := len(slice), cap(slice)
	if c < 64 {
		return c, false
	}
	if c >= 1024 && c/l >= 2 {
		return int(float32(c) * LoadFactor), true
	}
	if c < 1024 && c/l >= 4 {
		return c / 2, true
	}
	return c, false
}
