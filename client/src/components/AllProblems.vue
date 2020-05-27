<template>
  <div>
    <p v-for="(prob) in problems" :key=prob>
      <router-link :to="{ name: 'Problem', params: {id: prob.Id}}">
        {{ prob }}
      </router-link>
    </p>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import {getAllProblems} from '../api/GetProblem';

import {Problem} from '../models/Problem';

export default Vue.extend({

data: () => ({
  problems: {}
}),

async created() {
  const {data} = await getAllProblems();
  this.problems = data;
},

});
</script>
