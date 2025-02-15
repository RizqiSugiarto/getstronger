<script setup lang="ts">
import { ref } from 'vue'
import { signup } from '@/http/requests'
import { RouterLink, useRouter } from 'vue-router'
import AppButton from '@/ui/components/AppButton.vue'
import { type SignupRequest } from '@/proto/api/v1/auth_service_pb.ts'
import { useAlertStore } from '@/stores/alerts.ts'

const router = useRouter()
const alertStore = useAlertStore()
const req = ref<SignupRequest>({
  $typeName: 'api.v1.SignupRequest',
  email: '',
  firstName: '',
  lastName: '',
  password: '',
  passwordConfirmation: '',
})

const onSignup = async () => {
  const res = await signup(req.value)
  if (!res) return
  alertStore.setSuccess('Please check your inbox to verify your email')
  await router.push('/login')
}
</script>

<template>
  <form class="space-y-6" method="POST" @submit.prevent="onSignup">
    <div>
      <label for="email" class="block font-medium text-gray-900">First name</label>
      <div class="mt-2">
        <input
          id="firstname"
          v-model="req.firstName"
          name="firstname"
          type="text"
          required
          class="block w-full rounded-md border-0 bg-white/5 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600"
        />
      </div>
    </div>

    <div>
      <label for="email" class="block font-medium text-gray-900">Last name</label>
      <div class="mt-2">
        <input v-model="req.lastName" name="lastname" type="text" required />
      </div>
    </div>

    <div>
      <label for="email" class="block font-medium text-gray-900">Email address</label>
      <div class="mt-2">
        <input
          id="email"
          v-model="req.email"
          name="email"
          type="email"
          autocomplete="email"
          required
        />
      </div>
    </div>

    <div>
      <div class="flex items-center justify-between">
        <label for="password" class="block font-medium text-gray-900">Password</label>
      </div>
      <div class="mt-2">
        <input
          id="password"
          v-model="req.password"
          name="password"
          type="password"
          autocomplete="current-password"
          required
        />
      </div>
    </div>

    <div>
      <div class="flex items-center justify-between">
        <label for="password" class="block font-medium text-gray-900"> Confirm Password </label>
      </div>
      <div class="mt-2">
        <input
          id="passwordConfirmation"
          v-model="req.passwordConfirmation"
          name="passwordConfirmation"
          type="password"
          autocomplete="current-password"
          required
        />
      </div>
    </div>

    <div>
      <AppButton type="submit" colour="primary">Sign up</AppButton>
    </div>
  </form>

  <p class="mt-6 text-center text-gray-400">
    Already a member?
    <RouterLink to="/login" class="font-semibold text-indigo-600 hover:text-indigo-500">
      Login
    </RouterLink>
  </p>
</template>

<style scoped>
input {
  @apply block w-full rounded-md border-0 bg-white py-3 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600;
}
</style>
