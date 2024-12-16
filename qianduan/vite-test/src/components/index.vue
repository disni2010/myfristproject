<template>
    <div class="bg-img"></div>
    <div class="page-container">
        <div id="shouye">
            <h3 class="title">首页</h3>
            <div class="button-container">
                <router-link to="/login" id="login" class="btn">登录</router-link>
                <router-link to="/register" id="resiger" class="btn">注册</router-link>
            </div>
            <router-view />
        </div>
    </div>
</template>

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

.button-container {
    display: flex;
    justify-content: space-around;
    margin-top: 100px;
}

.btn {
    padding: 12px 40px;
    background: linear-gradient(145deg, #3ca8df, #2980b9);
    color: white;
    text-decoration: none;
    border-radius: 25px;
    font-weight: 500;
    font-size: 18px;
    transition: all 0.3s ease;
    box-shadow: 0 4px 15px rgba(60, 168, 223, 0.2);
}

.btn:hover {
    background: linear-gradient(145deg, #2980b9, #3ca8df);
    transform: scale(1.05);
    box-shadow: 0 6px 20px rgba(60, 168, 223, 0.3);
}

#login,
#resiger {
    position: static;
}
</style>


<!-- <script setup>
import { useMainStore } from "../store"
import { storeToRefs } from 'pinia'
const mainstore = useMainStore();
console.log(mainstore.count);

这是有问题的 这拿到的数据不是响应式 一次型的 不会随着数据更改而更改里面数据
pinia 其实就是把state数据作为reactive处理了
const { count } = mainstore;
const { count } = storeToRefs(mainstore)

const handleChangestate = () => {

    方式一 最简单方法
    mainstore.count++

    方式二 修改多个数据
     mainstore.$patch({
        count: mainstore.count + 1,
       foo: 'hello',
       arr: [...mainstore.arr, 4]
     })

    方式三 $patch 一个函数
    mainstore.$patch(state => {
        state.count++
        state.foo = 'hello'
        state.arr.push(4)
    })

    方法四 调用mainstore里面的方法
    mainstore.increment()
} 

</script>


 pinia 需要在script后面加上setup才能使用不然就是在生成实例前使用错误
** Composition API 风格 **: 使用 setup 脚本可以让组件直接利用 Vue 3 的 Composition API，而无需显式地定义 export default 和 setup
函数。这使得组件的逻辑更容易阅读和组织。
 ** 自动导入 **: 在 < script setup> 中声明的变量、导入的组件或函数会自动成为组件的反应性属性或方法，无需通过 return 来暴露。
 ** 更少的样板代码 **: 省去了在 data, computed, methods 等选项中声明属性和方法的步骤，使得组件代码更加简洁。
 ** 编译器优化 **: Vue 编译器会对 < script setup> 块进行特别处理，可以进行一些优化，比如静态分析来提升性能。 

 传统写法:
    < script>
         export default {
         data() {
         return {
         message: 'Hello Vue!'
         };
         },
         methods: {
         sayHello() {
         console.log(this.message);
         }
         }
        }
         </> 


 <script setup> 写法:

    Vue
    深色版本
    <script setup>
        import {ref} from 'vue';
        const message = ref('Hello Vue!');
const sayHello = () => {
            console.log(message.value);
}
</script>  
-->