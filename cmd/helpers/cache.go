package helpers

var cache = make(map[string]string)

func SetCache(k, v string) bool {

	if HasLenInArray(cache) {
		cache[k] = v
		return true
	}

	for key := range cache {

		if key == k {
			cache[key] = v

			return true
		} else {
			cache[k] = v
			return true
		}
	}

	return false

}

func GetCache(k string) (string, bool) {

	if HasLenInArray(cache) {
		return "", false
	}

	v := cache[k]

	if v == "" {
		return "", false
	}

	return v, true
}

func ReturnAll() map[string]string {

	return cache
}

func ClearCache() bool {
	cache = make(map[string]string)
	return true
}

func DeleteValueOnCache(k string) bool {

	if HasLenInArray(cache) {
		delete(cache, k)
		return true
	}
	return false
}
