package main

import (
	"fmt"
	Model "github.com/Lucas-Palomo/go-mycms-marsutils/examples/internal"
	"github.com/Lucas-Palomo/go-mycms-marsutils/marsutils"
	"log"
)

func main() {
	conf := Model.Conf{
		DB: Model.DatabaseConf{
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "toor",
		},
	}

	str, err := marsutils.ToString(&conf, marsutils.TOML)

	if err != nil {
		log.Fatal(err)
	}

	conf2 := Model.Conf{}
	err = marsutils.FromString(str, &conf2, marsutils.TOML)

	if err != nil {
		log.Fatal(err)
	}

	// Now conf2 has all populated values
	fmt.Printf("%#v\n", conf2)
}
