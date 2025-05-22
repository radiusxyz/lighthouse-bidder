package anvilclient

//
//import (
//	"github.com/radiusxyz/lighthouse-bidder/httpclient"
//	"os/exec"
//)
//
//type Anvil struct {
//	url    string
//	client *httpclient.HttpClient
//}
//
//func New(anvilUrl string, rpcHttpUrl string) (*Anvil, error) {
//	if err := executeAnvilFork(rpcHttpUrl); err != nil {
//		return nil, err
//	}
//
//	return &Anvil{
//		url:    anvilUrl,
//		client: httpclient.New(),
//	}, nil
//}
//
//func executeAnvilFork(rpcHttpUrl string) error {
//	cmd := exec.Command("anvil",
//		"--fork-url", rpcHttpUrl,
//		"--port", "8547",
//	)
//	return cmd.Start()
//}
//
//func (a *Anvil) GetBlockByNumber(blockNumber uint64) {
//
//}
