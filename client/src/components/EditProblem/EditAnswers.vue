<template>
  <div>
      Answers

    <v-tabs v-model="value" color="primary" slider-color="primary">
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

export default Vue.extend({
  computed: {},

  data: () => ({
      value: 0,

  }),

  methods: {}
});
</script>
