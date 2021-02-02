<template>
  <v-sheet>
    <v-main>
      <v-card>
        <statement></statement>

        <answers></answers>
      </v-card>
      <v-card>
        <v-toolbar dense flat>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <v-btn text @click="submitProblem">submit</v-btn>
          </v-toolbar-items>
        </v-toolbar>
      </v-card>
    </v-main>
  </v-sheet>
</template>
<script lang="ts">
import Vue from "vue";

import Statement from "./SolveProblem/Statement.vue";
import Answers from "./SolveProblem/Answers.vue";
import {saveProblemAttempt, getProblemAttemptById, updateProblemAttempt} from '@/api/ProblemAttempt/ProblemAttempt'
import { ProblemAttemptResponse } from '@/models/UserInteraction';

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
    submitProblem() {
      console.log("submit problem");

      //TODO: put this in the proper place
      saveProblemAttempt()
      .then( (data) => console.log(data))
    }
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