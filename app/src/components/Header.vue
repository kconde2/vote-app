<template>
  <fragment>
    <nav class="navbar navbar-expand-md fixed-top navbar-dark bg-dark">
      <div class="container">
        <router-link :to="{ name: 'dashboard'}" class="navbar-brand">VoteApp</router-link>
        <button
          class="navbar-toggler"
          type="button"
          data-toggle="collapse"
          data-target="#menu-collapse"
          aria-controls="menu-collapse"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="menu-collapse">
          <ul class="navbar-nav mr-auto">
            <li class="nav-item" v-if="user.access_level == 1">
              <router-link :to="{ name: 'user-list'}" class="nav-link">Users</router-link>
            </li>

            <li class="nav-item">
              <router-link :to="{ name: 'topic-list'}" class="nav-link">Topics</router-link>
            </li>

            <li class="nav-item">
              <router-link :to="{ name: 'vote-list'}" class="nav-link">Votes</router-link>
            </li>
          </ul>
        </div>

        <ul class="navbar-nav mr-auto">
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              href="/"
              id="user-account"
              data-toggle="dropdown"
              aria-haspopup="true"
              aria-expanded="false"
            >Hi {{ user.first_name }}</a>
            <div class="dropdown-menu" aria-labelledby="user-account">
              <a class="dropdown-item" href="#">Account</a>
              <a class="dropdown-item" href="#" @click="logout">Logout</a>
            </div>
          </li>
        </ul>
      </div>
    </nav>
  </fragment>
</template>

<script>
import store from "../store/index";

export default {
  data: () => ({
    user: {}
  }),
  methods: {
    logout: function() {
      // Call Vuex action
      store
        .dispatch("logout")
        .then(() => {
          // redirect to dashboard
          this.$router.push({
            name: "login"
          });
        })
        .catch(() => {
          // handle errors
          this.error = true;
        });
    },
    getUser: function() {
      this.user = JSON.parse(localStorage.getItem("user")) || {};
    }
  },
  beforeMount() {
    this.getUser();
  }
};
</script>

<style lang="scss" scoped>
.va-header {
  background-color: #fff;
  box-shadow: 0 0 30px rgba(28, 39, 60, 0.08);
  border-bottom: 1px solid rgba(28, 39, 60, 0.12);

  &__wrapper {
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  &__logo {
    color: #57b846;
    font-weight: bold;
  }
}
</style>
