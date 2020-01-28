<template>
  <div class="container">
    <div class="row">
      <div class="col-12 mx-auto">
        <div class="card border-0">
          <div class="card-header">Let's Vote !</div>
          <div class="card-body">
            <div>
              <strong>Title</strong>
            </div>
            <div>{{ votes.title }}</div>
            <div>
              <strong>Descrpition</strong>
            </div>
            <div>{{ votes.desc }}</div>
            <hr />
            <div class="mt-4">
              <button type="submit" v-on:click="updateVote()" class="btn btn-primary">Je vote</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row mt-4" v-if="user.access_level == 1">
      <div class="col-12 mx-auto">
        <div class="card border-0">
          <div class="card-header">Liste of voting users : {{ votes.uuid_votes.length }} votes</div>
          <div class="card-body">
            <p v-if="!votes.uuid_votes.length">No votes found</p>
            <table v-if="votes.uuid_votes.length" style="width:100%" class="table table-bordered">
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
    this.getTopic();
    this.getUser();
  },
  methods: {
    getTopic: function() {
      store
        .dispatch("getVotes", this.uuid)
        .then(votes => {
          this.votes = votes;
        })
        .catch(() => {});
    },
    getUser: function() {
      let uuid = JSON.parse(localStorage.getItem("user")).uuid;
      store
        .dispatch("getUserInfo", uuid)
        .then(user => {
          this.user = user;
        })
        .catch(() => {});
    },
    updateVote: function() {
      this.votes.uuid_votes = "";
      store
        .dispatch("updateTopic", this.votes)
        .then(() => {
          this.$router.push({
            name: "topic-list"
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
