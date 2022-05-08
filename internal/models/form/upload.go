package form

type PreUpload struct {
	Size uint64 `form:"size"`
	Md5  string `form:"md5"`
}

// 秒传
type Instant struct {
	Md5      string `form:"md5"`
	Name     string `form:"name"`
	TargetId uint   `form:"targetId"`
}
