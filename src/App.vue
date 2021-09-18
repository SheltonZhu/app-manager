<template>
  <div id="app">
    <Nav />
    <router-view />
    <el-backtop target="#app" />
  </div>
</template>
<script>
import { mapState } from "vuex";
import Nav from "./components/Nav";

export default {
  name: "App",
  computed: {
    ...mapState(["token"]),
  },
  components: { Nav },
  mounted() {
    if (this.token === "") {
      this.$router.replace({
        path: "login",
        query: { redirect: this.$router.currentRoute.fullPath },
      });
    }
  },
};
</script>
<style>
body {
  margin: 0;
  padding: 0;
  background: #102a43;
  background-image: url("assets/dots.png");
  background-size: contain;
}

#app {
  font-family: Roboto, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 0;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}
</style>
