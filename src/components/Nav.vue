<template>
  <div id="nav">
    <el-menu
      default-active="1-4-1"
      mode="horizontal"
      @open="handleOpen"
      @close="handleClose"
      :router="true"
    >
      <el-submenu index="1">
        <template slot="title">
          <i class="el-icon-menu" />
          <span slot="title">Docker</span>
        </template>
        <el-menu-item index="/docker">
          容器
        </el-menu-item>
      </el-submenu>
      <el-menu-item index="/home">
        <i class="el-icon-setting" />
        <span slot="title">设置</span>
      </el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "Nav",
  data() {
    return {
      isCollapse: true,
    };
  },
  methods: {
    handleOpen(key, keyPath) {
      console.log(key, keyPath);
    },
    handleClose(key, keyPath) {
      console.log(key, keyPath);
    },
  },
};
</script>

<style scoped>
#nav {
  margin-bottom: 5px;
}
</style>
