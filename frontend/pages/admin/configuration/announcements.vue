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
        <div class="max-w-5xl mx-auto">
          
          <div class="mb-12 relative">
            <div class="flex items-center gap-3 mb-4">
              <div class="h-[1px] w-12 bg-white/20"></div>
              <span class="text-xs font-bold tracking-widest text-gray-400 uppercase italic">Central_Comm_Hub</span>
            </div>
            <h1 class="text-5xl lg:text-7xl font-black italic uppercase tracking-tighter mb-4">
              SYSTEM_ANNOUNCEMENTS
            </h1>
            <p class="text-gray-500 text-sm max-w-2xl uppercase tracking-wider leading-relaxed">
              Real-time broadcast of system updates, mission directives, and portal maintenance logs.
            </p>
          </div>

          <div class="space-y-6 relative pb-20">
            <div class="absolute left-0 top-0 bottom-0 w-[1px] bg-white/10 ml-[11px]"></div>

            <div v-for="announcement in announcements" :key="announcement.id" class="relative pl-12 group">
              <div class="absolute left-0 top-6 w-6 h-6 bg-black border-2 z-10 flex items-center justify-center transition-all group-hover:scale-110" 
                   :style="{ borderColor: announcement.type === 'critical' ? '#ef4444' : themeStore.accentColor }">
                <div class="w-1.5 h-1.5" :style="{ backgroundColor: announcement.type === 'critical' ? '#ef4444' : themeStore.accentColor }"></div>
              </div>

              <div class="bg-white/[0.02] border border-white/5 p-6 lg:p-8 hover:bg-white/[0.04] hover:border-white/20 transition-all duration-300">
                <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 mb-4">
                  <div class="flex items-center gap-4">
                    <span class="text-[10px] font-black px-2 py-0.5 border border-white/20 uppercase tracking-widest text-gray-400">
                      ID_{{ announcement.id }}
                    </span>
                    <span :class="[
                      'text-[10px] font-black uppercase tracking-widest',
                      announcement.type === 'critical' ? 'text-red-500' : 'text-blue-400'
                    ]">
                      // {{ announcement.category }}
                    </span>
                  </div>
                  <span class="text-[10px] text-gray-500 font-bold tracking-widest uppercase">
                    TIMESTAMP: {{ announcement.date }}
                  </span>
                </div>

                <h2 class="text-xl lg:text-2xl font-black uppercase italic tracking-tight mb-4 group-hover:text-white transition-colors">
                  {{ announcement.title }}
                </h2>

                <p class="text-gray-400 text-sm leading-relaxed uppercase tracking-wide mb-6">
                  {{ announcement.content }}
                </p>

                <div class="flex items-center justify-between pt-4 border-t border-white/5">
                  <div class="flex items-center gap-2">
                    <div class="w-6 h-6 rounded-full bg-white/10 flex items-center justify-center">
                       <svg xmlns="http://www.w3.org/2000/svg" class="h-3 w-3 text-gray-400" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
                       </svg>
                    </div>
                    <span class="text-[10px] font-bold text-gray-500 uppercase">AUTH_BY: {{ announcement.author }}</span>
                  </div>
                  
                  <button class="text-[10px] font-black uppercase tracking-[0.2em] flex items-center gap-2 group/btn" :style="{ color: themeStore.accentColor }">
                    READ_FULL_LOG 
                    <span class="group-hover/btn:translate-x-1 transition-transform">»</span>
                  </button>
                </div>
              </div>
            </div>
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

// Mock data for announcements
const announcements = ref([
  {
    id: 'BC-902',
    type: 'critical',
    category: 'SECURITY_ALERT',
    date: '31_DEC_2025',
    author: 'SYS_ADMIN_01',
    title: 'Biometric Uplink Maintenance',
    content: 'All neural-link connections will be throttled for 120 minutes starting 04:00 UTC to upgrade decryption protocols.'
  },
  {
    id: 'BC-884',
    type: 'standard',
    category: 'PLATFORM_UPDATE',
    date: '28_DEC_2025',
    author: 'PROD_LEAD',
    title: 'New Talent Nodes Active',
    content: 'The Job Portal has integrated 500+ new verified blockchain developer nodes from the Neo-Tokyo sector.'
  },
  {
    id: 'BC-871',
    type: 'standard',
    category: 'REWARD_DIRECTIVE',
    date: '20_DEC_2025',
    author: 'FIN_EXEC',
    title: 'Token Incentive Program',
    content: 'Early adopters with over 100 hours of uplink time are now eligible for the Phase 2 token distribution.'
  }
])
</script>

<style scoped>
@reference 'tailwindcss';

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #1f2937; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: v-bind('themeStore.accentColor'); }

.group {
  animation: slideIn 0.5s cubic-bezier(0.2, 1, 0.3, 1) forwards;
  opacity: 0;
}

@keyframes slideIn {
  from { opacity: 0; transform: translateX(-10px); }
  to { opacity: 1; transform: translateX(0); }
}

/* Staggered animation for list items */
.group:nth-child(1) { animation-delay: 0.1s; }
.group:nth-child(2) { animation-delay: 0.2s; }
.group:nth-child(3) { animation-delay: 0.3s; }
</style>