<script setup lang="ts">
import { RouterView } from 'vue-router'
import AppAlert from '@/ui/components/AppAlert.vue'
import AppNavBottom from '@/ui/components/AppNavBottom.vue'
import AppNavTop from '@/ui/components/AppNavTop.vue'
</script>

<template>
  <div class="pb-12">
    <AppNavTop />
    <AppAlert />
    <main>
      <RouterView />
    </main>
    <AppNavBottom />
  </div>
</template>

<style scoped>
main {
  @apply mx-auto max-w-4xl bg-gray-100 pt-4 pb-8 px-2;
}
</style>
