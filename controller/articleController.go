package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv15/global"
	"github.com/liuhongdi/digv15/pkg/page"
	"github.com/liuhongdi/digv15/pkg/result"
	"github.com/liuhongdi/digv15/pkg/validCheck"
	"github.com/liuhongdi/digv15/request"
	"github.com/liuhongdi/digv15/service"
)

type ArticleController struct{}

func NewArticleController() ArticleController {
	return ArticleController{}
}
//得到一篇文章的详情
func (a *ArticleController) GetOne(c *gin.Context) {
	resultRes := result.NewResult(c)
	param := request.ArticleRequest{ID: validCheck.StrTo(c.Param("id")).MustUInt64()}
	valid, _ := validCheck.BindAndValid(c, &param)
	if !valid {
		//result.ErrorCode(global.ErrParam.Code,global.ErrParam.Msg+":"+errs.Error())
		resultRes.Error(global.ErrParam)
		return
	}

	articleOne,err := service.GetOneArticle(param.ID);
	if err != nil {
		//result.Error(404,"数据查询错误")
		//result.ErrorCode(global.ErrArticleNot.Code,global.ErrArticleNot.Msg)
		resultRes.Error(global.ErrArticleNot)
	} else {
		resultRes.Success(&articleOne);
	}
	return
}

//得到多篇文章，按分页返回
func (a *ArticleController) GetList(c *gin.Context) {
	resultRes := result.NewResult(c)
	pageInt := 0
	//is exist?
	curPage := c.Query("page")
    //if curPage not exist
    if (len(curPage) == 0) {
		pageInt = 1
	} else {
		param := request.ArticleListRequest{Page: validCheck.StrTo(c.Param("page")).MustInt()}
		valid, _ := validCheck.BindAndValid(c, &param)
		if !valid {
			//result.Error(400,errs.Error())
			resultRes.Error(global.ErrParam)
			return
		}
		pageInt = param.Page
	}

	pageSize := 2;
	pageOffset := (pageInt-1) * pageSize

	articles,err := service.GetArticleList(pageOffset,pageSize)
	if err != nil {
		//result.ErrorCode(global.ErrArticleS.Code,"数据查询错误");
		resultRes.Error(global.ErrArticleS);
		fmt.Println(err.Error())
	} else {
		//sum,_ := dao.SelectcountAll()
		sum,_ := service.GetArticleSum()
		pageInfo,_ := page.GetPageInfo(pageInt,pageSize,sum)
		resultRes.Success(gin.H{"list":&articles,"pageinfo":pageInfo});
	}
	return
}

