<template>
  <div class="container">
    <div class="row">
      <div class="col-12 d-flex mb-2">
        <div class="d-flex align-items-center">
          <h1 class="h3">List of users</h1>
        </div>

        <div class="ml-auto d-flex align-items-center">
          <router-link :to="{ name: 'register'}" href="#" class="btn btn-primary">Add user</router-link>
        </div>
      </div>
    </div>

    <div class="row">
      <div class="col-12">
        <div class="card border-0">
          <div class="card-body">
            <h5 class="card-title"></h5>
            <div v-if="message.status" class="alert alert-success fade show" role="alert">
              <strong>{{ message.content }}</strong>
              <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="card-text table-responsive">
              <p v-if="!users.length">No users found</p>
              <table v-if="users.length" style="width:100%" class="table table-bordered">
                <tr>
                  <th>Firstname</th>
                  <th>Lastname</th>
                  <th>Email</th>
                  <th class="text-center">Birthdate</th>
                  <th class="text-center">Rôle</th>
                  <th class="text-center">Edit</th>
                  <th class="text-center">Delete</th>
                </tr>
                <tr v-bind:key="user.uuid" v-for="user in users">
                  <td>{{ user.first_name }}</td>
                  <td>{{ user.last_name }}</td>
                  <td>{{ user.email }}</td>
                  <td class="text-center">{{ user.birth_date }}</td>
                  <td class="text-center">{{ getRole(user.access_level) }}</td>

                  <td class="text-center">
                    <router-link
                      class="btn btn-secondary w-100"
                      :to="{ name: 'edit-user', params: { uuid: user.uuid }}"
                    >Edit</router-link>
                  </td>
                  <td class="text-center">
                    <button class="btn btn-danger w-100" v-on:click="deleteUser(user.uuid)">Delete</button>
                  </td>
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
    users: {},
    message: {
      status: false,
      content: ""
    }
  }),
  methods: {
    getUserList: function() {
      store
        .dispatch("getUsers")
        .then(users => {
          this.users = users;
        })
        .catch(() => {});
    },
    deleteUser: function(uuid) {
      store
        .dispatch("deleteUser", uuid)
        .then(result => {
          result.success
            ? (this.message = { status: true, content: result.success })
            : (this.message = { status: true, content: result.error });

          // delete user from users list
          this.users = this.users.filter(user => user.uuid !== uuid);
        })
        .catch(() => {});
    },
    updateUser: function(uuid) {
      store
        .dispatch("updateUser", uuid)
        .then(result => {
          result.success
            ? (this.message = { status: true, content: result.success })
            : (this.message = { status: true, content: result.error });
        })
        .catch(() => {});
    },
    getRole: function(accessLevel) {
      if (accessLevel == 1) {
        return "Administrateur";
      } else if (accessLevel == 0) {
        return "Votant";
      } else {
        return "Unknown";
      }
    }
  },
  beforeMount() {
    this.getUserList();
  }
};
</script>

<style>
</style>
