package models

import (
	"strings"
)

type File struct {
	BaseModel
	Md5      string `json:"md5"`
	Path     string `json:"path"`
	Size     uint64 `json:"size"`
	Extend   string `json:"extend"`
	MimeType string `json:"mime_type"`
	OwnId    uint   `json:"own_id"`
	RefCount uint   `json:"ref_count"`
	Kind     uint8  `json:"kind"`
}

// 文件大分类
const (
	KIND_OTHER = 0 // 其他
	KIND_DOC   = 1 // 文档
	KIND_IMAGE = 2 // 图版
	KIND_AUDIO = 3 // 音频
	KIND_VIDEO = 4 // 视频
)

var Ext2kind = map[string]uint8{
	// Docment
	"txt":  KIND_DOC,
	"rtf":  KIND_DOC,
	"doc":  KIND_DOC,
	"docx": KIND_DOC,
	"ppt":  KIND_DOC,
	"pptx": KIND_DOC,
	"xls":  KIND_DOC,
	"xlsx": KIND_DOC,
	"pdf":  KIND_DOC,

	// Image
	"png":  KIND_IMAGE,
	"bmp":  KIND_IMAGE,
	"jpg":  KIND_IMAGE,
	"gif":  KIND_IMAGE,
	"jpeg": KIND_IMAGE,
	"psd":  KIND_IMAGE,
	"tiff": KIND_IMAGE,
	"svg":  KIND_IMAGE,
	"cdr":  KIND_IMAGE,
	"emf":  KIND_IMAGE,

	// Audio
	"mp3":  KIND_AUDIO,
	"aac":  KIND_AUDIO,
	"wma":  KIND_AUDIO,
	"wav":  KIND_AUDIO,
	"ape":  KIND_AUDIO,
	"flac": KIND_AUDIO,
	"midi": KIND_AUDIO,

	// Video
	"avi":  KIND_VIDEO,
	"mp4":  KIND_VIDEO,
	"flv":  KIND_VIDEO,
	"f4v":  KIND_VIDEO,
	"rmb":  KIND_VIDEO,
	"wmv":  KIND_VIDEO,
	"3gp":  KIND_VIDEO,
	"mkv":  KIND_VIDEO,
	"rw":   KIND_VIDEO,
	"rmvb": KIND_VIDEO,
	"rm":   KIND_VIDEO,
	"mov":  KIND_VIDEO,
	"mpeg": KIND_VIDEO,
}

func (f *File) GetFileByMd5(md5 string) bool {
	if err := f.DB().Where("md5 = ?", md5).First(f).Error; err != nil {
		return false
	}

	return true
}

func (f *File) IncRef() error {
	f.RefCount++
	return f.DB().Model(f).Update("ref_count", f.RefCount).Error
}

func (f *File) DeIncRef() error {
	if f.RefCount >= 1 {
		f.RefCount--
	}
	return f.DB().Model(f).Update("ref_count", f.RefCount).Error
}

// 已上传成功,  入库
func (f *File) Create(userId uint, size uint64, md5, path, ext, mimeType string) error {
	f.Size = size
	f.Md5 = md5
	f.Extend = ext
	f.OwnId = userId
	f.RefCount = 1
	f.MimeType = mimeType

	// 将windows下路径分隔符替换成Unix形式入库
	f.Path = strings.ReplaceAll(path, "\\", "/")
	// 没有值时将会使用空值也就是0，0即是Other
	f.Kind = Ext2kind[ext]

	return f.DB().Model(f).Create(f).Error
}
