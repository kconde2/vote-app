
import router from "./../router";
import store from "./../store";
import auth from "./../utils/auth"
import axios from "axios";

export default {
  loggedExpireID: null,
  isLoggedIn: function () {
    let isLoggedIn = localStorage.getItem('token') &&
      localStorage.getItem('user') &&
      localStorage.getItem('expiresDate');

    return isLoggedIn ? true : false;
  },
  runLoginDurationCheck: function () {
    if (auth.isLoggedIn()) {
      const token = localStorage.getItem('token');
      axios.defaults.headers.common.Authorization = `Bearer ${token}`;
      let expiresDate = localStorage.getItem('expiresDate');
      // let expiresDate = "2020-01-26T12:53:36.4006803+01:00";
      let expiresTimestamp = (new Date(expiresDate)).getTime();

      this.loggedExpireID = setInterval(() => {
        let now = Date.now();
        // let diff = expiresTimestamp - now;
        // console.log(diff)
        if (now >= expiresTimestamp) {
          if (auth.isLoggedIn()) {
            store
              .dispatch("logout")
              .then(() => {
                // redirect to dashboard
                router.push({
                  name: "login"
                });
              })
              .catch(() => {
                // handle errors
                this.error = true;
              });
          } else {
            console.log(this.loggedExpireID);
            clearInterval(this.loggedExpireID);
            this.loggedExpireID = null;
          }
        }
      }, 1000);
    }
  }
};
