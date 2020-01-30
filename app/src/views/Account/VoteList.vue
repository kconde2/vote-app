<template>
  <div class="container">
    <div class="row">
      <div class="col-12 mx-auto">
        <div class="card border-0">
          <div class="card-header">My votes topics' list</div>
          <div class="card-body">
            <p v-if="!votes.length">Empty votes list</p>
            <table v-if="votes.length" class="table table-bordered w-100">
              <tr>
                <th>Title</th>
                <th>Description</th>
              </tr>

              <tr v-for="vote in votes" :key="vote.uuid">
                <td>{{ vote.title }}</td>
                <td>{{ vote.desc }}</td>
              </tr>
            </table>
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
    votes: {},
    user: {},
    message: {
      status: false,
      content: ""
    }
  }),
  beforeMount() {
    this.getUser();
    this.getUserVotes();
  },
  methods: {
    getUserVotes: function() {
      store
        .dispatch("getUserVotes", this.user.uuid)
        .then(({ user, votes }) => {
          this.user = user;
          this.votes = votes;
        })
        .catch(() => {});
    },
    getUser: function() {
      this.user = JSON.parse(localStorage.getItem("user")) || {};
    }
  }
};
</script>

<style></style>
