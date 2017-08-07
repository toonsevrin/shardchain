package types


type Block struct {
	index uint
	previousHash string
	timestamp uint64
	data string//type?
	hash string//sha3
}

