package dirHelper

import (
	"regexp"

	"github.com/xstnet/starfire-cloud/internal/errors"
)

func CheckName(dirname string) error {
	if dirname == "" {
		return errors.New("文件夹名称不能为空")
	}
	matched, err := regexp.MatchString(`^[^/\\\\:\\*\\?\\<\\>\\|\"]{1,255}$`, dirname)
	if err != nil || !matched {
		return errors.New(`文件夹名称不能包含\/:*?"<>|`)
	}
	return nil
}
