<template>
  <div>
    <div class="auth-wrapper__header auth-wrapper__header--login">Connexion</div>
    <Formik
      @on-submit="handleSubmit"
      submit-label="Valider"
      :initial-values="form"
      class="auth-wrapper__form"
    >
      <div v-if="error" class="is-invalid">Invalid Credentials</div>
      <div class="form-group">
        <label for="email">Email</label>
        <Field
          type="email"
          name="email"
          class="form-control"
          placeholder="Entrez votre email"
          :class="{ 'is-invalid': $v.$anyDirty && $v.form.email.$error }"
          v-model.trim="$v.form.email.$model"
        />
        <div v-if="$v.$anyDirty && $v.form.email.$error" class="invalid-feedback">
          <span v-if="!$v.form.email.required">L'email est obligatoire</span>
          <span v-else-if="!$v.form.email.email">Ce champ doit être un email valide</span>
        </div>
      </div>

      <div class="form-group">
        <label for="password">Mot de passe</label>
        <Field
          type="password"
          name="password"
          class="form-control"
          placeholder="Entrez votre mot de passe"
          :class="{ 'is-invalid': $v.$anyDirty && $v.form.password.$error }"
          v-model.trim="$v.form.password.$model"
        />
        <div v-if="$v.$anyDirty && $v.form.password.$error" class="invalid-feedback">
          <span v-if="!$v.form.password.required">Le mot de passe est obligatoire</span>
        </div>
      </div>

      <template v-slot:submit-button>
        <button type="submit" class="btn btn-primary">Se connecter</button>
      </template>
    </Formik>
  </div>
</template>

<script>
import Field from "../../components/Form/Field.vue";
import Formik from "../../components/Form/Formik.vue";
import { required, email } from "vuelidate/lib/validators";
import store from "../../store/index";
import auth from "../../utils/auth";

export default {
  components: {
    Formik,
    Field
  },
  data: () => ({
    form: {
      email: "",
      password: ""
    },
    error: false
  }),
  methods: {
    handleSubmit: function(data) {
      this.$v.form.$touch();
      if (this.$v.form.$error) return;

      // API call
      store
        .dispatch("login", JSON.stringify(data))
        .then(() => {
          auth.runLoginDurationCheck();
          // redirect to dashboard
          this.$router.push({
            name: "dashboard"
          });
        })
        .catch(() => {
          // handle errors
          this.error = true;
        });
    }
  },
  validations: {
    form: {
      email: { required, email },
      password: { required }
    }
  }
};
</script>
<style scoped>
.is-invalid {
  color: red;
}
</style>
