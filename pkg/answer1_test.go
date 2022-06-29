package pkg

import (
	"fmt"
	"testing"
)

func TestAnswer1(t *testing.T) {
	arr := []int64{4,2,9,6,1,3}
	s := Run1(arr)
	fmt.Println(arr)
	fmt.Println(s)
}
