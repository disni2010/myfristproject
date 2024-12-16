<template>
    <div class="page-container">
        <div id="shouye">
            <h3 class="title">登录</h3>
            <form @submit.prevent="sumbitLogin" class="login-form">
                <div class="form-group">
                    <label>用户名</label>
                    <input type="text" v-model.lazy="$store.state.mainuser.user.username" class="form-input"
                        placeholder="请输入用户名">
                </div>
                <div class="form-group">
                    <label>密码</label>
                    <input type="password" v-model.lazy="$store.state.mainuser.user.userpwd" class="form-input"
                        placeholder="请输入密码">
                </div>
                <button type="submit" class="submit-btn">登录</button>
                <router-link to="/register" class="login-link">没有账号，前往注册</router-link>
            </form>
        </div>
    </div>
</template>

<script>

export default {

    data() {
        return {
            package: {
                userid: '',
                orderslistlen: 0,
                scenicspotslistlen: 0,
                ticketslistlen: 0,
                refundslistlen: 0,
            }
        }
    },

    methods: {
        async sumbitLogin() {
            try {
                //发送数据给后台并等待接收返回
                const response = await this.$axios.post('/api/upload/login', this.$store.state.mainuser.user)
                this.response = response.data
                // if (response.data.code == 200) {
                //     //这里将返回的用户数据进行持久化存储直到退出登录或重启服务器
                //     window.sessionStorage.setItem('mainUser', JSON.stringify(response.data.data))
                //     this.$store.state.mainuser.user = response.data.data
                //     this.getdata()
                // } else {
                //     console.log("登陆失败");
                //     console.log(response.data);
                // }
                //持久化存储
                //localStorage.setItem('zheng', str);
                if (response.data.code == 200) {
                    alert("登录成功")
                    window.sessionStorage.setItem('mainUser', JSON.stringify(response.data.data))
                    this.$store.state.mainuser.user = response.data.data
                    this.getdata()
                } else {
                    alert("登录失败:" + response.data.msg);
                    console.log(response.data);
                }
            }
            catch (error) {
                console.error('Error sending message:', error);
                this.response = 'An error occurred while sending the message.';
            }
        },
        async getdata() {
            this.package.userid = this.$store.state.mainuser.user.userid
            this.package.orderslistlen = this.$store.state.mainuser.orders?.length || 0
            this.package.scenicspotslistlen = this.$store.state.Scenicspots?.length || 0
            this.package.ticketslistlen = this.$store.state.Tickets?.length || 0
            this.package.refundslistlen = this.$store.state.Refunds?.length || 0
            try {
                //发送数据给后台并等待接收返回
                const response = await this.$axios.post('/api/upload/getotherdata', this.package)
                this.response = response.data
                console.log(response.data);
                if (response.data.code != 500) {
                    //这里需要将等待数据存入后再跳转页面
                    // window.sessionStorage.setItem("ordersList", JSON.stringify(response.data.msg.orderslist.orders))
                    // window.localStorage.setItem("ticketsList", JSON.stringify(response.data.msg.ticketslist.tickets))
                    // window.localStorage.setItem("scenicspotsList", JSON.stringify(response.data.msg.scenicspotlist.scenicspots))
                    // this.$router.push('/shopmain')
                    Promise.all([
                        //如果响应数据中的订单列表不为空，那么就将这些订单数据存储到sessionStorage中，并且当存储操作完成后
                        //s解决Promise。如果订单列表为空，那么直接返回一个已经解决的Promise，从而不执行任何存储操作。
                        (response.data.msg.orderslist?.orders || false && Array.isArray(response.data.msg.orderslist.orders) && response.data.msg.orderslist.orders.length > 0 ?
                            new Promise(resolve => {
                                window.sessionStorage.setItem("ordersList", JSON.stringify(response.data.msg.orderslist.orders));
                                resolve();
                            }) : Promise.resolve()),
                        (response.data.msg.ticketslist?.tickets || false && Array.isArray(response.data.msg.ticketslist.tickets) && response.data.msg.ticketslist.tickets.length > 0 ?
                            new Promise(resolve => {
                                window.localStorage.setItem("ticketsList", JSON.stringify(response.data.msg.ticketslist.tickets));
                                resolve();
                            }) : Promise.resolve()),
                        (response.data.msg.scenicspotlist?.scenicspots || false && Array.isArray(response.data.msg.scenicspotlist.scenicspots) && response.data.msg.scenicspotlist.scenicspots.length > 0 ?
                            new Promise(resolve => {
                                window.localStorage.setItem("scenicspotsList", JSON.stringify(response.data.msg.scenicspotlist.scenicspots));
                                resolve();
                            }) : Promise.resolve()),
                        (response.data.msg.refundslist?.refunds || false && Array.isArray(response.data.msg.refundslist.refunds) && response.data.msg.refundslist.refunds.length > 0 ?
                            new Promise(resolve => {
                                window.sessionStorage.setItem("refundsList", JSON.stringify(response.data.msg.refundslist.refunds));
                                resolve();
                            }) : Promise.resolve()),
                    ]).then(() => {
                        // 所有异步操作都完成后执行跳转页面
                        this.$router.push('/shopmain')
                    });
                } else {
                    alert("登录失败")
                    console.log(response.data);
                }
            }
            catch (error) {
                console.error('Error sending message:', error);
                this.response = 'An error occurred while sending the message.';
            }
        }
    }
}

</script>

<style scoped>
.page-container {
    position: absolute;
    display: flex;
    justify-content: center;
    align-items: center;
    bottom: 28%;
    left: 38%;
}

#shouye {
    border-radius: 30px;
    height: 400px;
    width: 350px;
    border: none;
    padding: 40px;
    background: rgba(255, 255, 255, 0.9);
    box-shadow: 0 8px 32px rgba(31, 38, 135, 0.15);
    backdrop-filter: blur(4px);
    transition: transform 0.3s ease;
}

#shouye:hover {
    transform: translateY(-5px);
}

.title {
    color: #1a5f7a;
    text-align: center;
    font-size: 32px;
    margin-bottom: 40px;
    font-weight: 600;
    text-shadow: 1px 1px 2px rgba(0, 0, 0, 0.1);
}

.login-form {
    display: flex;
    flex-direction: column;
    gap: 20px;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 8px;
}

.form-group label {
    color: #1a5f7a;
    font-weight: 500;
    font-size: 16px;
}

.form-input {
    padding: 12px;
    border: 2px solid #e0f7ff;
    border-radius: 12px;
    font-size: 16px;
    transition: all 0.3s ease;
    outline: none;
}

.form-input:focus {
    border-color: #3ca8df;
    box-shadow: 0 0 0 2px rgba(60, 168, 223, 0.2);
}

.submit-btn {
    margin-top: 20px;
    padding: 12px 40px;
    background: linear-gradient(145deg, #3ca8df, #2980b9);
    color: white;
    border: none;
    border-radius: 25px;
    font-weight: 500;
    font-size: 18px;
    cursor: pointer;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(60, 168, 223, 0.2);
}

.submit-btn:hover {
    background: linear-gradient(145deg, #2980b9, #3ca8df);
    transform: scale(1.05);
    box-shadow: 0 6px 20px rgba(60, 168, 223, 0.3);
}

.login-link {
    text-align: center;
    color: #1a5f7a;
    text-decoration: none;
    font-size: 14px;
    transition: all 0.3s ease;
}

.login-link:hover {
    color: #3ca8df;
    text-decoration: underline;
}
</style>