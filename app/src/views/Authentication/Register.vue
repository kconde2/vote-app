<template>
  <div class="auth-wrapper">
    <div class="auth-wrapper__content shadow-sm">
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
          <div v-if="errors.status" class="is-invalid">
            <ul v-bind:key="key" v-for="(item, key) in errors.message">
              <li>{{ item }}</li>
            </ul>
          </div>
          <div class="form-group">
            <label for="first_name">Prénom</label>
            <Field
              type="text"
              name="first_name"
              class="form-control"
              placeholder="Entrez votre prénom"
              :class="{ 'is-invalid': $v.$anyDirty && $v.form.first_name.$error }"
              v-model.trim="$v.form.first_name.$model"
            />
            <div
              v-if="$v.$anyDirty && !$v.form.first_name.required"
              class="invalid-feedback"
            >Votre prénom est obligatoire</div>
          </div>

          <div class="form-group">
            <label for="last_name">Nom</label>
            <Field
              type="text"
              name="last_name"
              class="form-control"
              placeholder="Entrez votre nom de famille"
              :class="{ 'is-invalid': $v.$anyDirty && $v.form.last_name.$error }"
              v-model.trim="$v.form.last_name.$model"
            />
            <div
              v-if="$v.$anyDirty && !$v.form.last_name.required"
              class="invalid-feedback"
            >Votre nom de famille est obligatoire</div>
          </div>
          <div class="form-group">
            <label for="birth_date">Date de naissance</label>
            <Field
              type="date"
              name="birth_date"
              class="form-control"
              placeholder="Entrez votre nom de famille"
              :class="{ 'is-invalid': $v.$anyDirty && $v.form.birth_date.$error }"
              v-model.trim="$v.form.birth_date.$model"
            />
            <div
              v-if="$v.$anyDirty && !$v.form.birth_date.required"
              class="invalid-feedback"
            >Votre date de naissance est obligatoire</div>
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
            <label for="pass">Mot de passe</label>
            <Field
              type="password"
              name="pass"
              class="form-control"
              placeholder="Entrez votre mot de passe"
              :class="{ 'is-invalid': $v.$anyDirty && $v.form.pass.$error }"
              v-model.trim="$v.form.pass.$model"
            />
            <div v-if="$v.$anyDirty && $v.form.pass.$error" class="invalid-feedback">
              <span v-if="!$v.form.pass.required">Le mot de passe est obligatoire</span>
              <span
                v-if="!$v.form.pass.minLength"
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
    </div>
  </div>
</template>

<script>
import Field from "../../components/Form/Field.vue";
import Formik from "../../components/Form/Formik.vue";
import { required, email, minLength, sameAs } from "vuelidate/lib/validators";
import store from "../../store/index";
import moment from "moment";

export default {
  components: {
    Formik,
    Field
  },
  data: () => ({
    form: {
      first_name: "",
      last_name: "",
      birth_date: "",
      email: "",
      pass: "",
      passwordConfirm: "",
      access_level: "0"
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

      // format date into correct format
      data.birth_date = moment(data.birth_date).format("DD-MM-YYYY");

      // API call
      store
        .dispatch("register", JSON.stringify(data))
        .then(() => {
          // redirect to dashboard
          this.$router.push({
            name: "dashboard"
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
      first_name: { required },
      last_name: { required },
      birth_date: { required },
      email: { required, email },
      pass: { required, minLength: minLength(1) },
      passwordConfirm: { required, sameAsPassword: sameAs("pass") }
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
