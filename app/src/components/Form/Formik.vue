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
  mounted() {
    this.formValues = this.initialValues;
  },
  provide: function() {
    return {
      initialValues: this.initialValues,
      updateFields: this.updateFields
    };
  },
  methods: {
    handleSubmit: function() {
      // remove vue metas objects and emit form data
      this.$emit("on-submit", Object.assign({}, this.formValues));
    },
    updateFields: function(type, name, value) {
      if (type == "checkbox") {
        if (this.formValues.hasOwnProperty(name)) {
          if (this.formValues[name].includes(value)) {
            this.formValues[name] = this.formValues[name].filter(
              formValue => value != formValue
            );
          } else {
            this.formValues[name] = [value, ...this.formValues[name]];
          }
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
