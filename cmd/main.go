package main

import (
	"flag"
	"fmt"
	"go-func-task/service"
	"go-func-task/service/presenter"
	"go-func-task/service/produser"
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

	svc := service.NewService(produce, present)

	err := svc.Run()
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
