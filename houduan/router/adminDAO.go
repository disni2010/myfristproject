package main

import (
	"Go_Work/demo/CreateDemo/ginvuedemo1/houduan/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type returnData struct {
	DataType string `json:"dataType"`
	Data     string `json:"data"`
}

func adminServer(c *gin.Engine) {
	adminLogin(c)
	adminIndex(c)
	adminList(c)
	commentList(c)
	tircketsList(c)
	scenicspotslist(c)
	refundlist(c)
	getData(c)
	AddTickets(c)
	AddScenicspot(c)
	AddUser(c)
}

func adminLogin(c *gin.Engine) {
	admin := model.Admin{Adminid: "admin", Adminpwd: "123456", IsUse: true}
	jsonadmin, err := json.Marshal(admin)
	if err != nil {
		panic(err)
	}
	//挂载网页并向里面传递参数
	c.GET("/login", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			//marshal后是[]byte 需要转为string
			"admin": string(jsonadmin),
		})
	})

	c.POST("/login/post", func(ctx *gin.Context) {
		err := ctx.ShouldBindJSON(&admin)
		if err != nil {
			fmt.Println(err)
			return
		}
		//做判断逻辑  这里没做逻辑 所以管理员直接可以登录

		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "登录成功",
		})
		// if adminid == admin.Adminid && adminpwd == admin.Adminpwd {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"code": 200,
		// 		"msg":  "登录成功",
		// 	})
		// } else {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"code": 400,
		// 		"msg":  "登录失败",
		// 	})
		// }
	})
}

func adminIndex(c *gin.Engine) {
	c.GET("/admin/index", func(ctx *gin.Context) {
		ticketslist, err := model.GetTicketsList()
		if err != nil {
			fmt.Println(err)
			return
		}
		//json.Marshal 返回字节切片:
		// json.Marshal 返回一个字节切片，表示 JSON 字符串的 UTF-8 编码形式。
		// 解析字节切片:
		// 字节切片中的每个数字代表一个字节的值，可以解析为实际的 JSON 字符串。
		// 转换为字符串:
		// 使用 string(jsonBytes) 将字节切片转换为字符串。
		ticketslistjson, _ := json.Marshal(ticketslist)
		scenicspotslist, err := model.GetScenicspotsList()
		if err != nil {
			fmt.Println(err)
			return
		}
		scenicspotslistjson, _ := json.Marshal(scenicspotslist)
		commentslist, err := model.GetAllComments()
		if err != nil {
			fmt.Println(err)
			return
		}
		commentslistjson, _ := json.Marshal(commentslist)
		allusers, err := model.GetAllUsaers()
		if err != nil {
			fmt.Println(err)
			return
		}
		allusersjson, _ := json.Marshal(allusers)
		refundlist, err := model.GetRefundList()
		if err != nil {
			fmt.Println(err)
			return
		}
		refundlistjson, _ := json.Marshal(refundlist)
		ctx.HTML(http.StatusOK, "adminMain.html", gin.H{
			"ticketslist":     string(ticketslistjson),
			"scenicspotslist": string(scenicspotslistjson),
			"commentslist":    string(commentslistjson),
			"allusers":        string(allusersjson),
			"refundlist":      string(refundlistjson),
		})
	})
}

func adminList(c *gin.Engine) {
	c.GET("/admin/list", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "adminControll.html", gin.H{})
	})
}
func commentList(c *gin.Engine) {
	c.GET("/admin/commentlist", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "commentList.html", gin.H{})
	})
}
func tircketsList(c *gin.Engine) {
	c.GET("/admin/tircketslist", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "tirckets.html", gin.H{})
	})
}
func scenicspotslist(c *gin.Engine) {
	c.GET("/admin/scenicSpotlist", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "scenicSpotlist.html", gin.H{})
	})
}
func refundlist(c *gin.Engine) {
	c.GET("/admin/refundlist", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "refundControll.html", gin.H{})
	})
}

func getData(c *gin.Engine) {
	c.POST("/admin/UpData", func(ctx *gin.Context) {
		var returndata returnData
		err := ctx.ShouldBindJSON(&returndata)
		if err != nil {
			fmt.Println(err)
		}
		switch returndata.DataType {
		case "TicketsList":
			var tickets model.TicketsList
			err := json.Unmarshal([]byte(returndata.Data), &tickets)
			if err != nil {
				fmt.Println(err)
			} else {
				//获取返回的数据后 插入数据库
				err = model.SetTicketsList(tickets)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "ScenicspotsList":
			var scenicspots model.ScenicspotsList
			err := json.Unmarshal([]byte(returndata.Data), &scenicspots)
			if err != nil {
				fmt.Println(err)
			} else {
				err = model.SetScenicspotsList(scenicspots)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "CommenstList":
			var comment model.CommenstList
			err := json.Unmarshal([]byte(returndata.Data), &comment)
			if err != nil {
				fmt.Println(err)
			} else {
				err = model.SetCommentList(comment)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "UserList":
			var userlist model.UserList
			err := json.Unmarshal([]byte(returndata.Data), &userlist)
			if err != nil {
				fmt.Println(err)
			} else {
				err = model.UpdataUAllUser(userlist)
				if err != nil {
					fmt.Println(err)
				}
			}
		case "RefundsList":
			var refundlist model.RefundsList
			err := json.Unmarshal([]byte(returndata.Data), &refundlist)
			if err != nil {
				fmt.Println(err)
			} else {
				err = model.SetRefundList(refundlist)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "更新成功",
			})
		}
	})
}

func AddTickets(c *gin.Engine) {
	c.POST("/upload/addtickets", func(ctx *gin.Context) {
		var ticket model.Tickets
		err := ctx.ShouldBindJSON(&ticket)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求参数错误",
			})
			return
		}
		ticketlist, err := model.GetTicketsList()
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  "获取门票列表失败",
			})
			return
		}
		for _, v := range ticketlist.Tickets {
			if v.Ticketname == ticket.Ticketname {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 402,
					"msg":  "门票已存在",
				})
				return
			}
		}
		err = model.AddTickets(ticket)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "添加成功",
		})
	})
}

func AddUser(c *gin.Engine) {
	c.POST("/upload/adduser", func(ctx *gin.Context) {
		var user model.UserData
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求参数错误",
			})
			return
		}
		err = model.Register(user)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err,
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "添加成功",
		})
	})
}

func AddScenicspot(c *gin.Engine) {
	c.POST("/upload/addscenicspot", func(ctx *gin.Context) {
		var scenicspot model.ScenicspotsList
		err := ctx.ShouldBindJSON(&scenicspot)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "请求参数错误",
			})
			return
		}
		err = model.SetScenicspotsList(scenicspot)
		if err != nil {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 500,
				"msg":  err,
			})
			return
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "添加成功",
			})
		}
	})
}
