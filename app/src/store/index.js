import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import Api from "./api";
import jwt_decode from 'jwt-decode';

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
    },
    expireTime: null
  },
  // Update token in local storage
  mutations: {
    setToken(currentState, token) {
      localStorage.setItem('token', token);
    },
    setCurrentUser(currentState, user) {
      localStorage.setItem('user', JSON.stringify(user));
    },
    setExpireTime(currentState, expireTime) {
      currentState.expireTime = expireTime
      localStorage.setItem('expiresDate', expireTime);
    }
  },
  actions: {
    login: (context, credentials) => {
      return Api.post("login", credentials, context.state.header).then((token) => {
        context.commit('setToken', token.jwt);
        context.commit('setExpireTime', token.time);
        axios.defaults.headers.common.Authorization = `Bearer ${token.jwt}`;
        context.commit('setCurrentUser', jwt_decode(token.jwt));
       return Promise.resolve();
      }).catch(error => {
        localStorage.removeItem('token');
        return Promise.reject(error);
      });
    },

    register: (context, credentials) => {
      return Api.post("users/", credentials, context.state.header).then(() => {
        return Promise.resolve();
      }).catch(error => {
        if (error.data["code"] == 401)
          return Promise.reject({ "errors": { "message": "Session timeout, please log in again" } });
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

    deleteUser: (context, uuid) => {
      return Api.delete("users/" + uuid, context.state.header).then((user) => {
        return Promise.resolve(user);
      }).catch(error => { return Promise.reject(error); });
    },

    updateUser: (context, uuid) => {
      return Api.put("users/" + uuid, context.state.header).then((user) => {
        return Promise.resolve(user);
      }).catch(error => { return Promise.reject(error); });
    },

    logout: () => {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      localStorage.removeItem('expiresDate');
      delete axios.defaults.headers.common['Authorization'];
      Promise.resolve();
    },
  },
  modules: {}
});
