<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)] relative',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-10" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-6 py-8 lg:px-12">
        <div class="max-w-6xl mx-auto">
          
          <div class="mb-12 flex justify-between items-end">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12 bg-white/20"></div>
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">Data_Extraction_Suite_v2</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                EXPORT_ARCHIVE
              </h1>
            </div>
            <div class="hidden md:block text-right">
              <p class="text-[10px] font-black text-gray-600 uppercase tracking-widest">System_Time</p>
              <p class="text-xl font-black italic text-white">{{ currentTime }}</p>
            </div>
          </div>

          <div class="grid grid-cols-1 lg:grid-cols-3 gap-8 mb-12">
            
            <div class="lg:col-span-2 space-y-6">
              <div class="bg-white/[0.02] border border-white/10 p-8 relative overflow-hidden">
                <div class="absolute top-0 left-0 w-full h-[1px] bg-gradient-to-r from-transparent via-cyan-500 to-transparent opacity-30"></div>

                <h3 class="text-xl font-black italic uppercase tracking-tighter mb-8 flex items-center gap-3">
                  <span class="w-2 h-2 rounded-full animate-pulse" :style="{ backgroundColor: themeStore.accentColor }"></span>
                  EXTRACT_CONFIG_BUFFER
                </h3>

                <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
                  <div class="space-y-6">
                    <div class="space-y-3">
                      <label class="text-[9px] font-black text-gray-500 uppercase tracking-[0.2em]">Target_Dataset</label>
                      <select v-model="selectedModule" class="hud-select">
                        <option v-for="m in dataModules" :key="m.id" :value="m.id">{{ m.name }}</option>
                      </select>
                    </div>

                    <div class="space-y-3">
                      <label class="text-[9px] font-black text-gray-500 uppercase tracking-[0.2em]">Output_Format</label>
                      <div class="flex gap-2">
                        <button v-for="fmt in ['CSV', 'JSON', 'XLSX']" :key="fmt"
                          @click="selectedFormat = fmt"
                          :class="[
                            'flex-1 py-3 text-[10px] font-black border transition-all',
                            selectedFormat === fmt ? 'bg-white text-black border-white' : 'border-white/10 text-gray-500 hover:border-white/30'
                          ]"
                        >
                          {{ fmt }}
                        </button>
                      </div>
                    </div>
                  </div>

                  <div class="space-y-6 border-l border-white/5 pl-0 md:pl-10">
                    <div class="space-y-3">
                      <label class="text-[9px] font-black text-cyan-500 uppercase tracking-[0.2em]">Recursive_Protocol</label>
                      <div class="flex flex-col gap-2">
                        <button 
                          @click="isScheduled = !isScheduled"
                          :class="[
                            'w-full p-4 border text-left transition-all flex justify-between items-center',
                            isScheduled ? 'border-cyan-500/50 bg-cyan-500/5' : 'border-white/10 bg-transparent opacity-40'
                          ]"
                        >
                          <span class="text-[10px] font-black uppercase">Enable_Auto_Schedule</span>
                          <div class="w-8 h-4 bg-gray-800 relative">
                            <div class="absolute top-1 left-1 w-2 h-2 transition-all duration-300" 
                                 :style="{ transform: isScheduled ? 'translateX(16px)' : 'translateX(0)', backgroundColor: isScheduled ? '#00f3ff' : '#444' }"></div>
                          </div>
                        </button>
                      </div>
                    </div>

                    <transition name="fade">
                      <div v-if="isScheduled" class="space-y-4 pt-2">
                        <div class="space-y-2">
                          <label class="text-[8px] font-black text-gray-600 uppercase">Frequency_Pattern</label>
                          <select v-model="scheduleFreq" class="hud-select-sm">
                            <option>EVERY_MONDAY_00:00</option>
                            <option>DAILY_RECAP_23:59</option>
                            <option>MONTHLY_FINANCIAL_DUMP</option>
                          </select>
                        </div>
                        <div class="space-y-2">
                          <label class="text-[8px] font-black text-gray-600 uppercase">Destination_Uplink</label>
                          <input type="text" placeholder="admin-email@portal.com" class="hud-input-sm">
                        </div>
                      </div>
                    </transition>
                  </div>
                </div>

                <div class="mt-12 flex gap-4">
                  <button 
                    @click="initiateExport"
                    class="flex-1 py-5 text-[11px] font-black uppercase tracking-[0.4em] text-black transition-all hover:brightness-110 active:scale-[0.98]"
                    :style="{ backgroundColor: themeStore.accentColor }"
                  >
                    {{ isScheduled ? 'SAVE_SCHEDULED_TASK' : 'RUN_INSTANT_EXTRACT' }}
                  </button>
                </div>
              </div>
            </div>

            <div class="space-y-6">
              <div class="bg-white/[0.02] border border-white/5 p-6">
                <p class="text-[9px] font-black text-gray-500 uppercase tracking-widest mb-6">Archive_Stability</p>
                <div class="space-y-4">
                  <div v-for="i in 3" :key="i" class="flex items-center justify-between">
                    <span class="text-[10px] font-bold text-gray-400 uppercase">NODE_0{{i}}_SYMMETRY</span>
                    <span class="text-[10px] font-black text-green-500 uppercase">99.9%</span>
                  </div>
                </div>
              </div>
              
              <div class="bg-[#050508] border border-white/10 p-6 relative">
                 <div class="flex items-center gap-2 mb-4 text-amber-500">
                   <svg class="w-4 h-4" fill="currentColor" viewBox="0 0 24 24"><path d="M12 2L1 21h22L12 2zm0 3.99L19.53 19H4.47L12 5.99zM11 16h2v2h-2zm0-6h2v4h-2z"/></svg>
                   <span class="text-[9px] font-black uppercase tracking-widest">Security_Protocol</span>
                 </div>
                 <p class="text-[10px] font-bold text-gray-500 uppercase leading-relaxed italic">
                   All scheduled tasks are encrypted with AES-256. Manual override required for deletions.
                 </p>
              </div>
            </div>
          </div>

          <div class="bg-black border border-white/10 overflow-hidden shadow-2xl">
            <div class="p-6 border-b border-white/10 bg-white/[0.02] flex justify-between items-center">
              <h4 class="text-xs font-black uppercase tracking-[0.3em]">Activity_Log</h4>
              <span class="text-[9px] font-black text-gray-600 uppercase">Filter: All_Events</span>
            </div>
            
            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="border-b border-white/10 bg-white/[0.01]">
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">ID_Hash</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Task_Name</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Type</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Status</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-right text-gray-500">Actions</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-white/5">
                  <tr v-for="job in exportHistory" :key="job.id" class="hover:bg-white/[0.03] transition-colors group">
                    <td class="p-6 font-mono text-[10px] text-gray-600 uppercase">#{{ job.id.toString().slice(-6) }}</td>
                    <td class="p-6">
                      <div class="flex flex-col">
                        <span class="text-[11px] font-black uppercase tracking-tight text-white">{{ job.name }}</span>
                        <span class="text-[9px] font-bold text-gray-600 uppercase">{{ job.time }}</span>
                      </div>
                    </td>
                    <td class="p-6">
                      <span class="text-[9px] font-black uppercase px-2 py-0.5 border" 
                        :class="job.isScheduled ? 'border-cyan-500/50 text-cyan-500' : 'border-white/10 text-gray-500'">
                        {{ job.isScheduled ? 'RECURSIVE' : 'INSTANT' }}
                      </span>
                    </td>
                    <td class="p-6">
                      <div class="flex items-center gap-2">
                        <div class="w-1.5 h-1.5 rounded-full" :class="job.status === 'Complete' ? 'bg-green-500' : 'bg-amber-500 animate-pulse'"></div>
                        <span class="text-[10px] font-black uppercase" :class="job.status === 'Complete' ? 'text-white' : 'text-amber-500'">{{ job.status }}</span>
                      </div>
                    </td>
                    <td class="p-6 text-right">
                      <button v-if="job.status === 'Complete'" class="text-[10px] font-black uppercase text-white border border-white/10 px-4 py-2 hover:bg-white/5 transition-all">
                        ACCESS_DATA
                      </button>
                      <button v-else-if="job.isScheduled" class="text-[10px] font-black uppercase text-red-500 border border-red-500/20 px-4 py-2 hover:bg-red-500/10 transition-all">
                        TERMINATE
                      </button>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
        </div>
      </div>
    </main>

    <transition name="toast">
      <div v-if="showToast" class="fixed bottom-8 right-8 z-[200] bg-black border-l-4 p-4 min-w-[320px] shadow-2xl" :style="{ borderColor: themeStore.accentColor }">
        <p class="text-[10px] font-black uppercase tracking-widest text-gray-500">System_Notification</p>
        <p class="text-xs font-bold mt-1 tracking-tight">{{ toastMessage }}</p>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const selectedModule = ref(1)
const selectedFormat = ref('CSV')
const isScheduled = ref(false)
const scheduleFreq = ref('EVERY_MONDAY_00:00')
const showToast = ref(false)
const toastMessage = ref('')
const currentTime = ref('00:00:00')

const dataModules = [
  { id: 1, name: 'CANDIDATE_DATABASE' },
  { id: 2, name: 'RECRUITER_LEDGER' },
  { id: 3, name: 'FINANCIAL_REPORTS' },
  { id: 4, name: 'APPLICATION_TRAFFIC' }
]

const exportHistory = ref([
  { id: 882191, name: 'CANDIDATE_LIST_Q4', isScheduled: false, status: 'Complete', time: '2025-12-30 14:02' },
  { id: 882192, name: 'WEEKLY_AUTO_SYNC', isScheduled: true, status: 'Scheduled', time: 'RECURRING' },
  { id: 882193, name: 'REVENUE_AUDIT_JAN', isScheduled: false, status: 'Processing', time: '2025-12-31 21:10' }
])

const initiateExport = () => {
  themeStore.initializeUplink()
  toastMessage.value = isScheduled.value ? 'AUTOMATED_TASK_SAVED_TO_CRON' : 'MANUAL_EXTRACTION_INITIALIZED'
  showToast.value = true
  
  const newJob = {
    id: Date.now(),
    name: isScheduled.value ? 'AUTO_' + dataModules.find(m => m.id === selectedModule.value)?.name : dataModules.find(m => m.id === selectedModule.value)?.name,
    isScheduled: isScheduled.value,
    status: isScheduled.value ? 'Scheduled' : 'Processing',
    time: isScheduled.value ? 'RECURRING' : new Date().toLocaleString()
  }
  
  exportHistory.value.unshift(newJob)
  setTimeout(() => { showToast.value = false }, 4000)
}

onMounted(() => {
  setInterval(() => {
    currentTime.value = new Date().toLocaleTimeString('en-GB')
  }, 1000)
})
</script>

<style scoped>
  @reference 'tailwindcss';
.hud-select {
  @apply w-full bg-white/[0.03] border border-white/10 p-4 text-[11px] font-black uppercase tracking-widest focus:outline-none focus:border-white/40 appearance-none;
}
.hud-select-sm {
  @apply w-full bg-white/[0.05] border border-white/10 px-3 py-2 text-[10px] font-black uppercase focus:outline-none appearance-none;
}
.hud-input-sm {
  @apply w-full bg-white/[0.05] border border-white/10 px-3 py-2 text-[10px] font-black focus:outline-none placeholder:text-gray-700;
}

.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease, transform 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; transform: translateY(-5px); }

.toast-enter-active, .toast-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-enter-from, .toast-leave-to { transform: translateX(100%) translateY(20px); opacity: 0; }
</style>