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
	startNode := flag.String("start", "A", "Start node name")
	endNode := flag.String("end", "B", "Finish node name")

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	flag.Parse()

	graph, err := datasource.Loader(path.Join(dir, "db", *filename))
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	nodes, err := logic.ShortestPath(*graph, *startNode, *endNode)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	formatter.PrintTravel(nodes)
}
