<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)]',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-10" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-6 py-8 lg:px-12">
        <div class="max-w-6xl mx-auto">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-12">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12 bg-white/20"></div>
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">System_Audit_Logs_v2.0</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                SIGNAL_FEED
              </h1>
            </div>

            <button 
              @click="themeStore.initializeUplink()"
              class="px-8 py-4 text-[10px] font-black uppercase tracking-widest text-black transition-all hover:brightness-110 flex items-center gap-3" 
              :style="{ backgroundColor: themeStore.accentColor }"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              REFRESH_UPLINK
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-4 gap-6 mb-8">
            <div v-for="stat in summaryStats" :key="stat.label" class="bg-white/[0.02] border border-white/5 p-6">
              <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest mb-2">{{ stat.label }}</p>
              <div class="flex items-end justify-between">
                <span class="text-3xl font-black italic tracking-tighter" :class="stat.colorClass">{{ stat.value }}</span>
                <span class="text-[10px] font-black uppercase tracking-widest opacity-50">UNIT_COUNT</span>
              </div>
            </div>
          </div>

          <div class="flex flex-wrap gap-4 mb-6">
            <select v-model="severityFilter" class="bg-black border border-white/10 px-4 py-2 text-[10px] font-bold uppercase tracking-widest focus:outline-none focus:border-white/30 text-gray-400">
              <option value="all">ALL_SEVERITY</option>
              <option value="critical text-red-500">CRITICAL_ONLY</option>
              <option value="warning">WARNINGS</option>
            </select>
            <select v-model="typeFilter" class="bg-black border border-white/10 px-4 py-2 text-[10px] font-bold uppercase tracking-widest focus:outline-none focus:border-white/30 text-gray-400">
              <option value="all">ALL_SOURCES</option>
              <option value="auth">AUTH_SYSTEM</option>
              <option value="system">CORE_CORE</option>
            </select>
          </div>

          <div class="bg-black border border-white/10 overflow-hidden shadow-2xl">
            <div class="overflow-x-auto">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="border-b border-white/10 bg-white/[0.02]">
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Timestamp</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Level</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Source</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Signal_Payload</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500 text-right">Encryption</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-white/5">
                  <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-white/[0.03] transition-colors group">
                    <td class="p-6 font-bold text-xs text-white/40 uppercase tracking-widest">{{ log.timestamp.split(' ')[1] }}</td>
                    <td class="p-6">
                      <span 
                        class="text-[10px] font-black uppercase border px-2 py-1"
                        :class="{
                          'border-red-500 text-red-500': log.severity === 'critical',
                          'border-amber-500 text-amber-500': log.severity === 'warning',
                          'border-cyan-500 text-cyan-500': log.severity === 'info'
                        }"
                      >
                        {{ log.severityLabel }}
                      </span>
                    </td>
                    <td class="p-6">
                      <span class="text-xs font-bold text-gray-400 uppercase tracking-widest">
                        {{ log.source }}
                      </span>
                    </td>
                    <td class="p-6">
                      <div class="text-sm font-medium text-white group-hover:text-cyan-400 transition-colors">
                        <span class="opacity-30 mr-2">></span>{{ log.message }}
                      </div>
                    </td>
                    <td class="p-6 text-right">
                      <span class="text-[9px] font-bold text-gray-600 uppercase tracking-[0.3em]">
                        {{ themeStore.encryptionLevel }}
                      </span>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            <div class="p-4 bg-white/[0.01] border-t border-white/5 flex justify-between items-center text-[9px] font-bold text-gray-600 tracking-widest uppercase">
              <span>Sync_Status: Verified</span>
              <span>Buffer_Size: 50_Entries</span>
            </div>
          </div>

        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()

const severityFilter = ref('all')
const typeFilter = ref('all')

const summaryStats = ref([
  { label: 'Critical_Events', value: '03', colorClass: 'text-red-500' },
  { label: 'Warning_Leaks', value: '12', colorClass: 'text-amber-500' },
  { label: 'Neural_Traffic', value: '48', colorClass: 'text-white' },
  { label: 'Total_Signals', value: '156', colorClass: 'text-white' }
])

const logs = ref([
  { id: 1, timestamp: '2025-12-31 14:32:18', severity: 'critical', severityLabel: 'CRIT', type: 'auth', message: 'Unauthorized access attempt: Node_7', source: 'SYS_CORE' },
  { id: 2, timestamp: '2025-12-31 14:28:45', severity: 'warning', severityLabel: 'WARN', type: 'job', message: 'Memory allocation at 92% capacity', source: 'JOB_ENG' },
  { id: 3, timestamp: '2025-12-31 14:25:12', severity: 'info', severityLabel: 'INFO', type: 'user', message: 'Neural link established for admin.01', source: 'AUTH_MOD' },
  { id: 4, timestamp: '2025-12-31 14:20:03', severity: 'info', severityLabel: 'INFO', type: 'system', message: 'Daily backup sequence finalized', source: 'BACKUP_S' },
])

const filteredLogs = computed(() => {
  let filtered = logs.value
  if (severityFilter.value !== 'all') filtered = filtered.filter(log => log.severity === severityFilter.value)
  if (typeFilter.value !== 'all') filtered = filtered.filter(log => log.type === typeFilter.value)
  return filtered
})
</script>

<style scoped>
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: rgba(255, 255, 255, 0.1);
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: var(--accent-glow, #00f3ff);
}
</style>