<template>
  <v-card color="green">
    <edit-statement :problem="problem"></edit-statement>

    <edit-answers :problem="problem"></edit-answers>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { getProblemById } from "../api/GetProblem";

import EditStatement from "./EditProblem/EditStatement.vue";
import EditAnswers from "./EditProblem/EditAnswers.vue";
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

  methods: {},

  components: {
    EditStatement,
    EditAnswers
  }
});
</script>
<style scoped>
.v-card {
  margin: 10px;
}
</style>