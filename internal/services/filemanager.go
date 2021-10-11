package services

import (
	"github.com/gin-gonic/gin"
	"github.com/xstnet/starfire-cloud/internal/common"
	"github.com/xstnet/starfire-cloud/internal/errors"
	"github.com/xstnet/starfire-cloud/internal/models"
	"github.com/xstnet/starfire-cloud/internal/models/form"
	"github.com/xstnet/starfire-cloud/pkg/convert"
)

// 创建文件夹
func Mkdir(c *gin.Context, userId uint) (*models.UserFile, error) {
	userFile := &models.UserFile{}
	err := c.ShouldBindJSON(&userFile)
	if err != nil {
		return nil, errors.InvalidParameter()
	}

	if err := common.CheckDirname(userFile.Name); err != nil {
		return nil, err
	}

	userFile.UserId = userId

	if err := userFile.Mkdir(); err != nil {
		return nil, errors.New("创建文件夹失败，原因: " + err.Error())
	}

	return userFile, nil
}

// 重命名
func Rename(c *gin.Context, userId uint) (*models.UserFile, error) {
	var data = make(gin.H, 4)
	err := c.ShouldBindJSON(&data)

	if err != nil {
		return nil, errors.InvalidParameter()
	}

	id, ok := convert.GetFloat64(data["id"])
	if !ok || id <= 0 {
		return nil, errors.InvalidParameter()
	}
	newname, ok := convert.GetString(data["newname"])
	if !ok {
		return nil, errors.InvalidParameter()
	}

	if err := common.CheckFilename(newname); err != nil {
		return nil, err
	}

	userFile := &models.UserFile{}
	if result := userFile.DB().First(userFile, uint(id)); result.Error != nil || userFile.UserId != userId {
		return nil, errors.New("操作对象不存在")
	}

	userFile.Name = newname
	if err := userFile.Rename(); err != nil {
		return nil, err
	}

	return userFile, nil
}

// 移动
func Move(c *gin.Context, userId uint) (*models.UserFile, error) {
	var data = make(gin.H, 4)
	err := c.ShouldBindJSON(&data)

	if err != nil {
		return nil, errors.InvalidParameter()
	}

	fromId, ok := convert.GetFloat64(data["from_id"])
	if !ok || fromId <= 0 {
		return nil, errors.InvalidParameter()
	}
	destId, ok := convert.GetFloat64(data["dest_id"])
	if !ok || destId <= 0 {
		return nil, errors.InvalidParameter()
	}

	userFile := &models.UserFile{}
	if result := userFile.DB().First(userFile, uint(fromId)); result.Error != nil || userFile.UserId != userId {
		return nil, errors.New("操作对象不存在")
	}

	userFile.ParentId = uint(destId)
	if err := userFile.Move(); err != nil {
		return nil, err
	}

	return userFile, nil
}

// 获取文件列表
func List(c *gin.Context, userId uint) (*gin.H, error) {
	listForm := &form.FileList{}
	if err := c.ShouldBind(listForm); err != nil {
		return nil, errors.InvalidParameter()
	}

	limit, offset := common.ProcessPageByList(listForm.Page, listForm.PageSize, 200)

	where := map[string]interface{}{
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
		return &gin.H{"list": &userFiles, "more": more}, nil
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

	fileIds = common.SliceUniqueUint(&fileIds)

	files := make([]models.File, len(fileIds))
	models.DB().Find(&files, fileIds)

	var mapId2File = make(map[uint]models.File, len(files))

	for _, v := range files {
		mapId2File[v.ID] = v
	}

	var listData = make([]map[string]interface{}, 0, len(userFiles))

	for _, v := range userFiles {
		item := common.Struct2Map(v)
		item["file"] = map[string]interface{}{
			"id": 0,
		}
		if v.IsDir == models.IS_DIR_NO {
			file, ok := mapId2File[v.FileId]
			if ok {
				item["file"] = map[string]interface{}{
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

	return &gin.H{
		"list": &listData,
		"more": more,
	}, nil
}

// 删除
// 只标记当前节点，不处理子元素，从回收站删除时再处理子元素
func Delete(c *gin.Context, userId uint) (int64, error) {
	var data = &form.FileIdList{}
	if err := c.ShouldBind(data); err != nil {
		return 0, errors.InvalidParameter()
	}

	result := models.DB().Model(&models.UserFile{}).
		Where(data.FIdList).
		Where("user_id = ? and is_delete=?", userId, models.IS_DELETE_NO).
		Update("is_delete", 1)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}

func DirList(c *gin.Context, userId uint) (*[]models.UserFile, error) {
	listForm := &form.FileList{}
	if err := c.ShouldBind(listForm); err != nil {
		return nil, errors.InvalidParameter()
	}

	userFile := &models.UserFile{
		UserId:   userId,
		ParentId: listForm.ParentId,
	}

	data := userFile.DirList()

	return data, nil
}
