package tree

type WalHeader struct {
	Key       uint64
	Size      uint64
	Timestamp int64
}

type WalEntry struct {
	WalHeader
	Value []byte
}
