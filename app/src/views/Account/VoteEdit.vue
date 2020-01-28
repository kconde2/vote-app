<template>
  <div class="container">
    <div class="row">
      <div class="col-8 mx-auto">
        <div class="card border-0">
          <div class="card-header">Let's Vote !</div>
          <div class="card card-body">
            <div class="card-header">{{ votes.title }}</div>
            <div class="card-body">{{ votes.desc }}</div>
          </div>
          <div class="card-footer">
            <button
              style="padding:2% 40% 2% 40%; margin-left:5%"
              type="submit"
              v-on:click="UpdateVote()"
              class="btn btn-primary"
            >
              JE VOTE
            </button>
          </div>
        </div>
      </div>
    </div>

    <div class="row" v-if="user.access_level == 1">
      <div class="col-8 mx-auto">
        <div class="card border-0">
          <div class="card-header">
            Liste of voting users : {{ votes.uuid_votes.length }} votes
          </div>
          <div class="card card-body">
            <p v-if="!votes.uuid_votes.length">No votes found</p>
            <table
              v-if="votes.uuid_votes.length"
              style="width:100%"
              class="table table-bordered"
            >
              <thead>
                <tr>
                  <th>User UUID</th>
                </tr>
              </thead>
              <tbody id="v-for-object">
                <tr v-for="vote in votes.uuid_votes" :key="vote.id">
                  <td>{{ vote }}</td>
                </tr>
              </tbody>
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
    votes: {
      title: "",
      uuid_votes: [],
      desc: ""
    },
    user: {},
    message: {
      status: false,
      content: ""
    }
  }),

  mounted() {
    this.uuid = this.$route.params.uuid; // id of the sujet
    this.getTopic(this.uuid);
    this.user = this.getUser();
  },
  methods: {
    getTopic: function(uuid) {
      store
        .dispatch("getVotes", uuid)
        .then(votes => {
          console.log(votes);
          this.votes = votes;
        })
        .catch(() => {});
    },

    getUser: function() {
      this.user = JSON.parse(localStorage.getItem("user")) || {};
      console.log("USER INFO : ", this.user.uuid);
      store.dispatch(user => {
        this.users = user;
        console.log("votes", user, this.user);
      });
    },

    UpdateVote: function() {
      this.votes.uuid_votes = "";
      store
        .dispatch("updateTopic", this.votes)
        .then(() => {
          this.$router.push({
            name: "dashboard"
          });
        })
        .catch(error => {
          // handle errors
          this.errors = { status: true, message: error.errors };
        });
    }
  }
};
</script>

<style scoped></style>
