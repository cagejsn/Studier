<template>
  <v-card style="height: 200px; width: 100%; background-color: #ADD8E6">
    <h1 v-if="loading">???</h1>

    <div v-else class="top-level-grid">
      <div class="card-title">
        <router-link
          :to="{ name: 'SolveProblem', params: { id: problemToSample.id } }"
        >
          <v-card-title primary-title>
            {{ problemToSample.statement.substring(0, 40) }}
          </v-card-title>
        </router-link>
      </div>
      <div class="actions-menu">
        <v-menu bottom offset-y>
          <template v-slot:activator="{ on, attrs }">
            <v-btn icon v-bind="attrs" v-on="on">
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </template>

          <v-list>
            <v-list-item @click="onClickEditProblem()">
              <v-list-item-title>Edit Problem</v-list-item-title>
            </v-list-item>
            <v-list-item @click="onClickShareProblem()">
              <v-list-item-title>Share</v-list-item-title>
            </v-list-item>
            <v-list-item @click="onClickAddToProblemSet()">
              <v-list-item-title>Add to Problem Set</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </div>

      <div class="tags">
        <p v-if="loading">loading: {{ loading }}</p>
        <v-chip-group v-else class="ma-3">
          <v-chip small v-for="tag in getTags()" :key="tag">
            {{ tag }}
          </v-chip>
        </v-chip-group>
      </div>
      <div class="rating">
        <v-rating v-model="rating"> </v-rating>
      </div>
      <div class="author">
        <v-card-subtitle style="text-align: right; width: 100%">
          Author
        </v-card-subtitle>
      </div>
    </div>
  </v-card>
</template>
<script lang="ts">
import Vue from "vue";
import { PropType } from "vue";
import { ProblemMetadata } from "@/models/ProblemSample";
import { Problem } from "@/models/Problem";

export default Vue.extend({
  data: () => ({
    rating: 1
  }),
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
  methods: {
    onClickEditProblem() {
      console.log("edit problem");
    },
    onClickShareProblem() {
      console.log("share problem");
    },
    onClickAddToProblemSet() {
      console.log("add to problem set");
    },
    getTags() {
      return ["this", "that", "other"];
      //return this.problemToSampleMetadata.tags
    }
  }
});
</script>
<style>
.top-level-grid {
  width: 100%;
  height: 100%;
  display: grid;
  grid-template-columns: 1fr 1fr 1fr 1fr 36px;
  grid-template-rows: 1fr 1fr 1fr 1fr 54px;
}

.card-title {
  grid-column-start: 1;
  grid-column-end: 4;
  grid-row-start: 1;
  grid-row-end: 3;
}

.rating {
  grid-column-start: 1;
  grid-column-end: 4;
  grid-row-start: 4;
  grid-row-end: 4;
}

.tags {
  grid-column-start: 1;
  grid-column-end: 2;
  grid-row-start: 5;
  grid-row-end: 6;
}

.actions-menu {
  grid-column-start: 5;
  grid-column-end: 6;
  grid-row-start: 1;
  grid-row-end: 4;
}

.author {
  grid-column-start: 4;
  grid-column-end: 6;
  grid-row-start: 5;
  grid-row-end: 6;
}
</style>