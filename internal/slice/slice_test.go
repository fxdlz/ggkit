package slice

import (
	"fmt"
	"log"
	"testing"
)

func TestSliceDelete(t *testing.T) {
	var sliceInt = []int{1, 2, 3, 4, 5}
	sliceInt, target, err := Delete(sliceInt, 1)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("val:%v,len:%v,cap:%v,target:%v\n", sliceInt, len(sliceInt), cap(sliceInt), target)

}
