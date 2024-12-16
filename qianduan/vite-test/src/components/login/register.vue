<template>
    <div class="page-container">
        <div id="shouye">
            <h3 class="title">注册</h3>
            <form @submit.prevent="sumbitRegister" class="login-form">
                <div class="form-group">
                    <label>用户名</label>
                    <input type="text" v-model.lazy="$store.state.mainuser.user.username" class="form-input"
                        placeholder="请输入想创建的用户名">
                </div>
                <div class="form-group">
                    <label>密码</label>
                    <input type="password" v-model.lazy="$store.state.mainuser.user.userpwd" class="form-input"
                        placeholder="请输入想创建的密码">
                </div>
                <div class="form-group">
                    <label>电话</label>
                    <input type="tel" v-model.lazy="$store.state.mainuser.user.phone" class="form-input"
                        placeholder="请输入想创建的手机号">
                </div>
                <div class="form-group">
                    <label>邮箱</label>
                    <input type="email" v-model.lazy="$store.state.mainuser.user.email" class="form-input"
                        placeholder="请输入想创建的邮箱">
                </div>
                <button type="submit" class="submit-btn">注册</button>
                <router-link to="/login" class="login-link">已有账号，返回登录</router-link>
            </form>
        </div>
    </div>
</template>

<script>
export default {
    methods: {
        async sumbitRegister() {
            try {
                //Vue 组件的方法中直接使用 $store.state.user 会报错，因为 $store 不是直接在组件实例上下文中可用的。Vue 组件中访问 Vuex store 的正确方式是通过 this.$store
                //发送数据给后台并等待接收返回
                const response = await this.$axios.post('/api/upload/register', this.$store.state.mainuser.user)
                this.response = response.data
                console.log(response);
                if (response.data.code == 200) {
                    //这里将返回的用户数据进行持久化存储直到退出登录或重启服务器
                    alert("注册成功")
                    this.$router.push('/')
                } else {
                    alert("注册失败:"+response.data.msg);
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
    bottom: 16%;
    left: 38%;
}

#shouye {
    border-radius: 30px;
    min-height: 500px;
    /* 增加高度以适应更多的表单项 */
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