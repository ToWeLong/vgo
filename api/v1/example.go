package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/towelong/vgo/biz"
)

func (s *Server) ping(ctx *gin.Context) {
	atc, err := biz.Article.GetArticleById(ctx, 1)
	if err != nil {
		_ = ctx.Error(err)
	}
	ctx.JSON(200, atc)
}
