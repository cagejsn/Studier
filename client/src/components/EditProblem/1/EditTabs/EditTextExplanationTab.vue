<template>
  <v-card>
    <div>{{ this.$store.getters["getSolution"] }}</div>
    <v-textarea
      name="input-7-1"
      label="Explanation"
      v-model="explanation"
      hint="Hint text"
    ></v-textarea>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  computed: {
    explanation: {
      get() {
        return this.$store.state.openProblem!.solution.explanation;
      },
      set(newValue) {
        this.$store.state.openProblem!.solution.explanation = newValue;
      }
    }
  }
});
</script>
