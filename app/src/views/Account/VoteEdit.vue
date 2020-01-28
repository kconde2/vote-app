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
            <div>{{ vote.title }}</div>
            <div>
              <strong>Descrpition</strong>
            </div>
            <div>{{ vote.desc }}</div>
            <hr />
            <div class="mt-4">
              <button
                type="submit"
                v-on:click="updateVote()"
                class="btn btn-primary w-100"
                :disabled="hasAlreadyVoted"
              >Je vote</button>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="row mt-4" v-if="user.access_level == 1">
      <div class="col-12 mx-auto">
        <div class="card border-0">
          <div class="card-header">Liste of voting users : {{ vote.uuid_votes.length }} votes</div>
          <div class="card-body">
            <p v-if="!vote.uuid_votes.length">No votes found</p>
            <table v-if="vote.uuid_votes.length" style="width:100%" class="table table-bordered">
              <thead>
                <tr>
                  <th>User UUID</th>
                  <th>Firstname</th>
                  <th>Lastname</th>
                </tr>
              </thead>
              <tbody id="v-for-object">
                <tr v-for="user in voteUsers" :key="user.id">
                  <td>{{ user.uuid }}</td>
                  <td>{{ user.first_name }}</td>
                  <td>{{ user.last_name }}</td>
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
    vote: {
      title: "",
      uuid_votes: [],
      desc: ""
    },
    voteUsers: [],
    user: {},
    message: {
      status: false,
      content: ""
    }
  }),

  mounted() {
    this.uuid = this.$route.params.uuid; // id of the sujet
    this.getTopicWithUsers();
    this.getUser();
  },
  computed: {
    hasAlreadyVoted: function() {
      return this.voteUsers.some(user => user.uuid === this.user.uuid);
    }
  },
  methods: {
    // get vote and associated users
    getTopicWithUsers: function() {
      store
        .dispatch("getTopicUsers", this.uuid)
        .then(({ users, vote }) => {
          this.voteUsers = users;
          this.vote = vote;
        })
        .catch(() => {});
    },
    // get current user info
    getUser: function() {
      let uuid = JSON.parse(localStorage.getItem("user")).uuid;
      store
        .dispatch("getUserInfo", uuid)
        .then(user => {
          this.user = user;
        })
        .catch(() => {});
    },
    // vote for topic
    updateVote: function() {
      this.vote.uuid_votes = "";
      store
        .dispatch("updateTopic", this.vote)
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
