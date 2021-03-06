package response

import (
	"github.com/gogf/gf/net/ghttp"
	"yj-app/app/model"
)

// 通用api响应
type TableResp struct {
	c *model.TableDataInfo
	r *ghttp.Request
}

//返回一个成功的消息体
func BuildTable(r *ghttp.Request, total int, rows interface{}) *TableResp {
	msg := model.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: total,
		Rows:  rows,
	}
	a := TableResp{
		c: &msg,
		r: r,
	}
	return &a
}

//输出json到客户端
func (resp *TableResp) WriteJsonExit() {
	resp.r.Response.WriteJsonExit(resp.c)
}
