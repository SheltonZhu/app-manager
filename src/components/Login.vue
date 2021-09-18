<template>
  <div class="login">
    <el-card>
      <h2>登 录</h2>
      <el-form
        class="login-form"
        :model="model"
        :rules="rules"
        ref="form"
        @submit.native.prevent="login"
      >
        <el-form-item prop="username">
          <el-input
            v-model="model.username"
            placeholder="账号"
            prefix-icon="fas fa-user"
          />
        </el-form-item>
        <el-form-item prop="password">
          <el-input
            v-model="model.password"
            placeholder="密码"
            type="password"
            prefix-icon="fas fa-lock"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            :loading="loading"
            class="login-button"
            type="primary"
            native-type="submit"
            block
          >
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import apis from "../router/request";
import { mapMutations } from "vuex";

export default {
  name: "Login",
  data() {
    return {
      model: {
        username: "",
        password: "",
      },
      loading: false,
      rules: {
        username: [
          {
            required: true,
            message: "用户名必填",
            trigger: "blur",
          },
          {
            min: 5,
            message: "用户名至少5个字符",
            trigger: "blur",
          },
        ],
        password: [
          { required: true, message: "密码必填", trigger: "blur" },
          {
            min: 5,
            message: "密码至少5个字符",
            trigger: "blur",
          },
        ],
      },
    };
  },
  methods: {
    ...mapMutations(["updateToken"]),
    simulateLogin() {
      return apis.login(this.model.username, this.model.password);
    },
    async login() {
      try {
        let valid = await this.$refs.form.validate();
        if (!valid) {
          return false;
        } else {
          this.loading = true;
          try {
            let res = await this.simulateLogin();
            let token = res.data.result.token;
            this.updateToken(token);
            this.$message.success("登录成功！");
            setTimeout(() => {
              // let redirect = this.$route.query.redirect;
              // if (redirect) {
              //   this.$router.push(redirect);
              // } else {
              //   this.$router.push("/");
              // }
              this.$router.push("/");
            }, 1000);
          } catch (e) {
            this.$message.error("账号或密码不正确！");
          } finally {
            this.loading = false;
          }
        }
      } catch (e) {
        // console.log(e)
      }
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.login {
  flex: 1;
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
}

.login-button {
  width: 100%;
  margin-top: 40px;
}

.login-form {
  width: 290px;
}

.forgot-password {
  margin-top: 10px;
}

.el-button--primary {
  background: rgb(0, 124, 137) !important;
  border-color: rgb(0, 124, 137) !important;
}

.el-button--primary:hover,
.el-button--primary.active,
.el-button--primary:focus {
  background: rgb(0, 124, 137) !important;
  border-color: rgb(0, 124, 137) !important;
}

.login .el-input__inner:hover {
  border-color: rgb(0, 124, 137) !important;
}

.login .el-input__prefix {
  background: rgb(238, 237, 234);
  height: calc(100% - 2px);
  left: 1px;
  top: 1px;
  border-radius: 3px;
}

.login .el-input__prefix .el-input__icon {
  width: 30px;
}

.login .el-input input {
  padding-left: 35px;
}

.login .el-card {
  padding-top: 0;
  padding-bottom: 30px;
}

h2 {
  letter-spacing: 1px;
  font-family: "Open Sans", Roboto, sans-serif;
  padding-bottom: 20px;
}

a {
  color: rgb(0, 124, 137) !important;
  text-decoration: none;
}

a:hover,
a:active,
a:focus {
  color: rgb(0, 124, 137) !important;
}

.login .el-card {
  width: 340px;
  display: flex;
  justify-content: center;
}
</style>
