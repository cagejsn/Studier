<template>
  <div>
      Problem Statement
    <v-tabs v-model="value" >
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
    <v-tabs-items v-model="value">
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

export default Vue.extend({
  computed: {},

  data: () => ({
    value: 0
  }),

  methods: {},
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
