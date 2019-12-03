<template>
  <div id="app">
    <Formik @onSubmit="handleSubmit" submitLabel="Valider">
      <Field type="text" name="firstname" label="Prénom" />
      <Field type="number" name="phone" label="Numéro de téléphone" />
      <Field type="select" name="age" label="Couleur préférée">
        <option value="red">Red</option>
        <option value="green">Green</option>
        <option value="blue">Blue</option>
      </Field>
      <Field type="textarea" name="comment" label="Commentaires" />
    </Formik>
  </div>
</template>

<script>
import Formik from "./components/Form/Formik.vue";
import Field from "./components/Form/Field.vue";

export default {
  name: "app",
  components: {
    Field,
    Formik
  },
  methods: {
    handleSubmit: function(values) {
      alert(JSON.stringify(values));
    },
    init() {
      console.log(this);
    }
  },
  mounted: () => {
    this.init();
  }
};
</script>
