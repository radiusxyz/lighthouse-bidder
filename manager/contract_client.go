package manager

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	common2 "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/radiusxyz/lighthouse-bidder/common"
	"github.com/radiusxyz/lighthouse-bidder/config"
	"github.com/radiusxyz/lighthouse-bidder/contracts/bindings"
	"math/big"
	"strings"
)

type ContractClient struct {
	ethClient    *ethclient.Client
	parsedABI    abi.ABI
	filterer     *bindings.BindingsFilterer
	contract     *bindings.Bindings
	transactOpts *bind.TransactOpts
}

func NewContractClient(conf *config.Config) (*ContractClient, error) {
	ethClient, err := ethclient.Dial(*conf.LighthouseChainUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum client: " + err.Error())
	}

	parsedABI, err := abi.JSON(strings.NewReader(bindings.BindingsMetaData.ABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse ABI: " + err.Error())
	}

	filterer, err := bindings.NewBindingsFilterer(common2.HexToAddress(*conf.LighthouseContractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create filterer: " + err.Error())
	}

	contract, err := bindings.NewBindings(common2.HexToAddress(*conf.LighthouseContractAddress), ethClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create contract: " + err.Error())
	}

	privateKey, err := common.LoadPrivateKey(*conf.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to load private key: " + err.Error())
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(int64(*conf.LighthouseChainId)))
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: " + err.Error())
	}
	transactOpts.Context = context.Background()
	transactOpts.GasLimit = *conf.GasLimit

	return &ContractClient{
		ethClient:    ethClient,
		parsedABI:    parsedABI,
		filterer:     filterer,
		contract:     contract,
		transactOpts: transactOpts,
	}, nil
}

func (c *ContractClient) EthClient() *ethclient.Client {
	return c.ethClient
}
func (c *ContractClient) ParsedABI() abi.ABI {
	return c.parsedABI
}
func (c *ContractClient) Filterer() *bindings.BindingsFilterer {
	return c.filterer
}
func (c *ContractClient) Contract() *bindings.Bindings {
	return c.contract
}
func (c *ContractClient) TransactOpts() *bind.TransactOpts {
	return c.transactOpts
}

func (c *ContractClient) GetNonce(bidderAddress common2.Address) (*big.Int, error) {
	callOpts := &bind.CallOpts{
		Context: context.Background(),
		Pending: true,
	}

	nonce, err := c.contract.GetNonce(callOpts, bidderAddress)
	if err != nil {
		return nil, err
	}
	return nonce, nil
}
