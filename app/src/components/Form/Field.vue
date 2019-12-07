<template>
  <div :class="'field field-' + type">
    <input
      v-if="this.regularFieldTypes.includes(type)"
      :type="type"
      :name="name"
      @input="updateFieldData"
      :value="this.initialValues.hasOwnProperty(name) ? this.initialValues[name] : ''"
      v-bind="$attrs"
    />

    <input
      v-if="type == 'checkbox'"
      :type="type"
      :value="value"
      :name="name"
      @input="updateFieldData"
      :checked="this.initialValues.hasOwnProperty(name) && this.initialValues[name].includes(value)"
      v-bind="$attrs"
    />

    <input
      v-if="type == 'radio'"
      :type="type"
      :value="value"
      :name="name"
      @input="updateFieldData"
      :checked="(this.initialValues.hasOwnProperty(name) && this.initialValues[name] == value)"
      v-bind="$attrs"
    />

    <select v-if="type == 'select'" :name="name" @input="updateFieldData" v-bind="$attrs">
      <slot></slot>
    </select>

    <textarea v-if="type == 'textarea'" :name="name" @input="updateFieldData"></textarea>
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
    label: String,
    value: String
  },
  inject: ["initialValues", "updateFields"],
  data: () => ({
    regularFieldTypes: [
      "text",
      "number",
      "button",
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
    fieldTypes: ["select", "textarea", "checkbox", "radio"]
  }),
  mounted() {
    this.fieldTypes = [...this.fieldTypes, ...this.regularFieldTypes];
    if (!this.fieldTypes.includes(this.type)) {
      throw new Error("[Fied Component] |" + this.type + " is not defined");
    }
  },
  methods: {
    updateFieldData: function(event) {
      this.updateFields(this.type, this.name, event.target.value);
    }
  }
};
</script>

<style></style>
