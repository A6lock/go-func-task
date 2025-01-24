package main

import (
	"flag"
	"fmt"
	"go-func-task/presenter"
	"go-func-task/produser"
	"go-func-task/service"
	"os"
)

func main() {
	inputFilePath := flag.String("input", "", "Path to input file")
	outputFile := flag.String("output", "output.txt", "Path to output file")
	flag.Parse()

	if *inputFilePath == "" {
		fmt.Println("Необходимо указать путь к входному файлу с помощью флага --input")
		os.Exit(1)
	}

	produce := produser.NewProducer(*inputFilePath)
	present := presenter.NewPresenter(*outputFile)

	srvce := service.NewService(produce, present)

	err := srvce.Run()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
