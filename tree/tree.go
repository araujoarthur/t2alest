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
	dbgh("Started inner following, current path: ", fmt.Sprintf("%s", path))
	if current_node == nil {
		current_node = t.Root()
		if path[0] == "." {
			path = path[1:]
		}
	}

	if len(path) == 0 {
		dbgh("Len of path was zero", fmt.Sprintf("returning the current node with name %s and children", current_node.Name()))
		return current_node, nil
	}

	if current_node.IsFile() {
		dbgh("Node was a file, but path len was not zero", fmt.Sprintf("node name: %s", current_node.Name()))
		return nil, ETIUnableToFollow
	}

	folder, err := current_node.AsFolder()
	if err != nil {
		dbgh("Error trying to cast the node as folder", fmt.Sprintf("node name: %s", current_node.Name()))
		return nil, err
	}

	if !folder.HasChildren() {
		dbgh("Node was a file, but path len was not zero", fmt.Sprintf("node name: %s, children: %s", folder.Name(), folder.children))
		return nil, ETIUnableToFollow
	}

	children, err := folder.GetChildren()
	if err != nil {
		dbgh("Error trying to get children of folder", fmt.Sprintf("node name: %s, error: %s", folder.Name(), err))
		return nil, err
	}

	evaluatedStep := path[0]
	nextSteps := path[1:]
	dbgh("Reached the loop phase", fmt.Sprintf("with evaluated step '%s' and next steps '%s'", evaluatedStep, nextSteps))
	for _, child := range children {
		dbgh("Inside loop phase", fmt.Sprintf("with evaluated step '%s' and child name '%s'", evaluatedStep, child.CleanName()))
		if child.CleanName() == strings.TrimSuffix(evaluatedStep, "/") {
			return t.followPath(nextSteps, child)
		}
	}
	return nil, ETIPathNotFound
}

/* Interface to the internal followPath funciton */
func (t *Tree) FollowPath(path string) (Node, error) {
	dbgh("Started Following a Path:", path)
	path = filepath.ToSlash(path)
	path = strings.TrimSuffix(path, "/")

	dbgh("Started Following a Path, After All Trims:", path)
	separatePath := strings.SplitAfter(path, "/")
	dbgh("Started Following a Path, Sparated One:", fmt.Sprintf("%s", separatePath))
	return t.followPath(separatePath, nil)
}

/* Interface to the internal explorePath function */
func (t *Tree) ExplorePath(path string) (Node, []string, error) {
	dbgh("Started Exploring a Path:", path)
	path = filepath.ToSlash(path)
	path = strings.TrimSuffix(path, "/")

	dbgh("Started Exploring a Path, After All Trims:", path)
	separatePath := strings.SplitAfter(path, "/")
	dbgh("Started Exploring a Path, Sparated One:", fmt.Sprintf("%s", separatePath))
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

	evaluatedStep := path[0]
	nextSteps := path[1:]

	for _, child := range children {
		if child.CleanName() == evaluatedStep {
			return t.explorePath(nextSteps, child)
		}
	}

	return current_node, path, nil

}

/*
C[UT] reates a file node at the given path.
*/
func (t *Tree) CreateFile(path string, name string) (*FileNode, error) {
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
[UT] Creates a folder at a given path. If recursive is false, the function will fail if any of the path's folders but the last does not exist.
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

		if furthestNode.IsFile() {
			return nil, ETIExpectedFolderFoundFile
		}

		furthestFolder, err := furthestNode.AsFolder()
		if err != nil {
			return nil, err
		}

		currentFolder := furthestFolder
		dbgh("Creating at recursive mode", "starting loop")
		for len(pathLeft) > 0 {
			dbgh("Creating at recursive mode", fmt.Sprintf("Path Left Len: %d", len(pathLeft)))
			creatingNow := strings.TrimSuffix(pathLeft[0], "/")
			dbgh("Creating at recursive mode, creating", fmt.Sprintf("Creating now: %s", creatingNow))
			pathLeft = pathLeft[1:]
			dbgh("Creating at recursive mode, creating", fmt.Sprintf("Path Left: %s", pathLeft))
			currentFolder, err = currentFolder.InsertFolder(creatingNow)

			if err != nil {
				return nil, err
			}

		}

		createAt = currentFolder
	}

	return createAt.InsertFolder(name)
}

func (t *Tree) RemoveFile(path string) error                   { return nil }
func (t *Tree) RemoveFolder(path string, recursive bool) error { return nil }
func (t *Tree) SearchAll(str string) []Node                    { return nil }
func (t *Tree) SearchFile(str string) []FileNode               { return nil }
func (t *Tree) SearchFolder(str string) []FolderNode           { return nil }

func dbgh(title string, msg string) {
	fmt.Printf("[DEBUG] %s: %s\n", title, msg)
}
