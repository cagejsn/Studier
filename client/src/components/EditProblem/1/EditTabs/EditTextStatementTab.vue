<template>
    <v-text-field filled label="Problem Statement" v-model="textStatement"></v-text-field>
</template>
<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    
    data: () => ({
        value: ""
    }),
    computed: {
        textStatement: {
            get() {
                return this.$store.state.openProblem!.statement
            },
            set(newValue){
                this.$store.state.openProblem!.statement = newValue
            }
           
        }
    }
})
</script>