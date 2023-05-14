package helpers_test

import (
	"testing"

	"github.com/IcaroSilvaFK/go_lang_infiscal/cmd/helpers"
	"github.com/stretchr/testify/assert"
)

func TestHasLenInArrayWithLen(t *testing.T) {
	arr := []string{"a", "b", "c"}

	assert.True(t, helpers.HasLenInArray(arr))

}

func TestHasLenInArrayNoLength(t *testing.T) {
	arr := []string{}

	assert.False(t, helpers.HasLenInArray(arr))

}
