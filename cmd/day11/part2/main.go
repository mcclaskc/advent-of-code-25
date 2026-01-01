package main

import (
	"advent-of-code-25/internal/utils"
	"fmt"
	"strings"
	"time"
)

type Node struct {
	name    string
	outputs []string
}

type Graph struct {
	nodes map[string]*Node
	head  *Node
}

func createGraph(inputPath string) Graph {
	var nodes []Node

	// 1. Parse all data into the slice
	utils.ProcessInput(inputPath, func(line string) error {
		parts := strings.Split(line, ": ")
		name := parts[0]
		outputs := strings.Split(parts[1], " ")
		nodes = append(nodes, Node{name, outputs})
		return nil
	})

	graph := Graph{
		nodes: make(map[string]*Node),
	}

	// 2. Wire up the map using stable pointers to slice indices
	for i := range nodes {
		n := &nodes[i]
		graph.nodes[n.name] = n
	}

	// 3. Set the head pointer from the map AFTER it is fully built
	// This ensures graph.head and graph.nodes["svr"] point to the exact same memory.
	if svr, ok := graph.nodes["svr"]; ok {
		graph.head = svr
	} else {
		panic("Head node 'svr' not found in input")
	}

	return graph
}

// Progress tracker to ensure it's not actually an infinite loop
var visitedCounter int64
var lastPrint time.Time

const (
	StateNeedAll = 0
	StateNeedDAC = 1
	StateNeedFFT = 2
	StateNeedOut = 3
)

func calculateNewState(name string, oldState int) int {
	if name == "dac" {
		if oldState == StateNeedAll {
			return StateNeedFFT
		} else if oldState == StateNeedDAC {
			return StateNeedOut
		}
	} else if name == "fft" {
		if oldState == StateNeedAll {
			return StateNeedDAC
		} else if oldState == StateNeedFFT {
			return StateNeedOut
		}
	}
	return oldState
}

type CacheKey struct {
	name  string
	state int
}

var countCache = make(map[CacheKey]int)

func countPathsFromNode(graph Graph, node *Node, visited map[string]bool, oldState int) int {
	newState := calculateNewState(node.name, oldState)

	visitedCounter++
	if visitedCounter%10_000_000 == 0 {
		if time.Since(lastPrint) > time.Second {
			fmt.Printf("Visited %d nodes... current: %d\n", visitedCounter, newState)
			lastPrint = time.Now()
		}
	}

	cacheKey := CacheKey{node.name, newState}
	if count, ok := countCache[cacheKey]; ok {
		return count
	}

	numPaths := 0
	visited[node.name] = true
	for _, output := range node.outputs {
		if visited[output] {
			continue
		}
		if output == "out" {
			if newState == StateNeedOut {
				numPaths++
			}
			continue
		}
		if nextNode, ok := graph.nodes[output]; ok {
			numPaths += countPathsFromNode(graph, nextNode, visited, newState)
		}
	}
	visited[node.name] = false
	countCache[cacheKey] = numPaths
	return numPaths
}

func countAllPaths(inputPath string) int {
	graph := createGraph(inputPath)
	visited := make(map[string]bool)
	return countPathsFromNode(graph, graph.head, visited, StateNeedAll)
}

func main() {
	println(countAllPaths("../input.txt"))
}
