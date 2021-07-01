<template>
  <div>
    Answers

    <v-tabs v-model="answersType">
      <v-tab>
        Multiple Choice
      </v-tab>
      <v-tab>
        Text Entry
      </v-tab>
      <v-tab>
        True/False
      </v-tab>
    </v-tabs>
    <v-tabs-items v-model="answersType">
      <v-tab-item>
        <edit-multiple-choice-answer-tab></edit-multiple-choice-answer-tab>
      </v-tab-item>
      <v-tab-item>
        <edit-text-entry-answer-tab></edit-text-entry-answer-tab>
      </v-tab-item>
      <v-tab-item>
        <edit-true-false-answer-tab></edit-true-false-answer-tab>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import EditMultipleChoiceAnswerTab from "./EditTabs/EditMultipleChoiceAnswerTab.vue";
import EditTrueFalseAnswerTab from "./EditTabs/EditTrueFalseAnswerTab.vue";
import EditTextEntryAnswerTab from "./EditTabs/EditTextEntryAnswerTab.vue";

import { AnswersType } from "@/models/Problem";

export default Vue.extend({
  data: () => ({
    AnswersType,
    answersType: AnswersType.MULTIPLE_CHOICE
  }),
  // computed: {
  //   answersType: {
  //     get(): AnswersType {
  //       return this.$store.state.openProblem.type.answersType;
  //     },
  //     set(value) {
  //       this.$store.commit("setProblemAnswersType", value);
  //     }
  //   }
  // },

  watch: {
    answersType(newVal, oldVal) {
      this.$store.commit("setProblemAnswersType", newVal);
    }
  },
  created() {
    this.answersType = this.$store.getters["getOpenProblem"].type.answersType;
  },

  components: {
    EditMultipleChoiceAnswerTab,
    EditTrueFalseAnswerTab,
    EditTextEntryAnswerTab
  }
});
</script>
