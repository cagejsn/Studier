<template>
  <v-card color="green">
    <statement></statement>

    <answers></answers>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { getProblemById } from "../api/GetProblem";

import Statement from "./SolveProblem/Statement.vue";
import Answers from "./SolveProblem/Answers.vue";
import { Problem } from "../models/Problem";

export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    },
    problem() {
      return this.$store.state.openProblem;
    }
  },

  data: () => ({}),

  methods: {
  },

  components: {
    Statement,
    Answers
  }
});
</script>
<style scoped>
.v-card {
  margin: 10px;
}
</style>