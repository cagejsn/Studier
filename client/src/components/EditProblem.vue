<template>
  <v-card color="green">
    <v-toolbar>
      <v-toolbar-items>
        <v-btn @click="saveProblem">save</v-btn>
        <v-switch :label="problemState" :input-value="isProblemActive" @change="problemStateChanged"></v-switch>
      </v-toolbar-items>
    </v-toolbar>

    <edit-statement></edit-statement>

    <edit-answers></edit-answers>
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
    },
    problemState(){
      return this.$store.getters["getProblemState"]
    },
    isProblemActive() {
      return this.$store.getters["getProblemState"] === "ACTIVED";
    }
  },

  data: () => ({}),

  methods: {
    saveProblem() {
      this.$store.dispatch("saveProblem");
    },
    problemStateChanged(value: boolean) {
      if(value){
        this.$store.dispatch('updateProblemState', 'ACTIVED')
      } else {
        this.$store.dispatch('updateProblemState', 'DRAFT')
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