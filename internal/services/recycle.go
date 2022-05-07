package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
)

func RecycleRestore(c *gin.Context, userId uint) (int64, error) {
	var data = &form.FileIdsItem{}
	if err := c.ShouldBind(data); err != nil {
		return 0, errors.InvalidParameter()
	}

	result := models.DB().Model(&models.UserFile{}).
		Where(data.FileIds).
		Where("user_id = ? and is_delete=?", userId, models.IsDeleteYes).
		Update("is_delete", 0)

	// c.ShouldBindBodyWith()
	// binding.JSON.BindBody()

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
