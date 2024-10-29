<template>
  <nav>
    <RouterLink to="/">Home</RouterLink> |
    <RouterLink to="/about">About</RouterLink> |
    <span v-if="isLoggedIn">
       <RouterLink to="/logs">Logs</RouterLink> |
       <button @click="appSignOut">Logout</button>
    </span>
    <span v-else>
      <RouterLink to="/register">Register New User</RouterLink> |
      <RouterLink to="/signin">Sign In</RouterLink>
    </span>
  </nav>
  <RouterView/>
</template>
<script setup lang="ts">
import { ref, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import {getAuth,onAuthStateChanged, signOut} from "firebase/auth";

console.log("Mjupp: ", )
const auth = getAuth()

const router = useRouter()
const isLoggedIn = ref(true)
onAuthStateChanged(auth, function(user) {
  if (user) {
    isLoggedIn.value = true
  } else {
    isLoggedIn.value = false
  }
})
const appSignOut = () => {
  signOut(auth)
  router.push('/')
}

</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #42b983;
}
</style>
