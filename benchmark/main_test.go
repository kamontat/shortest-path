package test_test

import (
	"path"
	"testing"

	datasource "challenge.kamontat.net/crea-datasource"
	logic "challenge.kamontat.net/crea-logic"
)

var filepath = path.Join("..", "app", "db", "test-1000.csv")

func BenchmarkLoadCsvFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		datasource.Loader(filepath)
	}
}

func BenchmarkFindShortestPath(b *testing.B) {
	graph, _ := datasource.Loader(filepath)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logic.ShortestPath(*graph, "A", "Z")
	}
}
