package sliceTools

import (
	"fmt"
	"strings"
)

func RemoveIndex[T comparable](arr *[]T, index int) {
	*arr = append((*arr)[:index], (*arr)[index+1:]...)
}

func RemoveFirst[T comparable](arr *[]T, value T) {
	for i, v := range *arr {
		if value == v {
			RemoveIndex(arr, i)
			return
		}
	}
}

func RemoveAll[T comparable](arr *[]T, value T) {
	tmp := Clone(*arr)
	for i, v := range *tmp {
		if value == v {
			RemoveIndex(arr, i)
		}
	}
}

func IndexOf[T comparable](arr []T, value T) int {
	for i, v := range arr {
		if value == v {
			return i
		}
	}
	return -1
}

func IsEmpty[T comparable](arr []T) bool {
	return len(arr) == 0
}

func Clone[T comparable](arr []T) *[]T {
	ArrClone := make([]T, len(arr))
	ArrClone = append(ArrClone, arr...)
	return &ArrClone
}

func ToString[T comparable](arr []T) (s string) {
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range arr {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	sb.WriteString("]")
	return sb.String()
}

func Contains[T comparable](arr []T, value T) bool {
	for _, v := range arr {
		if value == v {
			return true
		}
	}
	return false
}
