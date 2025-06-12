package tree

type Node interface {
	IsFile() bool 
	IsFolder() bool
	
	GetName() string
	GetParent() *FolderNode

	AsFile() *FileNode
	AsFolder() *FolderNode
}

type FolderNode struct {
	name string
	parent *FolderNode
	children []Node
}

type FileNode struct {
	name string
	extension string
	parent *FolderNode 
}

func (fn *FolderNode) IsFile { return false }

