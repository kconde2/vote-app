<template>
  <div class="container">
    <div class="row">
      <div class="col-12">
        <div class="card border-0">
          <div class="card-body">TOPIC EDIT</div>
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
      desc: ""
    },
    errors: {
      status: false,
      message: ""
    }
  }),
  methods: {
    handleSubmit: function(data) {
      this.$v.form.$touch();
      if (this.$v.form.$error) return;
      console.log(data);
      // API call
      store
        .dispatch("editTopic", JSON.stringify(data))
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
