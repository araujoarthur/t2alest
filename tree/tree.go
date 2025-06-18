package tree

import (
	"strings"
)

type FilePath string
type FilePathSteps []string

type Tree struct {
	root FolderNode
}

/*
Creates an empty tree with the root node already set.
*/
func CreateTree() *Tree {
	return &Tree{
		root: *createRootFolder(),
	}
}

// Field Accessors
func (t *Tree) Root() *FolderNode {
	return &t.root
}

/*
This function starts a navigation from the root path up to the final path in the string, returning it if it exists or an error if anything on the path
does not exist or is a file (except the last, which can be a file).
*/
func (t *Tree) followPath(path string, current_node Node) (Node, error) {
	

	if current_node == nil {
		current_node = t.Root()
	}

	if path
}

/*
Follows a path up to the point it's not possible anymore (i.e the rest of the path doesn't exist). Returns the most distant existent element
in the path.
*/
func (t *Tree) explorePath(path string) *Node {}

/*
Creates a file node at the given path.
*/
func (t *Tree) CreateFile(path string, name string) (*FileNode, error) {}

/*
Creates a folder at a given path. If recursive is false, the function will fail if any of the path's folders but the last does not exist.
*/
func (t *Tree) CreateFolder(path string, name string, recursive bool) (*FolderNode, error) {}
func (t *Tree) RemoveFile(path string) error                                               { return nil }
func (t *Tree) RemoveFolder(path string, recursive bool) error                             { return nil }
func (t *Tree) SearchAll(str string) []Node                                                {}
func (t *Tree) SearchFile(str string) []FileNode                                           {}
func (t *Tree) SearchFolder(str string) []FolderNode                                       {}


func (t FilePath) Steps() FilePathSteps {
	return strings.Split(t, "/")
}


func (t FilePathSteps) Rejoin() FilePath {
	return strings.Join(t, "/")
}