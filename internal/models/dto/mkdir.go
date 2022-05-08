package dto

import "github.com/xstnet/starfire-cloud/internal/models"

type MkdirDto struct {
	*models.UserFile
}

func LoadMkdir(uf *models.UserFile) *MkdirDto {
	return &MkdirDto{
		UserFile: uf,
	}
}
