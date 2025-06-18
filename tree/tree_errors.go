package tree

import (
	"fmt"
)

var (
	ETINameNotValid            = TIErrorNew(1, "the node's name is invalid")
	ETIFolderAsFile            = TIErrorNew(2, "cannot handle a folder as a file")
	ETIFileAsFolder            = TIErrorNew(3, "cannot handle a file as a folder")
	ETINoChildren              = TIErrorNew(4, "the selected folder has no children")
	ETINoFolderChildren        = TIErrorNew(5, "the node is a valid folder but has no children of type Folder")
	ETINoFileChildren          = TIErrorNew(6, "the node is a valid folder but has no children of type File")
	ETIDanglingFile            = TIErrorNew(7, "a file needs a parent")
	ETIChildNotFound           = TIErrorNew(8, "this specific child does not exist")
	ETIExpectedFileFoundFolder = TIErrorNew(9, "expected a file node, found a folder node")
	ETIExpectedFolderFoundFile = TIErrorNew(10, "expected a folder node, found a file node")
	ETIUnableToFollow          = TIErrorNew(11, "unable to follow path")
)

type ETreeIntrinsic struct {
	Code    int32
	Message string
}

func (e *ETreeIntrinsic) Error() string {
	return fmt.Sprintf("error(%d): %s", e.Code, e.Message)
}

// Creates the error type
func TIErrorNew(code int32, msg string) *ETreeIntrinsic {
	return &ETreeIntrinsic{Code: code, Message: msg}
}
