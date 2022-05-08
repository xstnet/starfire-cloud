package recycle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

func Restore(c *gin.Context) {
	rowsAffected, err := doRestore(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, fmt.Sprintf("还原成功, %d条数据已恢复", rowsAffected), nil)
}

func doRestore(c *gin.Context, userId uint) (int64, error) {
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
