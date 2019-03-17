package main

import (
	"fmt"
	"github.com/labstack/gommon/color"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

var checkPre = color.Yellow("[") + color.Green("✓") + color.Yellow("]")
var tildPre = color.Yellow("[") + color.Green("~") + color.Yellow("]")
var crossPre = color.Yellow("[") + color.Red("✗") + color.Yellow("]")

func main() {
	var worker sync.WaitGroup
	var count int
	var err error

	// Parse arguments
	parseArgs(os.Args)

	// Check if input folder exist
	if _, err := os.Stat(arguments.Input); os.IsNotExist(err) {
		fmt.Println(crossPre +
			color.Yellow(" [") +
			color.Red(arguments.Input) +
			color.Yellow("] ") +
			color.Red("Invalid input folder."))
	}

	err = readTrackerList()
	if err != nil {
		log.Fatal(err)
	}

	files, err := ioutil.ReadDir(arguments.Input)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		worker.Add(1)
		count++

		go func(f os.FileInfo) {
			err := work(f.Name(), &worker)

			if err != nil {
				fmt.Println(crossPre +
					color.Yellow(" [") +
					color.Red(f.Name()) +
					color.Yellow("] ") +
					color.Red("Error: ") +
					color.Yellow(err.Error()))
			}
		}(f)

		if count == arguments.Concurrency {
			worker.Wait()
			count = 0
		}
	}

	worker.Wait()
}
