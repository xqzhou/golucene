package store

type RAMDirectory struct {
	Directory
}

func NewRAMDirectory() *RAMDirectory {
	return &RAMDirectory{}
}
