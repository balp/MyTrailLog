<script setup lang="ts">
import {ref} from 'vue'
import {useRouter} from 'vue-router'
import {getAuth, signInWithEmailAndPassword} from 'firebase/auth'

const email = ref('')
const password = ref('')
const errMsg = ref()
const router = useRouter()
const signIn = () => {
  console.log("Mjupp: ",)
  const auth = getAuth()

  signInWithEmailAndPassword(auth, email.value, password.value) // THIS LINE CHANGED
      .then((data) => {
        console.log('Successfully logged in!');
        router.push('/logs')
      })
      .catch(error => {
        switch (error.code) {
          case 'auth/invalid-email':
            errMsg.value = 'Invalid email'
            break
          case 'auth/user-not-found':
            errMsg.value = 'No account with that email was found'
            break
          case 'auth/wrong-password':
            errMsg.value = 'Incorrect password'
            break
          default:
            errMsg.value = 'Email or password was incorrect'
            break
        }
      });
}
</script>

<template>
  <h1>Login to Your Account</h1>
  <p><input type="text" placeholder="Email" v-model="email"/></p>
  <p><input type="password" placeholder="Password" v-model="password"/></p>
  <p v-if="errMsg">{{ errMsg }}</p>
  <p>
    <button @click="signIn">Submit</button>
  </p>
</template>

<style scoped>

</style>