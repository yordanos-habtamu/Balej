<template>
  <div class="min-h-screen flex">
    <!-- Sidebar -->
    <!-- <aside class="flex-shrink-0 pt-8">
      <admin-sidebar />
    </aside> -->

    <!-- Main Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="max-w-4xl mx-auto px-6 py-10 lg:px-12 lg:py-16">
        <!-- Header -->
        <div class="text-center mb-12">
          <h1 class="text-4xl lg:text-5xl font-bold mb-4">Post a New Job</h1>
          <p class="text-xl text-gray-400">Create a job posting to attract top talent</p>
        </div>

        <!-- Progress Steps -->
        <div class="flex items-center justify-center mb-12">
          <div class="flex items-center">
            <div :class="[step >= 1 ? 'bg-[color:var(--primary-color)]' : 'bg-gray-700', 'w-12 h-12 rounded-full flex items-center justify-center text-lg font-bold transition']">
              1
            </div>
            <div class="w-32 h-1 mx-4" :class="step >= 2 ? 'bg-[color:var(--primary-color)]' : 'bg-gray-700'"></div>
            <div :class="[step >= 2 ? 'bg-[color:var(--primary-color)]' : 'bg-gray-700', 'w-12 h-12 rounded-full flex items-center justify-center text-lg font-bold transition']">
              2
            </div>
          </div>
        </div>

        <!-- Form Card -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-10 lg:p-14">
          <form @submit.prevent="step === 2 ? submitForm() : nextStep()">
            <!-- Step 1: Job Basics -->
            <div v-if="step === 1" class="space-y-8">
              <div>
                <label class="block text-lg font-medium text-gray-300 mb-3">Job Title</label>
                <input
                  v-model="formData.title"
                  type="text"
                  placeholder="e.g. Senior Full-Stack Developer"
                  class="w-full px-6 py-4 bg-gray-700 border border-gray-600 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] text-white placeholder-gray-500 transition"
                  required
                />
              </div>

              <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                <div>
                  <label class="block text-lg font-medium text-gray-300 mb-3">Job Type</label>
                  <select
                    v-model="formData.job_type"
                    class="w-full px-6 py-4 bg-gray-700 border border-gray-600 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] text-white"
                    required
                  >
                    <option value="" disabled>Select job type</option>
                    <option value="Full-time">Full-time</option>
                    <option value="Part-time">Part-time</option>
                    <option value="Contract">Contract</option>
                    <option value="Internship">Internship</option>
                    <option value="Remote">Remote</option>
                  </select>
                </div>

                <div>
                  <label class="block text-lg font-medium text-gray-300 mb-3">Category</label>
                  <select
                    v-model="formData.category"
                    class="w-full px-6 py-4 bg-gray-700 border border-gray-600 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] text-white"
                    required
                  >
                    <option value="" disabled>Select category</option>
                    <option value="Construction">Construction</option>
                    <option value="Repairing">Repairing</option>
                    <option value="HouseKeeping">House Keeping</option>
                    <option value="Management">Management</option>
                  </select>
                </div>
              </div>

              <div>
                <label class="block text-lg font-medium text-gray-300 mb-3">Job Description</label>
                <textarea
                  v-model="formData.description"
                  rows="8"
                  placeholder="Describe the role, responsibilities, requirements, and benefits..."
                  class="w-full px-6 py-4 bg-gray-700 border border-gray-600 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] text-white placeholder-gray-500 resize-none transition"
                  required
                ></textarea>
              </div>
            </div>

            <!-- Step 2: Additional Details -->
            <div v-if="step === 2" class="space-y-8">
              <div>
                <label class="block text-lg font-medium text-gray-300 mb-3">Location</label>
                <input
                  v-model="formData.location"
                  type="text"
                  placeholder="e.g. Remote, New York, NY, or On-site"
                  class="w-full px-6 py-4 bg-gray-700 border border-gray-600 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] text-white placeholder-gray-500 transition"
                  required
                />
              </div>

              <div>
                <label class="block text-lg font-medium text-gray-300 mb-3">Maximum Applications (optional)</label>
                <input
                  v-model.number="formData.applicationCount"
                  type="number"
                  min="0"
                  placeholder="Leave blank for unlimited"
                  class="w-full px-6 py-4 bg-gray-700 border border-gray-600 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] text-white placeholder-gray-500 transition"
                />
                <p class="text-sm text-gray-500 mt-3">Set a limit to close applications automatically when reached</p>
              </div>
            </div>

            <!-- Navigation Buttons -->
            <div class="flex justify-between pt-10">
              <button
                v-if="step > 1"
                type="button"
                @click="step--"
                class="px-8 py-4 bg-gray-700 hover:bg-gray-600 text-white rounded-2xl font-semibold transition"
              >
                Back
              </button>

              <button
                type="submit"
                :disabled="mutationLoading"
                class="px-10 py-4 bg-[color:var(--primary-color)] hover:bg-[color:var(--primary-color)/0.9] text-white rounded-2xl font-semibold text-lg transition shadow-lg disabled:opacity-70 disabled:cursor-not-allowed ml-auto flex items-center gap-3"
              >
                <span v-if="mutationLoading">
                  <svg class="w-6 h-6 animate-spin" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <circle class="opacity-25" cx="12" cy="12" r="10" stroke-width="4"></circle>
                    <path class="opacity-75" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                  </svg>
                  Posting...
                </span>
                <span v-else>
                  {{ step === 2 ? 'Post Job' : 'Next Step' }}
                </span>
              </button>
            </div>
          </form>

          <!-- Success Message -->
          <div v-if="success" class="mt-10 p-8 bg-green-600/20 border border-green-600 rounded-3xl text-center">
            <svg class="w-16 h-16 text-green-400 mx-auto mb-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
            </svg>
            <p class="text-2xl font-bold text-green-400">Job Posted Successfully!</p>
            <p class="text-gray-300 mt-3">Redirecting to dashboard...</p>
          </div>

          <!-- Error Message -->
          <div v-if="mutationError" class="mt-8 p-6 bg-red-600/20 border border-red-600 rounded-3xl text-center">
            <p class="text-xl text-red-400 font-medium">Error posting job</p>
            <p class="text-gray-300 mt-2">{{ errorMessage }}</p>
          </div>
        </div>
      </div>
    </main>
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
