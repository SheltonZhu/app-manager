import axios from "axios";
import store from "../store";
import router from "../router";
import storage from "../model/storage";

const apis = axios.create({
  baseURL: process.env.VUE_APP_BASE_URL,
  timeout: 6000,
});

apis.interceptors.request.use(
  (config) => {
    let token = store.state.token || storage.get("token");
    if (token) {
      config.headers.Authorization = token;
    } else {
      router.replace({
        path: "/login",
        query: { redirect: router.currentRoute.fullPath },
      });
    }
    return config;
  },
  (err) => {
    // vm.$loading.close();
    return Promise.reject(err);
  }
);
apis.interceptors.response.use(
  (response) => {
    // vm.$loading.close();
    return response;
  },
  (error) => {
    // 关闭loading
    // vm.$loading.close();
    if (error.response) {
      // vm.$toast.center("返回code:" + error.response.status + ";" + "请求错误");
      switch (error.response.status) {
        case 401:
          // 返回 401 清除token信息并跳转到登录页面
          // if (store.state.token || storage.get("token")) {
          //   // TODO: refresh token
          //   store.commit("updateToken", "");
          //   router.go(0);
          // } else {
          //   store.commit("updateToken", "");
          //   router
          //     .replace({
          //       path: "/login",
          //       query: { redirect: router.currentRoute.fullPath },
          //     })
          //     .then()
          //     .catch((e) => {
          //       console.log(e);
          //     });
          // }
          store.commit("updateToken", "");
          router
            .replace({
              path: "/login",
              query: { redirect: router.currentRoute.fullPath },
            })
            .then()
            .catch((e) => {
              console.log(e);
            });
      }
    }
    return Promise.reject(error); // 返回接口返回的错误信息
  }
);

const dockerBase = "/api/v1/app";

apis.healthBeat = () => {
  return apis.get("/ping");
};
apis.login = (account, password) => {
  return apis.post("/auth", {
    account: account,
    password: password,
  });
};
apis.ListContainers = (queryOption) => {
  let query = "";
  if (queryOption && Object.keys(queryOption).length > 0) {
    query =
      "?" +
      Object.keys(queryOption)
        .map((key) => {
          return (
            encodeURIComponent(key) + "=" + encodeURIComponent(queryOption[key])
          );
        })
        .join("&");
  }
  return apis.get(`${dockerBase}/dockers${query}`);
};
apis.restartContainer = (containerId) => {
  return apis.post(`${dockerBase}/docker/${containerId}/restart`);
};
apis.stopContainer = (containerId) => {
  return apis.post(`${dockerBase}/docker/${containerId}/stop`);
};
export default apis;
