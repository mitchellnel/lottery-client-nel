package main

import (
	"context"
	"fmt"

	// general purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosclient"
	// types from our blockchain blog
	"github.com/mitchellnel/lottery-chain-nel/x/lottery/types"
)

func main() {
	fmt.Println("Hello, World!")

	_ = &types.MsgStartLottery{
		Creator: "steve",
	}
}

func setupClient() {
	// create an instance of cosmosclient
	cosmos, err := cosmosclient.New(
		context.Background(),
		cosmosclient.WithHome("/Users/user/.lottery-chain-nel"),
	)

    _, _ = cosmos, err
}
