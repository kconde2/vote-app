<template>
  <div class="container">
    <div class="row">
      <div class="col-12 d-flex mb-2">
        <div class="d-flex align-items-center">
          <h1 class="h3">Liste des utilisateurs</h1>
        </div>

        <div class="ml-auto d-flex align-items-center">
          <router-link
            :to="{ name: 'user-add'}"
            href="#"
            class="btn btn-primary"
          >Ajouter un utilisateur</router-link>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <div class="card border-0">
          <div class="card-body">
            <h5 class="card-title">Title</h5>
            <div class="card-text table-responsive">
              <table style="width:100%" class="table table-bordered">
                <tr>
                  <th>Firstname</th>
                  <th>Lastname</th>
                  <th>Date de naissance</th>
                </tr>
                <tr v-bind:key="user.uuid" v-for="user in users">
                  <td>{{ user.first_name }}</td>
                  <td>{{ user.last_name }}</td>
                  <td>{{ user.birth_date }}</td>
                </tr>
              </table>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import store from "../../store/index";

export default {
  data: () => ({
    users: {}
  }),
  methods: {
    getUserList: function() {
      store
        .dispatch("getUsers")
        .then(users => {
          this.users = users;
          console.log(users);
        })
        .catch(() => {});
    }
  },
  beforeMount() {
    this.getUserList();
  }
};
</script>

<style>
</style>
