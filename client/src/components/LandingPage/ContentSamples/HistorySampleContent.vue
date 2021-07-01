<template>
  <div>
    History
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import { getSampleContent } from "@/api/SampleContent";
export default Vue.extend({
  data: () => ({
    contentToSample: "history",
    sampleContent: {},
    loading: true
  }),

  created() {
    getSampleContent(this.contentToSample).then((response) => {
      this.sampleContent = response.data
          this.loading = false;
    }).catch((error) => {
      console.log(error)
    })

  }
});
</script>