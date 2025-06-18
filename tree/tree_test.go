package tree

import (
	"fmt"
	"strings"
	"testing"
)

func TestCreateRoot(t *testing.T) {
	rootFolder := createRootFolder()
	if rootFolder.Name() != "./" {
		t.Fail()
	}
}

func TestAddChild(t *testing.T) {
	t.Skip()
	rootFolder := createRootFolder()

	c, err := rootFolder.GetChildren()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("Amount: %d\n", len(c))

	newFolder, err := NewFolderNode("test", rootFolder)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	rootFolder.addChildren(newFolder)

	c, err = rootFolder.GetChildren()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Printf("Successfully added a child, total amount: %d\n", len(c))

	newFile, err := NewFileNode("tf", rootFolder)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	rootFolder.addChildren(newFile)

	c, err = rootFolder.GetChildren()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	newFolder2, err := NewFolderNode("subfolder", newFolder)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	newFolder.addChildren(newFolder2)

	newFolder2, err = NewFolderNode("chance", newFolder)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	newFolder.addChildren(newFolder2)

	StructuredPrint(rootFolder, 0)
	fmt.Printf("Successfully added a child, total amount: %d\n", len(c))
}

func CreateDefaultTree() *Tree {
	rootFolder := createRootFolder()

	newFolder, err := NewFolderNode("test", rootFolder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rootFolder.addChildren(newFolder)

	newFile, err := NewFileNode("tf.txt", rootFolder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	rootFolder.addChildren(newFile)

	newFolder2, err := NewFolderNode("subfolder", newFolder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	newFolder.addChildren(newFolder2)

	newFolder2, err = NewFolderNode("chance", newFolder)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	newFolder.addChildren(newFolder2)

	return &Tree{
		root: *rootFolder,
	}
}

func TestStructuredPrintOnTree(t *testing.T) {
	tree := CreateDefaultTree()
	StructuredPrint(tree.Root(), 0)
}

func TestFollowPath(t *testing.T) {
	tree := CreateDefaultTree()
	path := "./tf/show"
	spath := strings.Split(path, "/")
	fmt.Println(spath, len(spath))
	res, err := tree.followPath(spath, nil)
	fmt.Println(res)
	fmt.Println(err)
}
