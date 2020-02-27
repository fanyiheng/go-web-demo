package v1

import (
	"github.com/fanyiheng/go-web-demo/er"
	"github.com/fanyiheng/go-web-demo/persist"
	"github.com/fanyiheng/go-web-demo/router/api"
	"github.com/fanyiheng/go-web-demo/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PageAccount(c *gin.Context) (api.Resp, error) {
	name := c.Query("name")
	account := &persist.Account{Name: name}
	page := util.NewPage(c)
	err := account.Find(page)
	if err == nil {
		return page, nil
	}
	return nil, er.Msg("账户分页查询异常").Src(err)
}

// account/:id
func GetAccount(c *gin.Context) (api.Resp, error) {
	id := c.Param("id")
	idi, err := strconv.Atoi(id)
	if err != nil || idi < 0 {
		return nil, er.Msgf("账户ID不合法[%s]", id)
	}
	account := &persist.Account{}
	account.ID = uint(idi)
	err = account.GetById()
	if err == nil {
		return account, nil
	}
	return nil, er.Msg("账户获取异常").Src(err)
}

// DELETE /account/:id
func Delete(c *gin.Context) (api.Resp, error) {
	id := c.Param("id")
	idi, err := strconv.Atoi(id)
	if err != nil || idi < 0 {
		return nil, er.Msgf("账户ID不合法[%s]", id)
	}
	account := &persist.Account{}
	account.ID = uint(idi)
	err = account.Delete()
	if err == nil {
		return 1, nil
	}
	return nil, er.Msg("账户删除异常").Src(err)
}
