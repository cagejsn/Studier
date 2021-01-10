<template>
  <div>
    Problem Statement
    <v-tabs v-model="statementType">
      <v-tab>
        Text Problem
      </v-tab>
      <v-tab>
        Image Problem
      </v-tab>
      <v-tab>
        Sound Problem
      </v-tab>
    </v-tabs>
    <v-tabs-items v-model="statementType">
      <v-tab-item>
        <edit-text-statement-tab></edit-text-statement-tab>
      </v-tab-item>
      <v-tab-item>
        <edit-image-statement-tab></edit-image-statement-tab>
      </v-tab-item>
      <v-tab-item>
        <edit-sound-statement-tab></edit-sound-statement-tab>
      </v-tab-item>
    </v-tabs-items>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import EditImageStatementTab from "./EditTabs/EditImageStatementTab.vue";
import EditTextStatementTab from "./EditTabs/EditTextStatementTab.vue";
import EditSoundStatementTab from "./EditTabs/EditSoundStatementTab.vue";
import { Problem } from "../../models/Problem";

export default Vue.extend({
  computed: {
    statementType: {
      get(): number {
        console.log(this.$store.getters.getOpenProblem);
        return this.$store.getters.getOpenProblem.type.statementType;
      },
      set(value) {
        this.$store.commit("setProblemStatementType", value);
      }
    }
  },

  data: () => ({
    // statementType: 0
  }),

  methods: {},
  watch: {
    // statementType(old, newer) {
    //   console.log(old, newer);
    // }
  },
  created() {
    // if(this.$store.state.openProblem.type){
    //   this.statementType = this.$store.state.openProblem.type.statementType
    // }

  },
  components: {
    EditImageStatementTab,
    EditTextStatementTab,
    EditSoundStatementTab
  }
});
</script>
<style scoped>
v-card {
  color: aliceblue;
}
</style>
