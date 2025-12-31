<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <div :class="['flex-1 flex flex-col h-full relative transition-all duration-500', themeStore.terminalCollapsed ? 'ml-20' : 'ml-72']">
      
      <header class="h-32 flex flex-col justify-center px-12 shrink-0 border-b border-white/5 bg-black/60 relative overflow-hidden z-20">
        <div class="absolute inset-0 bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.25)_50%),linear-gradient(90deg,rgba(255,0,0,0.06),rgba(0,255,0,0.02),rgba(0,0,255,0.06))] bg-[length:100%_2px,3px_100%] pointer-events-none"></div>
        <span class="text-[10px] font-black tracking-[0.4em] uppercase opacity-40 text-cyan-400">Registry_v4.0 // Secure_Isolation</span>
        <h1 class="text-5xl font-black italic uppercase tracking-tighter">SHADOW_REGISTRY</h1>
      </header>

      <main class="flex-1 overflow-hidden p-8 lg:p-12 bg-[#020205] z-10">
        <div class="h-full border border-white/10 rounded-sm flex flex-col bg-white/[0.01] backdrop-blur-sm overflow-hidden shadow-[0_0_50px_rgba(0,0,0,1)]">
          
          <div class="grid grid-cols-12 bg-white/[0.04] border-b border-white/10 py-5 px-8">
            <div class="col-span-1 text-[9px] font-black uppercase text-gray-500 tracking-widest text-center">Sig.</div>
            <div class="col-span-4 text-[9px] font-black uppercase text-gray-500 tracking-widest">Entity_Profile</div>
            <div class="col-span-2 text-[9px] font-black uppercase text-gray-500 tracking-widest text-center">Stealth_Lvl</div>
            <div class="col-span-2 text-[9px] font-black uppercase text-gray-500 tracking-widest text-center">Intercepts</div>
            <div class="col-span-3 text-[9px] font-black uppercase text-gray-500 tracking-widest text-right">Protocols</div>
          </div>

          <div class="flex-1 overflow-y-auto custom-scrollbar">
            <div 
              v-for="entity in shadowList" :key="entity.id" 
              @click="openDossier(entity)"
              class="grid grid-cols-12 items-center py-5 px-8 border-b border-white/5 hover:bg-white/[0.03] cursor-crosshair transition-all group relative overflow-hidden"
            >
              <div class="col-span-1 flex justify-center">
                <div class="relative">
                  <div class="w-2 h-2 rounded-full bg-amber-500 animate-pulse"></div>
                  <div class="absolute inset-0 w-2 h-2 rounded-full bg-amber-500 animate-ping opacity-40"></div>
                </div>
              </div>
              
              <div class="col-span-4">
                <div class="flex items-center gap-3">
                  <div class="h-8 w-[2px]" :style="{ backgroundColor: themeStore.accentColor }"></div>
                  <div>
                    <p class="text-xs font-black uppercase tracking-wider text-white group-hover:text-cyan-400 transition-colors">{{ entity.name }}</p>
                    <p class="text-[9px] text-gray-500 font-bold uppercase mt-0.5">{{ entity.uid }} // <span class="text-gray-400">{{ entity.origin }}</span></p>
                  </div>
                </div>
              </div>

              <div class="col-span-2 px-6">
                <div class="w-full h-1 bg-white/5 rounded-full overflow-hidden">
                  <div class="h-full bg-amber-500/60" :style="{ width: entity.stealthLevel + '%' }"></div>
                </div>
                <p class="text-[8px] font-black text-gray-600 mt-2 text-center uppercase">{{ entity.stealthLevel }}% Ghosting</p>
              </div>

              <div class="col-span-2 text-center">
                <span class="text-xl font-black italic text-white tracking-tighter">{{ entity.ghostActions }}</span>
              </div>

              <div class="col-span-3 flex justify-end gap-3" @click.stop>
                <button @click="handleRehab(entity)" class="group/btn px-4 py-2 bg-white/[0.03] border border-white/10 text-[9px] font-black uppercase hover:border-white/40 transition-all">
                  <span class="group-hover/btn:tracking-[0.1em] transition-all">Restore</span>
                </button>
                <button @click="handleHardBan(entity)" class="px-4 py-2 bg-red-600/10 text-red-500 border border-red-500/20 text-[9px] font-black uppercase hover:bg-red-600 hover:text-white transition-all">
                  Term
                </button>
              </div>
            </div>
          </div>
        </div>
      </main>

      <transition name="dossier-slide">
        <div v-if="isDossierOpen" class="fixed inset-y-0 right-0 w-full max-w-2xl z-[100] bg-[#050508] border-l border-white/20 shadow-[-50px_0_100px_rgba(0,0,0,0.9)] flex flex-col">
          
          <div class="h-32 px-10 border-b border-white/10 flex justify-between items-center shrink-0 bg-white/[0.02] backdrop-blur-md">
            <div>
              <p class="text-[9px] font-black text-cyan-500 uppercase tracking-[0.3em] mb-1">Dossier_ID // {{ selectedEntity?.uid }}</p>
              <h2 class="text-4xl font-black italic uppercase tracking-tighter">{{ selectedEntity?.name }}</h2>
            </div>
            <button @click="isDossierOpen = false" class="group p-4 text-gray-400 hover:text-white transition-all">
              <svg class="w-6 h-6 group-hover:rotate-90 transition-transform" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M6 18L18 6M6 6l12 12" stroke-width="3"/></svg>
            </button>
          </div>

          <div class="flex-1 overflow-y-auto custom-scrollbar p-10 space-y-12 bg-[radial-gradient(circle_at_top_right,rgba(255,255,255,0.02),transparent_40%)]">
            
            <section>
              <h4 class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-6 border-l-2 border-cyan-500 pl-3">I. Request_Telemetry</h4>
              <div class="p-6 bg-black border border-white/5 rounded-sm shadow-inner">
                <div class="h-56">
                  <canvas id="behaviorChart"></canvas>
                </div>
              </div>
            </section>

            <section>
              <h4 class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-6 border-l-2 border-red-500 pl-3">II. Technical_Anomalies</h4>
              <div class="grid grid-cols-2 gap-4">
                <div v-for="tag in ['VPN_EXIT_NODE', 'SPOOFED_OS', 'LATENCY_BURST', 'HIGH_CONCURRENCY']" :key="tag"
                     class="p-4 bg-white/[0.02] border border-white/5 flex items-center gap-3">
                  <div class="w-1.5 h-1.5 bg-red-500"></div>
                  <span class="text-[9px] font-black text-gray-400 uppercase tracking-tighter">{{ tag }}</span>
                </div>
              </div>
            </section>

            <section>
              <h4 class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-6 border-l-2 border-white/40 pl-3">III. Heuristic_Violations</h4>
              <div class="space-y-4">
                <div v-for="reason in selectedEntity?.reasons" :key="reason.label" 
                     class="p-6 bg-red-500/5 border border-red-500/10 relative overflow-hidden group">
                  <div class="absolute top-0 right-0 p-1 bg-red-500/10 text-[7px] font-black text-red-500 uppercase">{{ reason.severity }}</div>
                  <h5 class="text-xs font-black uppercase text-gray-300 mb-2">{{ reason.label }}</h5>
                  <p class="text-[10px] text-gray-500 leading-relaxed font-bold uppercase">Target system triggered auto-suppression based on high-similarity match with known bot-pattern database v8.2.</p>
                </div>
              </div>
            </section>

            <div class="h-10"></div>
          </div>

          <div class="p-10 border-t border-white/10 flex gap-4 bg-black/80 backdrop-blur-lg">
            <button @click="handleRehab(selectedEntity)" class="flex-1 py-5 border border-white/10 text-[10px] font-black uppercase hover:bg-white/10 transition-all tracking-widest">Execute_Rehabilitation</button>
            <button @click="handleHardBan(selectedEntity)" class="flex-1 py-5 bg-red-600 text-white text-[10px] font-black uppercase hover:bg-red-500 transition-all tracking-widest">Hard_Ban_Protocol</button>
          </div>
        </div>
      </transition>

      <transition name="toast-pop">
        <div v-if="toast.show" class="fixed bottom-12 right-12 z-[200] bg-black border-l-4 p-6 min-w-[340px] shadow-2xl" :style="{ borderColor: toast.type === 'danger' ? '#ef4444' : themeStore.accentColor }">
          <p class="text-[10px] font-black uppercase opacity-50 tracking-widest">Security_Signal_Intercepted</p>
          <p class="text-xs font-black mt-1 uppercase">{{ toast.message }}</p>
        </div>
      </transition>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue'
import { useThemeStore } from '~/stores/theme'
import Chart from 'chart.js/auto'

const themeStore = useThemeStore()
const isDossierOpen = ref(false)
const selectedEntity = ref(null)
const toast = ref({ show: false, message: '', label: '', type: 'success' })

const shadowList = ref([
  { id: 1, name: 'Quantum_Solutions_LLC', uid: 'EMP-9921', origin: 'Singapore', stealthLevel: 85, ghostActions: 421, reasons: [{ label: 'Bulk_Job_Scraping', severity: 'CRITICAL' }, { label: 'VPN_Node', severity: 'HIGH' }] },
  { id: 2, name: 'Dev_Shadow_X', uid: 'USR-1102', origin: 'Unknown_Proxy', stealthLevel: 40, ghostActions: 88, reasons: [{ label: 'Automated_Apply_Bot', severity: 'CRITICAL' }] }
])

let chartInstance = null

const initChart = () => {
  const ctx = document.getElementById('behaviorChart')
  if (!ctx || chartInstance) return
  
  chartInstance = new Chart(ctx, {
    type: 'line',
    data: {
      labels: ['00:00', '04:00', '08:00', '12:00', '16:00', '20:00', '23:59'],
      datasets: [{
        label: 'API Calls',
        data: [12, 19, 300, 280, 240, 400, 390], 
        borderColor: '#ef4444',
        backgroundColor: 'rgba(239, 68, 68, 0.05)',
        borderWidth: 2,
        fill: true,
        tension: 0.2,
        pointBackgroundColor: '#ef4444',
        pointRadius: 4,
        pointHoverRadius: 8
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: { legend: { display: false } },
      scales: {
        x: { grid: { color: 'rgba(255,255,255,0.02)' }, ticks: { color: '#666', font: { size: 9, family: 'monospace' } } },
        y: { grid: { color: 'rgba(255,255,255,0.02)' }, ticks: { color: '#666', font: { size: 9, family: 'monospace' } } }
      }
    }
  })
}

const openDossier = async (entity) => {
  selectedEntity.value = entity
  isDossierOpen.value = true
  await nextTick()
  initChart()
}

const handleRehab = (entity) => {
  shadowList.value = shadowList.value.filter(e => e.id !== entity.id)
  toast.value = { show: true, message: `${entity.name} REHABILITATED`, type: 'success' }
  isDossierOpen.value = false
  setTimeout(() => toast.value.show = false, 3000)
}

const handleHardBan = (entity) => {
  shadowList.value = shadowList.value.filter(e => e.id !== entity.id)
  toast.value = { show: true, message: `${entity.name} HARD BANNED`, type: 'danger' }
  isDossierOpen.value = false
  setTimeout(() => toast.value.show = false, 3000)
}
</script>

<style scoped>
@reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }

.dossier-slide-enter-active, .dossier-slide-leave-active { transition: all 0.6s cubic-bezier(0.16, 1, 0.3, 1); }
.dossier-slide-enter-from, .dossier-slide-leave-to { transform: translateX(100%); }

@keyframes scanline { 0% { bottom: 100%; } 80% { bottom: 100%; } 100% { bottom: 0%; } }
</style>