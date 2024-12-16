package model

import "errors"

var (
	//200成功 202更新数据 500失败
	ERROR_USER_NOT_EXIST      = errors.New("用户不存在")
	ERROR_USER_EXIST          = errors.New("用户已存在")
	ERROR_USER_PASSWORD       = errors.New("密码错误")
	ERROR_USER_UPDATA         = errors.New("更新用户信息")
	ERROR_TICKETS_NOT_ENOUGH  = errors.New("余票不足")
	ERROR_ORDER_HAS_EXIST     = errors.New("订单已存在")
	ERROR_TICKETS_NOT_EXIST   = errors.New("票不存在")
	ERROR_ORDER_HAS_NOT_EXIST = errors.New("订单不存在")
)
