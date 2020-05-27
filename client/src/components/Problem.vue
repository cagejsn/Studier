<template>
  <div>{{ id }}</div>
</template>
<script lang="ts">
import Vue from "vue";
export default Vue.extend({
  computed: {
    id() {
      return this.$route.params.id;
    }
  }
});
</script>
