package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"github.com/tsingsun/woocoo/pkg/conf"
)

var (
	defaultNode *snowflake.Node
)

func SetDefaultNode(cnf *conf.AppConfiguration) {
	if nb := cnf.Int("snowflake.nodeBits"); nb > 0 {
		snowflake.NodeBits = uint8(nb)
	}
	if sb := cnf.Int("snowflake.stepBits"); sb > 0 {
		snowflake.StepBits = uint8(sb)
	}
	nid := cnf.Int("snowflake.nodeID")
	if nid <= 0 {
		nid = 1
	}
	node, err := snowflake.NewNode(int64(nid))
	if err != nil {
		panic(err)
	}
	defaultNode = node
}

func New() snowflake.ID {
	return defaultNode.Generate()
}
