package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
	// feel free to add fields as you see fit
}

type Node struct {
	ID       int
	Children []*Node
	// feel free to add fields as you see fit
}

func Build(records []Record) (*Node, error) {
	nodes := make(map[int]*Node)
	processedNodes := make(map[int]bool)

	for _, record := range records {
		if record.ID < 0 || record.ID >= len(records) {
			return nil, errors.New("off of limits")
		}

		if record.ID != 0 && (record.ID == record.Parent || record.Parent > record.ID) {
			return nil, errors.New("wrong parent")
		}

		if _, ok := processedNodes[record.ID]; ok {
			return nil, errors.New("already processed")
		}
		processedNodes[record.ID] = true

		newNodeToInclude := &Node{ID: record.ID}
		if node, ok := nodes[record.ID]; ok {
			newNodeToInclude = node
		} else {
			nodes[record.ID] = newNodeToInclude
		}

		if record.ID == 0 {
			if record.Parent != 0 {
				return nil, errors.New("root with parent")
			}

			if _, ok := nodes[record.ID]; !ok {
				nodes[record.ID] = newNodeToInclude
			}
			continue
		}

		if node, ok := nodes[record.Parent]; ok {
			nodes := append(node.Children, newNodeToInclude)
			sort.Slice(nodes, func(i, j int) bool { return nodes[i].ID < nodes[j].ID })
			node.Children = nodes
		} else {
			nodes[record.Parent] = &Node{
				ID:       record.Parent,
				Children: []*Node{newNodeToInclude},
			}
		}
	}

	for i := 0; i < len(records); i++ {
		if _, ok := processedNodes[i]; !ok {
			return nil, errors.New("no nodes for index")
		}
	}

	return nodes[0], nil
}
