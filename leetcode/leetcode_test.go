package leetcode

import (
	"strconv"
	"testing"
)

var (
	tree = treeNode.Deserialize("1,2,3,null,5")
	ln   = listNode.Deserialize("1,2,2,1")
)

func TestT(t *testing.T) {
	num, _ := strconv.ParseUint("10010110111001001101001111110101", 2, 64)
	t.Log(reverseBits(uint32(num)))
}
