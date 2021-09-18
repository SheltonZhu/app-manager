<template>
  <div>
    <div class="header">
      <h2 class="logo">
        <img
          width="128"
          alt="Vue logo"
          src="../assets/logo.png"
        >
      </h2>
    </div>
    <Login />

    <div class="footer">
      <div class="links">
        <span>All rights </span>
        <span>reserved by </span>
        <span>&copy;Shelton Zhu. </span>
      </div>
      <div class="version">
        Version 1.0.0-beta-1
      </div>
    </div>
  </div>
</template>

<script>
import Login from "../components/Login";

export default {
  name: "LoginView",
  components: {
    Login,
  },
};
</script>
<style scoped>
.footer,
.header h1,
h2,
h3 {
  color: #f0f4f8;
  padding: 0;
  margin: 0;
}

.footer,
.header .links {
  display: flex;
  font-family: "Open Sans";
}

.footer,
.header .links span {
  padding: 0 10px;
  font-size: 18px;
  border-right: 1px solid #9fb3c8;
}

.footer,
.header .links span:last-child {
  border-right: none;
}

.footer,
.header {
  padding: 20px 20px;
  color: #f0f4f8;
  display: flex;
  flex-direction: column;
  align-items: center;
}

.footer,
.header .version {
  font-family: "Open Sans";
  padding: 0 10px;
  color: #9fb3c8;
  font-size: 12px;
  margin-top: 5px;
}

.header {
  padding: 10px 20px;
}

.header .logo {
  font-family: "Open Sans";
  letter-spacing: 3px;
  padding-top: 15px;
  padding-bottom: 15px;
}

.header .logo .part-2 {
  color: #d64545;
}
</style>
