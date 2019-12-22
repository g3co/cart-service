package main

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"main/cart-service/manager"
	"main/cart-service/structs"
	"os"
	"path/filepath"
)

func main() {

	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile | log.LUTC)

	c := flag.String("c", "", "")
	flag.Parse()

	if string(*c) == "" {
		log.Println("config file is not defined")
		os.Exit(1)
	}

	fmt.Println("C is ", *c)

	config := loadConfig(*c)

	m := manager.Manager{
		Config: config,
	}

	m.Run()
}

func loadConfig(path string) structs.Config {
	absPath, _ := filepath.Abs(fmt.Sprintf("%s", path))

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var config structs.Config

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	return config
}
