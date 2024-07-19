package models

type FindAllPage struct {
	Size   uint64
	Offset uint64
}

type FindResult struct {
	Data   []string
	Cursor uint64
}
