<template>
  <div v-if="!loading">
    <h1>{{ id }}</h1>
    <edit-problem v-if="problem.state == 'DRAFT'" :problem="problem"></edit-problem>
    <solve-problem v-else-if="problem.state == 'ACTIVE'" :problem="problem"></solve-problem>
  </div>
</template>
<script lang="ts">
import Vue from "vue";
import { getProblemById } from "../api/GetProblem";

import EditProblem from "./EditProblem.vue";
import SolveProblem from "./SolveProblem.vue";

export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    },
    problem() {
      return this.$store.state.openProblem;
    }
  },
  data: () => ({
    loading: true,
  }),

  async created() {
    getProblemById(this.id).then((problem) => {
      this.loading = false;
      this.$store.commit('setOpenProblem', problem.data)
    });
  },

  methods: {},
  components: {
    EditProblem,
    SolveProblem,
  },
});
</script>
