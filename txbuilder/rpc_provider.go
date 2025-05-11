package txbuilder

type RpcProvider interface {
	GetPendingNonce(address string) (uint64, error)
}
