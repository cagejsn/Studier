<template>
  <v-sheet>
    <v-main>
      <v-card>
        <statement></statement>

        <answers @user-selected-answer="setUserInput"></answers>
      </v-card>

      <v-card
        v-if="
          problemAttemptSubmissionState ==
            ProblemAttemptSubmissionState.PENDING_SUBMISSION
        "
      >
        <v-toolbar dense flat>
          <v-spacer></v-spacer>
          <v-toolbar-items>
            <!-- this section hinges on whether the user has submitted their
            attempt for grading
            graded user submission / user can submit -->
            <v-btn
              text
              :disabled="!problemAttemptUserInput"
              @click="submitProblemAttempt"
              >submit</v-btn
            >
          </v-toolbar-items>
        </v-toolbar>
      </v-card>

      <v-card
        v-if="
          problemAttemptSubmissionState ==
            ProblemAttemptSubmissionState.SUBMITTED_PENDING_GRADE
        "
      >
        You Submitted: {{ problemAttemptUserInput }}
      </v-card>

      <v-card
        v-if="
          problemAttemptSubmissionState ==
            ProblemAttemptSubmissionState.SUBMITTED_GRADED
        "
      >
        <problem-attempt-outcome></problem-attempt-outcome>
      </v-card>
    </v-main>
  </v-sheet>
</template>
<script lang="ts">
import Vue from "vue";

import Statement from "./SolveProblem/Statement.vue";
import Answers from "./SolveProblem/Answers.vue";
import { ProblemAttemptSubmissionState } from "@/models/ProblemAttempt";
import ProblemAttemptOutcome from "./SolveProblem/ProblemAttemptOutcome.vue";

export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    },
    problem() {
      return this.$store.state.openProblem;
    },
    problemAttempt() {
      return this.$store.getters["getProblemAttempt"];
    },
    problemAttemptSubmissionState(): ProblemAttemptSubmissionState {
      return this.$store.getters["getProblemAttemptSubmissionState"];
    },
    problemAttemptUserInput() {
      return this.$store.getters["getProblemAttemptUserInput"];
    }
  },

  data: () => ({
    ProblemAttemptSubmissionState
  }),

  methods: {
    beginAttemptProblem() {
      this.$store.dispatch("beginAttemptProblem", this.id);
    },
    setUserInput(userInput: any) {
      this.$store.dispatch("updateProblemAttempt", userInput);
    },
    submitProblemAttempt() {
      this.$store.dispatch("submitProblemAttempt");
    }
  },

  created() {
    this.beginAttemptProblem();
  },

  components: {
    Statement,
    Answers,
    ProblemAttemptOutcome
  }
});
</script>
<style scoped>
.v-card {
  margin: 10px;
}
</style>
