package main

import (
	"fmt"
	"io"
	"log"
	"os"

	sio "github.com/thanxuanvinh/mesage/io"
	_ "github.com/thanxuanvinh/mesage/server/hub"
	_ "github.com/thanxuanvinh/mesage/server/service"
	"github.com/thanxuanvinh/mesage/services"
)

func main() {
	hubMan := newHubManager()
	serviceMan := newServiceManager()
	fmt.Println(hubMan)
	fmt.Println(serviceMan)
}

func newHubManager() *sio.Manager {
	configReader := make(chan io.Reader, 16)
	for _, s := range []string{"lemon", "mongo"} {
		f, err := os.Open(fmt.Sprintf("./configs/hubs/%sHub.yaml", s))
		if err != nil {
			log.Fatal("error opening config file")
		}
		configReader <- f
	}
	close(configReader)

	man, err := sio.NewManager(configReader)
	if err != nil {
		log.Fatalln("error creating new hub manager: ", err)
	}
	return man
}

func newServiceManager() *services.Manager {
	configReaders := make(chan io.Reader, 16)
	f, err := os.Open("./configs/services/exampleService.yaml")
	if err != nil {
		log.Fatal("error opening config file ", err)
	}

	configReaders <- f
	close(configReaders)

	man, err := services.NewManager(configReaders)
	if err != nil {
		log.Fatalln("error creating new service manager: ", err)
	}

	return man
}
