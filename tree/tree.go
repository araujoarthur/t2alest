/*
The tree package implements everything needed under the hood as well as the user interface to operate a file tree.

It currently provides:
- the Node interface
- the FolderNode and FileNode implementations of the Node interface
- Tree editing and traversal functions

Usage:

t := tree.CreateTree()
*/
package tree

import (
	"fmt"
	"path/filepath"
	"strings"
)

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
does not exist or is a file (except the last, which can bea file). The received path must be sanitized to remove the trailing bar
*/
func (t *Tree) followPath(path []string, current_node Node) (Node, error) {
	if current_node == nil {
		current_node = t.Root()
		if path[0] == "." {
			path = path[1:]
		}
	}

	if len(path) == 0 {
		return current_node, nil
	}

	if current_node.IsFile() {
		return nil, ETIUnableToFollow
	}

	folder, err := current_node.AsFolder()
	if err != nil {
		return nil, err
	}

	if !folder.HasChildren() {
		return nil, ETIUnableToFollow
	}

	children, err := folder.GetChildren()
	if err != nil {
		return nil, err
	}

	evaluatedStep := path[0]
	nextSteps := path[1:]
	for _, child := range children {
		if child.CleanName() == strings.TrimSuffix(evaluatedStep, "/") {
			return t.followPath(nextSteps, child)
		}
	}

	return nil, ETIPathNotFound
}

/* Interface to the internal followPath funciton */
func (t *Tree) FollowPath(path string) (Node, error) {
	path = filepath.ToSlash(path)
	path = strings.TrimSuffix(path, "/")

	separatePath := strings.SplitAfter(path, "/")
	return t.followPath(separatePath, nil)
}

/* Interface to the internal explorePath function */
func (t *Tree) ExplorePath(path string) (Node, []string, error) {
	path = filepath.ToSlash(path)
	path = strings.TrimSuffix(path, "/")

	separatePath := strings.SplitAfter(path, "/")
	return t.explorePath(separatePath, nil)
}

/*
[UT] Follows a path up to the point it's not possible anymore (i.e the rest of the path doesn't exist). Returns the most distant existent element
in the path and the portion of the path that is missing.
*/
func (t *Tree) explorePath(path []string, current_node Node) (Node, []string, error) {
	if current_node == nil {
		current_node = t.Root()
		if path[0] == "." {
			path = path[1:]
		}
	}

	if len(path) == 0 { // All locations exist
		return current_node, path, nil
	}

	if current_node.IsFile() {
		return nil, nil, ETIUnableToFollow
	}

	folder, err := current_node.AsFolder()
	if err != nil {
		return nil, nil, ETIUnableToFollow
	}

	if !folder.HasChildren() {
		return current_node, path, nil
	}

	children, err := folder.GetChildren()
	if err != nil {
		return nil, nil, err
	}

	evaluatedStep := strings.TrimSuffix(path[0], "/")
	nextSteps := path[1:]

	for _, child := range children {
		fmt.Printf("CleanName: %s | Evaluated step: %s\n", child.CleanName(), evaluatedStep)
		if child.CleanName() == evaluatedStep {
			return t.explorePath(nextSteps, child)
		}
	}

	return current_node, nextSteps, nil

}

/*
Creates a file node at the given path.
*/
func (t *Tree) CreateFile(path string, name string) (*FileNode, error) {
	path = filepath.ToSlash(path)
	path = strings.TrimSuffix(path, "/")
	pathSeparated := strings.Split(path, "/")

	node, err := t.followPath(pathSeparated, nil)
	if err != nil {
		return nil, err
	}

	if node.IsFile() {
		return nil, ETIExpectedFolderFoundFile
	}

	fnode, err := node.AsFolder()
	if err != nil {
		return nil, err
	}

	created, err := fnode.InsertFile(name)
	if err != nil {
		return nil, err
	}

	return created, nil
}

/*
Creates a folder at a given path. If recursive is false, the function will fail if any of the path's folders but the last does not exist.
*/
func (t *Tree) CreateFolder(path string, name string, recursive bool) (*FolderNode, error) {
	path = filepath.ToSlash(path)

	var createAt *FolderNode
	if !recursive {
		final, err := t.FollowPath(path)
		if err != nil {
			return nil, err
		}

		createAt, err = final.AsFolder()
		if err != nil {
			return nil, err
		}
	} else {
		furthestNode, pathLeft, err := t.ExplorePath(path)
		if err != nil {
			return nil, err
		}

		fmt.Printf("furthestNode: %s\n\npathLeft: %s\n", furthestNode.Name(), pathLeft)
		if furthestNode.IsFile() {
			return nil, ETIExpectedFolderFoundFile
		}

		furthestFolder, err := furthestNode.AsFolder()
		if err != nil {
			return nil, err
		}

		currentFolder := furthestFolder
		for len(pathLeft) > 0 {

			creatingNow := strings.TrimSuffix(pathLeft[0], "/")
			fmt.Printf("Creating now: %s, pathleft: %s\n\n", creatingNow, pathLeft)
			pathLeft = pathLeft[1:]
			currentFolder, err = currentFolder.InsertFolder(creatingNow)

			if err != nil {
				return nil, err
			}

		}

		createAt = currentFolder
	}

	return createAt.InsertFolder(name)
}

func (t *Tree) RemoveFile(path string) error {
	node, err := t.FollowPath(path)
	if err != nil {
		return err
	}

	fil, err := node.AsFile()
	if err != nil {
		return err
	}

	base := filepath.Base(path)
	filParent := fil.Parent()
	filParent.RemoveNode(base)
	return nil
}

func (t *Tree) RemoveFolder(path string, recursive bool) error {
	node, err := t.FollowPath(path)
	if err != nil {
		return err
	}

	folder, err := node.AsFolder()
	if err != nil {
		return err
	}

	if folder.HasChildren() && !recursive {
		return ETICannotRemoveParent
	}

	if folder == t.Root() {
		return ETICannotRemoveRoot
	}

	base := filepath.Base(path)
	folderParent := folder.Parent()

	if err := folderParent.RemoveNode(base); err != nil {
		return err
	}

	return nil
}

func (t *Tree) EvaluateNodePath(node Node) string {
	currNode := node
	currPath := currNode.Name()

	currNode = currNode.Parent()
	for currNode != nil {
		folder, ok := currNode.(*FolderNode)
		if !ok || folder == nil {
			break
		}

		currPath = currNode.Name() + currPath
		currNode = currNode.Parent()
	}

	return currPath
}

func (t *Tree) SearchAll(str string) ([]Node, error) {
	res, err := t.Root().DFS(str)
	if err != nil {
		return nil, err
	}

	return res, nil
}

//func (t *Tree) SearchFile(str string) []FileNode     { return nil }
//func (t *Tree) SearchFolder(str string) []FolderNode { return nil }
