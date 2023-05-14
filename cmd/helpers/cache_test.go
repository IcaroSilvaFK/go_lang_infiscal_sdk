package helpers_test

import (
	"testing"

	"github.com/IcaroSilvaFK/go_lang_infiscal/cmd/helpers"
	"github.com/stretchr/testify/assert"
)

func TestSetCache(t *testing.T) {

	key := "secret"
	value := "abcd"
	expect := true
	ok := helpers.SetCache(key, value)

	assert.Equal(t, expect, ok)

}

func TestGetCache(t *testing.T) {

	key := "go"
	v := "go"
	expect := "go"

	ok := helpers.SetCache(key, v)

	assert.Equal(t, true, ok)

	got, ok := helpers.GetCache(key)

	assert.Equal(t, ok, true)
	assert.Equal(t, got, expect)
}

func TestRemoveAllOnCache(t *testing.T) {

	ok := helpers.ClearCache()

	assert.Equal(t, ok, true)
}

func TestGetAllOnUnexistentCache(t *testing.T) {

	helpers.ClearCache()
	v := helpers.ReturnAll()

	assert.Equal(t, len(v), 0)
}

func TestGetAllOnCache(t *testing.T) {

	helpers.SetCache("key", "v1")
	helpers.SetCache("key2", "v2")

	v := helpers.ReturnAll()

	assert.Equal(t, len(v), 2)
}
