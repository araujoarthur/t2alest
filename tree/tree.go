package tree

import (
	"fmt"
	"strings"
)

type Node interface {
	IsFile() bool 
	IsFolder() bool
	
	Name() string
	Parent() *FolderNode

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
func (fn *FolderNode) Name() string { return fn.name }
func (fn *FolderNode) Parent() *FolderNode { return fn.parent }
func (fn *FolderNode) AsFile() (*FileNode, error) { return nil, ETIFolderAsFile }
func (fn *FolderNode) AsFolder() (*FolderNode, error) { return fn, nil }

// Node interface implementation for FileNode
func (fn *FileNode) IsFile() bool { return true }
func (fn *FileNode) IsFolder() bool { return false }
func (fn *FileNode) Name() string { return fn.name }
func (fn *FileNode) Parent() *FolderNode { return fn.parent }
func (fn *FileNode) AsFile() (*FileNode, error) { return fn, nil }
func (fn *FileNode) AsFolder() (*FolderNode, error) { return nil, ETIFileAsFolder}

// General functions
func ValidateNodeName(name string) error {
	if strings.ContainsRune(name, '\\') || strings.ContainsRune(name, '/') {
		return ETINameNotValid
	}

	return nil
}

/***********************/
/* FolderNode methods */
/*********************/

/* PRIVATE */

/*
	Inserts a new child node into children field.
*/
func (fn *FolderNode) addChildren(n Node) {
	fn.children = append(fn.children, n)
}

/* PUBLISHED */ 

// Constructor

/*
	Builds a new folder node with a parent and without children. By design the children slice is always non-null (but can be empty)
*/
func NewFolderNode(name string, parent *FolderNode) (FolderNode, error) {
	if err := ValidateNodeName(name); err != nil {
		return FolderNode{}, err
	}

	return FolderNode{
		name: name,
		parent: parent,
		children: []Node{},
	}, nil
}

/*
	Returns true if the current node has a parent, false otherwise.
*/
func (fn *FolderNode) HasParent() bool {
	return fn.parent != nil
}

/*
	Returns true if the current node has at least one children, false otherwise.
*/
func (fn *FolderNode) HasChildren() bool {
	return (fn.children != nil) && (len(fn.children) > 0)
}

/*
	Returns a copy of the children slice of the current node.
	It has an error handling armor but at the moment it cannot possibly fail.

	One edge case happens if the folder's children field is NIL. It isn't supposed to be, so it will trigger a panic.
*/
func (fn *FolderNode) GetChildren() ([]Node, error) {
	if fn.children == nil { panic("something is not right, children shouldn't be nil") }

	return append([]Node(nil), fn.children...), nil
}

/*
	Returns a slice containing all children that are of Folder type. It will fail with an error if the current Node has no Folder children.
*/
func (fn *FolderNode) GetFolderChildren() ([]FolderNode, error) {
	var resp []FolderNode = make([]FolderNode, 0)
	
	for _, item := range fn.children {
		folder, err := item.AsFolder()
		if err == nil {
			resp = append(resp, *folder)
		}
	}
	
	if !(len(resp) > 0) {
		return nil, ETINoFolderChildren
	}

	return resp, nil
}

/*
	Returns a slice containing all children that are of File type.  It will fail with an error if the current Node has no File children.
*/
func (fn *FolderNode) GetFileChildren() ([]FileNode, error) {
	var resp []FileNode = make([]FileNode, 0)

	for _, item := range fn.children {
		file, err := item.AsFile()
		if err == nil {
			resp = append(resp, *file)
		}
	}

	if !(len(resp) > 0) {
		return nil, ETINoFileChildren
	}

	return resp, nil
}


/*
	Adds a new Folder as children of the current folder.
*/
func (fn *FolderNode) InsertFolder(name string) (*FolderNode, error) {
	newFolder, err := NewFolderNode(name, fn)
	if err != nil {
		return nil, err;
	}
	
	fn.addChildren(&newFolder)

	return &newFolder, nil
}

/*
	Adds a new File as children of the current folder.
*/
func (fn *FolderNode) InsertFile(name string, extension string) (*FileNode, error) {
	newFile, err := NewFileNode(name, extension, fn)

	if err != nil {
		return nil, err
	}
	
	fn.addChildren(&newFile)

	return &newFile, nil
}

/********************/
/* FileNode methods */
/********************/

/* PRIVATE */

/* PUBLISHED */

// Constructor

func NewFileNode(name string, extension string, parent *FolderNode) (FileNode, error) {
	if parent == nil {
		return FileNode{}, ETIDanglingFile
	}

	if err := ValidateNodeName(name); err != nil {
		return FileNode{}, ETINameNotValid
	}
	
	return FileNode{
		name: name,
		extension: extension,
		parent: parent,
	}, nil
}


