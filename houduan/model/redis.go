package model

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/goccy/go-json"
)

//不能导入main程序

var (
	connRedis *redis.Client
)

func InitRedis() {
	connRedis = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}

func getUserById(userid string) (user UserData, err error) {
	res, err := connRedis.HGet("User"+userid, "UserData").Result()
	if err != nil {
		err = ERROR_USER_NOT_EXIST
		return
	}
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("数据反序列化失败")
		return
	}
	return
}

func getUserByName(username string) (user UserData, err error) {
	userlist, err := GetAllUsaers()
	if err != nil {
		return
	}
	err = ERROR_USER_NOT_EXIST
	for _, v := range userlist.UserList {
		if v.UserData.Username == username {
			user = v.UserData
			err = nil
			return
		}
	}
	return
}

// 登录处理函数
func Login(loginuser UserData) (user UserData, err error) {
	user, err = getUserByName(loginuser.Username)
	if err != nil {
		fmt.Println("用户不存在")
		err = ERROR_USER_NOT_EXIST
		return
	}
	if user.UserPwd != loginuser.UserPwd {
		fmt.Println(user.Username + ":" + loginuser.Username)
		fmt.Println("密码错误")
		err = ERROR_USER_PASSWORD
		return
	}
	return
}

// 注册处理函数
func Register(user UserData) (err error) {
	_, err = getUserById(user.Userid)
	if err == nil {
		fmt.Println("用户已存在")
		err = ERROR_USER_EXIST
		return
	}
	_, err = getUserByName(user.Username)
	if err == nil {
		fmt.Println("用户已存在")
		err = ERROR_USER_EXIST
		return
	}
	hasExit := false
	Member := 0
	for {
		if hasExit {
			break
		}
		_, err := getUserById(strconv.Itoa(Member))
		if err != nil {
			user.Userid = strconv.Itoa(Member)
			hasExit = true
		} else {
			Member++
		}
	}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	err = connRedis.HSet("User"+user.Userid, "UserData", data).Err()
	if err != nil {
		fmt.Println("设置用户数据失败")
		return
	}
	err = connRedis.HSet("User"+user.Userid, "OrdersList", "OrdersList"+user.Userid).Err()
	if err != nil {
		fmt.Println("设置用户列表失败")
		return
	}
	err = connRedis.Set("OrdersList"+user.Userid, "{}", 0).Err()
	if err != nil {
		fmt.Println("设置用户订单列表失败")
		return
	}
	return
}

// 获取订单列表函数
func GetData(packages *Package) (err error) {
	needUpdate := false
	ScenicspotsListlen, Scenerr := GetScenicspotListLen()
	if Scenerr != nil {
		fmt.Println("获取景点列表失败")
		return
	}
	if ScenicspotsListlen != packages.Scenicspotslistlen {
		needUpdate = true
		packages.Scenicspotslistlen = ScenicspotsListlen
		err = ERROR_USER_UPDATA
		data, dataerr := GetScenicspotsList()
		if dataerr != nil {
			fmt.Println("读取景点列表失败")
			return
		}
		packages.Scenicspotlist = data
	}
	TicketsListlen, Ticketserr := GetTicketsListLen()
	if Ticketserr != nil {
		fmt.Println("获取门票列表失败")
		return
	}
	if TicketsListlen != packages.Ticketslistlen {
		needUpdate = true
		packages.Ticketslistlen = TicketsListlen
		err = ERROR_USER_UPDATA
		data, dataerr := GetTicketsList()
		if dataerr != nil {
			fmt.Println("读取门票列表失败")
			return
		}
		packages.Ticketslist = data
	}
	orederlen, ordererr := GetOrdersListLen(packages.Userid)
	if ordererr != nil {
		fmt.Println("获取用户信息失败")
		return
	}
	if orederlen != packages.Orderslistlen {
		needUpdate = true
		packages.Orderslistlen = orederlen
		err = ERROR_USER_UPDATA
		data, dataerr := GetOrdersList(packages.Userid)
		if dataerr != nil {
			fmt.Println("读取订单列表失败")
			return
		}
		packages.Orderslist = data
	}
	RefundList, Refunderr := GetRefundListById(packages.Userid)
	if Refunderr != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	refundlen, err := GetRefundListLenByUserId(packages.Userid)
	if err != nil {
		fmt.Println("获取退款列表长度失败")
		return
	}
	if packages.Refundlistlen != refundlen {
		needUpdate = true
		err = ERROR_USER_UPDATA
		packages.Refundlist = RefundList
		packages.Refundlistlen = len(RefundList.Refunds)
	}
	if needUpdate {
		err = ERROR_USER_UPDATA
		return
	}
	return
}

// 获得景区列表的函数
func GetScenicspotsList() (scenicspotsList ScenicspotsList, err error) {
	res, err := connRedis.Get("ScenicspotsList").Result()
	if err != nil {
		fmt.Println("获取景点列表失败")
		return
	}
	err = json.Unmarshal([]byte(res), &scenicspotsList)
	if err != nil {
		fmt.Println("数据反序列化失败")
		return
	}
	return
}

func GetScenicspotListLen() (lenght int, err error) {
	res, err := connRedis.Get("ScenicspotsList").Result()
	if err != nil {
		fmt.Println("获取景点列表失败")
		return
	}
	var scenicspotsList ScenicspotsList
	err = json.Unmarshal([]byte(res), &scenicspotsList)
	if err != nil {
		fmt.Println("scenicspotsList数据反序列化失败")
		return
	}
	lenght = len(scenicspotsList.Scenicspots)
	return
}

func SetScenicspotsList(scenicspotsList ScenicspotsList) (err error) {
	data, err := json.Marshal(scenicspotsList)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	err = connRedis.Set("ScenicspotsList", data, 0).Err()
	if err != nil {
		fmt.Println("设置景点列表失败")
		return
	}
	return
}

// 获取门票列表的函数
func GetTicketsList() (tickets TicketsList, err error) {
	res, err := connRedis.Get("TicketsList").Result()
	if err != nil {
		fmt.Println("获取门票列表失败")
		return
	}
	err = json.Unmarshal([]byte(res), &tickets)
	if err != nil {
		fmt.Println("数据反序列化失败")
		return
	}
	return
}

func GetTicketsListLen() (lenght int, err error) {
	res, err := connRedis.Get("TicketsList").Result()
	if err != nil {
		fmt.Println("获取门票列表失败")
		return
	}
	var ticketslist TicketsList
	err = json.Unmarshal([]byte(res), &ticketslist)
	if err != nil {
		fmt.Println("TicketsList数据反序列化失败")
		return
	}
	lenght = len(ticketslist.Tickets)
	return
}

func SetTicketsList(tickets TicketsList) (err error) {
	data, err := json.Marshal(tickets)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	err = connRedis.Set("TicketsList", data, 0).Err()
	if err != nil {
		fmt.Println("设置门票列表失败")
		return
	}
	return
}

func AddTickets(ticket Tickets) (err error) {
	data, err := connRedis.Get("TicketsList").Result()
	if err != nil {
		fmt.Println("获取门票列表失败")
		return
	}
	var ticketslist TicketsList
	err = json.Unmarshal([]byte(data), &ticketslist)
	if err != nil {
		fmt.Println("ticketslist数据反序列化失败")
		return
	}
	ticketslist.Tickets = append(ticketslist.Tickets, ticket)
	str, err := json.Marshal(ticketslist)
	if err != nil {
		fmt.Println("ticketslist序列化失败")
		return
	}
	err = connRedis.Set("TicketsList", str, 0).Err()
	if err != nil {
		fmt.Println("设置门票列表失败")
		return
	}
	return
}

// 获取用户订单列表的函数
func GetOrdersList(userid string) (orderlist OrdersList, err error) {
	getordersliststring, err := connRedis.HGet("User"+userid, "OrdersList").Result()
	if err != nil {
		fmt.Println("获取订单列表id失败")
		return
	}
	data, err := connRedis.Get(getordersliststring).Result()
	if err != nil {
		fmt.Println("获取订单列表失败")
		return
	}
	err = json.Unmarshal([]byte(data), &orderlist)
	if err != nil {
		fmt.Println("数据反序列化失败")
		return
	}
	return
}

func GetOrdersListLen(userid string) (lenght int, err error) {
	getordersliststring, err := connRedis.HGet("User"+userid, "OrdersList").Result()
	if err != nil {
		fmt.Println("获取订单列表id失败")
		return
	}
	data, err := connRedis.Get(getordersliststring).Result()
	if err != nil {
		fmt.Println("获取订单列表失败")
	}
	var orderlist OrdersList
	err = json.Unmarshal([]byte(data), &orderlist)
	if err != nil {
		fmt.Println("orderlist数据反序列化失败")
		return
	}
	lenght = len(orderlist.Orders)
	return
}

func SetOrdersList(userid string, orderlist OrdersList) (err error) {
	err = connRedis.HSet("User"+userid, "OrdersList", "OrdersList"+userid).Err()
	if err != nil {
		fmt.Println("设置订单列表id失败")
		return
	}
	str, err := json.Marshal(orderlist)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	//放入redis数据库时先将数据序列化后再放入数据库中
	err = connRedis.Set("OrdersList"+userid, str, 0).Err()
	if err != nil {
		fmt.Println("设置订单列表失败" + err.Error())
		return
	}
	return
}

func AddOrderList(userid string, orderprice string, ordernum string, scenicspotid string, ordertickname string) (err error) {
	orderid, err := GenerateSecureRelatedID(userid)
	if err != nil {
		fmt.Println("生成订单id失败")
		return
	}
	var order = Orders{
		Orderid:         orderid,
		Ordernum:        ordernum,
		Orderprice:      orderprice,
		Orderstatus:     "1",
		Orderuserid:     userid,
		Scenicspotid:    scenicspotid,
		OrderTicketname: ordertickname,
	}
	orderlist, err := GetOrdersList(userid)
	if err != nil {
		fmt.Println("获取用户订单列表失败")
		return
	}
	orderlist.Orders = append(orderlist.Orders, order)
	err = SetOrdersList(userid, orderlist)
	if err != nil {
		fmt.Println("设置用户订单列表失败")
		return
	}
	return
}

func GetOrderById(userid string, orderId string) (err error) {
	orderlist, err := GetOrdersList(userid)
	if err != nil {
		fmt.Println("获取用户订单列表失败")
		return
	}
	for _, v := range orderlist.Orders {
		if v.Orderid == orderId {
			err = ERROR_ORDER_HAS_EXIST
			return
		}
	}
	return
}

func GenerateSecureRelatedID(userID string) (string, error) {
	// 生成8字节的随机数
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	// 转换为uint64
	randomNum := binary.BigEndian.Uint64(b)

	// 获取时间戳
	timestamp := time.Now().UnixNano()

	// 组合ID：用户ID + 时间戳后6位 + 随机数后6位
	return fmt.Sprintf("%s%06d%06d",
		userID,
		timestamp%1000000,
		randomNum%1000000), nil
}

// 管理退款相关
func GetRefundList() (refundlist RefundsList, err error) {
	refundliststr, err := connRedis.Get("RefundsList").Result()
	if err != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	err = json.Unmarshal([]byte(refundliststr), &refundlist)
	if err != nil {
		fmt.Println("refund数据反序列化失败")
		return
	}
	return
}

func SetRefundList(refundlist RefundsList) (err error) {
	data, err := json.Marshal(refundlist)
	if err != nil {
		fmt.Println("refund序列化失败")
		return
	}
	err = connRedis.Set("RefundsList", data, 0).Err()
	if err != nil {
		fmt.Println("设置退款列表失败")
		return
	}
	return
}
func AddRefund(refund Refund) (err error) {
	refundlist, err := GetRefundList()
	if err != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	refundlist.Refunds = append(refundlist.Refunds, refund)
	err = SetRefundList(refundlist)
	if err != nil {
		fmt.Println("设置退款列表失败")
	}
	return
}

func AddRefundByUser(userid string, order Orders, refund Refund) (err error) {
	var refund1 = Refund{
		Orderid:      order.Orderid,
		Refunduserid: userid,
		Ticketname:   order.OrderTicketname,
		Refundnum:    order.Ordernum,
		Refundprice:  order.Orderprice,
		Reason:       refund.Reason,
		Date:         time.Now().Format("2006-01-02"),
		Refundstatus: "1",
	}
	refundlist, err := GetRefundList()
	if err != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	refundlist.Refunds = append(refundlist.Refunds, refund1)
	err = SetRefundList(refundlist)
	if err != nil {
		fmt.Println("设置退款列表失败")
		return
	}
	return
}
func GetRefundListById(userid string) (refundlist RefundsList, err error) {
	list, err := GetRefundList()
	if err != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	for _, v := range list.Refunds {
		if v.Refunduserid == userid {
			refundlist.Refunds = append(refundlist.Refunds, v)
		}
	}
	return
}

func GetRefundListLenByUserId(userid string) (num int, err error) {
	refundlist, err := GetRefundListById(userid)
	if err != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	num = len(refundlist.Refunds)
	return
}

// 获取评论列表的函数
func GetAllComments() (comments CommenstList, err error) {
	res, err := connRedis.Get("Comments").Result()
	if err != nil {
		fmt.Println("获取评论列表失败")
		return
	}
	err = json.Unmarshal([]byte(res), &comments)
	if err != nil {
		fmt.Println("数据反序列化失败")
		return
	}
	return
}

func SetCommentList(commentlist CommenstList) (err error) {
	data, err := json.Marshal(commentlist)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	err = connRedis.Set("Comments", data, 0).Err()
	if err != nil {
		fmt.Println("设置评论列表失败")
		return
	}
	return
}

func AddComment(comment Comments) (err error) {
	commentlist, err := GetAllComments()
	if err != nil {
		fmt.Println("获取评论列表失败")
		return
	} else {
		commentlist.Comments = append(commentlist.Comments, comment)
		err = SetCommentList(commentlist)
		if err != nil {
			fmt.Println("设置评论列表失败")
			return
		}
	}
	return
}

// 获取所有用户和订单信息
func GetAllUsaers() (userlist UserList, err error) {
	list, err := connRedis.Keys("User*").Result()
	if err != nil {
		fmt.Println("获取用户列表失败")
		return
	}
	for _, v := range list {
		var userdata UserData
		var orderslist OrdersList
		userdata, err = getUserById(v[4:])
		if err != nil {
			fmt.Println("获取用户信息失败")
			return
		}
		orderslist, err = GetOrdersList(v[4:])
		if err != nil {
			fmt.Println("获取订单列表失败")
			return
		}
		userlist.UserList = append(userlist.UserList, UserDetal{userdata, orderslist})
	}

	return
}

func ChangeUser(newuser NewUser) (err error) {

	user, err := getUserById(newuser.Userid)
	if err != nil {
		fmt.Println("获取用户信息失败")
		return err
	}
	switch newuser.Method {
	case "changeusername":
		user.Username = newuser.Username
	case "changepassword":
		user.UserPwd = newuser.UserPwd
	case "changeemail":
		user.Email = newuser.Email
	case "changephone":
		user.Phone = newuser.Phone
	default:
		fmt.Println("未知信息")
	}
	//将用户放入数据库时先将数据序列化后再放入数据库中
	userstr, err := json.Marshal(&user)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	err = connRedis.HSet("User"+user.Userid, "UserData", userstr).Err()
	if err != nil {
		fmt.Println("更新用户信息失败:", err)
	}
	return
}

func UpdataUAllUser(userlist UserList) (err error) {
	//获取当前数据库里所有user并且进行删除
	// list, err := connRedis.Keys("User*").Result()
	// if err != nil {
	// 	fmt.Println("获取用户列表失败")
	// 	return
	// }
	// for _, v := range list {
	// 	err := connRedis.Del(v).Err()
	// 	if err != nil {
	// 		fmt.Println("删除用户失败")
	// 	}
	// 	err = connRedis.Del("OrdersList" + v[4:]).Err()
	// 	if err != nil {
	// 		fmt.Println("删除订单列表失败")
	// 	}
	// }
	//删除后根据将服务端发送的用户列表以及订单列表存入数据库
	for _, v := range userlist.UserList {
		// err := Register(v.UserData)
		// if err != nil {
		// 	fmt.Println("更新用户失败")
		// }
		userstr, err := json.Marshal(&v.UserData)
		if err != nil {
			fmt.Println("序列化失败")
			return err
		}
		err = connRedis.HSet("User"+v.UserData.Userid, "UserData", userstr).Err()
		if err != nil {
			fmt.Println("更新用户信息失败:", err)
		}
		err = SetOrdersList(v.UserData.Userid, v.OrdersList)
		if err != nil {
			fmt.Println("设置订单列表失败")
		}
	}
	return
}

// 管理员登录
func AdminLogin(loginadmin Admin) (err error) {
	data, err := connRedis.Get("Admin").Result()
	if err != nil {
		fmt.Println("获取管理员信息失败")
	}
	var admin Admin
	err = json.Unmarshal([]byte(data), &admin)
	if err != nil {
		fmt.Println("数据反序列化失败")
	}
	if admin.Adminid == loginadmin.Adminid && admin.Adminpwd == loginadmin.Adminpwd {
		return nil
	} else {
		return ERROR_USER_PASSWORD
	}
}

// 实现购票逻辑
func BuyTickets(msg Msg) (orderlist OrdersList, err error) {
	var user UserData
	var tickets Tickets
	err = json.Unmarshal([]byte(msg.Field1), &user)
	if err != nil {
		fmt.Println("用户反序列化失败")
		return
	}
	err = json.Unmarshal([]byte(msg.Field2), &tickets)
	if err != nil {
		fmt.Println("门票反序列化失败")
		return
	}
	_, err = getUserById(user.Userid)
	if err != nil {
		fmt.Println("用户不存在")
		err = ERROR_USER_NOT_EXIST
		return
	}
	ticketslist, err := GetTicketsList()
	if err != nil {
		fmt.Println("获取门票列表失败")
		return
	}
	isFound := false
	ticketname := ""
	//在golang中for range是拷贝数据在进行循环并不会通过修改进行更改到原本的数组
	for i, v := range ticketslist.Tickets {
		if v.Ticketname == tickets.Ticketname {
			//可以添加其他判定 比如时间过期判定以及是否已经售罄判定
			//这里不能使用字符串的类型进行判断 使用字符串只会对头进行比较 比如 39 < 4
			// if v.Ticketnum < tickets.Ticketnum {
			// 	fmt.Println("门票数量不足")
			// 	err = ERROR_TICKETS_NOT_ENOUGH
			// 	return
			// }
			if numa, erra := strconv.Atoi(v.Ticketnum); erra != nil {
				if numb, errb := strconv.Atoi(tickets.Ticketnum); errb != nil {
					if numa < numb {
						fmt.Println("门票数量不足")
						err = ERROR_TICKETS_NOT_ENOUGH
						return
					}
				}
			}
			// v.Ticketnum = fmt.Sprint(func() int { a, _ := strconv.Atoi(v.Ticketnum); b, _ := strconv.Atoi(tickets.Ticketnum); return a - b }())

			ticketslist.Tickets[i].Ticketnum = strconv.Itoa(func() int { a, _ := strconv.Atoi(v.Ticketnum); b, _ := strconv.Atoi(tickets.Ticketnum); return a - b }())
			ticketname = v.Ticketname
			isFound = true
			if v.Ticketnum == "0" {
				v.Ticketrefund = "true"
			}
			break
		}
	}
	if !isFound {
		fmt.Println("门票不存在")
		err = ERROR_TICKETS_NOT_EXIST
		return
	}
	err = SetTicketsList(ticketslist)
	if err != nil {
		fmt.Println("设置门票列表失败")
		return
	}
	err = AddOrderList(user.Userid, tickets.Ticketprice, tickets.Ticketnum, tickets.Scenicspotid, ticketname)
	if err != nil {
		fmt.Println("添加订单失败")
		return
	}
	orderlist, err = GetOrdersList(user.Userid)
	if err != nil {
		fmt.Println("获取订单列表失败")
	}
	return
}

// 实现支付功能
func PayOrder(msg Msg) (orderlist OrdersList, err error) {
	var userid string
	var order Orders
	err = json.Unmarshal([]byte(msg.Field1), &userid)
	if err != nil {
		fmt.Println("userid反序列化失败")
		return
	}
	err = json.Unmarshal([]byte(msg.Field2), &order)
	if err != nil {
		fmt.Println("订单反序列化失败")
		return
	}
	orderlist, err = GetOrdersList(userid)
	if err != nil {
		fmt.Println("获取订单列表失败")
		return
	}
	hasexit := false
	for i, v := range orderlist.Orders {
		if v.Orderid == order.Orderid {
			orderlist.Orders[i].Orderstatus = "2"
			hasexit = true
			break
		}
	}
	if !hasexit {
		fmt.Println("订单不存在")
		orderlist = OrdersList{}
		err = ERROR_ORDER_HAS_NOT_EXIST
		return
	}
	err = SetOrdersList(userid, orderlist)
	if err != nil {
		fmt.Println("设置订单列表失败")
	}
	return
}

// 实现申请退款功能
func RefundOrder(msg Msg) (refundlist RefundsList, err error) {
	var userid string
	var order Orders
	var refund Refund
	err = json.Unmarshal([]byte(msg.Field1), &userid)
	if err != nil {
		fmt.Println("userid反序列化失败")
		return
	}
	if userid == "" {
		fmt.Println("userid为空")
		err = ERROR_USER_NOT_EXIST
		return
	}
	err = json.Unmarshal([]byte(msg.Field2), &order)
	if err != nil {
		fmt.Println("订单反序列化失败")
		return
	}
	err = json.Unmarshal([]byte(msg.Field3), &refund)
	if err != nil {
		fmt.Println("退款单反序列化失败")
		return
	}
	err = AddRefundByUser(userid, order, refund)
	if err != nil {
		fmt.Println("添加退款单失败")
		return
	}
	refundlist, err = GetRefundListById(userid)
	if err != nil {
		fmt.Println("获取退款列表失败")
	}
	orderlist, err := GetOrdersList(userid)
	if err != nil {
		fmt.Println("获取订单列表失败")
		return
	}
	for i, v := range orderlist.Orders {
		if v.Orderid == order.Orderid {
			orderlist.Orders[i].Orderstatus = "3"
			break
		}
	}
	err = SetOrdersList(userid, orderlist)
	if err != nil {
		fmt.Println("设置订单列表失败")
	}
	return
}

// 实现取消退款功能
func CancelRefundOrder(msg Msg) (refundlist RefundsList, err error) {
	var userid string
	var refund Refund
	err = json.Unmarshal([]byte(msg.Field1), &userid)
	if err != nil {
		fmt.Println("userid反序列化失败")
		return
	}
	err = json.Unmarshal([]byte(msg.Field2), &refund)
	if err != nil {
		fmt.Println("订单反序列化失败")
		return
	}
	refundalllist, err := GetRefundList()
	if err != nil {
		fmt.Println("获取退款列表失败")
		return
	}
	orderlist, err := GetOrdersList(refund.Refunduserid)
	if err != nil {
		fmt.Println("获取订单列表失败")
		return
	}
	for i, v := range refundalllist.Refunds {
		if v.Orderid == refund.Orderid {
			refundalllist.Refunds[i].Refundstatus = "3"
			break
		}
	}
	for i, v := range orderlist.Orders {
		if v.Orderid == refund.Orderid {
			orderlist.Orders[i].Orderstatus = "2"
			break
		}
	}
	err = SetRefundList(refundalllist)
	if err != nil {
		fmt.Println("设置退款列表失败")
		return
	}
	err = SetOrdersList(refund.Refunduserid, orderlist)
	if err != nil {
		fmt.Println("设置订单列表失败")
	}
	return
}
