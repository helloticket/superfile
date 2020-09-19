package helper

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCompareLess(t *testing.T) {
	assert.True(t, Less(1, 2))
	assert.False(t, Less(2, 2))
	assert.False(t, Less(3, 2))
	assert.True(t, Less("a", "aa"))
	assert.False(t, Less("aa", "aa"))
	assert.False(t, Less("aaa", "aa"))
	assert.False(t, Less(time.Now(), time.Now()))
	assert.True(t, Less(int64(1), int64(2)))
	assert.False(t, Less(int64(2), int64(2)))
	assert.False(t, Less(int64(3), int64(2)))
}
