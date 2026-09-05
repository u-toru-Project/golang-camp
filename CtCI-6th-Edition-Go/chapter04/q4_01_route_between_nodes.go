package chapter04

import "fmt"

var SampleGraph = map[string][]string{
	"A": {"B", "C"},
	"B": {"D"},
	"C": {"D", "E"},
	"D": {"B", "C"},
	"E": {"C", "F"},
	"F": {"E", "O", "I", "G"},
	"G": {"F", "H"},
	"H": {"G"},
	"I": {"F", "J"},
	"O": {"F"},
	"J": {"K", "L", "I"},
	"K": {"J"},
	"L": {"J"},
	"P": {"Q", "R"},
	"Q": {"P", "R"},
	"R": {"P", "Q"},
}

func validateGraph(graph map[string][]string, start, end string) error {
	if graph == nil {
		return fmt.Errorf("graph is nil")
	}
	if _, ok := graph[start]; !ok {
		return fmt.Errorf("start node '%s' not in graph", start)
	}
	if _, ok := graph[end]; !ok {
		return fmt.Errorf("end node '%s' not in graph", end)
	}
	return nil
}

// IsRoute reports DFS reachability on a directed graph.
// Time O(v + e), space O(v).
func IsRoute(graph map[string][]string, start, end string) (bool, error) {
	if err := validateGraph(graph, start, end); err != nil {
		return false, err
	}
	return isRouteDFS(graph, start, end, map[string]struct{}{}), nil
}

func isRouteDFS(graph map[string][]string, start, end string, visited map[string]struct{}) bool {
	if start == end {
		return true
	}
	for _, node := range graph[start] {
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}
		if node == end || isRouteDFS(graph, node, end, visited) {
			return true
		}
	}
	return false
}

// IsRouteBfs reports BFS reachability on a directed graph.
// Time O(v + e), space O(v).
func IsRouteBfs(graph map[string][]string, start, end string) (bool, error) {
	if err := validateGraph(graph, start, end); err != nil {
		return false, err
	}
	if start == end {
		return true, nil
	}
	visited := map[string]struct{}{start: {}}
	queue := []string{start}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		for _, adjacent := range graph[node] {
			if _, ok := visited[adjacent]; ok {
				continue
			}
			visited[adjacent] = struct{}{}
			if adjacent == end {
				return true, nil
			}
			queue = append(queue, adjacent)
		}
	}
	return false, nil
}

// IsRouteBidirectional reports bidirectional BFS reachability.
// Time O(v + e), space O(v).
func IsRouteBidirectional(graph map[string][]string, start, end string) (bool, error) {
	if err := validateGraph(graph, start, end); err != nil {
		return false, err
	}
	if start == end {
		return true, nil
	}
	reverse := buildReverse(graph)
	visitedStart := map[string]struct{}{start: {}}
	visitedEnd := map[string]struct{}{end: {}}
	queueStart := []string{start}
	queueEnd := []string{end}
	for len(queueStart) > 0 && len(queueEnd) > 0 {
		if expandRoute(graph, &queueStart, visitedStart, visitedEnd) {
			return true, nil
		}
		if expandRoute(reverse, &queueEnd, visitedEnd, visitedStart) {
			return true, nil
		}
	}
	return false, nil
}

func expandRoute(adjacency map[string][]string, queue *[]string, visited, otherVisited map[string]struct{}) bool {
	node := (*queue)[0]
	*queue = (*queue)[1:]
	neighbors, ok := adjacency[node]
	if !ok {
		return false
	}
	for _, neighbor := range neighbors {
		if _, seen := otherVisited[neighbor]; seen {
			return true
		}
		if _, seen := visited[neighbor]; !seen {
			visited[neighbor] = struct{}{}
			*queue = append(*queue, neighbor)
		}
	}
	return false
}

func buildReverse(graph map[string][]string) map[string][]string {
	reverse := make(map[string][]string)
	for node, neighbors := range graph {
		if _, ok := reverse[node]; !ok {
			reverse[node] = nil
		}
		for _, neighbor := range neighbors {
			reverse[neighbor] = append(reverse[neighbor], node)
		}
	}
	return reverse
}

func RunQ401() {
	queries := [][2]string{{"A", "L"}, {"A", "A"}, {"Q", "G"}, {"R", "A"}}
	for _, q := range queries {
		a, _ := IsRoute(SampleGraph, q[0], q[1])
		b, _ := IsRouteBfs(SampleGraph, q[0], q[1])
		c, _ := IsRouteBidirectional(SampleGraph, q[0], q[1])
		fmt.Printf("%s->%s: %t / %t / %t\n", q[0], q[1], a, b, c)
	}
}
