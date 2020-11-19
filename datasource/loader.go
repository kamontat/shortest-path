package datasource

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	model "challenge.kamontat.net/crea-model"
)

// Loader will load csv file and change to path[]
func Loader(filename string) (*model.Graph, error) {
	graph := model.NewGraph()
	csvFile, err := os.Open(filename)
	defer csvFile.Close()

	if err != nil {
		fmt.Println(err)
		return graph, err
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
		return graph, err
	}

	for _, line := range csvLines {
		number64, err := strconv.ParseInt(line[2], 10, 64)
		if err != nil {
			fmt.Println(err)
			return graph, err
		}

		node1 := graph.GetNode(line[0])
		node2 := graph.GetNode(line[1])

		node1.AddEdge(node2, number64)
	}

	return graph, nil
}
