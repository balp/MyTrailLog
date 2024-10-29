import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import { initializeApp } from 'firebase/app'

const config = {
    apiKey: "",
    authDomain: "",
};
export const firebase_app = initializeApp(config);
createApp(App).use(router).mount('#app')

