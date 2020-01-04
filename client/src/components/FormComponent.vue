<template>
  <b-form @submit.prevent="submit">
    <b-form-group
      :key="value.group.id"
      v-for="value in values"
      id="input-group-1"
      :label="value.group.label"
      :label-for="value.group.for"
      :description="value.group.description"
    >
      <b-form-input
        :id="value.id"
        v-model="value.value"
        :type="value.type"
        :required="value.required"
        :placeholder="value.placeholder"
      ></b-form-input>
    </b-form-group>

    <b-button class="b-button" type="submit" variant="primary">Submit</b-button>
    <b-button class="b-button" type="reset" variant="danger">Reset</b-button>
  </b-form>
</template>

<script>
export default {
  name: 'FormComponent',

  props: {
    initialValues: Array,
    onSubmit: Function
  },

  data() {
    return {
      values: []
    }
  },

  created() {
    this.values = this.initialValues;
  },

  methods: {
    submit() {
      const values = {};
      this.values.map(value => values[value.id] = value.value);
      this.onSubmit({ values });
    }
  }
}
</script>

<style>

</style>
