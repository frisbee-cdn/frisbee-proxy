package grpc

type FindValue struct {
	Key string `json:"key"`
}

type FindValueReply struct {
	Value []byte `json:"value"`
}

type Store struct{
	Key string `json:"key"`
	Value []byte `json:"content"`
}

type Error struct {
	Message string `json:"message"`
}

