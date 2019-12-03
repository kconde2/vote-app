<template>
  <form @submit="handleSubmit" :method="method">
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
      default: "POST",
      type: String
    }
  },
  data: () => ({}),
  methods: {
    handleSubmit: function() {
      this.$emit("onSubmit", this.formValues);
    }
  }
};
</script>

<style>
</style>
