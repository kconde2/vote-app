<template>
  <div class="home">
    <Formik @on-submit="handleSubmit" submit-label="Valider" :initial-values="initialValues">
      <label for="firstname">
        Prénom
        <Field type="text" name="firstname" />
      </label>

      <label for="lastname">
        Nom de famille
        <Field type="text" name="lastname" />
      </label>

      <label for="phone">
        Numéro de téléphone
        <Field type="number" name="phone" />
      </label>

      <label for="color">
        Couleur préférée
        <Field type="select" name="color">
          <option>Selectionnez</option>
          <option value="red">Red</option>
          <option value="green">Green</option>
          <option value="blue">Blue</option>
        </Field>
      </label>

      <label for="gender">
        Sexe
        <Field type="radio" name="gender" value="Homme" />
        <Field type="radio" name="gender" value="Femme" />
      </label>

      <label for="hobbies">
        Hobbies
        <Field type="checkbox" name="hobbies" value="Cinéma" />
        <Field type="checkbox" name="hobbies" value="Sports" />
        <Field type="checkbox" name="hobbies" value="Dance" />
      </label>

      <label for="details">
        Informations supplémentaires
        <Field type="textarea" name="details" />
      </label>
    </Formik>
  </div>
</template>

<script>
import Formik from "../components/Form/Formik.vue";
import Field from "../components/Form/Field.vue";

export default {
  name: "form-use-case",
  components: {
    Field,
    Formik
  },
  data: () => ({
    initialValues: {
      firstname: "Kaba",
      lastname: "CONDE",
      gender: "Homme",
      phone: "012345678",
      hobbies: ["Cinéma", "Dance"]
    }
  }),
  methods: {
    handleSubmit: function(data) {
      console.log("data => ", data);
    }
  }
};
</script>
