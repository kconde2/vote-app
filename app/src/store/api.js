// import Vue from 'vue';
import store from './index';
import axios from "axios";

export default {
  get(url, headers) {
    return axios
      .get(store.state.apiBaseUrl + url, headers)
      .then(response => Promise.resolve(response.data))
      .catch(error => Promise.reject(error));
  },
  post(url, data, headers) {
    return axios
      .post(store.state.apiBaseUrl + url, data, headers)
      .then(response => Promise.resolve(response))
      .catch(error => Promise.reject(error));
  },

};
