

<template>
  <div class="min-h-screen bg-gray-100 flex flex-col items-center justify-center py-12 px-4">
    <div class="max-w-lg w-full bg-white rounded-2xl shadow-lg p-8">
      <h1 class="text-2xl font-bold text-gray-800 mb-6">Register</h1>

      <!-- Step indicators -->
      <div class="flex justify-between mb-6">
        <span :class="step === 1 ? activeStep : stepStyle">Step 1</span>
        <span :class="step === 2 ? activeStep : stepStyle">Step 2</span>
      </div>

      <!-- Step 1 -->
      <div class="space-y-4" v-if="step === 1">
      <FormInput
  v-for="key in step1Keys"
  :key="key"
  :keyName="key"
  :modelValue="formData[key]"
  @update:modelValue="(value: string) => (formData[key] = value)"
  :error="errors[key]"
/>

      </div>

      <!-- Step 2 -->
      <div class="space-y-4" v-if="step === 2">
        <FormInput
  v-for="key in step2Keys"
  :key="key"
  :keyName="key"
  :modelValue="formData[key]"
  @update:modelValue="(value: string) => (formData[key] = value)"
  :error="errors[key]"
/>

      </div>

      <!-- Buttons -->
      <div class="flex justify-between pt-6">
        <button
          v-if="step > 1"
          @click="step--"
          class="bg-gray-300 text-gray-800 px-4 py-2 rounded hover:bg-gray-400"
        >
          Back
        </button>

        <button
          v-if="step < 2"
          @click="nextStep"
          class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 ml-auto"
        >
          Next
        </button>

        <button
          v-if="step === 2"
          @click="register"
          :disabled="mutationLoading"
          class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 ml-auto disabled:opacity-50"
        >
          {{ mutationLoading ? 'Registering...' : 'Register' }}
        </button>
      </div>

      <!-- Feedback -->
      <p v-if="mutationError" class="text-red-500 text-sm mt-2">{{ errorMessage }}</p>
      <p v-if="success" class="text-green-500 text-sm mt-2">Registration successful!</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
const client = useUrqlClient()

const router = useRouter()

const role ='talent'
const id='1'
// Step keys as const arrays
const step1Keys = ['first_name', 'last_name', 'email', 'password'] as const
const step2Keys = [ 'address', 'phone_number'] as const
type Step1Key = typeof step1Keys[number]
type Step2Key = typeof step2Keys[number]
type RegisterKey = Step1Key | Step2Key

// GraphQL mutation
const REGISTER = gql`
  mutation Register($input: NewUser!) {
    Register(input: $input) {
      id
      first_name
      last_name
      email
      role
      address
      phone_number
    }
  }
`

// Step state
const step = ref(1)
const mutationLoading = ref(false)
const mutationError = ref<CombinedError | null>(null)
const success = ref(false)

// Form state
const formData = reactive<Record<RegisterKey, string>>({
  first_name: '',
  last_name: '',
  email: '',


  password: '',
  address: '',
  phone_number: ''
})

const errors = reactive<Record<RegisterKey, string>>({
  first_name: '',
  last_name: '',
  email: '',
  password: '',
  address: '',
  phone_number: ''
})

const errorMessage = computed(() => mutationError.value?.message || '')

function validateStep(stepNum: number): boolean {
  const keys = stepNum === 1 ? step1Keys : step2Keys
  keys.forEach((k) => (errors[k] = ''))

  let isValid = true
  for (const k of keys) {
    const val = formData[k]
    if (!val) {
      errors[k] = `${k.replace(/_/g, ' ')} is required`
      isValid = false
    }

    if (k === 'email' && val && !/^\S+@\S+\.\S+$/.test(val)) {
      errors[k] = 'Valid email is required'
      isValid = false
    }

    if (k === 'password' && val && val.length < 6) {
      errors[k] = 'Password must be at least 6 characters'
      isValid = false
    }

    if (k === 'phone_number' && val && !/^\d{10,}$/.test(val)) {
      errors[k] = 'Valid phone number is required'
      isValid = false
    }
  }

  return isValid
}

function nextStep() {
  if (validateStep(step.value)) step.value++
}

async function register() {
  if (!validateStep(2)) return

  mutationLoading.value = true
  mutationError.value = null
  success.value = false


  try {
    const result = await client
      .mutation(REGISTER, { input: { ...formData,role,id } })
      .toPromise()

    if (result.data?.Register?.id) {
      success.value = true
      ;(Object.keys(formData) as RegisterKey[]).forEach((k) => (formData[k] = ''))
      step.value = 1
    } else if (result.error) {
      mutationError.value = result.error
    }
      if (result.data?.Register?.id) {
  success.value = true
  Object.keys(formData).forEach((k) => (formData[k as RegisterKey] = ''))
  step.value = 1

  // Navigate to login page after short delay
  setTimeout(() => {
    router.push('/signin')
  }, 1000)
}
  } catch (err: any) {
    mutationError.value = err
    console.error('GraphQL Error:', err)
  } finally {
    mutationLoading.value = false
  }

}



const activeStep = 'text-blue-600 font-bold underline'
const stepStyle = 'text-gray-400'
</script>

