<template>
  <div v-if="!loading">
    <router-view></router-view>
  </div>
</template>
<script lang="ts">
import Vue from "vue";
import { getProblemById } from "../api/GetProblem";

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
    loading: true
  }),

  async created() {
    getProblemById(this.id).then(problem => {
      this.loading = false;
      this.$store.commit("setOpenProblem", problem.data);
    });
  }
});
</script>
