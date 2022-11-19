package cast_anything

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToUnsignedE(t *testing.T) {
	e, err := ToUnsignedE[uint]("111")
	assert.Nil(t, err)
	assert.Equal(t, uint(111), e)
}
