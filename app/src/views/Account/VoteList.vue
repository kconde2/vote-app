<template>
  <div class="container">
    VoteList
    <div class="row">
      <div class="col-8 mx-auto">
        <div class="card border-0">
          <div class="card-header">Liste of voting topic : votes</div>
          <div class="card card-body">
            <table
              v-if="topics.length"
              style="width:100%"
              class="table table-bordered"
            >
              <tr>
                <th>name</th>
              </tr>
              <tr v-bind:key="topic.uuid" v-for="topic in topics">
                <td>
                  {{ topic }}
                </td>
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
    topics: {},
    user: {},
    message: {
      status: false,
      content: ""
    }
  }),
  beforeMount() {
    this.getTopicList();
  },
  methods: {
    getTopicList: function() {
      store
        .dispatch("getTopic")
        .then(topics => {
          this.topics = topics;
          console.log("topics", this.topics);
        })
        .catch(() => {});
    },

    getVotes: function(uuid) {
      store
        .dispatch("getVotes", uuid)
        .then(vote => {
          this.votes = vote;
          console.log("votes", this.votes);
        })
        .catch(() => {});
    },

    getUser: function() {
      this.user = JSON.parse(localStorage.getItem("user")) || {};
      console.log("USER INFO : ", this.user.uuid);
    }
  }
};
</script>

<style></style>
