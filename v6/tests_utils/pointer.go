package tests_utils

func Pointer[T any](v T) *T {
	return &v
}
