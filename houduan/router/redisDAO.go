package main

import (
	"Go_Work/demo/CreateDemo/ginvuedemo1/houduan/model"
	"fmt"

	"github.com/gin-gonic/gin"
)

func server(c *gin.Engine) {
	serverLogin(c)
	serverRegister(c)
	serverGetOtherData(c)
	serverComment(c)
	serverChangeUser(c)
	serverBuyTickets(c)
	serverPayOrder(c)
	serverRefundOrder(c)
	cancelRefundOrder(c)
}

func serverLogin(c *gin.Engine) {
	//通过域名获取传入参数 适合单参数
	// router.POST("/upload/:username", func(c *gin.Context) {
	//c.Get 用于获取 Gin 上下文中的数据，而 c.BindJSON 用于解析请求体中的 JSON 数据到 Go 结构体
	//username, _ := c.Get("username")
	// 	username := c.Param("username")
	// 	fmt.Println(username)
	// })
	c.POST("/upload/login", func(c *gin.Context) {
		user := model.UserData{}
		//获取从前端传过来的数据
		err := c.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println(err)
		} else {
			user, err := model.Login(user)
			if err != nil {
				if err == model.ERROR_USER_NOT_EXIST {
					c.JSON(200, gin.H{
						"code": 400,
						"msg":  "用户不存在",
					})
				} else if err == model.ERROR_USER_PASSWORD {
					c.JSON(200, gin.H{
						"code": 400,
						"msg":  "密码错误",
					})
				} else {
					c.JSON(200, gin.H{
						"code": 500,
						"msg":  "未知错误，请重新尝试",
					})
				}
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "登陆成功",
					"data": user,
				})
			}

		}
	})
}

func serverRegister(c *gin.Engine) {
	c.POST("/upload/register", func(c *gin.Context) {
		user := model.UserData{}
		err := c.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println(err)
		} else {
			err := model.Register(user)
			if err != nil {
				if err == model.ERROR_USER_EXIST {
					c.JSON(200, gin.H{
						"code": 500,
						"msg":  "用户已存在，请重新输入用户名以及密码",
					})
				} else {
					c.JSON(200, gin.H{
						"code": 500,
						"msg":  "注册失败,请重新输入",
					})
				}
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "注册成功",
				})
			}
		}
	})
}

func serverGetOtherData(c *gin.Engine) {
	c.POST("/upload/getotherdata", func(c *gin.Context) {
		packages := model.Package{}
		err := c.ShouldBindJSON(&packages)
		if err != nil {
			fmt.Println(err)
		} else {
			err := model.GetData(&packages)
			if err != nil {
				if err == model.ERROR_USER_UPDATA {
					c.JSON(200, gin.H{
						"code": 202,
						"msg":  packages,
					})
				} else {
					c.JSON(200, gin.H{
						"code": 500,
						"msg":  err,
					})
				}
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "无需更新",
				})
			}
		}
	})
}

func serverComment(c *gin.Engine) {
	c.POST("/upload/comment", func(c *gin.Context) {
		comment := model.Comments{}
		err := c.ShouldBindJSON(&comment)
		if err != nil {
			fmt.Println("获取评论失败,", err)
		} else {
			err = model.AddComment(comment)
			if err != nil {
				c.JSON(200, gin.H{
					"code": 500,
					"msg":  err,
				})
			} else {
				c.JSON(200, gin.H{
					"code": 200,
					"msg":  "评论成功",
				})
			}
		}
	})
}

func serverChangeUser(c *gin.Engine) {
	c.POST("/upload/changeuser", func(ctx *gin.Context) {
		user := model.NewUser{}
		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(200, gin.H{
				"code": 500,
				"msg":  err,
			})
		} else {
			err = model.ChangeUser(user)
			if err != nil {
				ctx.JSON(200, gin.H{
					"code": 500,
					"msg":  err,
				})
			} else {
				ctx.JSON(200, gin.H{
					"code": 200,
					"user": user,
				})
			}

		}
	})
}

func serverBuyTickets(c *gin.Engine) {
	c.POST("/upload/buytickets", func(ctx *gin.Context) {
		msg := model.Msg{}
		err := ctx.ShouldBindJSON(&msg)
		if err != nil {
			fmt.Println(err)
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		} else {
			orderlist, err := model.BuyTickets(msg)
			if err != nil {
				ctx.JSON(400, gin.H{
					"code": 500,
					"msg":  err,
				})
			} else {
				ctx.JSON(200, gin.H{
					"code": 200,
					"msg":  orderlist,
				})
			}
		}
	})
}

func serverPayOrder(c *gin.Engine) {
	c.POST("/upload/payorder", func(ctx *gin.Context) {
		msg := model.Msg{}
		err := ctx.ShouldBindJSON(&msg)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		}
		orderlist, err := model.PayOrder(msg)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		} else {
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  orderlist,
			})
		}
	})
}

func serverRefundOrder(c *gin.Engine) {
	c.POST("/upload/refundorder", func(ctx *gin.Context) {
		msg := model.Msg{}
		err := ctx.ShouldBindJSON(&msg)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		}
		refundlist, err := model.RefundOrder(msg)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		} else {
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  refundlist,
			})
		}
	})
}

func cancelRefundOrder(c *gin.Engine) {
	c.POST("/upload/cancelrefundorder", func(ctx *gin.Context) {
		msg := model.Msg{}
		err := ctx.ShouldBindJSON(&msg)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		}
		refundlist, err := model.CancelRefundOrder(msg)
		if err != nil {
			ctx.JSON(400, gin.H{
				"code": 500,
				"msg":  err,
			})
		} else {
			ctx.JSON(200, gin.H{
				"code": 200,
				"msg":  refundlist,
			})
		}
	})
}
