package main

import (
	_ "embed"
	"fmt"
	"path"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputData string

const (
	sizeLimit           = 100000
	totalAvailableSpace = 70000000
	updateRequiredSpace = 30000000
)

type node struct {
	fullPath string
	name     string
	size     int
	isDir    bool
	children []*node
	parent   *node
}

func (n *node) getOrCreateChildDir(name string) *node {
	for _, c := range n.children {
		if c.name == name && c.isDir == true {
			return c
		}
	}
	// Child doesn't exist
	childNode := &node{
		fullPath: path.Join(n.fullPath, name),
		name:     name,
		isDir:    true,
		parent:   n,
	}
	n.children = append(n.children, childNode)
	return childNode
}

func (n *node) getOrCreateChildFile(name string, fsize int) *node {
	for _, c := range n.children {
		// already exists
		if c.name == name && c.isDir == false {
			return c
		}
	}
	// child doesn't exist
	childNode := &node{
		fullPath: path.Join(n.fullPath, name),
		name:     name,
		size:     fsize,
		parent:   n,
	}
	n.children = append(n.children, childNode)
	return childNode
}

func main() {
	root := &node{isDir: true, fullPath: "/", name: ""}
	currNode := root
	// Build the tree from the commands
	for _, line := range strings.Split(inputData, "\n") {
		parts := strings.Split(line, " ")

		if strings.HasPrefix(line, "$") {
			if parts[1] == "cd" {
				switch parts[2] {
				case "..":
					currNode = currNode.parent
				case "/":
					currNode = root
				default:
					currNode = currNode.getOrCreateChildDir(parts[2])
				}
			}
			continue
		}
		// list output
		if parts[0] == "dir" {
			currNode.getOrCreateChildDir(parts[1])
		} else {
			fsize, err := strconv.Atoi(parts[0])
			if err != nil {
				panic(fmt.Errorf("could not convert file size to int %s: %v", parts[0], err))
			}
			currNode.getOrCreateChildFile(parts[1], fsize)
		}
	}

	// Populate the sizes
	totalUsedSpace := populateFileSizes(root)
	fmt.Println("totalUsedSpace", totalUsedSpace)

	// Find all directories with AT MOST the size limit
	// fmt.Println("answer", findDirSizeSum(root))

	spaceAvailable := totalAvailableSpace - totalUsedSpace
	spaceNeeded := updateRequiredSpace - spaceAvailable
	fmt.Println("spaceAvailable", spaceAvailable)
	fmt.Println("spaceNeeded", spaceNeeded)

	dirSizes := getDirectorySizes(root)
	sort.Ints(dirSizes)
	for _, dsize := range dirSizes {
		if dsize >= spaceNeeded {
			fmt.Println("answer", dsize)
			return
		}
	}
}

func findDirSizeSum(n *node) int {
	if n == nil {
		return 0
	}
	var sum int
	if n.isDir && n.size <= sizeLimit {
		sum = n.size
	}
	for _, c := range n.children {
		sum += findDirSizeSum(c)
	}
	return sum
}

func populateFileSizes(n *node) int {
	if n == nil {
		return 0
	}
	if len(n.children) == 0 {
		return n.size
	}
	for _, c := range n.children {
		n.size += populateFileSizes(c)
	}
	return n.size
}

func getDirectorySizes(n *node) []int {
	if n == nil {
		return nil
	}
	var dirSizes []int
	if n.isDir {
		dirSizes = append(dirSizes, n.size)
	}
	for _, c := range n.children {
		dirSizes = append(dirSizes, getDirectorySizes(c)...)
	}
	return dirSizes
}
