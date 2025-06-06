package blockchain_ethereum

type IContract[T any] interface {
	GetContract(chainID uint64) (T, error)
	GetAllAddress() ([]T, error)
}
