<template>
  <v-container>
    <v-row>
      <v-col class="d-flex flex-column" :cols="12">
        <p v-for="problem in problems" :key="problem.id">

            <problem-preview :loading="loading" :problemToSample="problem">
            </problem-preview>

        </p>
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import Vue from "vue";
import { PropType } from "vue";
import { Problem } from "@/models/Problem";
import ProblemPreview from "@/components/ProblemPreview.vue";
export default Vue.extend({
  props: {
    loading: Boolean,
    problems: {
      required: true,
      type: Array as PropType<Problem[]>
    }
  },

  components: {
    ProblemPreview
  }
});
</script>