package cast_anything

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestToBool(t *testing.T) {

	// 字符串
	r, err := ToBoolE("false")
	assert.Nil(t, err)
	assert.False(t, r)
	r, err = ToBoolE("")
	assert.Nil(t, err)
	assert.False(t, r)

	// 整数
	r, err = ToBoolE(1)
	assert.Nil(t, err)
	assert.True(t, r)
	r, err = ToBoolE(0)
	assert.Nil(t, err)
	assert.False(t, r)

	// 浮点数
	r, err = ToBoolE(0.0001)
	assert.Nil(t, err)
	assert.True(t, r)
	r, err = ToBoolE(float64(0))
	assert.Nil(t, err)
	assert.False(t, r)

	// slice
	r, err = ToBoolE([]int{})
	assert.Nil(t, err)
	assert.False(t, r)
	r, err = ToBoolE([]int{1, 2})
	assert.Nil(t, err)
	assert.True(t, r)

	// map
	r, err = ToBoolE(map[string]struct{}{})
	assert.Nil(t, err)
	assert.False(t, r)
	r, err = ToBoolE(map[string]struct{}{"1": {}})
	assert.Nil(t, err)
	assert.True(t, r)

	// struct
	type foo struct{}
	var f *foo
	r, err = ToBoolE(f)
	assert.Nil(t, err)
	assert.False(t, r)
	f = &foo{}
	r, err = ToBoolE(f)
	assert.Nil(t, err)
	assert.True(t, r)

	// 指针
	r, err = ToBoolE(reflect.ValueOf(f).Pointer())
	assert.Nil(t, err)
	assert.True(t, r)

	// 函数
	var f2 func()
	r, err = ToBoolE(f2)
	assert.Nil(t, err)
	assert.False(t, r)
	f2 = func() {}
	r, err = ToBoolE(f2)
	assert.Nil(t, err)
	assert.True(t, r)

}

func TestToBoolE(t *testing.T) {

}

func TestToBoolOrDefault(t *testing.T) {

}
