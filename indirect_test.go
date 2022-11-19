package cast_anything

import (
	"fmt"
	"testing"
)

func Test_indirect(t *testing.T) {

	v1 := "this is some string content"
	v2 := &v1
	v3 := &v2
	v4 := &v3
	v5 := &v4
	v6 := &v5
	a := indirect(v6)
	fmt.Println(fmt.Sprintf("%T, %v", a, a))

	v3 = nil
	b := indirect(v6)
	fmt.Println(fmt.Sprintf("%T, %v", b, b))
	t.Log(b == nil)

}
