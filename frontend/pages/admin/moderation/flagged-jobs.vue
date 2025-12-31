<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden">
    <admin-sidebar />

    <main :class="['flex-1 flex flex-col h-full transition-all duration-500 relative', themeStore.terminalCollapsed ? 'pl-20' : 'pl-72']">
      <div class="z-10 px-6 py-8 lg:px-12">
        <div class="max-w-7xl mx-auto flex justify-between items-end">
          <div>
            <span class="text-[10px] font-black tracking-widest uppercase opacity-60">Neural_Defense // Marketplace</span>
            <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">FLAGGED_LISTINGS</h1>
          </div>
          <button @click="showRiskModal = true" class="px-6 py-4 bg-white/[0.03] border border-white/10 hover:border-white/40 transition-all text-[9px] font-black uppercase">
            + Define_New_Risk
          </button>
        </div>
      </div>

      <div class="flex-1 px-6 lg:px-12 pb-8 overflow-hidden">
        <div class="h-full bg-black border border-white/10 flex flex-col">
          <div class="overflow-y-auto custom-scrollbar">
            <table class="w-full text-left">
              <thead class="bg-white/[0.02] sticky top-0 border-b border-white/10">
                <tr>
                  <th class="p-6 text-[10px] font-black uppercase text-gray-500">Posting_Entity</th>
                  <th class="p-6 text-[10px] font-black uppercase text-gray-500">Risk_Category</th>
                  <th class="p-6 text-[10px] font-black uppercase text-gray-500">Suppressed_Reach</th>
                  <th class="p-6 text-[10px] font-black uppercase text-gray-500 text-right">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="job in flaggedJobs" :key="job.id" class="hover:bg-white/[0.02]">
                  <td class="p-6">
                    <p class="text-xs font-black uppercase">{{ job.title }}</p>
                    <p class="text-[9px] text-gray-600 uppercase">{{ job.company }}</p>
                  </td>
                  <td class="p-6">
                    <span class="text-[8px] font-black px-2 py-1 bg-red-500/10 text-red-500 border border-red-500/20 uppercase">
                      {{ job.category }}
                    </span>
                  </td>
                  <td class="p-6">
                    <p class="text-xs font-black text-white italic">{{ job.reachBlocked }}</p>
                    <p class="text-[8px] font-bold text-gray-600 uppercase">Impressions_Shielded</p>
                  </td>
                  <td class="p-6 text-right">
                    <div class="flex justify-end gap-3">
                      <button @click="openInspect(job)" class="px-4 py-2 border border-white/10 text-[9px] font-black uppercase hover:bg-white/5 transition-all">INSPECT</button>
                      <button @click="handleAction(job, 'PURGE')" class="px-4 py-2 bg-red-600/20 text-red-500 text-[9px] font-black uppercase">PURGE</button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <div v-if="isInspectOpen" class="fixed inset-0 z-[1000] flex justify-end">
        <div @click="isInspectOpen = false" class="absolute inset-0 bg-black/80 backdrop-blur-md"></div>
        <div class="relative w-full max-w-2xl bg-[#050508] border-l border-white/10 p-10 flex flex-col h-full shadow-2xl">
          <div class="mb-10 flex justify-between items-start">
            <div>
              <p class="text-[10px] font-black text-gray-500 uppercase tracking-widest">Diagnostic_Audit</p>
              <h2 class="text-3xl font-black italic uppercase tracking-tighter">{{ selectedJob?.title }}</h2>
            </div>
            <button @click="isInspectOpen = false" class="text-white hover:text-red-500">✕</button>
          </div>

          <div class="flex-1 overflow-y-auto space-y-8 custom-scrollbar pr-4">
             <div class="grid grid-cols-2 gap-4">
               <div class="p-4 bg-white/[0.02] border border-white/5">
                 <p class="text-[8px] font-black text-gray-600 uppercase mb-2">Suppression_Rate</p>
                 <p class="text-xl font-black italic text-amber-500">92%_HEURISTIC_MATCH</p>
               </div>
               <div class="p-4 bg-white/[0.02] border border-white/5">
                 <p class="text-[8px] font-black text-gray-600 uppercase mb-2">Candidate_Protection</p>
                 <p class="text-xl font-black italic text-green-500">{{ selectedJob?.reachBlocked }} Saved</p>
               </div>
             </div>

             <div class="p-6 bg-white/[0.02] border border-white/10">
               <p class="text-[9px] font-black text-gray-500 uppercase mb-4">Original_Description_Intercept</p>
               <p class="text-sm font-serif italic text-gray-400 leading-relaxed">{{ selectedJob?.description }}</p>
             </div>
          </div>

          <div class="pt-8 border-t border-white/10 flex gap-4">
            <button @click="handleAction(selectedJob, 'VALIDATE')" class="flex-1 py-5 text-black font-black uppercase tracking-widest text-xs" :style="{ backgroundColor: themeStore.accentColor }">VALIDATE_POSTING</button>
            <button @click="handleAction(selectedJob, 'PURGE')" class="flex-1 py-5 bg-red-600 text-white font-black uppercase tracking-widest text-xs">PURGE_FROM_MARKET</button>
          </div>
        </div>
      </div>

      <div v-if="showRiskModal" class="fixed inset-0 z-[1001] flex items-center justify-center p-6">
        <div @click="showRiskModal = false" class="absolute inset-0 bg-black/90 backdrop-blur-xl"></div>
        <div class="relative w-full max-w-lg bg-[#0a0a0c] border border-white/10 p-10 shadow-2xl">
          <h3 class="text-2xl font-black italic uppercase tracking-tighter mb-8">Define_Risk_Vector</h3>
          <div class="space-y-6">
            <input v-model="newRiskName" placeholder="CATEGORY_NAME..." class="w-full bg-white/[0.03] border border-white/10 p-4 text-xs font-black uppercase focus:border-white outline-none" />
            <button @click="addRisk" class="w-full py-4 text-black font-black uppercase" :style="{ backgroundColor: themeStore.accentColor }">DEPLOY_PROTOCOL</button>
          </div>
        </div>
      </div>

      <transition name="toast">
        <div v-if="toast.show" class="fixed bottom-8 right-8 z-[2000] bg-black border-l-4 p-5 min-w-[340px] shadow-2xl" :style="{ borderColor: toast.type === 'danger' ? '#ef4444' : themeStore.accentColor }">
          <p class="text-[10px] font-black uppercase tracking-widest opacity-60">{{ toast.label }}</p>
          <p class="text-xs font-black mt-1 uppercase">{{ toast.message }}</p>
        </div>
      </transition>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const isInspectOpen = ref(false)
const showRiskModal = ref(false)
const selectedJob = ref(null)
const newRiskName = ref('')
const toast = ref({ show: false, message: '', label: '', type: 'success' })

const flaggedJobs = ref([
  { id: 1, title: 'Senior Wealth Manager', company: 'Nexus_Fintech', category: 'SALARY_ANOMALY', reachBlocked: 1402, description: 'Earn $5000 weekly from home. No experience needed. Join our Telegram group.' },
  { id: 2, title: 'Remote Data Entry', company: 'Unknown_04', category: 'OFF_PLATFORM_REDIRECT', reachBlocked: 569, description: 'Apply via http://unverified-link.com to start immediately.' }
])

const openInspect = (job) => {
  selectedJob.value = job
  isInspectOpen.value = true
}

const addRisk = () => {
  if (!newRiskName.value) return
  showRiskModal.value = false
  toast.value = { show: true, label: 'Neural_Update', message: `PROTOCOL_${newRiskName.value.toUpperCase()}_DEPLOYED`, type: 'success' }
  newRiskName.value = ''
  setTimeout(() => { toast.value.show = false }, 3000)
}

const handleAction = (job, type) => {
  isInspectOpen.value = false
  toast.value = { show: true, label: 'Security_Action', message: type === 'PURGE' ? `JOB_${job.id}_TERMINATED` : `JOB_${job.id}_VALIDATED`, type: type === 'PURGE' ? 'danger' : 'success' }
  flaggedJobs.value = flaggedJobs.value.filter(j => j.id !== job.id)
  setTimeout(() => { toast.value.show = false }, 3000)
}
</script>