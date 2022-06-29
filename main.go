package main

import (
	"Corgi/attack"
	"Corgi/common"
	"Corgi/configs"
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	//target := "https://httpbin.org/get"
	filePath := "C:\\Users\\Swift001\\GolandProjects\\Corgi\\urls.txt"

	file, err := common.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	configs.Banner()
	logrus.Info("Corgi Start Running! wow wow\n")
	fmt.Println("urls:", file)
	for _, target := range file {
		attack.Soldier(target)
	}

}
