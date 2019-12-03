<template>
  <div :class="'field field-' + type">
    <label v-if="this.inputTypes.includes(type)">
      <span>{{ label }}</span>
      <input :type="type" :name="name" />
    </label>

    <label v-if="type == 'select'">
      <span>{{ label }}</span>
      <select :name="name">
        <slot></slot>
      </select>
    </label>

    <label v-if="type == 'textarea'">
      <span>{{ label }}</span>
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
    fieldTypes: ["input", "select", "textarea"]
  }),
  mounted: () => {
    //console.log("beforeCreate");
    // if (this.fieldTypes.indexOf(type) < 0) {
    //   throw new Error("[Fied Component]");
    // }
  }
};
</script>

<style>
</style>
