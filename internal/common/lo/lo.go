package lo

import "fmt"

func Map[T any, U any](s []T, f func(T) U) []U {
	var result []U
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}

func Filter[T any](s []T, f func(T) bool) []T {
	var result []T
	for _, v := range s {
		if f(v) {
			result = append(result, v)
		}
	}
	return result
}

func Every[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if !f(v) {
			return false
		}
	}
	return true
}

func Some[T any](s []T, f func(T) bool) bool {
	for _, v := range s {
		if f(v) {
			return true
		}
	}
	return false
}

func AsStrings[T fmt.Stringer](s []T) []string {
	res := make([]string, len(s))
	for i, v := range s {
		res[i] = v.String()
	}
	return res
}
