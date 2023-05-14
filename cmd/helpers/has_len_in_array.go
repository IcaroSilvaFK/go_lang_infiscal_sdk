package helpers

func HasLenInArray[T map[string]string | []string](arr T) bool {

	return len(arr) > 0
}
