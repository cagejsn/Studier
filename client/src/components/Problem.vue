<template>
<div>
  <h1> {{ id }} </h1>
  <edit-problem v-if="problem.state == 'DRAFT'"></edit-problem>
  <h1 v-else>Hello</h1>
          </div>

</template>
<script lang="ts">

import Vue from "vue";
import {getProblemById} from '../api/GetProblem'

import EditProblem from './EditProblem.vue'

export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    },
  },
  data: () => ({
    loading: true,
    problem: {}
  }),

  async created(){
    getProblemById(this.id).then(problem => {
      this.loading = false
      this.problem = problem.data
    })
    
  },

  methods: {
    


  },
  components:{
    EditProblem,

  }

  
});
</script>
