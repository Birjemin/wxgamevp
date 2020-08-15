package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQuerySortByKeyStr(t *testing.T) {
	ast := assert.New(t)
	ast.Equal("key1=val1&key2=val2", QuerySortByKeyStr(map[string]string{"key1": "val1", "key2": "val2"}))
}
