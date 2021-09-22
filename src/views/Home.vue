<template>
  <div class="home" >
    <Nav />
    <router-view />
  </div>
</template>

<script>
import Nav from "../components/Nav";
export default {
  name: "Home",
  mounted() {},
  components: { Nav },
  data() {
    return {};
  },
  methods: {},
};
</script>
<style>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 400px;
}
</style>
