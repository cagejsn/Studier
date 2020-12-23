<template>
  <v-card color="green" >
    <edit-statement></edit-statement>

    <edit-answers></edit-answers>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { getProblemById } from "../api/GetProblem";

import EditStatement from "./EditProblem/EditStatement.vue";
import EditAnswers from "./EditProblem/EditAnswers.vue";

export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    }
  },

  data: () => ({
    problem: {}
  }),

  async created() {
    getProblemById(this.id).then(problem => (this.problem = problem));
  },

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