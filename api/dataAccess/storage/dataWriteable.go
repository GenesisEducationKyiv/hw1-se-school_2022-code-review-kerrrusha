package storage

type DataWriteable interface {
	Write(content string, append bool) int
}
