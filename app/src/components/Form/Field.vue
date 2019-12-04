<template>
  <div :class="'field field-' + type">
    <label v-if="this.inputTypes.includes(type)">
      <span v-if="label">{{ label }}</span>
      <input :type="type" :name="name" />
    </label>

    <label v-if="type == 'select'">
      <span v-if="label">{{ label }}</span>
      <select :name="name">
        <slot></slot>
      </select>
    </label>

    <label v-if="type == 'textarea'">
      <span v-if="label">{{ label }}</span>
      <textarea :name="name"></textarea>
    </label>
  </div>
</template>

<script>
export default {
  props: {
    name: String,
    type: {
      default: "input",
      type: String
    },
    label: String
  },
  data: () => ({
    values: {},
    inputTypes: [
      "text",
      "number",
      "button",
      "checkbox",
      "color",
      "date",
      "datetime-local",
      "email",
      "file",
      "hidden",
      "image",
      "month",
      "number",
      "pasword",
      "radio",
      "range",
      "reset",
      "search",
      "submit",
      "tel",
      "text",
      "time",
      "url",
      "week"
    ],
    fieldTypes: ["select", "textarea"]
  }),
  mounted() {
    this.fieldTypes = [...this.fieldTypes, ...this.inputTypes];

    if (!this.fieldTypes.includes(this.type)) {
      throw new Error("[Fied Component] |" + this.type + " is not defined");
    }
  }
};
</script>

<style></style>
