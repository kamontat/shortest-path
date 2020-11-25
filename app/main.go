package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	datasource "challenge.kamontat.net/crea-datasource"
	formatter "challenge.kamontat.net/crea-formatter"
	logic "challenge.kamontat.net/crea-logic"
)

func main() {
	filename := flag.String("db", "example.csv", "csv file name in db directory")
	start := flag.String("start", "", "Start node name")
	end := flag.String("end", "", "Finish node name")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	flag.Parse()

	if *start == "" {
		fmt.Printf("Enter start node: ")
		fmt.Scanln(start)
	}

	if *end == "" {
		fmt.Printf("Enter end node: ")
		fmt.Scanln(end)
	}

	graph, err := datasource.Loader(path.Join(dir, "db", *filename))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	nodes, err := logic.ShortestPath(*graph, *start, *end)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	formatter.PrintTravel(nodes)
}
