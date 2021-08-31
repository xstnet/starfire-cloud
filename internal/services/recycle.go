package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
)

func RecycleRestore(c *gin.Context, userId uint) (int64, error) {
	var data = &form.FileIdList{}
	if err := c.ShouldBind(data); err != nil {
		return 0, errors.InvalidParameter()
	}

	result := models.DB().Model(&models.UserFile{}).
		Where(data.FIdList).
		Where("user_id = ? and is_delete=?", userId, models.IS_DELETE_YES).
		Update("is_delete", 0)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
