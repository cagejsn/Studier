<template>
  <component @user-selected-answer="userSelectedAnswer" v-bind:is="answersComponent"></component>
</template>
<script lang="ts">
import Vue from "vue";
import {AnswersType} from "@/models/Problem"

import * as v1Templates from "../Templates/Answer/1"
export default Vue.extend({
  data: () => ({
    AnswersType,
  }),
  computed: {
    answersComponent(): string {
      const answersType = this.$store.state.openProblem.type.answersType
      return Object.values(AnswersType)[answersType] + "_" + this.$store.state.openProblem.version;
    },

  },

  methods:{
    userSelectedAnswer(payload: any){
      this.$emit('user-selected-answer', payload)
    }
  },

  // is there a way to ...expand the object of the v1Templates and reduce the amount of code?
  components: {
    MULTIPLE_CHOICE_1: v1Templates.MultipleChoice,
    TRUE_FALSE_1: v1Templates.TrueFalse,
    TEXT_ENTRY_1: v1Templates.TextEntry,
  }
});
</script>