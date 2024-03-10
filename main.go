package main

import (
	"fmt"
	"math"
)

func main() {
	var paths [3]float64
	for i := 0; i < 3; i++ {
		fmt.Scan(&paths[i])
	}

	graph := map[rune]map[rune]float64{
		'a': {'b': paths[0], 'c': paths[1]},
		'b': {'a': paths[0], 'c': paths[2]},
		'c': {'a': paths[1], 'b': paths[2]},
	}

	var res float64
	res += dijkstra(graph, 'a', 'b')
	res += dijkstra(graph, 'b', 'c')
	res += dijkstra(graph, 'c', 'a')

	fmt.Println(res)

	var q, w, e float64
	_, _ = fmt.Scan(&q, &w, &e)
	a := math.Min(2*(q+e), 2*(w+e))
	b := math.Min(2*(q+w), q+e+w)
	fmt.Print(int(math.Min(a, b)))
}

func dijkstra(graph map[rune]map[rune]float64, start rune, end rune) float64 {
	var distances = make(map[rune]float64)
	visited := make(map[rune]bool)
	var path = make(map[rune]rune)

	for key := range graph {
		if key != start {
			distances[key] = math.Inf(1)
		} else {
			distances[key] = 0
		}
	}

	for visited[end] == false {
		var lowestDistance = math.Inf(1)
		var node rune

		for key := range distances {
			if lowestDistance > distances[key] && !visited[key] {
				lowestDistance = distances[key]
				node = key
			}
		}

		var neighbors = graph[node]
		for key := range neighbors {
			var newDistance = distances[node] + neighbors[key]
			if newDistance < distances[key] {
				distances[key] = newDistance
				path[key] = node
			}
		}

		visited[node] = true
	}
	return distances[end]
}
