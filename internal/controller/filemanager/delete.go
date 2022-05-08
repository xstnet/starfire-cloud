package filemanager

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/response"
)

func Delete(c *gin.Context) {
	rowsAffected, err := doDelete(c, c.GetUint("userId"))
	if err != nil {
		response.Error(c, err.Error())
		return
	}
	response.Success(c, fmt.Sprintf("删除成功, %d条数据已放入回收站", rowsAffected), nil)
}

// Delete 删除
// 只标记当前节点，不处理子元素，从回收站删除时再处理子元素
func doDelete(c *gin.Context, userId uint) (int64, error) {
	params, err := form.GetJsonForm[form.FileIdsItem](c)
	if err != nil {
		return 0, errors.InvalidParameter()
	}

	result := models.DB().Model(new(models.UserFile)).
		Where(params.FileIds).
		Where("user_id = ? and is_delete=?", userId, models.IsDeleteNo).
		Update("is_delete", 1)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
