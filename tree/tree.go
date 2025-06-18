package tree

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
Follows a path and return the final node or an error if it's not possible to follow through the end.
*/
func (t *Tree) followPath(path string) (*Node, error) {}

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
