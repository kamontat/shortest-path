package formatter

import (
	"fmt"

	model "challenge.kamontat.net/crea-model"
)

// PrintTravel print result to console
func PrintTravel(nodes []*model.Node) {
	current := nodes[0]
	last := nodes[1:]

	stops := len(last) - 1
	total := int64(0)

	realCurrent := current.ToString()
	fmt.Printf("Start at %s\n", realCurrent)

	for _, next := range last {
		edge, err := current.GetEdge(next.GetName())

		if err == nil {
			total = total + edge.GetDistance()
			fmt.Printf("Move %s to %s take %d minutes\n", current.ToString(), next.ToString(), edge.GetDistance())

			current = next
		} else {
			fmt.Printf("%s", err.Error())
		}
	}

	fmt.Printf("Summary: %s to %s has %d stops in %d minutes", realCurrent, current.ToString(), stops, total)
}
