package chapter07

import (
	"fmt"
	"strings"
)

type FileSystemNode struct {
	Name     string
	IsDir    bool
	Children map[string]*FileSystemNode
	Content  string
}

func NewFileSystemNode(name string, isDir ...bool) *FileSystemNode {
	dir := true
	if len(isDir) > 0 {
		dir = isDir[0]
	}
	if name == "" || strings.Contains(name, "/") {
		panic("invalid name")
	}
	return &FileSystemNode{Name: name, IsDir: dir, Children: map[string]*FileSystemNode{}}
}

func (n *FileSystemNode) AddChild(node *FileSystemNode) {
	if node == nil {
		panic("node is nil")
	}
	if !n.IsDir {
		panic("not a directory")
	}
	n.Children[node.Name] = node
}

func (n *FileSystemNode) Resolve(path string) *FileSystemNode {
	if path == "" || path == "/" {
		return n
	}
	node := n
	for _, part := range splitPath(path) {
		child, ok := node.Children[part]
		if !node.IsDir || !ok {
			panic("path not found")
		}
		node = child
	}
	return node
}

func Mkdir(root *FileSystemNode, path string) {
	if root == nil {
		panic("root is nil")
	}
	node := root
	for _, part := range splitPath(path) {
		if _, ok := node.Children[part]; !ok {
			node.AddChild(NewFileSystemNode(part))
		}
		child := node.Children[part]
		if !child.IsDir {
			panic("not a directory")
		}
		node = child
	}
}

func WriteFile(root *FileSystemNode, path, content string) {
	if root == nil {
		panic("root is nil")
	}
	parts := splitPath(path)
	if len(parts) == 0 {
		panic("invalid path")
	}
	node := root
	for i := 0; i < len(parts)-1; i++ {
		part := parts[i]
		if _, ok := node.Children[part]; !ok {
			node.AddChild(NewFileSystemNode(part))
		}
		node = node.Children[part]
		if !node.IsDir {
			panic("not a directory")
		}
	}
	filename := parts[len(parts)-1]
	if existing, ok := node.Children[filename]; ok && existing.IsDir {
		panic("is a directory")
	}
	fileNode := NewFileSystemNode(filename, false)
	fileNode.Content = content
	node.AddChild(fileNode)
}

func splitPath(path string) []string {
	raw := strings.Split(path, "/")
	parts := make([]string, 0, len(raw))
	for _, p := range raw {
		if p != "" {
			parts = append(parts, p)
		}
	}
	return parts
}

func RunQ711() {
	root := NewFileSystemNode("root")
	Mkdir(root, "/home/user")
	WriteFile(root, "/home/user/readme.txt", "hi")
	fmt.Println(root.Resolve("/home/user/readme.txt").Content)
}
