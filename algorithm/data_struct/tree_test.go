package data_struct

import (
	"fmt"
	"testing"
)

func TestTree(t *testing.T) {
	root := InitTree()
	fmt.Println(PreOrder(root))
	fmt.Println(MidOrder(root))
	fmt.Println(AfterOrder(root))
	fmt.Println(LayerOrder(root))
}
