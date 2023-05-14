package helpers

import (
	"fmt"
)

var cache = make(map[string]string)

func SetCache(k, v string) bool {

	if len(cache) == 0 {
		cache[k] = v
		return true
	}

	for key := range cache {

		if key == k {
			cache[key] = v

			return true
		} else {
			fmt.Println("else :", k, v)
			cache[k] = v
			return true
		}
	}

	return false

}

func GetCache(k string) (string, bool) {

	if len(cache) == 0 {
		return "", false
	}

	v := cache[k]

	if v == "" {
		return "", false
	}

	return v, true
}
