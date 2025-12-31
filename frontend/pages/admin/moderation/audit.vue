<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <div :class="['flex-1 flex flex-col h-full relative transition-all duration-500', themeStore.terminalCollapsed ? 'ml-20' : 'ml-72']">
      
      <header class="h-32 flex flex-col justify-center px-12 shrink-0 border-b border-white/5 bg-black/60 relative overflow-hidden z-20">
        <div class="absolute inset-0 bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.1)_50%)] pointer-events-none"></div>
        <span class="text-[10px] font-black tracking-[0.4em] uppercase opacity-40 text-cyan-400">Security_Protocol // Audit_Chain</span>
        <h1 class="text-5xl font-black italic uppercase tracking-tighter">SYSTEM_AUDIT</h1>
      </header>

      <div class="px-12 py-4 bg-white/[0.02] border-b border-white/5 flex gap-6 items-center shrink-0 z-20">
        <div class="flex flex-col flex-1">
          <label class="text-[8px] font-black text-gray-500 uppercase mb-1 tracking-widest">Database_Query</label>
          <input type="text" placeholder="SEARCH_BY_IP_ACTOR_OR_TRACE..." class="bg-black border border-white/10 text-[10px] p-2 uppercase outline-none focus:border-cyan-500 transition-all w-full" />
        </div>
        <button @click="triggerExport" :disabled="isExporting" class="mt-4 px-6 py-2 bg-white/5 border border-white/10 text-[10px] font-black uppercase hover:bg-white/10 min-w-[140px] relative">
          <span :class="{ 'opacity-0': isExporting }">Export_CSV</span>
          <div v-if="isExporting" class="absolute inset-0 flex items-center justify-center">
            <div class="w-3 h-3 border-2 border-cyan-500 border-t-transparent rounded-full animate-spin"></div>
          </div>
        </button>
      </div>

      <main class="flex-1 overflow-hidden p-8 lg:p-12 bg-[#020205] z-10">
        <div class="h-full border border-white/10 rounded-sm flex flex-col bg-white/[0.01] backdrop-blur-sm overflow-hidden">
          <div class="grid grid-cols-12 bg-white/[0.04] border-b border-white/10 py-5 px-8">
            <div class="col-span-2 text-[9px] font-black uppercase text-gray-500 tracking-widest">Timestamp</div>
            <div class="col-span-2 text-[9px] font-black uppercase text-gray-500 tracking-widest">Principal</div>
            <div class="col-span-2 text-[9px] font-black uppercase text-gray-500 tracking-widest">Operation</div>
            <div class="col-span-5 text-[9px] font-black uppercase text-gray-500 tracking-widest">Payload</div>
            <div class="col-span-1 text-[9px] font-black uppercase text-gray-500 text-right">Trace</div>
          </div>

          <div class="flex-1 overflow-y-auto custom-scrollbar">
            <div v-for="log in auditLogs" :key="log.id" @click="openLogDetails(log)" class="grid grid-cols-12 items-center py-5 px-8 border-b border-white/5 hover:bg-white/[0.03] cursor-pointer transition-all group">
              <div class="col-span-2">
                <p class="text-[10px] font-black text-white leading-none">{{ log.time }}</p>
                <p class="text-[8px] font-bold text-gray-600 mt-1 uppercase">{{ log.date }}</p>
              </div>
              <div class="col-span-2 flex items-center gap-3">
                <div class="w-1 h-4" :style="{ backgroundColor: themeStore.accentColor }"></div>
                <div>
                  <p class="text-[10px] font-black text-white uppercase">{{ log.actor }}</p>
                  <p class="text-[8px] text-gray-500 font-bold">{{ log.actorType }}</p>
                </div>
              </div>
              <div class="col-span-2">
                <span :class="['text-[8px] font-black px-2 py-1 border uppercase', log.severity === 'CRITICAL' ? 'bg-red-500/10 border-red-500/30 text-red-500' : 'bg-cyan-500/10 border-cyan-500/30 text-cyan-400']">
                  {{ log.action }}
                </span>
              </div>
              <div class="col-span-5">
                <code class="text-[9px] text-gray-500 font-mono truncate block max-w-md bg-black/40 px-2 py-1 rounded">{{ log.payload }}</code>
              </div>
              <div class="col-span-1 text-right">
                <svg class="w-3 h-3 text-gray-600 group-hover:text-cyan-400 ml-auto" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M9 5l7 7-7 7" stroke-width="4"/></svg>
              </div>
            </div>
          </div>
        </div>
      </main>

      <transition name="drawer-slide">
        <div v-if="isDrawerOpen" class="fixed inset-y-0 right-0 w-full max-w-2xl z-[100] bg-[#08080a] border-l border-white/10 shadow-[-40px_0_80px_rgba(0,0,0,0.9)] flex flex-col">
          
          <div class="h-32 px-10 border-b border-white/10 flex justify-between items-center shrink-0 bg-white/[0.01]">
            <div>
              <p class="text-[9px] font-black text-cyan-500 uppercase tracking-[0.3em] mb-1">Log_Forensics // {{ selectedLog?.id }}</p>
              <h2 class="text-3xl font-black italic uppercase tracking-tighter">DATA_INSPECT</h2>
            </div>
            <button @click="isDrawerOpen = false" class="text-gray-500 hover:text-white transition-all hover:rotate-90">✕</button>
          </div>

          <div class="flex-1 overflow-y-auto p-10 custom-scrollbar space-y-12">
            
            <section>
              <h4 class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-8 border-l-2 border-cyan-500 pl-3">I. Chain_Of_Custody</h4>
              <div class="space-y-6 relative before:absolute before:left-[7px] before:top-2 before:bottom-2 before:w-[1px] before:bg-white/10">
                <div v-for="(step, idx) in forensicSteps" :key="idx" class="flex items-start gap-6 pl-0">
                  <div class="w-4 h-4 rounded-full bg-black border-2 border-cyan-500 z-10 shrink-0 mt-1"></div>
                  <div class="bg-white/[0.02] border border-white/5 p-4 flex-1 rounded-sm">
                    <div class="flex justify-between items-start mb-2">
                      <p class="text-[10px] font-black text-white uppercase">{{ step.title }}</p>
                      <span class="text-[8px] font-mono text-gray-600">OFFSET: +{{ idx * 12 }}ms</span>
                    </div>
                    <p class="text-[10px] font-bold text-gray-500 uppercase leading-relaxed">{{ step.desc }}</p>
                    <div class="mt-3 flex items-center gap-2">
                      <div class="px-2 py-0.5 bg-green-500/10 text-green-500 text-[7px] font-black rounded-full uppercase">Verified</div>
                      <code class="text-[7px] text-gray-700 font-mono">{{ step.hash }}</code>
                    </div>
                  </div>
                </div>
              </div>
            </section>

            <section>
              <h4 class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-6 border-l-2 border-amber-500 pl-3">II. Raw_Technical_Payload</h4>
              <div class="bg-black p-6 border border-white/5 rounded-sm relative group overflow-hidden">
                <div class="absolute top-0 left-0 w-full h-[1px] bg-gradient-to-r from-transparent via-amber-500/50 to-transparent"></div>
                <pre class="text-[10px] text-amber-500/80 font-mono leading-relaxed whitespace-pre-wrap">
{
  "trace_id": "{{ selectedLog?.id }}",
  "principal": "{{ selectedLog?.actor }}",
  "network": { "ip": "{{ selectedLog?.ip }}", "cluster": "EU-WEST-1" },
  "action": "{{ selectedLog?.action }}",
  "data": {{ selectedLog?.payload }},
  "integrity_hash": "sha256:77ea28b..."
}
                </pre>
              </div>
            </section>
          </div>

          <div class="p-10 border-t border-white/10 flex gap-4 bg-black/80">
            <button @click="triggerReportDownload" :disabled="isDownloadingReport" class="flex-1 py-5 border border-white/10 text-[10px] font-black uppercase hover:bg-white/10 transition-all relative">
              <span :class="{ 'opacity-0': isDownloadingReport }">Download_Forensic_PDF</span>
              <div v-if="isDownloadingReport" class="absolute inset-0 flex items-center justify-center">
                <div class="w-4 h-4 border-2 border-white/40 border-t-white rounded-full animate-spin"></div>
              </div>
            </button>
          </div>
        </div>
      </transition>

      <transition name="toast-pop">
        <div v-if="toast.show" class="fixed bottom-12 right-12 z-[200] bg-black border-l-4 p-6 min-w-[340px] shadow-2xl" :style="{ borderColor: toast.type === 'danger' ? '#ef4444' : (toast.type === 'process' ? '#fbbf24' : '#06b6d4') }">
          <div class="flex items-center gap-4">
            <div v-if="toast.type === 'process'" class="w-4 h-4 border-2 border-amber-500 border-t-transparent rounded-full animate-spin"></div>
            <div>
              <p class="text-[10px] font-black uppercase opacity-50 tracking-widest">{{ toast.label }}</p>
              <p class="text-xs font-black mt-1 uppercase">{{ toast.message }}</p>
            </div>
          </div>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const isDrawerOpen = ref(false)
const selectedLog = ref(null)
const isExporting = ref(false)
const isDownloadingReport = ref(false)
const toast = ref({ show: false, message: '', label: '', type: 'success' })

const forensicSteps = [
  { title: 'Event_Ingress', desc: 'Request received via secondary cluster gateway.', hash: '88b1...af01' },
  { title: 'Identity_Validation', desc: 'Actor principal authenticated via MFA token.', hash: '99x2...cc12' },
  { title: 'Payload_Persistence', desc: 'Audit entry written to immutable ledger.', hash: '44z8...ee99' }
]

const auditLogs = ref([
  { id: 'TRC-9921X', date: '2025-12-31', time: '22:15:02', actor: 'NEURAL_FILTER', actorType: 'SYSTEM', ip: 'internal_01', action: 'GHOST_RESTRICTION', severity: 'CRITICAL', payload: '{"target": "EMP-9921", "heuristic": "SCAM_PATTERN"}' },
])

const openLogDetails = (log) => { selectedLog.value = log; isDrawerOpen.value = true; }

const triggerExport = async () => {
  isExporting.value = true
  toast.value = { show: true, label: 'IO_PROTOCOL', message: 'COLLECTING_DATABASE_SHARDS...', type: 'process' }
  await new Promise(r => setTimeout(r, 2000))
  isExporting.value = false
  toast.value = { show: true, label: 'IO_PROTOCOL', message: 'CSV_EXPORT_SUCCESSFUL', type: 'success' }
  setTimeout(() => toast.value.show = false, 3000)
}

const triggerReportDownload = async () => {
  isDownloadingReport.value = true
  toast.value = { show: true, label: 'FORENSIC_ENGINE', message: 'SIGNING_FORENSIC_OBJECT...', type: 'process' }
  await new Promise(r => setTimeout(r, 2000))
  isDownloadingReport.value = false
  toast.value = { show: true, label: 'FORENSIC_ENGINE', message: 'REPORT_DOWNLOAD_STARTED', type: 'success' }
  setTimeout(() => toast.value.show = false, 3000)
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }
.drawer-slide-enter-active, .drawer-slide-leave-active { transition: transform 0.6s cubic-bezier(0.16, 1, 0.3, 1); }
.drawer-slide-enter-from, .drawer-slide-leave-to { transform: translateX(100%); }
.toast-pop-enter-active, .toast-pop-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-pop-enter-from, .toast-pop-leave-to { transform: translateY(30px); opacity: 0; }
</style>