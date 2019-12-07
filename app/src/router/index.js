import Vue from "vue";
import VueRouter from "vue-router";
import FormUseCase from "../views/FormUseCase.vue";
import Home from "../views/Home.vue";
import Login from "../views/Authentication/Login.vue";
import Register from "../views/Authentication/Register.vue";
import Auth from "../views/Authentication/Auth.vue";
import PageNotFound from "../views/PageNotFound.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Home,
    meta: {
      requiresAuth: true,
      isAdmin: true
    }
  },
  {
    path: "/auth",
    component: Auth,
    children: [
      {
        path: "login",
        name: "login",
        component: Login
      },
      {
        path: "register",
        name: "register",
        component: Register
      }
    ]
  },
  {
    path: "/formik",
    name: "formik-use-case",
    component: FormUseCase
  },
  {
    path: "/about",
    name: "about",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue")
  },
  { path: "*", component: PageNotFound }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});


router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requiresAuth) && !Auth.loggedIn) {
    // next({ path: '/auth/login', query: { redirect: to.fullPath } });
    next({ path: '/auth/login' });
  } else {
    next();
  }
});

export default router;
