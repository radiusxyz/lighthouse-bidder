package erigonclient

//
//import (
//	"context"
//	"github.com/ethereum/go-ethereum/rpc"
//)
//
//type Client struct {
//	c *rpc.Client
//}
//
//func DialContext(ctx context.Context, rawurl string) (*Client, error) {
//	c, err := rpc.DialContext(ctx, rawurl)
//	if err != nil {
//		return nil, err
//	}
//	return NewClient(c), nil
//}
//
//func NewClient(c *rpc.Client) *Client {
//	return &Client{c}
//}
//
//func (ec *Client) Close() {
//	ec.c.Close()
//}
//
//func (c *Client) ChainId
