package main

import (
	"log"
	"sort"
	"strconv"
	"strings"
)

const diskSize = 70000000
const updateSize = 30000000

func main() {
	lines := strings.Split(input, "\n")
	root := Parse(lines)
	log.Println("Root size:", root.Size())

	freeSpace := diskSize - root.Size()
	spaceToFind := updateSize - freeSpace

	matches := root.WhereDir(func(dir *Dir) bool {
		return dir.Size() > spaceToFind
	})

	sort.SliceStable(matches, func(i, j int) bool {
		return matches[i].Size() < matches[j].Size()
	})

	log.Println("Smallest match:", matches[0].Size())
}

func Parse(input []string) *Dir {
	root := Mkdir()
	var stack []*Dir

	var currentDir *Dir
	var prefix = ""
	for _, line := range input {
		if strings.HasPrefix(line, "$ cd ") {
			// Move
			where := strings.TrimPrefix(line, "$ cd ")
			switch where {
			case "/":
				currentDir = root
				stack = append(stack, root)

			case "..":
				//log.Printf("Going up (%v)", stack)
				currentDir = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				prefix = prefix[0 : len(prefix)-2]

			default:
				//move into
				stack = append(stack, currentDir)
				newDir, _ := currentDir.dirs[where]
				currentDir = newDir
				prefix = prefix + "  "
				//log.Printf("Going into %q (%v)", where, stack)
			}

			continue
		}

		if strings.HasPrefix(line, "$ ls") {
			//log.Println("Current dir:", currentDir)
			continue
		}

		parts := strings.Split(line, " ")
		name, p := parts[1], parts[0]
		if p == "dir" {
			//log.Printf("%vD: %v", prefix, name)

			currentDir.dirs[name] = Mkdir()
		} else {

			size, _ := strconv.Atoi(p)
			//log.Printf("%vF: %v (%v)", prefix, name, size)
			currentDir.items[name] = size
		}
	}

	return root
}

func Mkdir() *Dir {
	return &Dir{
		dirs:  map[string]*Dir{},
		items: map[string]int{},
	}
}

type Dir struct {
	dirs  map[string]*Dir
	items map[string]int
}

func (d Dir) Size() int {
	total := 0
	for _, size := range d.items {
		total += size
	}
	for _, dir := range d.dirs {
		total += dir.Size()
	}
	return total
}

func (d Dir) WhereDir(fn func(dir *Dir) bool) []*Dir {
	var matches []*Dir
	for _, dir := range d.dirs {
		if fn(dir) {
			matches = append(matches, dir)
		}
		matches = append(matches, dir.WhereDir(fn)...)
	}
	return matches
}
