<template>
  <form @submit.prevent="handleSubmit" :method="method">
    <slot></slot>
    <button type="submit">{{ submitLabel }}</button>
  </form>
</template>

<script>
export default {
  props: {
    initialValues: Object,
    submitLabel: String,
    method: {
      default: "post",
      type: String
    }
  },
  data: () => ({
    formValues: {}
  }),
  provide: function() {
    return {
      updateFields: this.updateFields
    };
  },
  methods: {
    handleSubmit: function() {
      this.$emit("onSubmit", this.formValues);
    },
    updateFields: function(type, name, value) {
      if (type == "checkbox") {
        if (this.formValues.hasOwnProperty(name)) {
          this.formValues[name] = [value, ...this.formValues[name]];
        } else {
          this.formValues[name] = [value];
        }
      } else {
        this.formValues[name] = value;
      }
    }
  }
};
</script>

<style></style>
