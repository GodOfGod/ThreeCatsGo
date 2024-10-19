package tools

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

func GenerateId() snowflake.ID  {
	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		panic("Generate id failed")
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id
}