<template>
  <v-sheet>
    <v-content>
      <v-toolbar>
        <v-toolbar-items>
          <v-btn @click="saveProblem">save</v-btn>
          <v-switch
            :label="problemState"
            :input-value="isProblemActive"
            @change="problemStateChanged"
          ></v-switch>
        </v-toolbar-items>
      </v-toolbar>

      <v-card v-if="problemVersion == 1">
        <edit-statement></edit-statement>
        <edit-answers></edit-answers>
      </v-card>
    </v-content>
  </v-sheet>
</template>
<script lang="ts">
import Vue from "vue";
import { getProblemById } from "../api/GetProblem";

import EditStatement from "./EditProblem/1/EditStatement.vue";
import EditAnswers from "./EditProblem/1/EditAnswers.vue";
import { Problem } from "../models/Problem";

export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    },
    problem() {
      return this.$store.state.openProblem;
    },
    problemState() {
      return this.$store.getters["getProblemState"];
    },
    isProblemActive() {
      return this.$store.getters["getProblemState"] === "ACTIVE";
    },
    problemVersion() {
      return this.$store.getters["getProblemVersion"];
    }
  },

  data: () => ({}),

  methods: {
    saveProblem() {
      this.$store.dispatch("saveProblem");
    },
    problemStateChanged(value: boolean) {
      if (value) {
        this.$store.dispatch("updateProblemState", "ACTIVE");
      } else {
        this.$store.dispatch("updateProblemState", "DRAFT");
      }
    }
  },

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