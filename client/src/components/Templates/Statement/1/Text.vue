<template>
  <v-text-field
    filled
    readonly
    label="Problem Statement"
    v-model="textStatement"
  ></v-text-field>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  data: () => ({
    value: ""
  }),
  computed: {
    textStatement() {
      return this.$store.state.openProblem!.statement;
    }
  }
});
</script>