<template>
  <div class="container">
    <div class="row">
      <div class="col-8 mx-auto">
        <div class="card border-0" v-if="!unauthorized">
          <div class="card-body">
            Edit user #{{ form.uuid }}
            <Formik
              @on-submit="handleSubmit"
              submit-label="Valider"
              :initial-values="form"
              class="auth-wrapper__form"
            >
              <div
                v-if="validation.success"
                class="alert alert-success"
                role="alert"
              >{{ validation.message }}</div>

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
                <label for="access_level">Type d'utilisateur</label>
                <Field
                  type="select"
                  name="access_level"
                  class="form-control"
                  v-model.trim="$v.form.access_level.$model"
                >
                  <option>Selectionnez</option>
                  <option value="0">Votant</option>
                  <option value="1">Administrateur</option>
                </Field>
                <div
                  v-if="$v.$anyDirty && !$v.form.access_level.required"
                  class="invalid-feedback"
                >Le niveau d'accès est obligatoire</div>
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

              <template v-slot:submit-button>
                <button type="submit" class="btn btn-success w-100 mt-2">Update</button>
              </template>
            </Formik>
          </div>
        </div>
        <div v-else-if="unauthorized" class="is-invalid">
          Aucun utilisateur trouvé
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import Field from "../../components/Form/Field.vue";
import Formik from "../../components/Form/Formik.vue";
import { email, minLength, required, between } from "vuelidate/lib/validators";
import store from "../../store/index";
import moment from "moment";

export default {
  components: {
    Formik,
    Field
  },
  data: () => ({
    user: {},
    form: {
      uuid: "",
      first_name: "",
      last_name: "",
      birth_date: "",
      access_level: "",
      email: ""
    },
    errors: {
      status: false,
      message: ""
    },
    unauthorized: false,
    validation: {
      success: false,
      message: "Informations modifiée avec succès"
    }
  }),
  beforeMount() {
    this.getUserInfo();
    this.form = {
      first_name: this.user.first_name,
      last_name: this.user.last_name,
      birth_date: this.user.birth_date,
      access_level: this.user.access_level,
      email: this.user.email
    };
  },
  methods: {
    handleSubmit: function(data) {
      this.$v.form.$touch();
      if (this.$v.form.$error) return;

      // format date into correct format
      data.birth_date = moment(data.birth_date).format("DD-MM-YYYY");

      // API call
      store
        .dispatch("updateUser", data)
        .then(() => {
          this.validation.success = true;
          setTimeout(() => {
            this.$router.push({
              name: "user-list"
            });
          }, 3000);
        })
        .catch(error => {
          // handle errors
          this.errors = { status: true, message: error.errors };
        });
    },
    getUserInfo: function() {
      moment.suppressDeprecationWarnings = true;
      store
        .dispatch("getUserInfo", this.$route.params.uuid)
        .then(user => {
          this.form.uuid = user.uuid;
          this.form.first_name = user.first_name;
          this.form.last_name = user.last_name;
          this.form.access_level = user.access_level.toString();
          this.form.birth_date = moment(user.birth_date).format("YYYY-MM-DD");
          this.form.email = user.email;
        })
        .catch(error => {
          if (error.response.data.error == "record not found") {
            this.errors = {
              status: true,
              message: [error.response.data.error]
            };

            this.unauthorized = true;
          }
        });
    }
  },
  validations: {
    form: {
      first_name: { required, minLength: minLength(3) },
      last_name: { required, minLength: minLength(3) },
      birth_date: { required },
      email: { email },
      access_level: { between: between(0, 1) }
    }
  }
};
</script>
