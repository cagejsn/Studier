<template>
  <div>
      Answers

    <v-tabs v-model="value">
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
    <v-tabs-items v-model="value">
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

import EditMultipleChoiceAnswerTab from './EditTabs/EditMultipleChoiceAnswerTab.vue'
import EditTrueFalseAnswerTab from './EditTabs/EditTrueFalseAnswerTab.vue';
import EditTextEntryAnswerTab from './EditTabs/EditTextEntryAnswerTab.vue';

export default Vue.extend({
  computed: {},

  data: () => ({
      value: 0,

  }),

  methods: {},
  components:{
      EditMultipleChoiceAnswerTab,
      EditTrueFalseAnswerTab,
      EditTextEntryAnswerTab
  }
});
</script>
