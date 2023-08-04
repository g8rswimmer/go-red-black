package rb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func equalNode(t *testing.T, expected *Node, actual *Node) {
	switch {
	case expected == nil && actual == nil:
		return
	case actual == nil:
		t.Errorf("expected node is persent actual [nil] %v", expected)
		return
	case expected == nil:
		t.Errorf("actual node is persent expected [nil] %v", actual)
		return
	default:
	}
	assert.Equal(t, expected.Pair, actual.Pair)
	assert.Equalf(t, expected.black, actual.black, "%+v :: %+v", expected, actual)
	equalNode(t, expected.left, actual.left)
	equalNode(t, expected.right, actual.right)
}
