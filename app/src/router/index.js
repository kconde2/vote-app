import Vue from "vue";
import VueRouter from "vue-router";
import FormUseCase from "../views/FormUseCase.vue";
import Admin from "../views/Admin.vue";
import Login from "../views/Authentication/Login.vue";
import Register from "../views/Authentication/Register.vue";
import Auth from "../views/Authentication/Auth.vue";
import PageNotFound from "../views/PageNotFound.vue";
import Dashboard from "../views/Account/Dashboard.vue";
import UserList from "../views/Account/UserList.vue";
import TopicList from "../views/Account/TopicList.vue";
import VoteList from "../views/Account/VoteList.vue";
import UserAdd from "../views/Account/UserAdd.vue";
import auth from "../utils/auth";

Vue.use(VueRouter);

const routes = [
  {
    path: "/account",
    component: Admin,
    meta: {
      requiresAuth: true,
    },
    children: [
      {
        path: "dashboard",
        name: "dashboard",
        component: Dashboard,
      },
      {
        path: "user/list",
        name: "user-list",
        component: UserList,
      },
      {
        path: "user/add",
        name: "user-add",
        component: UserAdd,
      },
      {
        path: "topic/list",
        name: "topic-list",
        component: TopicList,
      },
      {
        path: "vote/list",
        name: "vote-list",
        component: VoteList,
      },
      {
        path: "register",
        name: "register",
        component: Register
      }
    ]
  },
  {
    path: "/auth",
    component: Auth,
    children: [
      {
        path: "login",
        name: "login",
        component: Login
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
  if (to.matched.some(record => record.meta.requiresAuth) && auth.loggedIn == '') {
    next({ path: '/auth/login' });
  } else {
    if (to.path == '/') {
      next({ path: '/account/dashboard' });
    } else {
      next();
    }
  }
});

export default router;
