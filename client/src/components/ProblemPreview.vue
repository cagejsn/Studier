<template>
  <v-card style="height: 200px; width: 100%; background-color: #ADD8E6">
    <h1 v-if="loading">???</h1>
    <div v-else style="width: 100%; display: inline-flex">
      <div style="width: 80%">
        <router-link
          :to="{ name: 'SolveProblem', params: { id: problemToSample.id } }"
        >
          <v-card-title primary-title>
            {{ problemToSample.statement.substring(0, 40) }}
          </v-card-title>
        </router-link>

        <p v-if="loading">loading: {{ loading }}</p>
        <v-card-subtitle v-else>
          {{ problemToSample.id }}
        </v-card-subtitle>
      </div>
      <div
        style="width: 20% ; display: flex; justify-content: flex-end; align-items: flex-start;"
      >
        <v-menu bottom offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn icon v-bind="attrs" v-on="on">
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </template>

          <v-list>
            <v-list-item>
              <v-list-item-title>Edit Problem</v-list-item-title>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>Share</v-list-item-title>
            </v-list-item>
            <v-list-item>
              <v-list-item-title>Add to Problem Set</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>
    </div>

    <v-card-text> </v-card-text>

    <v-card-actions> </v-card-actions>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { PropType } from "vue";
import { ProblemMetadata } from "@/models/ProblemSample";
import { Problem } from "@/models/Problem";

export default Vue.extend({
  props: {
    loading: Boolean,
    problemToSample: {
      type: Object as PropType<Problem>,
      required: true
    },
    problemToSampleMetadata: {
      type: Object as PropType<ProblemMetadata>,
      required: false
    }
  },
  methods: {}
});
</script>