package model

//用户数据 User username(用户名) userpwd(用户密码) userid(用户id)  phone(手机) email(邮箱)
//票务数据 Tickets ticketname(票务名称) ticketprice(票务价格) ticketnum(票务数量) ticketrefund(是否售罄) scenicspotid(景区id) enddate(结束日期)
//订单数据 Orders orderid(订单id)  orderprice(订单价格) ordernum(订单数量) orderuserid(订单用户id) orderstatus(订单状态:1未完成2已完成3退款中4退款完成) ticketname(票务名称)
//景点数据 Scenicspots scenicspotid(景区id) scenicspotname(景区名称) scenicspotdetails(景区介绍)  opentime(开放时间) closetime(闭馆时间)
//反馈数据 Comments userid(用户id) comment(评论) date(评论日期) rating (评分)
//退款数据 Refunds orderid(订单id) refundprice(退款价格) refundnum(退款数量) refunduserid(退款用户id) refundstatus(退款状态:1申请退款2退款成功3取消退款4退款驳回) ticketname(票务名称) Data(日期) reason(退款原因) Note(驳回原因)

//数据结构 用户id{用户数据 订单数据key} 订单数据 景点数据  票务数据 反馈数据

//获取数据 hget 用户id 用户数据/订单数据

type UserData struct {
	Userid   string `json:"userid"`   //用户id
	Username string `json:"username"` //用户名
	UserPwd  string `json:"userpwd"`  //用户密码
	Phone    string `json:"phone"`    //手机号
	Email    string `json:"email"`    //邮箱
}

type Tickets struct {
	Ticketname   string `json:"ticketname"`   //票务名称
	Ticketprice  string `json:"ticketprice"`  //票务价格
	Ticketnum    string `json:"ticketnum"`    //票务数量
	Ticketrefund string `json:"ticketrefund"` //是否售罄
	Scenicspotid string `json:"scenicspotid"` //景区id
	StartDate    string `json:"startDate"`    //开始日期
	EndDate      string `json:"endDate"`      //结束日期
}

type Orders struct {
	Orderid         string `json:"orderid"`         //订单id
	OrderTicketname string `json:"orderticketname"` //票务名称
	Orderprice      string `json:"orderprice"`      //订单价格
	Ordernum        string `json:"ordernum"`        //订单数量
	Orderuserid     string `json:"orderuserid"`     //订单用户id
	Orderstatus     string `json:"orderstatus"`     //订单状态
	Scenicspotid    string `json:"scenicspotid"`    //景区id
}
type Scenicspots struct {
	Scenicspotid      string `json:"scenicspotid"`      //景区id
	Scenicspotname    string `json:"scenicspotname"`    //景区名称
	Scenicspotdetails string `json:"scenicspotdetails"` //景区介绍
	OpenTime          string `json:"opentime"`          //开放时间
	CloseTime         string `json:"closetime"`         //闭馆时间
}
type Refund struct {
	Orderid      string `json:"orderid"`      //订单id
	Refundprice  string `json:"refundprice"`  //退款价格
	Refundnum    string `json:"refundnum"`    //退款数量
	Refunduserid string `json:"refunduserid"` //退款用户id
	Refundstatus string `json:"refundstatus"` //退款状态
	Ticketname   string `json:"ticketname"`   //票务名称
	Date         string `json:"date"`         //日期
	Reason       string `json:"reason"`       //退款原因
	Note         string `json:"note"`         //备注
}

type ScenicspotsList struct {
	Scenicspots []Scenicspots `json:"scenicspots"`
}
type TicketsList struct {
	Tickets []Tickets `json:"tickets"`
}
type OrdersList struct {
	Orders []Orders `json:"orders"`
}
type RefundsList struct {
	Refunds []Refund `json:"refunds"`
}

type Package struct {
	Userid             string          `json:"userid"`
	Orderslist         OrdersList      `json:"orderslist"`
	Scenicspotlist     ScenicspotsList `json:"scenicspotlist"`
	Ticketslist        TicketsList     `json:"ticketslist"`
	Refundlist         RefundsList     `json:"refundslist"`
	Orderslistlen      int             `json:"orderslistlen"`
	Scenicspotslistlen int             `json:"scenicspotslistlen"`
	Ticketslistlen     int             `json:"ticketslistlen"`
	Refundlistlen      int             `json:"refundlistlen"`
}

type CommenstList struct {
	Comments []Comments `json:"comments"`
}

type Comments struct {
	Userid  string `json:"userid"`  //用户id
	Comment string `json:"comment"` //评论
	Date    string `json:"date"`    //评论日期
	Rating  string `json:"rating"`  //评分
}

type NewUser struct {
	Method   string `json:"method"` //方法
	Userid   string `json:"userid"`
	Username string `json:"username"`
	UserPwd  string `json:"userpwd"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

//添加管理员
type Admin struct {
	Adminid  string `json:"adminid"`  //管理员id
	Adminpwd string `json:"adminpwd"` //管理员密码
	IsUse    bool   `json:"isuse"`    //是否启用
}

//所有用户列表
type UserList struct {
	UserList []UserDetal `json:"userlist"`
}

type UserDetal struct {
	UserData   UserData   `json:"userdata"`
	OrdersList OrdersList `json:"orderslist"`
}

type Msg struct {
	Field1 string `json:"field1"`
	Field2 string `json:"field2"`
	Field3 string `json:"field3"`
	Field4 string `json:"field4"`
	Field5 string `json:"field5"`
}

//  package: {
//                 userid: '',
//                 orderslist: [],
//                 scenicspotsList: [],
//                 ticketslist: [],
//                 orderslistlen: 0,
//                 scenicspotslistlen: 0,
//                 ticketslistlen: 0
//             }

// 数据结构这里创建就错了
// hgetall User3
// 1) "UserData"
// 2) "{\"userid\":\"3\",\"username\":\"Charlie\",\"userpwd\":\"password3\"}"
// 3) "OrdersList"
// 4) "OrdersList3"
// 在redis里数据是这样 但浪费了hash的结构 可以使用列表结构同时方便进行管理
