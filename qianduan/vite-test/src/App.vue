<template>
  <!-- <form @submit.prevent="submitForm">
    <input type="text" v-model.lazy="username">
    <input type="submit" value="提交">
  </form> -->
  <router-view></router-view>
</template>

<script>

// import { useMainStore } from "./store"
// const mainstore = useMainStore();
// console.log(mainstore.count);
export default {

  data() {
    return {
      username: ''
    }
  },
  methods: {
    submitForm() {
      //1 通过组合域名发送
      // this.$axios.post('/api/upload/' + this.username).then({
      // })
      this.$axios.post('/api/upload', this.username)
    }
  },

}
</script>

<style scoped></style>
