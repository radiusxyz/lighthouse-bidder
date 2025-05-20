package anvilclient

import (
	"github.com/ethereum/go-ethereum/rpc"
	"os/exec"
)

type Anvil struct {
	rpcClient *rpc.Client
}

func New(rpcHttpUrl string) (*Anvil, error) {
	if err := startAnvilFork(rpcHttpUrl); err != nil {
		return nil, err
	}

	rpcClient, err := rpc.Dial("http://localhost:8547")
	if err != nil {
		panic("failed to connect to anvilclient: " + err.Error())
	}

	return &Anvil{
		rpcClient: rpcClient,
	}, nil
}

func startAnvilFork(rpcHttpUrl string) error {
	cmd := exec.Command("anvil",
		"--fork-url", rpcHttpUrl,
		"--port", "8547",
	)
	return cmd.Start()
}
