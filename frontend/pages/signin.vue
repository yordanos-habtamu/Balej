<template>
  <!-- <div class="flex justify-evenly fixed top-0 z-50 w-full bg-white shadow">
    <Logo />
    <Menu />
    <Action />
  </div> -->

  <div class="min-h-screen bg-gray-100 flex flex-col items-center justify-center py-12 px-4">
    <div class="max-w-lg w-full bg-white rounded-2xl shadow-lg p-8">
      <h1 class="text-2xl font-bold text-gray-800 mb-6">Sign In</h1>
      <div class="space-y-4">
        <input
          type="email"
          v-model="formData.email"
          placeholder="Enter your email"
          class="mt-1 p-3 w-full border border-gray-300outes — Nuxt does it based on file names.
 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <p v-if="errors.email" class="text-sm text-red-600">{{ errors.email }}</p>

        <input
          type="password"
          v-model="formData.password"
          placeholder="Enter your password"
          class="mt-1 p-3 w-full border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
        />
        <p v-if="errors.password" class="text-sm text-red-600">{{ errors.password }}</p>
       <div>
        <button
          @click="login"
          :disabled="mutationLoading"
          class="w-full bg-blue-600 text-white px-4 py-3 rounded-lg hover:bg-blue-700 transition duration-300 disabled:opacity-60"
        >
          {{ mutationLoading ? 'Signing In...' : 'Sign In' }}
        </button>
        <nuxt-link to="/accounts">
          <button class= " m-2 w-full bg-white border-2 border-black text-black py-3 rounded-lg hover:bg-black hover:text-white transition-all duration-100">Create Account</button>
        </nuxt-link>
        
       </div>
        

        <p v-if="mutationError" class="text-red-500 text-sm mt-2">{{ errorMessage }}</p>
        <p v-if="success" class="text-green-600 text-sm mt-2">Login successful!</p>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {jwtDecode} from 'jwt-decode'
import { useAuthStore } from '~/stores/auth'
import { ref, reactive, computed } from 'vue'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'

const tokenCookie = useCookie('token')
const client = useUrqlClient()
interface LoginForm {
  email: string
  password: string
}

const LoginMutation = gql`
  mutation LoginUser($input: LoginInput!) {
    LoginUser(input: $input) {
      token
      user {
        id
        first_name
        last_name
        email
        role
      }
    }
  }
`

const formData = reactive<LoginForm>({
  email: '',
  password: '',
})

const errors = reactive<Record<keyof LoginForm, string>>({
  email: '',
  password: '',
})

const mutationLoading = ref(false)
const mutationError = ref<CombinedError | null>(null)
const success = ref(false)
const token = ref('')
const errorMessage = computed(() => mutationError.value?.message || '')

const validateForm = (): boolean => {
  Object.keys(errors).forEach((key) => (errors[key as keyof LoginForm] = ''))
  let isValid = true

  if (!formData.email || !/^\S+@\S+\.\S+$/.test(formData.email)) {
    errors.email = 'Valid email is required'
    isValid = false
  }

  if (!formData.password || formData.password.length < 6) {
    errors.password = 'Password must be at least 6 characters'
    isValid = false
  }

  return isValid
}

const login = async () => {
  if (!validateForm()) return

  mutationLoading.value = true
  mutationError.value = null
  success.value = false

  try {
    const result = await client
      .mutation(LoginMutation, { input: { ...formData } })
      .toPromise()

    if (result.data?.LoginUser?.token) {
      token.value = result.data.LoginUser.token
      tokenCookie.value = token.value

      
      const decoded: any = jwtDecode(token.value)
      const auth = useAuthStore()
      console.log(decoded)
      auth.setUser(decoded)
      console.log(auth.user?.user_id)

      // ✅ Navigate based on role
      if (decoded.role === 'admin') {
        await navigateTo('/admin')
      } else if (decoded.role === 'talent') {
        await navigateTo('/talent')
      }else if(decoded.role ==='client'){
        await navigateTo('/client')
      }

      // Clear form and show success
      Object.keys(formData).forEach((key) => (formData[key as keyof LoginForm] = ''))
      success.value = true
    } else if (result.error) {
      mutationError.value = result.error
    }
  } catch (err: any) {
    mutationError.value = err
  } finally {
    mutationLoading.value = false
  }
}

</script>
