import Vue from "vue";
import Vuex from "vuex";
// import axios from "axios";
import Api from "./api";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    apiBaseUrl: 'http://localhost:4000/',
    currentUser: null,
    header: {
      "headers":
      {
        "content-type": "application/json",
      }
    }
  },
  // change the state tree
  mutations: {
    setCurrentUser(currentState, user) {
      currentState.currentUser = user;
    }
  },
  actions: {
    login: (context, credentials) => {
      return Api.post("login", credentials, context.state.header).then((jwt) => {
        context.commit('setCurrentUser', jwt);
        console.log(context.state.currentUser);
      }).catch(error => { return Promise.reject(error); });
    },

    register: (context, credentials) => {
      context.state.header = {
        "headers":
        {
          "content-type": "application/json",
          "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfbGV2ZWwiOjEsImV4cCI6MTU3ODg3ODUzNywib3JpZ19pYXQiOjE1Nzg4NzQ5MzcsInV1aWQiOiI4YTM1YzMwYi00NmVkLTQyN2QtYWU4Yi1jNzQ3OTIzNjQ1MzYifQ.FHch528hinD7vd_wWh5pevj8sLjbJ8Ss437swYNZPU8"
        }
      };
      return Api.post("users/", credentials, context.state.header).then(() => {

      }).catch(error => {
        // Display error messages about bad password
        return Promise.reject(error);
      });
    },
    getUsers: (context) => {
      context.state.header = {
        "headers":
        {
          "content-type": "application/json",
          "Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhY2Nlc3NfbGV2ZWwiOjEsImV4cCI6MTU3ODg3ODUzNywib3JpZ19pYXQiOjE1Nzg4NzQ5MzcsInV1aWQiOiI4YTM1YzMwYi00NmVkLTQyN2QtYWU4Yi1jNzQ3OTIzNjQ1MzYifQ.FHch528hinD7vd_wWh5pevj8sLjbJ8Ss437swYNZPU8"
        }
      };
      console.log(context.state.currentUser);
      return Api.get("users/", context.state.header).then((users) => {
        return Promise.resolve(users);
      }).catch(error => { return Promise.reject(error); });
    }
  },
  modules: {}
});
