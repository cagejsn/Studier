<template>
  <div>
    <v-toolbar color="primary" dense>
      <v-spacer>
      </v-spacer>
      <!-- <h1 @click="createNewProblem">Create a New Problem</h1> -->
      <v-toolbar-items>
        <v-menu bottom offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn color="white" icon v-bind="attrs" v-on="on">
              <v-icon>mdi-plus-circle</v-icon>
            </v-btn>
          </template>

          <v-list>
            <v-list-item>
              <v-list-item-title>New Problem</v-list-item-title>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>New Problem Set</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </v-toolbar-items>
    </v-toolbar>

    <problem-list :problems="problems"> </problem-list>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import { getAllProblems } from "../api/GetProblem";

import { Problem } from "../models/Problem";
import { createProblem } from "../api/CreateProblem";
import { AxiosResponse } from "axios";
import ProblemList from "@/components/ProblemList.vue";

export default Vue.extend({
  data: () => ({
    problems: [] as Problem[],
    loading: true
  }),

  async created() {
    const { data } = await getAllProblems();
    this.problems = data;
    this.loading = false;
  },
  methods: {
    createNewProblem() {
      createProblem().then((response: AxiosResponse<Problem>) => {
        const id = response.data.id;
        this.$router.push({ name: "EditProblem", params: { id } });
      });
    }
  },
  components: {
    ProblemList
  }
});
</script>
