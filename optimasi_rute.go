package main

import (
	"fmt"
	"math"
)

const INF = math.MaxInt32

func dijkstra(graph map[string]map[string]int, start string) map[string]int {
	dist := make(map[string]int)
	visited := make(map[string]bool)

	for node := range graph {
		dist[node] = INF
	}

	dist[start] = 0

	for i := 0; i < len(graph); i++ {
		minDist := INF
		var minNode string

		for node, d := range dist {
			if !visited[node] && d < minDist {
				minDist = d
				minNode = node
			}
		}

		if minNode == "" {
			break
		}

		visited[minNode] = true

		for neighbor, weight := range graph[minNode] {
			if dist[minNode]+weight < dist[neighbor] {
				dist[neighbor] = dist[minNode] + weight
			}
		}
	}

	return dist
}

func main() {
	graph := map[string]map[string]int{
		"A": {"B": 4, "C": 2},
		"B": {"C": 1, "D": 5},
		"C": {"D": 8, "E": 10},
		"D": {"E": 2},
		"E": {},
	}

	start := "A"
	result := dijkstra(graph, start)

	fmt.Println("Waktu tercepat dari Pos Satpam:")
	for node, time := range result {
		fmt.Printf("Ke %s : %d menit\n", node, time)
	}
}
