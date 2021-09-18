import Vue from "vue";
import Vuex from "vuex";
import storage from "../model/storage";
Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    token: "",
  },
  mutations: {
    updateToken(state, token) {
      state.token = token;
      storage.set("token", token);
      console.log("更新token成功！");
    },
  },
  actions: {},
  modules: {},
});
