<template>
  <div class="container">
    <div class="row">
      <div class="col-12 d-flex mb-2">
        <div class="d-flex align-items-center">
          <h1 class="h3">Sujets</h1>
        </div>

        <div class="ml-auto d-flex align-items-center" v-if="user.access_level == 1">
          <router-link :to="{ name: 'topic-edit'}" class="btn btn-primary">Ajouter un sujet</router-link>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12">
        <div class="card border-0">
          <div class="card-body">
            <h5 class="card-title">Topics' list</h5>
            <div v-if="message.status" class="alert alert-success fade show" role="alert">
              <strong>{{ message.content }}</strong>
              <button type="button" class="close" data-dismiss="alert" aria-label="Close">
                <span aria-hidden="true">&times;</span>
              </button>
            </div>
            <div class="card-text table-responsive">
              <p v-if="!votes.length">No votes found</p>
              <table v-if="votes.length" style="width:100%" class="table table-bordered">
                <thead>
                  <tr>
                    <th>Title</th>
                    <th>Description</th>
                    <th v-if="user.access_level == 1">Edition</th>
                    <th>Votes</th>
                  </tr>
                </thead>
                <tbody id="v-for-object">
                  <tr v-for="vote in votes" :key="vote.uuid">
                    <td>{{ vote.title }}</td>
                    <td>{{ vote.desc }}</td>
                    <td v-if="user.access_level == 1">
                      <router-link
                        :to="{ name: 'topic-edit', params: { uuid: vote.uuid } }"
                        class="btn btn-secondary w-100"
                      >Edit</router-link>
                    </td>

                    <td>
                      <router-link
                        :to="{ name: 'vote-edit', params: { uuid: vote.uuid } }"
                        class="btn btn-info w-100"
                      >Details</router-link>
                    </td>
                  </tr>
                </tbody>
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
    votes: [],
    user: {},
    message: {
      status: false,
      content: ""
    }
  }),
  methods: {
    getTopicList: function() {
      store
        .dispatch("getTopic")
        .then(votes => {
          this.votes = votes;
        })
        .catch(() => {});
    },

    showTopic: function(uuid) {
      store
        .dispatch("getVotes", uuid)
        .then(result => {
          result.success
            ? (this.message = { status: true, content: result.success })
            : (this.message = { status: true, content: result.error });
        })
        .catch(() => {});
    },

    updateTopic: function(uuid) {
      store
        .dispatch("updateTopic", uuid)
        .then(result => {
          result.success
            ? (this.message = { status: true, content: result.success })
            : (this.message = { status: true, content: result.error });
        })
        .catch(() => {});
    },

    getUser: function() {
      this.user = JSON.parse(localStorage.getItem("user")) || {};
    }
  },
  beforeMount() {
    this.getTopicList(), this.getUser();
  }
};
</script>
