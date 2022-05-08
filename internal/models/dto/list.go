package dto

import "github.com/xstnet/starfire-cloud/internal/models"

type FileList struct {
	models.UserFile
	File models.File `json:"file"`
}

func LoadFileList(uf models.UserFile) []FileList {

}
