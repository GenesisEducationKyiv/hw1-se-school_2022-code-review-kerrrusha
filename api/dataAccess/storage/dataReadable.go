package storage

type DataReadable interface {
	Read() []byte
}
