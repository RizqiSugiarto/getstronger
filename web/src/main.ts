import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import piniaPluginPersistedstate from 'pinia-plugin-persistedstate'

import App from './App.vue'
import router from './router/router'
import { RefreshAccessTokenOrLogout, ScheduleTokenRefresh } from '@/jwt/jwt'
import { useAuthStore } from '@/stores/auth'

const app = createApp(App)
const pinia = createPinia()
pinia.use(piniaPluginPersistedstate)

app.use(router)
app.use(pinia)

const init = async () => {
  const authStore = useAuthStore()
  if (authStore.accessToken) {
    await RefreshAccessTokenOrLogout()
    authStore.setAccessTokenRefreshInterval(ScheduleTokenRefresh())
  }

  app.mount('#app')
}

init()
