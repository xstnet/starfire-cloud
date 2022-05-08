package dto

import "github.com/xstnet/starfire-cloud/internal/models"

type SingleUploadDto struct {
	*models.UserFile
	File *models.File `json:"file"`
}

func LoadSingleUpload(uf *models.UserFile, file *models.File) *SingleUploadDto {
	return &SingleUploadDto{
		UserFile: uf,
		File:     file,
	}
}
