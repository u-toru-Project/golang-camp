package chapter04

import "fmt"

// DetermineBuildOrder returns a topological order; prerequisite appears before project.
// Time O(p + d), space O(p + d).
func DetermineBuildOrder(projects []string, dependencies [][2]string) ([]string, error) {
	if projects == nil {
		return nil, fmt.Errorf("projects is nil")
	}
	if dependencies == nil {
		dependencies = [][2]string{}
	}
	dependencyTree := make(map[string]map[string]struct{}, len(projects))
	for _, project := range projects {
		dependencyTree[project] = make(map[string]struct{})
	}
	for _, dep := range dependencies {
		prerequisite, project := dep[0], dep[1]
		if _, ok := dependencyTree[project]; !ok {
			return nil, fmt.Errorf("unknown project in dependency: '%s'", project)
		}
		if _, ok := dependencyTree[prerequisite]; !ok {
			return nil, fmt.Errorf("unknown dependency: '%s'", prerequisite)
		}
		dependencyTree[project][prerequisite] = struct{}{}
	}
	buildOrder := make([]string, 0, len(projects))
	unbuilt := make(map[string]struct{}, len(projects))
	for _, project := range projects {
		unbuilt[project] = struct{}{}
	}
	for len(unbuilt) > 0 {
		somethingBuilt := false
		for _, project := range projects {
			if _, still := unbuilt[project]; !still {
				continue
			}
			blocked := false
			for prereq := range dependencyTree[project] {
				if _, pending := unbuilt[prereq]; pending {
					blocked = true
					break
				}
			}
			if blocked {
				continue
			}
			buildOrder = append(buildOrder, project)
			delete(unbuilt, project)
			somethingBuilt = true
		}
		if !somethingBuilt {
			return nil, fmt.Errorf("No valid build order exists")
		}
	}
	return buildOrder, nil
}

func RunQ407() {
	projects := []string{"a", "b", "c", "d", "e", "f", "g"}
	dependencies := [][2]string{
		{"d", "g"},
		{"a", "e"},
		{"b", "e"},
		{"c", "a"},
		{"f", "a"},
		{"b", "a"},
		{"f", "c"},
		{"f", "b"},
	}
	order, _ := DetermineBuildOrder(projects, dependencies)
	fmt.Println(order)
}
