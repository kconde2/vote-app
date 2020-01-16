export default {
  loggedIn: localStorage.getItem('token'),
  login: function () { this.loggedIn = true },
  logout: function () { this.loggedIn = false }
};
