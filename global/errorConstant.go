package global

import (
	"github.com/liuhongdi/digv15/pkg/result"
)
var (
	// OK
	OK = result.NewError(0, "OK")

	//参数模块
	ErrParam = result.NewError(400, "参数不合法")

	//文章模块报错
	ErrArticleNot = result.NewError(10001, "文章不存在")
	ErrArticleS = result.NewError(10002, "文章查询出错")

	//用户模块
	ErrUserNot = result.NewError(20001, "用户不存在")

	// ...
)
