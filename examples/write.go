package main

import (
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
	err := marsutils.WriteFile("./data/config.json", &conf, marsutils.JSON)
	if err != nil {
		log.Fatal(err)
	}
	// Now conf content is in the config.json
}
