<template>
  <div class="triplog">
    <h1>Logs</h1>
  </div>
</template>

<script setup lang="ts">
  import {ref} from 'vue'

  var tripData = ref(null)

  async function fetchTrips() {
    tripData.value = null
    const res = await fetch(
        'http://localhost:8081/api/v1/trips'
    )
    tripData.value = await res.json()
  }

  fetchTrips()
</script>

<style scoped>

</style>