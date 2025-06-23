package tree

import (
	"fmt"
	"strings"
)

type Node interface {
	IsFile() bool
	IsFolder() bool

	Name() string
	CleanName() string
	Parent() *FolderNode

	AsFile() (*FileNode, error)
	AsFolder() (*FolderNode, error)

	fmt.Stringer
}

type FolderNode struct {
	name     string
	parent   *FolderNode
	children []Node
}

type FileNode struct {
	name   string
	parent *FolderNode
}

// Node interface implementation for FolderNode
func (fn *FolderNode) IsFile() bool                   { return false }
func (fn *FolderNode) IsFolder() bool                 { return true }
func (fn *FolderNode) Name() string                   { return fn.name }
func (fn *FolderNode) Parent() *FolderNode            { return fn.parent }
func (fn *FolderNode) AsFile() (*FileNode, error)     { return nil, ETIFolderAsFile }
func (fn *FolderNode) AsFolder() (*FolderNode, error) { return fn, nil }
func (fn *FolderNode) CleanName() string              { return fn.Name()[:len(fn.Name())-1] }

// Stringer interface implementation for FolderNode
func (fn *FolderNode) String() string {
	var childlist string = ""
	for _, c := range fn.children {
		childlist = fmt.Sprintf("%s\n%s", childlist, c)
	}

	var str string = fmt.Sprintf("[Name: %s\nChild Count: %d\n\tChildList: %s]", fn.name, len(fn.children), childlist)
	return str
}

// Node interface implementation for FileNode
func (fn *FileNode) IsFile() bool                   { return true }
func (fn *FileNode) IsFolder() bool                 { return false }
func (fn *FileNode) Name() string                   { return fn.name }
func (fn *FileNode) Parent() *FolderNode            { return fn.parent }
func (fn *FileNode) AsFile() (*FileNode, error)     { return fn, nil }
func (fn *FileNode) AsFolder() (*FolderNode, error) { return nil, ETIFileAsFolder }
func (fn *FileNode) CleanName() string              { return fn.Name() }

// Stringer interface implementation for FileNode
func (fn *FileNode) String() string {
	var str string = fmt.Sprintf("=FILE=\nName: %s\nParent:%s\n==", fn.Name(), fn.Parent().Name())
	return str
}

// General functions
func ValidateNodeName(name string) error {
	// Cannot check for starts with . bc there are valid hidden folders that start with .
	if strings.ContainsRune(name, '\\') || strings.ContainsRune(name, '/') || name == "." || name == ".." {
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
func NewFolderNode(name string, parent *FolderNode) (*FolderNode, error) {
	if err := ValidateNodeName(name); err != nil {
		return nil, err
	}

	return &FolderNode{
		name:     name + "/",
		parent:   parent,
		children: []Node{},
	}, nil
}

/*
Creates a special root folder.
*/
func createRootFolder() *FolderNode {
	return &FolderNode{
		name:     "./",
		parent:   nil,
		children: []Node{},
	}
}

/*
Checks if there's a specific child. Returns it if it exists or an error if it doesn't.
*/
func (fn *FolderNode) SearchChild(name string) (Node, error) {
	if !fn.HasChildren() {
		return nil, ETINoChildren
	}

	foldersChildren, err := fn.GetChildren()
	if err != nil {
		return nil, err
	}

	for _, child := range foldersChildren {
		if child.Name() == name {
			return child, nil
		}
	}

	return nil, ETIChildNotFound
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
	return (len(fn.children) > 0)
}

/*
Returns a copy of the children slice of the current node.
It has an error handling armor but at the moment it cannot possibly fail.

One edge case happens if the folder's children field is NIL. It isn't supposed to be, so it will trigger a panic.
*/
func (fn *FolderNode) GetChildren() ([]Node, error) {
	if fn.children == nil {
		panic("something is not right, children shouldn't be nil")
	}

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

func (fn *FolderNode) RemoveNode(name string) error {
	itemPos := -1
	for i, n := range fn.children {
		if n.CleanName() == name {
			itemPos = i
			break
		}
	}

	if itemPos == -1 {
		return ETIChildNotFound
	}

	fn.children = append(fn.children[:itemPos], fn.children[itemPos+1:]...)
	return nil
}

/*
Adds a new Folder as children of the current folder.
*/
func (fn *FolderNode) InsertFolder(name string) (*FolderNode, error) {
	for _, child := range fn.children {
		if child.CleanName() == name {
			return nil, ETIDuplicatedName
		}
	}
	newFolder, err := NewFolderNode(name, fn)
	if err != nil {
		return nil, err
	}

	fn.addChildren(newFolder)

	return newFolder, nil
}

/*
Adds a new File as children of the current folder.
*/
func (fn *FolderNode) InsertFile(name string) (*FileNode, error) {
	for _, child := range fn.children {
		if child.CleanName() == name {
			return nil, ETIDuplicatedName
		}
	}

	newFile, err := NewFileNode(name, fn)

	if err != nil {
		return nil, err
	}

	fn.addChildren(newFile)

	return newFile, nil
}

/*
Performs a depth first search on the current node's children.
*/
func (fn *FolderNode) DFS(name string) ([]Node, error) {
	var results []Node

	for _, c := range fn.children {
		if c.CleanName() == name {
			results = append(results, c)
		}

		if c.IsFolder() {
			cAF, err := c.AsFolder()
			if err != nil {
				panic("should not have a non-nil error at this point")
			}

			childResults, errDFS := cAF.DFS(name)
			if errDFS != nil {
				return results, errDFS
			}

			results = append(results, childResults...)
		}
	}

	return results, nil
}

/********************/
/* FileNode methods */
/********************/

/* PRIVATE */

/* PUBLISHED */

// Constructor

func NewFileNode(name string, parent *FolderNode) (*FileNode, error) {
	if parent == nil {
		return nil, ETIDanglingFile
	}

	if err := ValidateNodeName(name); err != nil {
		return nil, ETINameNotValid
	}

	return &FileNode{
		name:   name,
		parent: parent,
	}, nil
}

func StructuredPrint(n Node, il int) {
	var istr string = ""
	for i := 0; i <= 2*il; i++ {
		istr = istr + " "
	}

	if n.IsFile() {
		fmt.Println(istr + n.Name())
	} else {
		fmt.Println(istr + n.Name())
		folder, err := n.AsFolder()
		if err == nil && folder.HasChildren() {
			folderChildren, err := folder.GetChildren()
			if err == nil {
				for _, c := range folderChildren {
					StructuredPrint(c, il+1)
				}
			}
		}
	}
}
