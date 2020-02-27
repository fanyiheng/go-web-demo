package util

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
)

const (
	DefaultPageSize = 10
	DefaultPageNum  = 0
)

type Page struct {
	// 每页记录数
	PageSize    int         `json:"page_size"`
	PageNum     int         `json:"page_num"`
	TotalPage   int         `json:"total_page"`
	TotalRecord int         `json:"total_record"`
	Data        interface{} `json:"data"`
}

func NewPage(c *gin.Context) *Page {
	pageNum := c.Query("page_num")
	pageSize := c.Query("page_size")
	pageNumI, _ := strconv.Atoi(pageNum)
	pageSizeI, _ := strconv.Atoi(pageSize)
	return &Page{PageNum: pageNumI, PageSize: pageSizeI}
}

func PageFind(db *gorm.DB, page *Page) error {
	var count int
	if err := db.Count(&count).Error; err != nil || count == 0 {
		return err
	}
	page.TotalRecord = count
	if page.PageSize <= 0 {
		page.PageSize = DefaultPageSize
	}
	if page.PageNum <= 0 {
		page.PageNum = DefaultPageNum
	}
	page.TotalPage = count / page.PageSize
	if count%page.PageSize > 0 {
		page.TotalPage++
	}
	limit := page.PageSize * (page.PageNum + 1)
	offset := page.PageSize * page.PageNum
	return db.Limit(limit).Offset(offset).Find(page.Data).Error
}
