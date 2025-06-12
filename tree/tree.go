package tree

import (
	"fmt"
)

type Node interface {
	IsFile() bool 
	IsFolder() bool
	
	GetName() string
	GetParent() *FolderNode

	AsFile() (*FileNode, error)
	AsFolder() (*FolderNode, error)
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

// Node interface implementation for FolderNode
func (fn *FolderNode) IsFile() bool { return false }
func (fn *FolderNode) IsFolder() bool { return true }
func (fn *FolderNode) GetName() string { return fn.name }
func (fn *FolderNode) GetParent() *FolderNode { return fn.parent }
func (fn *FolderNode) AsFile() (*FileNode, error) { return nil, fmt.Errorf("cannot return a folder as file") }
func (fn *FolderNode) AsFolder() (*FolderNode, error) { return fn, nil }

// Node interface implementation for FileNode
func (fn *FileNode) IsFile() bool { return true }
func (fn *FileNode) IsFolder() bool { return false }
func (fn *FileNode) GetName() string { return fn.name }
func (fn *FileNode) GetParent() *FolderNode { return fn.parent }
func (fn *FileNode) AsFile() (*FileNode, error) { return fn, nil }
func (fn *FileNode) AsFolder() (*FolderNode, error) { return nil, fmt.Errorf("cannot return a file as folder")}


// FolderNode methods

// FileNode methods
