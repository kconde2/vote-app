<template>
  <div class="container">
    <div class="row">
      <div class="col-8 mx-auto">
        <div class="card border-0">
          <div class="card-header">Let's Vote !</div>
          <div class="card card-body">
            <div class="card-header">{{votes.title}}</div>
            <div class="card-body">{{votes.desc}}</div>
          </div>
          <div class="card-footer">
            <Formik submit-label="Valider" :initial-values="votes" class="auth-wrapper__form">
              <div
                class="form-group"
                style="background: white; color: black; padding: 1em; border-radius: 5px;"
              >
                <label for="uuid_votes">
                  <Field type="radio" name="uuid_votes" value="OUI" @change="upvote()" />OUI
                </label>
                <span style="padding-left: 60%;">
                  <i class="fas fa-thumbs-up"></i>
                  ({{ votes.uuid_votes.length }})
                </span>
              </div>
              <hr />
              <div
                class="form-group"
                style="background: white; color: black; padding: 1em; border-radius: 5px;"
              >
                <label for="uuid_votes">
                  <Field type="radio" name="uuid_votes" value="NON" @change="upvote()" />NON
                </label>
                <span style="padding-left: 60%;">
                  <i class="fas fa-thumbs-up"></i>
                  ({{ counter }})
                </span>
              </div>

              <template v-slot:submit-button>
                <button
                  style="padding:2% 40% 2% 40%; margin-left:5%"
                  type="submit"
                  class="btn btn-primary"
                  v-on:click="UpdatePoll()"
                >JE VOTE</button>
              </template>
            </Formik>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>


<script>
import store from "../../store/index";
import Field from "../../components/Form/Field.vue";
import Formik from "../../components/Form/Formik.vue";

export default {
  components: {
    Formik,
    Field
  },
  data: () => ({
    votes: {
      uuid_votes: []
    },
    counter: 0,
    user: {},
    click: [],
    message: {
      status: false,
      content: ""
    }
  }),

  mounted() {
    this.uuid = this.$route.params.uuid; // id of the sujet
    this.getTopic(this.uuid);
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
    },
    // a revoir
    upvote: function() {
      if (this.click == true) {
        this.counter -= 1;
        this.votes.uuid_votes += this.user.uuid;
      } else {
        this.votes.uuid_votes += this.user.uuid;
        this.click = true;
      }
    },
    downvote: function() {
      if (this.click == true) {
        this.votes.uuid_votes -= 1;
        this.counter += 1;
      } else {
        this.votes.uuid_votes += this.user.uuid;
        this.click = true;
      }
    },

    UpdatePoll: function() {
      (this.votes.uuid_votes += this.user.uuid),
        store
          .dispatch("UpdateVote", JSON.stringify(this.votes))
          .then(votes => {
            console.log(votes);
            this.votes = votes;
          })
          .catch(() => {});
    }
  }
};
</script>

<style scoped>
</style>