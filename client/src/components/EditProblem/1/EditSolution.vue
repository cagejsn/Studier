<template>
  <div>
    Solution
    <v-tabs v-model="solutionType">
      <v-tab>
        No Explanation
      </v-tab>
      <v-tab>
        Text Explanation
      </v-tab>
    </v-tabs>
    <v-tabs-items v-model="solutionType">
      <v-tab-item>
        <div></div>
      </v-tab-item>
      <v-tab-item>
        <edit-text-explanation-tab></edit-text-explanation-tab>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import EditTextExplanationTab from "./EditTabs/EditTextExplanationTab.vue";


import {SolutionType} from '@/models/Problem'

export default Vue.extend({
  data: () => ({
    SolutionType,
    solutionType: SolutionType.NO_EXPLANATION,
  }),

  watch: {
    solutionType(newVal, oldVal){
      this.$store.commit("setProblemSolutionType", newVal);
    }
  },
  created() {
    this.solutionType = this.$store.getters["getOpenProblem"].type.solutionType
  },

  components: {
    EditTextExplanationTab
  }
});
</script>
