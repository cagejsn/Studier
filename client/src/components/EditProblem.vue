<template>

    <div>
        <edit-statement></edit-statement>

        <edit-answers></edit-answers>

    </div>


</template>
<script lang="ts">

import Vue from "vue";
import {getProblemById} from '../api/GetProblem'

import EditStatement from './EditProblem/EditStatement.vue';
import EditAnswers from './EditProblem/EditAnswers.vue';



export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    }
  },

  data: () => ({
    problem: {}
  }),

  async created(){
    getProblemById(this.id).then(problem => this.problem = problem)
    
  },

  methods: {
    

    
  },

  components: {
      EditStatement,
      EditAnswers
  }
  
});
</script>
