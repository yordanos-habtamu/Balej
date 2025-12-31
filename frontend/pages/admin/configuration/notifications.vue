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
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">Alert_Monitor_v2.0</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                NOTIFICATIONS
              </h1>
            </div>

            <div class="flex gap-4">
              <button class="px-6 py-3 border border-white/10 text-[10px] font-black uppercase tracking-widest hover:bg-white hover:text-black transition-all">
                [ CLEAR_READ_LOGS ]
              </button>
              <button class="px-6 py-3 text-[10px] font-black uppercase tracking-widest text-black transition-all" :style="{ backgroundColor: themeStore.accentColor }">
                MARK_ALL_AS_SEEN
              </button>
            </div>
          </div>

          <div class="grid grid-cols-1 gap-4">
            <div v-for="note in notifications" :key="note.id" 
                 class="relative overflow-hidden bg-black border border-white/5 group hover:border-white/20 transition-all duration-300">
              
              <div class="absolute left-0 top-0 bottom-0 w-1" :class="getSeverityColor(note.type)"></div>

              <div class="p-6 flex flex-col lg:flex-row lg:items-center gap-6">
                <div class="min-w-[140px]">
                  <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest mb-1">{{ note.timestamp }}</p>
                  <p class="text-xs font-black text-white/40 uppercase">#REF_{{ note.id }}</p>
                </div>

                <div class="w-12 h-12 flex-shrink-0 flex items-center justify-center border border-white/10 bg-white/[0.02]">
                  <svg v-if="note.type === 'alert'" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                  </svg>
                  <svg v-else xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" :style="{ color: themeStore.accentColor }" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16h-1v-4h-1m1-4h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z" />
                  </svg>
                </div>

                <div class="flex-1">
                  <div class="flex items-center gap-3 mb-1">
                    <span class="text-[10px] font-black uppercase tracking-widest" :class="getSeverityText(note.type)">
                      [{{ note.category }}]
                    </span>
                    <span v-if="!note.read" class="w-1.5 h-1.5 rounded-full bg-cyan-400 animate-pulse"></span>
                  </div>
                  <h3 class="text-sm lg:text-base font-bold text-white uppercase tracking-tight mb-1">
                    {{ note.title }}
                  </h3>
                  <p class="text-xs text-gray-400 uppercase tracking-wide leading-relaxed">
                    {{ note.message }}
                  </p>
                </div>

                <div class="flex items-center gap-4 lg:border-l lg:border-white/5 lg:pl-8">
                  <button class="text-[10px] font-bold text-gray-500 hover:text-white uppercase tracking-widest transition-colors">
                    DISMISS
                  </button>
                  <button class="px-4 py-2 bg-white/5 hover:bg-white/10 text-[10px] font-black uppercase tracking-widest transition-all">
                    VIEW_LOG
                  </button>
                </div>
              </div>

              <div class="absolute inset-0 pointer-events-none opacity-0 group-hover:opacity-10 transition-opacity bg-gradient-to-b from-transparent via-white to-transparent h-20 -top-20 animate-scanline"></div>
            </div>
          </div>

          <div v-if="notifications.length === 0" class="py-20 text-center border border-dashed border-white/10">
            <p class="text-xs font-bold text-gray-500 uppercase tracking-[0.5em]">System_All_Clear</p>
          </div>

        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()

const notifications = ref([
  {
    id: '7721',
    type: 'alert',
    category: 'SECURITY',
    timestamp: '14:02:31 UTC',
    title: 'Unauthorized Access Attempt',
    message: 'Multiple failed login attempts detected from IP: 192.168.1.104. Region: Eastern_Sector.',
    read: false
  },
  {
    id: '7690',
    type: 'info',
    category: 'TALENT',
    timestamp: '12:45:10 UTC',
    title: 'New Unit Registration',
    message: 'User "Cypher_Protocol" has completed biometric verification and joined the registry.',
    read: false
  },
  {
    id: '7685',
    type: 'info',
    category: 'FINANCE',
    timestamp: '09:15:00 UTC',
    title: 'Payout Dispatched',
    message: 'Mission payment for Project_Genesis has been finalized and sent to 12 assigned nodes.',
    read: true
  }
])

const getSeverityColor = (type: string) => {
  if (type === 'alert') return 'bg-red-500'
  return `bg-[${themeStore.accentColor}]` || 'bg-cyan-500'
}

const getSeverityText = (type: string) => {
  if (type === 'alert') return 'text-red-500'
  return 'text-gray-400'
}
</script>

<style scoped>
@reference 'tailwindcss';

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #1f2937; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: v-bind('themeStore.accentColor'); }

@keyframes scanline {
  0% { transform: translateY(0); }
  100% { transform: translateY(500px); }
}

.animate-scanline {
  animation: scanline 3s linear infinite;
}

/* Subtle staggered enter animation */
.grid > div {
  animation: slideLeft 0.4s cubic-bezier(0.2, 1, 0.3, 1) forwards;
  opacity: 0;
}

@keyframes slideLeft {
  from { opacity: 0; transform: translateX(10px); }
  to { opacity: 1; transform: translateX(0); }
}
</style>