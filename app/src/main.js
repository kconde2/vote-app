import Vue from "vue";
import App from "./App.vue";
import Vuelidate from "vuelidate";
import router from "./router";
import store from "./store";
import jQuery from 'jquery';
import { Plugin } from 'vue-fragment';
import axios from "axios";

window.$ = window.jQuery = jQuery;

import 'popper.js';
import 'bootstrap';
import './assets/app.scss';

import '@fortawesome/fontawesome-free/css/all.css'
import '@fortawesome/fontawesome-free/js/all.js'

Vue.config.productionTip = false;

Vue.use(Plugin);
Vue.use(Vuelidate);

Vue.prototype.$http = axios;
const token = localStorage.getItem('token');

if (token) {
  // Vue.prototype.$http.defaults.headers.common['Autorization'] = token;
  axios.defaults.headers.common.Authorization = `Bearer ${token}`;
}

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
