package form

type PreUpload struct {
	Size uint64 `json:"size"`
	Md5  string `json:"md5"`
}

// 秒传
type Instant struct {
	Md5      string `json:"md5"`
	Name     string `json:"name"`
	TargetId uint   `json:"target_id"`
}
