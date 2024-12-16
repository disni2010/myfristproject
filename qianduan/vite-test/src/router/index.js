import {
    createRouter,
    createWebHashHistory,
    createWebHistory,
} from "vue-router";

const routers = [
    {
        //渲染时必须有一个匹配路由不能没有渲染页面
        path: '/',
        component: () => import('../components/index.vue')
    }, {
        path: '/login',
        component: () => import('../components/login/login.vue')
    }, {
        path: '/register',
        component: () => import('../components/login/register.vue')
    }, {
        path: '/shopmain',
        component: () => import('../components/shop/shopmain.vue'),
        children: [{
            path: 'UserControll',
            component: () => import('../components/shop/UserControll.vue')
        }, {
            path: 'shopTickets/:method',
            component: () => import('../components/shop/shopTickets .vue'),
        }, {
            path: 'ScenicSpots',
            component: () => import('../components/shop/ScenicSpots .vue')
        }, {
            path: 'Commet',
            component: () => import('../components/shop/Commet.vue')
        },]
    }

]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routers
})

export default router;