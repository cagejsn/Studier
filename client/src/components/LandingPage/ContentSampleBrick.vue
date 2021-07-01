<template>
  <router-link :to="getRouterLink()">
    <div style="height: 200px; width: 100%; background-color: #ADD8E6">
      <h1>{{ title }}</h1>
      <component v-bind:is="sampleContentComponent"></component>
    </div>
  </router-link>
</template>
<script lang="ts">
import Vue from "vue";
import { PropType } from "vue";
import { ContentToSample } from "@/models/LandingPage";
import * as ContentSamples from "./ContentSamples";

export default Vue.extend({
  computed: {
    sampleContentComponent() {
      return this.contentToSample.sampleContentType;
    }
  },
  props: {
    title: String,
    contentToSample: {
      type: Object as PropType<ContentToSample>,
      required: true
    }
  },
  methods: {
    getRouterLink(): string {
      return this.contentToSample.routerDest;
    }
  },
  components: {
    "history": ContentSamples.HistorySampleContent,
    "last-problem": ContentSamples.LastProblemSampleContent,
    "saved-problems": ContentSamples.SavedProblemsSampleContent,
    "my-problems": ContentSamples.MyProblemsSampleContent
  }
});
</script>
<style lang="scss">
</style>