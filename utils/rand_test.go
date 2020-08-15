package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRand(t *testing.T) {
	ast := assert.New(t)

	billNo := GenerateBillNo()
	ast.Equal(32, len(billNo))
}
