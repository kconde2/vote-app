import Vue from "vue";
import Vuex from "vuex";
import axios from "axios";
import Api from "./api";
import jwt_decode from "jwt-decode";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    apiBaseUrl: "http://localhost:4000/",
    currentUser: localStorage.getItem("user") || {},
    token: localStorage.getItem("token") || "",
    header: {
      headers: {
        "content-type": "application/json"
      }
    }
  },
  // Update token in local storage
  mutations: {
    setToken(currentState, token) {
      localStorage.setItem("token", token);
    },
    setCurrentUser(currentState, user) {
      localStorage.setItem("user", JSON.stringify(user));
    }
  },
  actions: {
    login: (context, credentials) => {
      return Api.post("login", credentials, context.state.header)
        .then(token => {
          context.commit("setToken", token.jwt);
          axios.defaults.headers.common.Authorization = `Bearer ${token.jwt}`;
          context.commit("setCurrentUser", jwt_decode(token.jwt));
        })
        .catch(error => {
          localStorage.removeItem("token");
          return Promise.reject(error);
        });
    },

    register: (context, credentials) => {
      return Api.post("users/", credentials, context.state.header)
        .then(() => {
          return Promise.resolve();
        })
        .catch(error => {
          if (error.data["code"] == 401)
            return Promise.reject({
              errors: { message: "Session timeout, please log in again" }
            });
          return Promise.reject(error.data);
        });
    },

    getUsers: context => {
      if (!context.state.token) return Promise.reject("No users found.");

      return Api.get("users/", context.state.header)
        .then(users => {
          return Promise.resolve(users);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },

    deleteUser: (context, uuid) => {
      return Api.delete("users/" + uuid, context.state.header)
        .then(user => {
          return Promise.resolve(user);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },

    updateUser: (context, uuid) => {
      return Api.put("users/" + uuid, context.state.header)
        .then(user => {
          return Promise.resolve(user);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },

    editTopic: (context, credentials) => {
      return Api.post("votes/", credentials, context.state.header)
        .then(() => {
          console.log(Promise);
          return Promise.resolve();
        })
        .catch(error => {
          if (error.data["code"] == 401)
            return Promise.reject({
              errors: { message: "error please try again" }
            });
          return Promise.reject(error.data);
        });
    },

    getTopic: context => {
      return Api.get("votes/", context.state.header)
        .then(votes => {
          console.log("hello", votes);
          return Promise.resolve(votes);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },
    getVotes: (context, uuid) => {
      return Api.get("votes/" + uuid, context.state.header)
        .then(votes => {
          console.log("votes", uuid, " avant uid", votes);
          return Promise.resolve(votes);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },

    updateTopic: (context, uuid) => {
      return Api.put("votes/" + uuid, context.state.header)
        .then(vote => {
          return Promise.resolve(vote);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },

    UpdateVote: function(uuid, credentials, context) {
      return Api.put("/votes/" + uuid, context.state.header, credentials)
        .then(vote => {
          return Promise.resolve(vote);
        })
        .catch(error => {
          return Promise.reject(error);
        });
    },

    logout: () => {
      localStorage.removeItem("token");
      delete axios.defaults.headers.common["Authorization"];
      Promise.resolve();
    }
  },
  modules: {}
});
