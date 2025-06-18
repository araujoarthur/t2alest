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

func (t *Tree) CreateFile(path string, name string) (*FileNode, error)                     {}
func (t *Tree) CreateFolder(path string, name string, recursive bool) (*FolderNode, error) {}
func (t *Tree) RemoveFile(path string) error                                               { return nil }
func (t *Tree) RemoveFolder(path string, recursive bool) error                             { return nil }
func (t *Tree) SearchAll(str string) []Node                                                {}
func (t *Tree) SearchFile(str string) []FileNode                                           {}
func (t *Tree) SearchFolder(str string) []FolderNode                                       {}
