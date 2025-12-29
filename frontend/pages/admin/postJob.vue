<template>
  <div class="min-h-screen bg-gray-100 flex flex-col items-center justify-center py-12 px-4">
    <div class="max-w-lg w-full bg-white rounded-2xl shadow-lg p-8">
      <h1 class="text-2xl font-bold text-gray-800 mb-6">Post Job</h1>

      <!-- Step indicators -->
      <div class="flex justify-between mb-6">
        <span :class="step === 1 ? activeStep : stepStyle">Step 1</span>
        <span :class="step === 2 ? activeStep : stepStyle">Step 2</span>
      </div>

      <!-- Step 1 -->
    <!-- Step 1 -->
<div class="space-y-4" v-if="step === 1">
  <div v-for="field in step1Fields" :key="field.key">
    <label :for="field.key" class="block text-sm font-medium text-gray-700">
      {{ field.label }}
    </label>

    <!-- Dropdowns for category & job_type -->
    <select
      v-if="field.key === 'category' || field.key === 'job_type'"
      :id="field.key"
      class="mt-1 p-2 w-full border rounded"
      v-model="formData[field.key]"
    >
      <option value="" disabled>Select {{ field.label }}</option>
      <option
        v-for="option in field.options"
        :key="option"
        :value="option"
      >
        {{ option }}
      </option>
    </select>

    <!-- Default input -->
    <input
      v-else
      :id="field.key"
      :type="field.type === 'int' ? 'number' : 'text'"
      class="mt-1 p-2 w-full border rounded"
      :value="formData[field.key]"
      @input="(e) => handleInput(field.key, (e.target as HTMLInputElement).value, field.type)"
    />

    <p v-if="errors[field.key]" class="text-red-500 text-sm">
      {{ errors[field.key] }}
    </p>
  </div>
</div>

      <!-- Step 2 -->
      <div class="space-y-4" v-if="step === 2">
        <div v-for="field in step2Fields" :key="field.key">
          <label :for="field.key" class="block text-sm font-medium text-gray-700">{{ field.label }}</label>
          <input
            :id="field.key"
            :type="field.type === 'int' ? 'number' : 'text'"
            class="mt-1 p-2 w-full border rounded"
            :value="formData[field.key]"
            @input="(e) => handleInput(field.key, (e.target as HTMLInputElement).value, field.type)"
          />
          <p v-if="errors[field.key]" class="text-red-500 text-sm">{{ errors[field.key] }}</p>
        </div>
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
          @click="submitForm"
          :disabled="mutationLoading"
          class="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 ml-auto disabled:opacity-50"
        >
          {{ mutationLoading ? 'Registering...' : 'Submit' }}
        </button>
      </div>

      <!-- Feedback -->
      <p v-if="mutationError" class="text-red-500 text-sm mt-2">{{ errorMessage }}</p>
      <p v-if="success" class="text-green-500 text-sm mt-2">Job posted successfully!</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { useAuthStore } from '@/stores/auth'
const client = useUrqlClient()

const REGISTER = gql`
  mutation AddJobOffer($input: NewJobOffer!) {
    AddJobOffer(input: $input) {
      title
      job_type
      category
      description
      location
      applicationCount
      offeredByuser {
        id
        email
        role
      }
    }
  }
`

const step1Fields = [
  { key: 'title', label: 'Job Title', type: 'string' },
  { key: 'job_type', label: 'Job Type', type: 'string', options: ['Full-time', 'Part-time', 'Contract', 'Internship', 'Remote'] },
  { key: 'category', label: 'Category', type: 'string', options: ['Construction', 'Repairing', 'HouseKeeping', 'Management'] },
  { key: 'description', label: 'Description', type: 'string' }
] as const


const step2Fields = [
  { key: 'location', label: 'Location', type: 'string' },
  { key: 'applicationCount', label: 'Application Count', type: 'int' }
] as const

const allFields = [...step1Fields, ...step2Fields]

type FieldKey = typeof allFields[number]['key']
type FieldType = typeof allFields[number]['type']

const formData = reactive<Record<FieldKey, string | number>>({
  title: '',
  job_type: '',
  category: '',
  description: '',
  location: '',
  applicationCount: 0
})

const errors = reactive<Record<FieldKey, string>>({
  title: '',
  job_type: '',
  category: '',
  description: '',
  location: '',
  applicationCount: ''
})

const handleInput = (key: FieldKey, value: string, type: FieldType) => {
  formData[key] = type === 'int' ? parseInt(value) || 0 : value
}

const validateStep = (
  fields: readonly { key: FieldKey; label: string; type: FieldType }[]
): boolean => {
  let isValid = true
  fields.forEach(({ key }) => {
    if (!formData[key]) {
      errors[key] = `${key.replace(/_/g, ' ')} is required`
      isValid = false
    } else {
      errors[key] = ''
    }
  })
  return isValid
}

const step = ref(1)
const mutationLoading = ref(false)
const mutationError = ref<CombinedError | null>(null)
const success = ref(false)
const router = useRouter()
const auth = useAuthStore()

const errorMessage = computed(() => mutationError.value?.message || '')

function nextStep() {
  if (validateStep(step1Fields)) {
    step.value++
  }
}

async function submitForm() {
  if (!validateStep(step2Fields)) return

  if (!auth.user?.user_id) {
    mutationError.value = new CombinedError({
      networkError: new Error('User not authenticated.')
    })
    return
  }

  mutationLoading.value = true
  mutationError.value = null
  success.value = false

  try {
    const result = await client
      .mutation(REGISTER, {
        input: {
          ...formData,
          offeredByuser: parseInt(auth.user.user_id)
        }
      })
      .toPromise()

    if (result.data?.AddJobOffer) {
      success.value = true
      Object.keys(formData).forEach((k) => {
        const key = k as FieldKey
        formData[key] = typeof formData[key] === 'number' ? 0 : ''
      })
      step.value = 1
      setTimeout(() => router.push('/admin'), 1000)
    } else if (result.error) {
      mutationError.value = result.error
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
