package db

type PKModel interface {
	PK() uint64
	SetPK(id uint64)
}
