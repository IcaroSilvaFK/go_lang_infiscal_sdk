package main

import (
	"fmt"
	"log"

	"github.com/IcaroSilvaFK/go_lang_infiscal/cmd/helpers"
)

func main() {
	ok := helpers.SetCache("k1", "v1")

	if !ok {
		log.Fatal("cache not set")
	}

	ok = helpers.SetCache("k2", "v2")

	if !ok {
		log.Fatal("cache not set")
	}

	v, _ := helpers.GetCache("k2")

	fmt.Println(v)

}
