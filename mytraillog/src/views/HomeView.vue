<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/my-trail-log-logo-transparent.png">
    <SiteIntro/>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import SiteIntro from '@/components/SiteIntro.vue';

export default defineComponent({
  name: 'HomeView',
  components: {
    SiteIntro
  },
});
</script>
