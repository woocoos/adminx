package snowflake_test

import (
	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSnowflake(t *testing.T) {
	snowflake.NodeBits = 1
	snowflake.StepBits = 8
	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Errorf("snowflake.NewNode(1) error: %v", err)
	}
	id := node.Generate()
	t.Log(id.Int64())
	assert.IsType(t, int64(0), id.Int64())
}
