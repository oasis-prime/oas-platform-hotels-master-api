package internal

type Internal struct {
}

var internal Internal

func GetInternal() *Internal {
	return &internal
}
