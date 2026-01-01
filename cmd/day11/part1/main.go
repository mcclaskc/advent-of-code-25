package main

import (
	"advent-of-code-25/internal/utils"
	"strings"
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
	var you Node
	utils.ProcessInput(inputPath, func(line string) error {
		parts := strings.Split(line, ": ")
		name := parts[0]
		outputs := strings.Split(parts[1], " ")
		node := Node{name, outputs}
		nodes = append(nodes, node)
		if name == "you" {
			you = node
		}
		return nil
	})
	graph := Graph{
		nodes: make(map[string]*Node),
		head:  &you,
	}
	for _, node := range nodes {
		graph.nodes[node.name] = &node
	}
	return graph
}

func countPathsFromNode(graph Graph, node *Node, visited map[string]bool) int {
	numPaths := 0
	visited[node.name] = true
	for _, output := range node.outputs {
		if output == "out" {
			numPaths++
		} else if visited[output] {
			continue
		} else {
			numPaths += countPathsFromNode(graph, graph.nodes[output], visited)
		}
	}
	visited[node.name] = false
	return numPaths
}

func countAllPaths(inputPath string) int {
	graph := createGraph(inputPath)
	visited := make(map[string]bool)
	return countPathsFromNode(graph, graph.head, visited)
}

func main() {
	println(countAllPaths("../input.txt"))
}
