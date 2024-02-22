package main

import (
	"fmt"
	Model "github.com/Lucas-Palomo/go-mycms-marsutils/examples/internal"
	"github.com/Lucas-Palomo/go-mycms-marsutils/marsutils"
	"log"
)

func main() {
	conf := Model.Conf{}
	err := marsutils.ReadFile("./data/config.toml", &conf, marsutils.TOML)
	if err != nil {
		log.Fatal(err)
	}
	// Now conf has all populated values
	fmt.Printf("%#v\n", conf)
}
