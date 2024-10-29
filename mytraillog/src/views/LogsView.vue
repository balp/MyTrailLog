<script setup lang="ts">
import {useRouter} from 'vue-router'
import {onBeforeUnmount} from 'vue'
import {getAuth, onAuthStateChanged} from "firebase/auth";

const router = useRouter()
const auth = getAuth()
const authListener = onAuthStateChanged(auth, function (user) {
  if (!user) {
    alert("You need to be logged in")
    router.push('/signin')
  }
})
onBeforeUnmount(() => {
  authListener()
})
</script>

<template>
  <h1>Trails</h1>
  <h3>This page is for users only</h3>
</template>

<style scoped>

</style>