<template>
  <component v-bind:is="statementComponent"></component>
</template>
<script lang="ts">
import Vue from "vue";
import { StatementType } from "@/models/Problem";

import * as v1Templates from "../Templates/Statement/1";
export default Vue.extend({
  data: () => ({
    StatementType
  }),
  computed: {
    statementComponent(): string {
      const statementType = this.$store.state.openProblem.type.statementType;
      return (
        Object.values(StatementType)[statementType] +
        "_" +
        this.$store.state.openProblem.version
      );
    }
  },

  // is there a way to ...expand the object of the v1Templates and reduce the amount of code?
  components: {
    AUDIO_1: v1Templates.Audio,
    IMAGE_1: v1Templates.Image,
    TEXT_1: v1Templates.Text
  }
});
</script>
