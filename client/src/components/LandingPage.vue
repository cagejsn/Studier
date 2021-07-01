<template>
  <div>
    <v-container>
      <v-row>
        <v-col class="d-flex flex-column" :cols="12">
          <p v-for="entry in Object.entries(contentSamples)" :key="entry[0]">
            <content-sample-brick :title="entry[0]" :contentToSample="entry[1]">
            </content-sample-brick>
            <!-- <router-link :to="{ name: contentSamples[key]', params: { id: prob.id } }">
        {{ prob }}
      </router-link> -->
          </p>
        </v-col>
      </v-row>
    </v-container>
  </div>
</template>
<script lang="ts">
import Vue from "vue";

import ContentSampleBrick from "@/components/LandingPage/ContentSampleBrick.vue";
import { ContentToSample } from "@/models/LandingPage";

export default Vue.extend({
  data: () => ({
    problems: {},
    contentSamples: {
      "Last Problem": {
        routerDest: "last-problem",
        sampleContentType: "last-problem"
      },
      "Saved Problems": {
        routerDest: "saved-problems",
        sampleContentType: "saved-problems"
      },
      "My Problems": {
        routerDest: "my-problems",
        sampleContentType: "my-problems"
      },
      History: {
        routerDest: "history",
        sampleContentType: "history"
      }
    } as Record<string, ContentToSample>
  }),

  components: {
    ContentSampleBrick
  }
});
</script>
