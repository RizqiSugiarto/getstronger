<script setup lang="ts">
import { onMounted } from 'vue'
import { logout } from '@/http/requests.ts'
import { useAuthStore } from '@/stores/auth.ts'
import { useNotificationStore } from '@/stores/notifications.ts'
import router from '@/router/router.ts'

onMounted(async () => {
  await logout()
  const authStore = useAuthStore()
  authStore.logout()

  const notificationStore = useNotificationStore()
  notificationStore.unreadCount = 0

  await router.push('/login')
})
</script>

<template></template>

<style scoped></style>
