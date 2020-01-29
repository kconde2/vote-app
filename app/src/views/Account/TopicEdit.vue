<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <div class="card border-0">
          <div class="card-body">Edit topic #{{ topic.uuid }}</div>
          <Formik
            @on-submit="handleSubmit"
            submit-label="Valider"
            :initial-values="form"
            class="auth-wrapper__form"
          >
            <div v-if="errors.status" class="is-invalid">
              <ul v-bind:key="key" v-for="(item, key) in errors.message">
                <li>{{ item }}</li>
              </ul>
            </div>
            <div class="form-group">
              <label for="title">Titre</label>
              <Field
                type="text"
                name="title"
                class="form-control"
                placeholder="Entrez votre Sujet"
                :class="{ 'is-invalid': $v.$anyDirty && $v.form.title.$error }"
                v-model.trim="$v.form.title.$model"
              />
              <div
                v-if="$v.$anyDirty && !$v.form.title.required"
                class="invalid-feedback"
              >Votre titre est obligatoire</div>
            </div>

            <div class="form-group">
              <label for="desc">Description</label>
              <Field
                type="text"
                name="desc"
                class="form-control"
                placeholder="Entrez votre description"
                :class="{ 'is-invalid': $v.$anyDirty && $v.form.desc.$error }"
                v-model.trim="$v.form.desc.$model"
              />
              <div
                v-if="$v.$anyDirty && !$v.form.desc.required"
                class="invalid-feedback"
              >Une description est obligatoire est obligatoire</div>
            </div>

            <template v-slot:submit-button>
              <button type="submit" class="btn btn-primary">Publier</button>
            </template>
          </Formik>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Field from "../../components/Form/Field";
import Formik from "../../components/Form/Formik.vue";
import { required } from "vuelidate/lib/validators";
import store from "../../store";

export default {
  components: {
    Formik,
    Field
  },
  data: () => ({
    form: {
      title: "",
      desc: "",
    },
    topic: {
      title: "",
      desc: "",
      uuid: "",
      uuid_votes: ""
    },
    errors: {
      status: false,
      message: ""
    }
  }),

  mounted() {
    this.uuid = this.$route.params.uuid;
    if (this.uuid) {
      this.getTopic(this.uuid);
    }
  },
  methods: {
    getTopic: function(uuid) {
      store
        .dispatch("getVotes", uuid)
        .then(topic => {
          this.form.title = topic.title;
          this.form.desc = topic.desc;

          this.topic = topic;
        })
        .catch(() => {});
    },

    handleSubmit: function(data) {
      this.topic.uuid_votes = "";
      // API call

      if (this.uuid) {
        this.topic.title = this.form.title;
        this.topic.desc = this.form.desc;

        // update an existing topic
        store
          .dispatch("updateTopic", this.topic)
          .then(() => {
            // redirect to TopicList
            this.$router.push({
              name: "topic-list"
            });
          })
          .catch(error => {
            // handle errors
            this.errors = { status: true, message: error.errors };
          });
      } else {
        // create new topic
        store
          .dispatch("addTopic", data)
          .then(() => {
            // redirect to TopicList
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
  },
  validations: {
    form: {
      title: { required },
      desc: { required }
    }
  }
};
</script>
<style scoped>
.is-invalid {
  color: red;
}
ul {
  list-style: none;
}
</style>
