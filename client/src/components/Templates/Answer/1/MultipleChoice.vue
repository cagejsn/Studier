<template>
  <v-container>
    <v-row>
      <v-col class="d-flex flex-column" :cols="1">
        <v-radio-group
          hide-details="true"
          class="radio-group"
          v-model="correctAnswer"
        >
          <v-radio
            class=" radio-item"
            v-for="answerEntry in answersEntries"
            :key="answerEntry[0]"
            :value="answerEntry[0]"
          >
          </v-radio>
        </v-radio-group>
      </v-col>

      <v-col :cols="10">
        <v-row v-for="answerEntry in answersEntries" :key="answerEntry[0]">
          <v-col style="display: flex; align-items: center;">
            <h1 style="text-align: center">{{ answerEntry[0] }}</h1>
          </v-col>

          <v-col :cols="11">
            <v-text-field readonly v-model="answers[answerEntry[0]]"> </v-text-field
          ></v-col>
        </v-row>
      </v-col>
      <v-col class="d-flex flex-column" :cols="1">
        <div></div>
      </v-col>
    </v-row>

  </v-container>
</template>
<script lang="ts">
import Vue from "vue";

import { MultipleChoiceAnswerCandidates } from "@/models/Problem";

export default Vue.extend({
  data: () => ({
    correctAnswer: "",
  
  }),

  computed: {
    answersEntries() {
      return Object.entries(this.answers);
    },
    answers(){
        return this.$store.getters["getAnswers"]
    }
  },

});
</script>
<style scoped>
.radio-group {
  height: 100%;
  margin: 0%;
}

.radio-group >>> .v-input__control {
  height: 100%;
}

.radio-group >>> .v-input__slot {
  height: 100% !important;
}

.radio-group >>> .v-input--radio-group__input {
  height: 100%;
  justify-content: space-around;
}

.radio-item {
  height: 100%;
  margin: 0px !important;
  justify-content: flex-end;
}
</style>