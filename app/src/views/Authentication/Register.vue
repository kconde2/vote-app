<template>
  <div>
    <div class="auth-wrapper__header auth-wrapper__header--register">
      <span>Inscription</span>
    </div>
    <Formik
      @on-submit="handleSubmit"
      submit-label="Valider"
      :initial-values="form"
      class="auth-wrapper__form"
    >
      <div class="form-group">
        <label for="email">Prénom</label>
        <Field
          type="text"
          name="firstname"
          class="form-control"
          placeholder="Entrez votre prénom"
          :class="{ 'is-invalid': $v.$anyDirty && $v.form.firstname.$error }"
          v-model.trim="$v.form.firstname.$model"
        />
        <div
          v-if="$v.$anyDirty && !$v.form.firstname.required"
          class="invalid-feedback"
        >Votre prénom est obligatoire</div>
      </div>

      <div class="form-group">
        <label for="email">Nom</label>
        <Field
          type="text"
          name="lastname"
          class="form-control"
          placeholder="Entrez votre nom de famille"
          :class="{ 'is-invalid': $v.$anyDirty && $v.form.lastname.$error }"
          v-model.trim="$v.form.lastname.$model"
        />
        <div
          v-if="$v.$anyDirty && !$v.form.lastname.required"
          class="invalid-feedback"
        >Votre nom de famille est obligatoire</div>
      </div>

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
          <span
            v-if="!$v.form.password.minLength"
          >Le mot de passe doit faire au minimum 8 caractères</span>
        </div>
      </div>

      <div class="form-group">
        <label for="passwordConfirm">Mot de passe de confirmation</label>
        <Field
          type="password"
          name="passwordConfirm"
          class="form-control"
          placeholder="Entrez votre mot de passe à nouveau"
          :class="{ 'is-invalid': $v.$anyDirty && $v.form.passwordConfirm.$error }"
          v-model.trim="$v.form.passwordConfirm.$model"
        />
        <div v-if="$v.$anyDirty && $v.form.passwordConfirm.$error" class="invalid-feedback">
          <span
            v-if="!$v.form.passwordConfirm.required"
          >Le mot de passe de confirmation est obligatoire</span>
          <span
            v-else-if="!$v.form.passwordConfirm.sameAsPassword"
          >Les deux mots de passe ne correspondent pas</span>
        </div>
      </div>

      <template v-slot:submit-button>
        <button type="submit" class="btn btn-primary">S'inscrire</button>
        <router-link to="login" class="alert-link ml-2">Se connecter</router-link>
      </template>
    </Formik>
  </div>
</template>

<script>
import Field from "../../components/Form/Field.vue";
import Formik from "../../components/Form/Formik.vue";
import { required, email, minLength, sameAs } from "vuelidate/lib/validators";

export default {
  components: {
    Formik,
    Field
  },
  data: () => ({
    form: {
      firstname: "",
      lastname: "",
      email: "",
      password: "",
      passwordConfirm: ""
    }
  }),
  methods: {
    handleSubmit: function(data) {
      this.$v.form.$touch();
      if (this.$v.form.$error) return;
      // to form submit after this
      alert("Send form to API", JSON.stringify(data));
    }
  },
  validations: {
    form: {
      firstname: { required },
      lastname: { required },
      email: { required, email },
      password: { required, minLength: minLength(8) },
      passwordConfirm: { required, sameAsPassword: sameAs("password") },
    }
  }
};
</script>
