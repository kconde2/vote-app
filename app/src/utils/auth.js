var Auth = {
  loggedIn: false,
  login: function () { this.loggedIn = true },
  logout: function () { this.loggedIn = false }
};

export default Auth;
