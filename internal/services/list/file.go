package list

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/helper/common"
	"github.com/xstnet/starfire-cloud/pkg/helper/d"
	"github.com/xstnet/starfire-cloud/pkg/helper/slice"
)

// FileList 获取文件列表
func FileList(c *gin.Context, userId uint) (*d.StringMap, error) {
	listForm, err := form.GetForm[form.FileList](c)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	limit, offset := common.ProcessPage(listForm.Page, listForm.PageSize, 200)

	where := d.StringMap{
		"user_id":   userId,
		"parent_id": listForm.ParentId,
		"is_delete": models.IS_DELETE_NO,
	}

	userFiles := make([]models.UserFile, limit+1)

	// todo, 处理排序， 搜索
	result := models.DB().Where(where).Order("is_dir desc").Order("updated_at desc").Limit(limit + 1).Offset(offset).Find(&userFiles)
	if result.Error != nil {
		return nil, errors.New("获取列表失败，原因：" + result.Error.Error())
	}

	more := 0

	if result.RowsAffected <= 0 {
		return &d.StringMap{"list": &userFiles, "more": more}, nil
	}

	// 如果获取到的元素大于pageSize,说明还有下一页
	if result.RowsAffected > int64(limit) {
		more = 1
		userFiles = userFiles[:result.RowsAffected-1]
	}

	// 获取File信息写入到返回结果中
	var fileIds = make([]uint, 0, len(userFiles))
	for _, v := range userFiles {
		if v.IsDir == models.IS_DIR_NO {
			fileIds = append(fileIds, v.FileId)
		}
	}

	fileIds = slice.Unique(fileIds)

	files := make([]models.File, len(fileIds))
	models.DB().Find(&files, fileIds)

	var mapId2File = make(map[uint]models.File, len(files))

	for _, v := range files {
		mapId2File[v.ID] = v
	}

	var listData = make([]d.StringMap, 0, len(userFiles))

	for _, v := range userFiles {
		item := common.Struct2Map(v)
		item["file"] = d.StringMap{
			"id": 0,
		}
		if v.IsDir == models.IS_DIR_NO {
			file, ok := mapId2File[v.FileId]
			if ok {
				item["file"] = d.StringMap{
					"id":   file.ID,
					"ext":  file.Extend,
					"size": file.Size,
					"md5":  file.Md5,
					"kind": file.Kind,
				}
			}
		}

		delete(item, "is_delete")
		delete(item, "user_id")
		listData = append(listData, item)
	}

	return &d.StringMap{
		"list": &listData,
		"more": more,
	}, nil
}
