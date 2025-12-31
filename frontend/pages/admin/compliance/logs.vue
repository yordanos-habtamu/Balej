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

      <div class="flex-1 flex flex-col relative z-10 px-6 py-8 lg:px-12 overflow-hidden">
        <div class="max-w-7xl mx-auto w-full h-full flex flex-col">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-8 shrink-0">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12 bg-white/20"></div>
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">System_Telemetry_v4.0</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                SIGNAL_FEED
              </h1>
            </div>

            <div class="flex gap-4">
              <button 
                @click="clearLogs"
                class="px-6 py-4 text-[10px] font-black uppercase tracking-widest border border-white/10 hover:bg-white/5 transition-all flex items-center gap-3"
              >
                PURGE_BUFFER
              </button>
              <button 
                class="px-8 py-4 text-[10px] font-black uppercase tracking-widest text-black transition-all hover:brightness-110" 
                :style="{ backgroundColor: themeStore.accentColor }"
              >
                EXPORT_CSV
              </button>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-6 shrink-0 text-white">
            <div v-for="stat in logStats" :key="stat.label" class="bg-white/[0.02] border border-white/5 p-4 relative overflow-hidden">
              <p class="text-[9px] font-black text-gray-500 uppercase tracking-widest">{{ stat.label }}</p>
              <p class="text-xl font-black italic tracking-tighter mt-1">{{ stat.value }}</p>
              <div class="absolute bottom-0 left-0 h-[1px]" :style="{ width: stat.width, backgroundColor: stat.color }"></div>
            </div>
          </div>

          <div class="bg-black/40 border border-white/10 p-4 mb-4 flex flex-wrap gap-4 items-center shrink-0">
            <input 
              v-model="search" 
              placeholder="FILTER_BY_EVENT_ID..." 
              class="bg-transparent border border-white/10 px-4 py-2 text-[10px] font-bold uppercase tracking-widest focus:outline-none focus:border-cyan-500/50 w-64"
            />
            <div class="flex gap-2">
              <button v-for="lvl in ['ALL', 'INFO', 'WARN', 'CRIT']" :key="lvl" 
                @click="filterLevel = lvl"
                class="px-3 py-1 text-[9px] font-black border transition-all"
                :class="filterLevel === lvl ? 'bg-white text-black border-white' : 'border-white/10 text-gray-500 hover:text-white'"
              >
                {{ lvl }}
              </button>
            </div>
          </div>

          <div class="flex-1 flex gap-4 overflow-hidden min-h-0">
            
            <div class="flex-1 bg-black border border-white/10 overflow-hidden flex flex-col relative">
              <div class="absolute inset-0 pointer-events-none bg-[linear-gradient(rgba(18,16,16,0)_50%,rgba(0,0,0,0.1)_50%),linear-gradient(90deg,rgba(255,0,0,0.02),rgba(0,255,0,0.01),rgba(0,0,255,0.02))] z-20 bg-[length:100%_2px,3px_100%]"></div>
              
              <div class="flex-1 overflow-y-auto custom-scrollbar p-6 font-mono text-[11px] space-y-1 relative z-10">
                <div 
                  v-for="log in filteredLogs" 
                  :key="log.id" 
                  @click="selectedLog = log"
                  :class="[
                    'flex gap-6 group/item cursor-pointer p-1 border-l-2 transition-all',
                    selectedLog?.id === log.id ? 'bg-white/[0.05] border-white' : 'border-transparent hover:bg-white/[0.02]'
                  ]"
                >
                  <span class="text-gray-600 shrink-0">[{{ log.timestamp }}]</span>
                  <span :class="getLevelColor(log.level)" class="w-12 shrink-0 font-black">[{{ log.level }}]</span>
                  <span class="text-gray-400 shrink-0 font-bold uppercase tracking-tighter w-24">{{ log.origin }}</span>
                  <span class="text-gray-300 tracking-tight flex-1 truncate">{{ log.message }}</span>
                </div>
              </div>

              <div class="p-3 border-t border-white/10 bg-white/[0.02] flex justify-between items-center relative z-10">
                <div class="flex gap-4 items-center">
                  <span class="flex h-2 w-2 rounded-full bg-green-500 animate-ping"></span>
                  <span class="text-[9px] font-black text-gray-500 uppercase tracking-widest">Uplink: Live</span>
                </div>
                <span class="text-[9px] font-black text-gray-600 uppercase tracking-widest">Logs_Cached: {{ logs.length }}</span>
              </div>
            </div>

            <transition name="slide">
              <div v-if="selectedLog" class="w-96 bg-[#050508] border border-white/10 flex flex-col shadow-2xl overflow-hidden">
                <div class="p-4 border-b border-white/10 bg-white/[0.02] flex justify-between items-center">
                  <span class="text-[10px] font-black uppercase tracking-widest">Event_Trace_Analysis</span>
                  <button @click="selectedLog = null" class="text-gray-500 hover:text-white">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M6 18L18 6M6 6l12 12" stroke-width="2"/></svg>
                  </button>
                </div>
                <div class="flex-1 p-6 overflow-y-auto custom-scrollbar font-mono">
                  <div class="space-y-6">
                    <div>
                      <label class="text-[9px] font-black text-gray-600 uppercase">Packet_Header</label>
                      <p class="text-xs text-white mt-1 break-all">{{ selectedLog.id }}</p>
                    </div>
                    <div>
                      <label class="text-[9px] font-black text-gray-600 uppercase">Origin_Node</label>
                      <p class="text-xs text-cyan-400 mt-1">{{ selectedLog.origin }}</p>
                    </div>
                    <div>
                      <label class="text-[9px] font-black text-gray-600 uppercase">Raw_Payload</label>
                      <div class="mt-2 bg-black p-4 border border-white/5 text-[10px] text-green-500/80 leading-relaxed whitespace-pre-wrap">
{
  "timestamp": "{{ selectedLog.timestamp }}",
  "level": "{{ selectedLog.level }}",
  "message": "{{ selectedLog.message }}",
  "node_id": "XRS-{{ Math.floor(Math.random() * 9999) }}",
  "status": "Verified",
  "protocol": "Web3_v2"
}
                      </div>
                    </div>
                  </div>
                </div>
                <div class="p-4 bg-white/[0.02] border-t border-white/10">
                  <button class="w-full py-3 text-[9px] font-black uppercase tracking-widest border border-cyan-500/50 text-cyan-500 hover:bg-cyan-500 hover:text-black transition-all">
                    DOWNLOAD_TRACE_LOG
                  </button>
                </div>
              </div>
            </transition>
          </div>

        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const search = ref('')
const filterLevel = ref('ALL')
const selectedLog = ref(null)

const logStats = ref([
  { label: 'Total_Signals', value: '142,801', width: '100%', color: '#333' },
  { label: 'Critical_Errors', value: '02', width: '5%', color: '#ef4444' },
  { label: 'Security_Events', value: '28', width: '25%', color: '#f59e0b' },
  { label: 'Data_Injections', value: '842', width: '60%', color: '#00f3ff' }
])

const logs = ref([
  { id: crypto.randomUUID(), timestamp: '2025-12-31 20:44:01', level: 'INFO', origin: 'AUTH_SRV', message: 'User_Admin session initialized from IP 192.168.1.1' },
  { id: crypto.randomUUID(), timestamp: '2025-12-31 20:45:12', level: 'WARN', origin: 'DB_CORE', message: 'Latency spike detected in CATEGORY_MATRIX query' },
  { id: crypto.randomUUID(), timestamp: '2025-12-31 20:46:00', level: 'INFO', origin: 'JOB_INJECT', message: 'New listing "Lead Neural Engineer" indexed' },
  { id: crypto.randomUUID(), timestamp: '2025-12-31 20:47:33', level: 'CRIT', origin: 'PAY_GATE', message: 'Handshake timeout on Stripe_Node_04' }
])

const filteredLogs = computed(() => {
  return logs.value.filter(l => {
    const matchesSearch = l.message.toLowerCase().includes(search.value.toLowerCase()) || 
                          l.origin.toLowerCase().includes(search.value.toLowerCase());
    const matchesLevel = filterLevel.value === 'ALL' || l.level === filterLevel.value;
    return matchesSearch && matchesLevel;
  });
})

const getLevelColor = (level: string) => {
  switch (level) {
    case 'CRIT': return 'text-red-500';
    case 'WARN': return 'text-amber-500';
    case 'INFO': return 'text-cyan-400';
    default: return 'text-gray-400';
  }
}

const clearLogs = () => {
  themeStore.initializeUplink();
  logs.value = [];
  selectedLog.value = null;
}

// Simulated real-time logging daemon
onMounted(() => {
  setInterval(() => {
    if(logs.value.length > 100) logs.value.pop();
    const newLog = {
      id: crypto.randomUUID(),
      timestamp: new Date().toISOString().replace('T', ' ').slice(0, 19),
      level: Math.random() > 0.9 ? 'WARN' : 'INFO',
      origin: 'SYS_DAEMON',
      message: `Heartbeat_Protocol_X${Math.floor(Math.random() * 99)}_Active_Sync`
    }
    logs.value.unshift(newLog);
  }, 4000);
})
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.05); }

/* Sidebar Slide Animation */
.slide-enter-active, .slide-leave-active { transition: transform 0.4s cubic-bezier(0.2, 1, 0.3, 1); }
.slide-enter-from, .slide-leave-to { transform: translateX(100%); }

/* CRT Scanline Pulse */
@keyframes pulse-opacity {
  0% { opacity: 0.1; }
  50% { opacity: 0.15; }
  100% { opacity: 0.1; }
}
</style>