package main

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"main/cart-service/manager"
	"main/cart-service/repository"
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

	config := loadConfig(*c)

	router := mux.NewRouter()

	db, err := bolt.Open(config.DbAddress, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	m := manager.Manager{
		Config: config,
		Router: router,
	}

	r := new(repository.Repository)
	r.Db = db
	m.Repo = r

	defer db.Close()

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
