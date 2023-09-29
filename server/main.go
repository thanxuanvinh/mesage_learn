package main

import (
	"fmt"
	"io"
	"log"
	"os"

	sio "github.com/thanxuanvinh/mesage/io"
	_ "github.com/thanxuanvinh/mesage/server/hub"
	_ "github.com/thanxuanvinh/mesage/server/service"
)

func main() {
	hubMan := newHubManager()
	fmt.Println(hubMan)
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
