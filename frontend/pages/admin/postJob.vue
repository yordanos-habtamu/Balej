<template>
  <div class="min-h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <main
      :class="[
        'flex-1 relative flex flex-col min-h-screen transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)]',
        themeStore.terminalCollapsed ? 'ml-20' : 'ml-72'
      ]"
    >
      <div
        class="absolute inset-0 pointer-events-none opacity-10"
        :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }"
      ></div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-6 py-10 lg:px-12">
        <div class="max-w-4xl mx-auto">
          
          <div class="mb-12 border-l-4 pl-6 py-2" :style="{ borderColor: themeStore.accentColor }">
            <h1 class="text-4xl lg:text-5xl font-black italic uppercase tracking-tighter">
              POST_<span :style="{ color: themeStore.accentColor }">MISSION_OFFER</span>
            </h1>
            <p class="text-[10px] text-gray-500 tracking-[0.4em] mt-2 uppercase">
              Uplink Status: Ready // Node: {{ auth.user?.user_id || 'NULL' }}
            </p>
          </div>

          <div class="flex items-center gap-4 mb-10">
            <div v-for="n in 2" :key="n" class="flex items-center gap-4">
              <div
                class="text-[10px] font-bold transition-colors duration-500"
                :style="{ color: step >= n ? themeStore.accentColor : '#374151' }"
              >
                STEP_0{{ n }}
              </div>
              <div v-if="n < 2" class="w-20 h-[1px] bg-gray-800 relative">
                <div
                  class="absolute inset-0 transition-all duration-700"
                  :style="{ width: step > n ? '100%' : '0%', backgroundColor: themeStore.accentColor }"
                ></div>
              </div>
            </div>
          </div>

          <div class="bg-black/40 backdrop-blur-md border border-white/10 relative overflow-hidden group">
            <div
              class="absolute top-0 right-0 w-12 h-12 border-r border-t border-white/20 transition-colors group-hover:border-cyan-500"
              :style="{ borderColor: themeStore.accentColor }"
            ></div>

            <form @submit.prevent="step === 2 ? submitForm() : nextStep()" class="p-8 lg:p-12">
              
              <div v-show="step === 1" class="space-y-8 animate-in fade-in slide-in-from-right-8 duration-500">
                <div class="space-y-3">
                  <label class="neural-label">>> DESIGNATION_TITLE</label>
                  <input
                    v-model="formData.title"
                    type="text"
                    placeholder="E.G. SYSTEMS_ARCHITECT"
                    class="neural-input"
                    :class="{ 'border-red-500/50': errors.title }"
                  />
                  <p v-if="errors.title" class="text-[7px] text-red-500 tracking-widest uppercase">{{ errors.title }}</p>
                </div>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
                  <NeuroDrop
                    v-model="formData.job_type"
                    :options="step1Fields[1].options"
                    placeholder="SELECT_MODALITY"
                    label="CONTRACT_MODALITY"
                  />
                  <NeuroDrop
                    v-model="formData.category"
                    :options="step1Fields[2].options"
                    placeholder="SELECT_SECTOR"
                    label="SECTOR_CATALOGUE"
                  />
                </div>

                <div class="space-y-3">
                  <label class="neural-label">>> DIRECTIVE_REQUISITES</label>
                  <textarea
                    v-model="formData.description"
                    rows="5"
                    placeholder="SPECIFY_PARAMETERS..."
                    class="neural-input resize-none"
                    :class="{ 'border-red-500/50': errors.description }"
                  ></textarea>
                </div>
              </div>

              <div v-show="step === 2" class="space-y-8 animate-in fade-in slide-in-from-right-8 duration-500">
                <div class="space-y-3">
                  <label class="neural-label">>> GEOSPATIAL_COORDS</label>
                  <input
                    v-model="formData.location"
                    type="text"
                    placeholder="E.G. NEURAL_HUB_REMOTE"
                    class="neural-input"
                    :class="{ 'border-red-500/50': errors.location }"
                  />
                </div>
                <div class="space-y-3">
                  <label class="neural-label">>> MAX_RECEPTION_THRESHOLD</label>
                  <input v-model.number="formData.applicationCount" type="number" class="neural-input" />
                  <p class="text-[8px] text-gray-600 uppercase">Limit concurrent uplink attempts.</p>
                </div>
              </div>

              <div class="flex justify-between items-center pt-10 mt-10 border-t border-white/5">
                <button v-if="step > 1" type="button" @click="step--" class="neural-btn-secondary">
                  [ REVERT_STEP ]
                </button>
                <div v-else></div>
                <button
                  type="submit"
                  :disabled="mutationLoading"
                  class="neural-btn-primary"
                  :style="{ backgroundColor: themeStore.accentColor }"
                >
                  <span v-if="mutationLoading" class="animate-spin mr-2">⚬</span>
                  {{ step === 2 ? 'DISPATCH_OFFER' : 'INITIALIZE_NEXT_PHASE' }}
                </button>
              </div>
            </form>

            <div
              v-if="success"
              class="absolute inset-0 bg-[#020205]/95 backdrop-blur-xl z-50 flex flex-col items-center justify-center p-10 text-center"
            >
              <div
                class="w-16 h-16 border-2 mb-6 flex items-center justify-center animate-bounce"
                :style="{ borderColor: themeStore.accentColor, color: themeStore.accentColor }"
              >
                <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path d="M5 13l4 4L19 7" stroke-width="3" />
                </svg>
              </div>
              <h3 class="text-2xl font-black italic uppercase tracking-tighter" :style="{ color: themeStore.accentColor }">Uplink_Successful</h3>
              <p class="text-[10px] text-gray-500 uppercase tracking-[0.4em] mt-2">Redirecting to terminal...</p>
            </div>
          </div>

          <div v-if="mutationError" class="mt-6 bg-red-500/10 border border-red-500/50 p-4 font-mono text-[9px] text-red-500 uppercase tracking-widest">
            ERROR_LOG: {{ errorMessage }}
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
import { useThemeStore } from '@/stores/theme'
import { useUrqlClient } from '@/composables/useUrqlClient'
import NeuroDrop from '@/components/NeuroDrop.vue'

const themeStore = useThemeStore()
const client = useUrqlClient()
const router = useRouter()
const auth = useAuthStore()

const REGISTER = gql`
mutation AddJobOffer($input: NewJobOffer!) {
  AddJobOffer(input: $input) {
    title job_type category description location applicationCount
  }
}
`

const step1Fields = [
  { key: 'title', label: 'TITLE', type: 'string' },
  { key: 'job_type', label: 'JOB_TYPE', type: 'string', options: ['Full-time', 'Part-time', 'Contract', 'Internship', 'Remote'] },
  { key: 'category', label: 'CATEGORY', type: 'string', options: ['Construction', 'Repairing', 'HouseKeeping', 'Management'] },
  { key: 'description', label: 'DESCRIPTION', type: 'string' }
] as const

const step2Fields = [
  { key: 'location', label: 'LOCATION', type: 'string' },
  { key: 'applicationCount', label: 'APP_COUNT', type: 'int' }
] as const

type FieldKey = 'title' | 'job_type' | 'category' | 'description' | 'location' | 'applicationCount'

const formData = reactive<Record<FieldKey, string | number>>({
  title: '', job_type: '', category: '', description: '', location: '', applicationCount: 0
})

const errors = reactive<Record<FieldKey, string>>({
  title: '', job_type: '', category: '', description: '', location: '', applicationCount: ''
})

const validateStep = (fields: readonly any[]): boolean => {
  let valid = true
  fields.forEach((field) => {
    const key = field.key as FieldKey
    if (!formData[key] && key !== 'applicationCount') {
      errors[key] = `${field.label}_IS_REQUIRED`
      valid = false
    } else {
      errors[key] = ''
    }
  })
  return valid
}

const step = ref(1)
const mutationLoading = ref(false)
const mutationError = ref<CombinedError | null>(null)
const success = ref(false)
const errorMessage = computed(() => mutationError.value?.message || '')

function nextStep() {
  if (validateStep(step1Fields)) {
    step.value = 2
  }
}

async function submitForm() {
  if (!validateStep(step2Fields)) return
  if (!auth.user?.user_id) {
    mutationError.value = new CombinedError({ networkError: new Error('UNAUTHORIZED_ACCESS') })
    return
  }

  mutationLoading.value = true
  try {
    const result = await client.mutation(REGISTER, {
      input: { ...formData, offeredByuser: parseInt(auth.user.user_id) }
    }).toPromise()

    if (result.data) {
      success.value = true
      setTimeout(() => router.push('/admin/dashboard'), 1500)
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

<style scoped>
@reference 'tailwindcss';

.neural-label { @apply text-[9px] font-mono text-gray-500 uppercase tracking-[0.4em] block italic; }
.neural-input { @apply w-full bg-[rgba(255,255,255,0.03)] border border-white/10 p-4 text-xs font-mono text-white transition-all uppercase outline-none; }
.neural-input:focus { 
  @apply bg-[rgba(255,255,255,0.07)]; 
  border-color: v-bind('themeStore.accentColor');
  box-shadow: 0 0 15px v-bind('themeStore.accentColor + "22"');
}
.neural-btn-primary { @apply px-10 py-4 font-black uppercase text-[10px] tracking-[0.3em] text-black transition-all hover:scale-105 active:scale-95 disabled:opacity-50; }
.neural-btn-secondary { @apply px-8 py-3 border border-white/10 text-[10px] uppercase font-bold tracking-[0.3em] hover:bg-white hover:text-black transition-all; }
.custom-scrollbar::-webkit-scrollbar { width: 3px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: v-bind('themeStore.accentColor'); }
</style>