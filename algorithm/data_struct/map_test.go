package data_struct

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMap(t *testing.T) {
	// 空结构体：内存地址一样，并且不占用任何内存空间
	a := struct{}{}
	b := struct{}{}
	if a == b {
		fmt.Printf("right %p \n", &a)
	}
	fmt.Println(unsafe.Sizeof(a))
}
