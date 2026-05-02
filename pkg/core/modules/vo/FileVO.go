package vo

type FileHashVO struct {
	FileName string `validate:"required" json:"fileName" form:"fileName"`
	FileHash string `validate:"required" json:"fileHash" form:"fileHash"`
}

type FileChunkHashVO struct {
	FileName string `validate:"required" json:"fileName" form:"fileName"`
	FileHash string `validate:"required" json:"fileHash" form:"fileHash"`
	Hash     string `validate:"required" json:"hash" form:"hash"`
}

type FileChunkMergeVO struct {
	FileName  string `validate:"required" json:"fileName" form:"fileName"`
	FileHash  string `validate:"required" json:"fileHash" form:"fileHash"`
	ChunkSize int64  `json:"chunkSize" form:"chunkSize"`
}
