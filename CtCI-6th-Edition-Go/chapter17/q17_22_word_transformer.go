package chapter17

import "fmt"

func WordTransformer(source, target string, words []string) []string {
	if len(source) != len(target) {
		panic("source and target must have the same length")
	}
	dictionary := make([]string, 0, len(words)+1)
	seen := make(map[string]struct{}, len(words)+1)
	for _, w := range words {
		if _, ok := seen[w]; !ok {
			seen[w] = struct{}{}
			dictionary = append(dictionary, w)
		}
	}
	if _, ok := seen[source]; !ok {
		dictionary = append(dictionary, source)
		seen[source] = struct{}{}
	}
	if source == target {
		return []string{source}
	}
	queue := []string{source}
	visited := map[string]struct{}{source: {}}
	parents := map[string]string{source: ""}
	for len(queue) > 0 {
		word := queue[0]
		queue = queue[1:]
		for _, neighbor := range neighbors(word, dictionary) {
			if _, ok := visited[neighbor]; ok {
				continue
			}
			visited[neighbor] = struct{}{}
			parents[neighbor] = word
			if neighbor == target {
				return reconstructPath(parents, target)
			}
			queue = append(queue, neighbor)
		}
	}
	return nil
}

func neighbors(word string, dictionary []string) []string {
	out := make([]string, 0)
	for _, candidate := range dictionary {
		if candidate == word || len(candidate) != len(word) {
			continue
		}
		unequal := 0
		for i := 0; i < len(word); i++ {
			if word[i] != candidate[i] {
				unequal++
				if unequal > 1 {
					break
				}
			}
		}
		if unequal == 1 {
			out = append(out, candidate)
		}
	}
	return out
}

func reconstructPath(parents map[string]string, end string) []string {
	path := []string{end}
	current := end
	for {
		parent, ok := parents[current]
		if !ok || parent == "" {
			break
		}
		path = append(path, parent)
		current = parent
	}
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func RunQ1722() {
	path := WordTransformer("damp", "like", []string{"damp", "lime", "limp", "lamp", "like"})
	fmt.Println(path)
}
