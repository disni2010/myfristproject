import { createStore } from 'vuex'
//1. 不加花括号：导入整个模块对象。例如，import axios from 'axios'会导入整个axios模块，可以通过axios.get()等方法来使用它。
// 2. 加上花括号：只导入模块中的指定变量或函数。例如，import { get, post } from
//     'axios'只导入了axios模块中的get和post方法，可以直接使用get()和post()调用
// 它们。
// 需要注意的是，如果导入的模块没有默认导出（即没有export default语句），则必须使用花括号
// 来指定导入的变量或函数。如果导入的模块有默认导出，则可以使用不加花括号的语法来导入整个
// 模块对象
import { defineStore } from 'pinia'


//用户数据 User username(用户名) userpwd(用户密码) userid(用户id)  phone(手机) email(邮箱)
//票务数据 Tickets ticketname(票务名称) ticketprice(票务价格) ticketnum(票务数量) ticketrefund(是否售罄) scenicspotid(景区id) enddate(结束日期)
//订单数据 Orders orderid(订单id)  orderprice(订单价格) ordernum(订单数量) orderuserid(订单用户id) orderstatus(订单状态:1未完成2已完成3退款中4退款完成) ticketname(票务名称)
//景点数据 Scenicspots scenicspotid(景区id) scenicspotname(景区名称) scenicspotdetails(景区介绍)  opentime(开放时间) closetime(闭馆时间)
//反馈数据 Comments userid(用户id) comment(评论) date(评论日期) rating (评分)
//退款数据 Refunds orderid(订单id) refundprice(退款价格) refundnum(退款数量) refunduserid(退款用户id) refundstatus(退款状态:1申请退款2退款成功3取消退款4退款驳回) ticketname(票务名称) Data(日期) reason(退款原因) Note(备注)
export default createStore({
    // 定义： 存储全局的状态。
    // 调用方式： 在 Vue 组件中可以通过 this.$store.state 访问
    state: {
        mainuser: {
            user: {
                userid: '',
                userpwd: '',
                username: '',
                phone: '',
                email: '',
            },
            orders: [],
        },
        Scenicspots: [],
        Tickets: [],
        Refunds: []
    },
    // 定义： 用于同步地更新 state。
    // 调用方式： 在组件中使用 this.$store.commit 调用
    // 为什么要这样声明： mutations 设计为同步操作，便于调试工具追踪状态的变化。
    mutations: {
        setUser(state, user) {
            state.mainuser.user = user;
        },
        setOrdersList(state, ordersList) {
            state.mainuser.orders = ordersList;
        },
        setScenicspotsList(state, Scenicspots) {
            state.Scenicspots = Scenicspots;
        },
        setTicketsList(state, Tickets) {
            state.Tickets = Tickets;
        },
        setRefundsList(state, Refunds) {
            state.Refunds = Refunds;
        },
    },
    // 定义： 用于处理异步操作（如网络请求），并最终通过 mutations 修改状态。
    // 调用方式： 在组件中使用 this.$store.dispatch 调用
    // 为什么要这样声明： actions 的异步特性允许将复杂的逻辑与状态修改分离。最终的状态修改依然通过 mutations，保持单一数据流的设计。
    actions: {
        //actions 主要用于处理异步操作，比如网络请求、延时任务等。actions 不应该直接修改状态，而是应该通过提交 mutations 来更新状态。
        loadUserDataFromStorage({ commit }) {
            const storedUser = sessionStorage.getItem('mainUser');
            if (storedUser) {
                commit('setUser', JSON.parse(storedUser));
            }
        },
        loadOrdersListFromStorage({ commit }) {
            const storedOrdersList = sessionStorage.getItem('ordersList');
            if (storedOrdersList) {
                commit('setOrdersList', JSON.parse(storedOrdersList));
            }
        },
        loadRefundsListFromStorage({ commit }) {
            const Refunds = sessionStorage.getItem('refundsList');
            if (Refunds) {
                commit('setRefundsList', JSON.parse(Refunds));
            } else {
                console.log("空数据refundsList");
            }
        },
        loadScenicspotsListFromStorage({ commit }) {
            const Scenicspots = localStorage.getItem('scenicspotsList');
            if (Scenicspots) {
                commit('setScenicspotsList', JSON.parse(Scenicspots));
            } else {
                console.log("空数据scenicspotsList");
            }
        },
        loadTicketsListFromStorage({ commit }) {
            const Tickets = localStorage.getItem('ticketsList');
            if (Tickets) {
                commit('setTicketsList', JSON.parse(Tickets));
            } else {
                console.log("空数据ticketsList");
            }
        },
        //在 Vuex 的 actions 中，你可以直接访问 this._vm（即 store 的内部 Vue 实例）来修改状态，但这并不是推荐的做法。尽管技术上可行，但这样做违反了 Vuex 的设计原则和最佳实践。
        saveUserDataToStorage({ state }) {
            sessionStorage.setItem('mainUser', JSON.stringify(state.mainuser.user));
        },
        saveOrdersListToStorage({ state }) {
            sessionStorage.setItem('ordersList', JSON.stringify(state.mainuser.orders));
        },
        saveRefundsListToStorage({ state }) {
            sessionStorage.setItem('refundsList', JSON.stringify(state.Refunds));
        },
        saveScenicspotsListToStorage({ state }) {
            localStorage.setItem('scenicspotsList', JSON.stringify(state.Scenicspots));
        },
        saveTicketsListToStorage({ state }) {
            localStorage.setItem('ticketsList', JSON.stringify(state.Tickets));
        },
    }
})
//定义并导出容器 参数1容器id 必须唯一 将来pinia内部会挂载到根容器上 参数2 选项对象
//返回值 一个函数 调用得到容器实例
export const useMainStore = defineStore('main', {
    //类似于组件的data 用来存储全局状态的
    //必须是函数：这样是为了服务端渲染时候避免交叉请求导致的数据状态污染
    //必须是箭头函数 这是为了更好的ts类型推导
    state: () => {
        return {
            count: 129,
            foo: 'bar'
        }
    },
    //类似于组件 computed 用来封装计算属性 有缓存功能的
    getters: {
        count10(state) {
            return state.count + 10
        },
        //如果在getter中使用this 则必须手动指定返回值的类型 否则类型推导不出来
        // count10(): number {
        //     return this.count + 10
        // }
        GetTicketsLen() {
            return this.Tickets.length
        },
        GetScenicspotsLen() {
            return this.Scenicspots.length
        },
        GetOrderLen() {
            return this.mainuser.orders.length
        },
        GetUserID() {
            return this.mainuser.user.userid
        }

    },
    //类似组件的methods 封装业务逻辑 修改state
    actions: {
        inc() {
            this.count++
        },
        //不能使用箭头函数修饰 箭头函数指向外部this
        changestate(num) {
            this.count += num
            this.foo = 'hello'
        },
    }
}, { persist: true, })