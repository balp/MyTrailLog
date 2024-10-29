<script setup lang="ts">
import { ref } from 'vue'
import {useRouter} from "vue-router";
import {getAuth, createUserWithEmailAndPassword} from 'firebase/auth'
const email = ref('')
const password = ref('')
const router = useRouter()
const register = () => {
  console.log("Mjupp: ", )
  const auth = getAuth()
  createUserWithEmailAndPassword(auth, email.value, password.value)
      .then((data) => {
        console.log('Successfully registered!');
        router.push('/logs')
      })
      .catch(error => {
        console.log(error.code);
        alert(error.message)
      })
}
</script>

<template>
  <h1>Create an Account</h1>
  <p><input type="text" placeholder="Email" v-model="email" /></p>
  <p><input type="password" placeholder="Password" v-model="password" /></p>
  <p><button @click="register">Submit</button></p>
</template>

<style scoped>

</style>