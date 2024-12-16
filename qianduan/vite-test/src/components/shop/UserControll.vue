<template>
    <div class="user-management">
        <h1>用户管理</h1>
        <div class="user-info">
            <p><strong>用户ID:</strong> {{ store.state.mainuser.user.userid }}</p>
            <p><strong>用户名:</strong> {{ store.state.mainuser.user.username }}</p>
        </div>

        <div class="action-selection">
            <el-radio-group v-model="showView">
                <el-radio label="changeusername" border>更改姓名</el-radio>
                <el-radio label="changepassword" border>更改密码</el-radio>
                <el-radio label="changephone" border>更改电话</el-radio>
                <el-radio label="changeemail" border>更改邮箱</el-radio>
            </el-radio-group>
        </div>

        <div class="change-form" v-if="showView === 'changeusername'">
            <h2>更改姓名</h2>
            <el-input v-model="changeUser.username" placeholder="请输入新名字" prefix-icon="el-icon-user" />
            <el-button type="primary" @click="submitChange">提交更改</el-button>
        </div>

        <div class="change-form" v-if="showView === 'changepassword'">
            <h2>更改密码</h2>
            <el-input v-model="changeUser.userpwd" type="password" placeholder="请输入新密码" show-password
                prefix-icon="el-icon-lock" />
            <el-button type="primary" @click="submitChange">提交更改</el-button>
        </div>

        <div class="change-form" v-if="showView === 'changephone'">
            <h2>更改电话</h2>
            <el-input v-model="changeUser.phone" type="tel" placeholder="请输入新电话" prefix-icon="el-icon-phone" />
            <el-button type="primary" @click="submitChange">提交更改</el-button>
        </div>

        <div class="change-form" v-if="showView === 'changeemail'">
            <h2>更改邮箱</h2>
            <el-input v-model="changeUser.email" type="email" placeholder="请输入新邮箱" prefix-icon="el-icon-email" />
            <el-button type="primary" @click="submitChange">提交更改</el-button>
        </div>

        <el-dialog :title="dialogTitle" :visible.sync="dialogVisible" width="30%" :before-close="handleClose">
            <span>{{ dialogMessage }}</span>
            <span slot="footer" class="dialog-footer">
                <el-button @click="dialogVisible = false">关闭</el-button>
            </span>
        </el-dialog>
    </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { useStore } from 'vuex';
import axios from 'axios';

const store = useStore();
const showView = ref('1');

const changeUser = ref({
    method: '',
    userid: '',
    userpwd: '',
    username: '',
    phone: '',
    email: ''
});

const dialogVisible = ref(false);
const dialogMessage = ref('');
let dialogTitle = "";

const submitChange = async () => {
    dialogTitle = showView.value === '1' ? '更改姓名' : '更改密码';
    changeUser.value.method = showView.value;
    changeUser.value.userid = store.state.mainuser.user.userid;
    try {
        const response = await axios.post('/api/upload/changeuser', changeUser.value);
        console.log(response);

        if (response.data.code === 200) {
            dialogMessage.value = '更改成功！';
            dialogVisible.value = true;
            // 如果是更改用户名，更新store中的用户名
            console.log(response.data);
            switch (response.data.user.method) {
                case "changeusername":
                    store.state.mainuser.user.username = response.data.user.username;
                    break;
                case "changepassword":
                    store.state.mainuser.user.userpwd = response.data.user.userpwd;
                    break;
                case "changephone":
                    store.state.mainuser.user.phone = response.data.user.phone;
                    break;
                case "changeemail":
                    store.state.mainuser.user.email = response.data.user.email;
            }
            store.state.mainuser.user = response.data.user;
            store.commit('saveUserDataToStorage');
            // 清空输入
            changeUser.value.username = '';
            changeUser.value.userpwd = '';
        }
    } catch (error) {
        console.error('Error sending message:', error);
        dialogMessage.value = '更改失败，请稍后重试。';
        dialogVisible.value = true;
    }
};

const handleClose = (done: () => void) => {
    done();
};
</script>

<style scoped>
.user-management {
    max-width: 500px;
    margin: 0 auto;
    padding: 20px;
    background-color: #f5f7fa;
    border-radius: 8px;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
}

h1 {
    color: #303133;
    border-bottom: 2px solid #409EFF;
    padding-bottom: 10px;
    margin-bottom: 20px;
}

.user-info {
    background-color: #0e29c2;
    padding: 15px;
    border-radius: 4px;
    margin-bottom: 20px;
}

.user-info p {
    margin: 5px 0;
}

.action-selection {
    margin-bottom: 20px;
}

.change-form {
    background-color: #fff;
    padding: 20px;
    border-radius: 4px;
}

h2 {
    color: #606266;
    margin-bottom: 15px;
}

.el-input {
    margin-bottom: 15px;
}

.el-button {
    width: 100%;
}
</style>