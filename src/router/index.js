import Vue from "vue";
import VueRouter from "vue-router";
// import Home from "../views/Home.vue";
import Login from "../views/Login";
import storage from "../model/storage";
import store from "../store";
import Container from "../views/Container";
import Home from "../views/Home";
Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    redirect: "/docker",
    meta: {
      requireAuth: true,
      title: "主页",
    },
    children: [
      {
        path: "/docker",
        name: "docker",
        meta: {
          requireAuth: true,
          title: "容器列表",
        },
        component: Container,
      },
    ]
  },
  {
    path: "/login",
    name: "login",
    component: Login,
  },
  {
    path: "/home",
    name: "home",
    meta: {
      requireAuth: true,
      title: "主页",
    },
    component: Home,
  },
  {
    path: "*",
    component: Container,
  },
];

const router = new VueRouter({
  routes,
});
if (storage.get("token")) {
  store.commit("updateToken", storage.get("token"));
}
router.beforeEach((to, from, next) => {
  if (to.matched.some((r) => r.meta.requireAuth)) {
    // if (to.meta.requireAuth)) {
    if (store.state.token) {
      next();
    } else {
      next({
        path: "/login",
        query: { redirect: to.fullPath },
      });
    }
  } else {
    next();
  }
});
export default router;
