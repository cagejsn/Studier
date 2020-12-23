<template>
  <div>
    <h1 @click="createNewProblem">Create a New Problem</h1>
    <p v-for="prob in problems" :key="prob.id">
      <router-link :to="{ name: 'Problem', params: { id: prob.id } }">
        {{ prob }}
      </router-link>
    </p>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import { getAllProblems } from "../api/GetProblem";

import { Problem } from "../models/Problem";
import { createProblem } from "../api/CreateProblem";
import { AxiosResponse } from "axios";

export default Vue.extend({
  data: () => ({
    problems: {}
  }),

  async created() {
    const { data } = await getAllProblems();
    this.problems = data;
  },
  methods: {
    createNewProblem() {

      this.$store.

      createProblem().then((response: AxiosResponse<Problem>) => {
        const id = response.data.id;
        this.$router.push({ name: "Problem", params: { id } });
      });
    }
  }
});
</script>
