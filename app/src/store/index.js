import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import Api from "./api";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    apiBaseUrl: 'http://localhost:4000/',
    currentUser: localStorage.getItem('user') || {},
    token: localStorage.getItem('token') || '',
    header: {
      "headers":
      {
        "content-type": "application/json",
      }
    }
  },
  // Update token in local storage
  mutations: {
    setToken(currentState, token) {
      localStorage.setItem('token', token);
    },
    setCurrentUser(currentState, user) {
      localStorage.setItem('user', user);
    }
  },
  actions: {
    login: (context, credentials) => {
      return Api.post("login", credentials, context.state.header).then((token) => {
        context.commit('setToken', token.jwt);
        axios.defaults.headers.common.Authorization = `Bearer ${token.jwt}`;
      }).catch(error => {
        localStorage.removeItem('token');
        return Promise.reject(error);
      });
    },

    register: (context, credentials) => {
      return Api.post("users/", credentials, context.state.header).then((user) => {
        // context.commit('setToken', user.jwt);
        console.log(user);
        context.commit('setCurrentUser', user);
        axios.defaults.headers.common.Authorization = `Bearer ${user.jwt}`;
      }).catch(error => {
        if (error.data["code"] == 401) {
          return Promise.reject("Session timeout.");
        }
        return Promise.reject(error.data);
      });
    },

    getUsers: (context) => {
      if (!context.state.token)
        return Promise.reject("No users found.");

      return Api.get("users/", context.state.header).then((users) => {
        return Promise.resolve(users);
      }).catch(error => { return Promise.reject(error); });
    },

    logout: () => {
      localStorage.removeItem('token');
      delete axios.defaults.headers.common['Authorization'];
      Promise.resolve();
    },
  },
  modules: {}
});
