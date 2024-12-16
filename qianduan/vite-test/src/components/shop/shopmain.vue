<template>
    <nav id="shopmain">
        <div class="nav-left">
            <h2 id="shagnchengshouye">商城首页</h2>
            <div class="nav-links">
                <div class="nav-item">
                    <router-link to="/shopmain/UserControll">用户管理</router-link>
                </div>

                <div class="nav-item">
                    <el-dropdown>
                        <span class="el-dropdown-link">
                            <h4>购票系统</h4>
                        </span>
                        <template #dropdown>
                            <el-dropdown-menu>
                                <el-dropdown-item>
                                    <router-link to="/shopmain/shopTickets/goupiaoliebiao">购票列表</router-link>
                                </el-dropdown-item>
                                <el-dropdown-item>
                                    <router-link to="/shopmain/shopTickets/goupiaochaxun">购票查询</router-link>
                                </el-dropdown-item>
                                <el-dropdown-item>
                                    <router-link to="/shopmain/shopTickets/tuikuanchaxun">退款查询</router-link>
                                </el-dropdown-item>
                            </el-dropdown-menu>
                        </template>
                    </el-dropdown>
                </div>

                <div class="nav-item">
                    <router-link to="/shopmain/ScenicSpots">景区介绍</router-link>
                </div>

                <div class="nav-item">
                    <router-link to="/shopmain/Commet">留言反馈</router-link>
                </div>
            </div>
        </div>

        <div class="nav-right">
            <div class="user-welcome">
                欢迎您, {{ this.$store.state.mainuser.user.username }}
            </div>

            <form id="search-form">
                <input type="text" placeholder="搜索..." class="search-input">
                <button type="submit" class="search-button">搜索</button>
            </form>
        </div>
    </nav>

    <router-view></router-view>
</template>


<script>
import { ArrowDown } from '@element-plus/icons-vue'
import { mapActions } from 'vuex'

export default {
    mounted() {
        //等待数据加载完成再 执行
        // console.log("读取存放在storage后并转为数据的");
        // console.log(this.$store.state.mainuser.user);
        // console.log(this.$store.state.orderslist);
        // console.log(this.$store.state.scenicspotlist);
        // console.log(this.$store.state.ticketslist);
        Promise.all([
            this.$store.dispatch('loadUserDataFromStorage'),
            this.$store.dispatch('loadOrdersListFromStorage'),
            this.$store.dispatch('loadScenicspotsListFromStorage'),
            this.$store.dispatch('loadTicketsListFromStorage'),
            this.$store.dispatch('loadRefundsListFromStorage')
        ]).then(() => {
            //console.log("读取存放在storage后并转为数据的");
            console.log(this.$store.state.mainuser.user);
            console.log(this.$store.state.mainuser.orders);
            console.log(this.$store.state.Scenicspots);
            console.log(this.$store.state.Tickets);
            console.log(this.$store.state.refundslist);
        }).catch(error => {
            console.error("数据加载错误: ", error);
        });
        //这样写法在数据写入前就已经调用了
        // this.$store.dispatch('loadUserDataFromStorage')
        // this.$store.dispatch('loadOrdersListFromStorage')
        // this.$store.dispatch('loadScenicspotsListFromStorage')
        // this.$store.dispatch('loadTicketsListFromStorage')
        // console.log("读取存放在storage后并转为数据的");
        // console.log(this.$store.state.mainuser.user);
        // console.log(this.$store.state.mainuser.orderslist);
        // console.log(this.$store.state.scenicspotlist);
        // console.log(this.$store.state.ticketslist);
    },
}



// const { getDataFromStorage } = mapActions(['getDataFromStorage'])
// // 在 mounted 钩子中调用方法
// const mounted = () => {
//     getDataFromStorage()
//     console.log(this.$store.state.mainuser.user);
//     console.log(this.$store.state.orderslist);
//     console.log(this.$store.state.scenicspotlist);
//     console.log(this.$store.state.ticketslist);
// }
// // 在组件挂载时执行
// mounted()
// export default {
//     methods: {
//         ...mapActions(['getDataFromStorage']),
//     },
//     mounted() {
//         this.getDataFromStorage()
//         console.log(this.$store.state.mainuser.user);
//         console.log(this.$store.state.orderslist);
//         console.log(this.$store.state.scenicspotlist);
//         console.log(this.$store.state.ticketslist);
//     },
// }
</script>
<style scoped>
#shopmain {
    width: 95vw;
    /* 使用视口宽度 */
    position: relative;
    left: 49%;
    right: 50%;
    margin-left: -50vw;
    margin-right: -50vw;
    padding: 1rem 4rem;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #fff;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
    top: -2vw;
}

.nav-left {
    display: flex;
    align-items: center;
    gap: 2rem;
}

#shagnchengshouye {
    margin: 0;
    font-size: 1.5rem;
    color: #333;
    font-weight: 600;
}

.nav-links {
    display: flex;
    gap: 1.5rem;
    align-items: center;
}

.nav-item a {
    text-decoration: none;
    color: #666;
    font-size: 1rem;
    padding: 0.5rem 1rem;
    border-radius: 4px;
    transition: all 0.3s ease;
}

.nav-item a:hover {
    color: #409EFF;
    background-color: rgba(64, 158, 255, 0.1);
}

.el-dropdown-link {
    cursor: pointer;
    color: #666;
    display: flex;
    align-items: center;
    padding: 0.5rem 1rem;
    transition: all 0.3s ease;
}

.el-dropdown-link:hover {
    color: #409EFF;
    background-color: rgba(64, 158, 255, 0.1);
}

.el-dropdown-link h4 {
    margin: 0;
    font-size: 1rem;
    font-weight: bold;
}

.nav-right {
    display: flex;
    align-items: center;
    gap: 1.5rem;
}

.user-welcome {
    color: #666;
    font-size: 0.9rem;
}

#search-form {
    display: flex;
    gap: 0.5rem;
}

.search-input {
    padding: 0.5rem 1rem;
    border: 1px solid #dcdfe6;
    border-radius: 4px;
    outline: none;
    transition: all 0.3s ease;
    width: 200px;
}

.search-input:focus {
    border-color: #409EFF;
    box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
}

.search-button {
    padding: 0.5rem 1rem;
    background-color: #409EFF;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    transition: all 0.3s ease;
}

.search-button:hover {
    background-color: #66b1ff;
}



/* 响应式设计 */
@media (max-width: 1024px) {
    #shopmain {
        padding: 1rem 2rem;
    }

    .nav-left {
        gap: 1rem;
    }

    .nav-links {
        gap: 1rem;
    }

    .search-input {
        width: 150px;
    }
}

@media (max-width: 768px) {
    #shopmain {
        flex-direction: column;
        gap: 1rem;
        padding: 1rem;
    }

    .nav-right {
        width: 100%;
        justify-content: center;
    }
}
</style>
