<style>
.my-component h2 {
  color: red;
}
</style>

<template>
  <div class="my-component">
    <h2>Hello from {{msg}}</h2>    
  </div>
</template>

<script>
// 导出组件定义
module.exports = {
  data: function () {
    return {
      msg: 'vue-loader'
    }
  },
}
</script>