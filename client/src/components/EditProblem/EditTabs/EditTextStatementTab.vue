<template>
    <v-text-field filled label="Problem Statement" v-model="value"></v-text-field>
</template>
<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
    
    data: () => ({
        value: ""
    })
})
</script>