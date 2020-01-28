import store from '../store/index';
import axios from "axios";

axios.defaults.headers.common.Authorization = `Bearer ${localStorage.getItem('token')}`;

export default {
  get(url) {
    return axios
      .get(store.state.apiBaseUrl + url)
      .then(response => Promise.resolve(response.data))
      .catch(error => Promise.reject(error));
  },
  post(url, data) {
    return axios
      .post(store.state.apiBaseUrl + url, data)
      .then(response => Promise.resolve(response.data))
      .catch(error => Promise.reject(error.response));
  },
  put(url, data) {
    return axios
      .put(store.state.apiBaseUrl + url, data)
      .then(response => Promise.resolve(response.data))
      .catch(error => Promise.reject(error.response));
  },
  delete(url) {
    return axios
      .delete(store.state.apiBaseUrl + url)
      .then(response => Promise.resolve(response.data))
      .catch(error => Promise.reject(error.response));
  }
};
