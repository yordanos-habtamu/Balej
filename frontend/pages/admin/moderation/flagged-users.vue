<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-white/10">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)] relative',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-[0.05]" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 flex flex-col relative z-10 px-6 py-8 lg:px-12 overflow-hidden">
        <div class="max-w-7xl mx-auto w-full h-full flex flex-col">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10 shrink-0">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12" :style="{ backgroundColor: themeStore.accentColor + '66' }"></div>
                <span class="text-xs font-black tracking-widest uppercase opacity-60">Identity_Security // Oversight</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                FLAGGED_ENTITIES
              </h1>
            </div>

            <div class="flex gap-4">
              <div class="bg-white/[0.02] border border-white/10 p-4 min-w-[140px]">
                <p class="text-[9px] font-black text-gray-500 uppercase">Critical_Threats</p>
                <p class="text-xl font-black text-red-500 mt-1">08</p>
              </div>
              <div class="bg-white/[0.02] border border-white/10 p-4 min-w-[140px]">
                <p class="text-[9px] font-black text-gray-500 uppercase">Auto_Isolated</p>
                <p class="text-xl font-black mt-1" :style="{ color: themeStore.accentColor }">124</p>
              </div>
            </div>
          </div>

          <div class="bg-black/40 border border-white/10 p-4 mb-6 flex flex-wrap gap-4 items-center shrink-0">
            <div class="relative flex-1 max-w-md">
              <input 
                v-model="search" 
                placeholder="SEARCH_BY_UID_OR_EMAIL..." 
                class="w-full bg-transparent border border-white/10 px-4 py-3 text-[10px] font-black uppercase tracking-widest focus:outline-none focus:border-white transition-all"
              />
            </div>
            <div class="flex gap-2">
              <button v-for="risk in ['ALL', 'CRITICAL', 'ELEVATED', 'LOW']" :key="risk"
                @click="filterRisk = risk"
                class="px-4 py-2 text-[9px] font-black border transition-all"
                :class="filterRisk === risk ? 'bg-white text-black border-white' : 'border-white/10 text-gray-500 hover:text-white'"
              >
                {{ risk }}
              </button>
            </div>
          </div>

          <div class="flex-1 bg-black border border-white/10 overflow-hidden shadow-2xl flex flex-col">
            <div class="overflow-x-auto custom-scrollbar flex-1">
              <table class="w-full text-left border-collapse">
                <thead>
                  <tr class="border-b border-white/10 bg-white/[0.02] sticky top-0 z-20">
                    <th class="p-6 text-[10px] font-black uppercase tracking-widest text-gray-500">Identity_Meta</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-widest text-gray-500">Risk_Vectors</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-widest text-gray-500">Threat_Index</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-widest text-gray-500">Status</th>
                    <th class="p-6 text-[10px] font-black uppercase tracking-widest text-gray-500 text-right">Actions</th>
                  </tr>
                </thead>
                <tbody class="divide-y divide-white/5">
                  <tr v-for="user in filteredUsers" :key="user.id" class="hover:bg-white/[0.01] transition-colors group">
                    <td class="p-6">
                      <div class="flex items-center gap-4">
                        <div class="w-10 h-10 bg-white/[0.03] border border-white/10 flex items-center justify-center font-black text-[10px] group-hover:border-white/40 transition-all">
                          UID
                        </div>
                        <div>
                          <p class="text-xs font-black uppercase text-white">{{ user.name }}</p>
                          <p class="text-[9px] font-bold text-gray-600 mt-0.5">{{ user.email }}</p>
                        </div>
                      </div>
                    </td>
                    <td class="p-6">
                      <div class="flex flex-wrap gap-1">
                        <span v-for="tag in user.vectors" :key="tag" class="text-[8px] font-black px-2 py-0.5 border border-white/5 bg-white/[0.03] text-gray-400">
                          {{ tag }}
                        </span>
                      </div>
                    </td>
                    <td class="p-6">
                      <div class="flex items-center gap-3">
                        <div class="w-24 h-1 bg-white/5 rounded-full overflow-hidden">
                          <div class="h-full" :class="getRiskBg(user.score)" :style="{ width: user.score + '%' }"></div>
                        </div>
                        <span class="text-[10px] font-black italic" :class="getRiskText(user.score)">{{ user.score }}%</span>
                      </div>
                    </td>
                    <td class="p-6">
                      <span class="text-[9px] font-black uppercase flex items-center gap-2">
                        <span class="w-1 h-1 rounded-full animate-pulse" :style="{ backgroundColor: user.active ? themeStore.accentColor : '#444' }"></span>
                        {{ user.active ? 'Monitoring' : 'Isolated' }}
                      </span>
                    </td>
                    <td class="p-6 text-right">
                      <div class="flex justify-end gap-2 opacity-0 group-hover:opacity-100 transition-opacity">
                        <button class="p-2 border border-white/10 hover:bg-white/5 text-[9px] font-black uppercase tracking-tighter">INTERROGATE</button>
                        <button 
                          class="p-2 text-black text-[9px] font-black uppercase tracking-tighter" 
                          :style="{ backgroundColor: themeStore.accentColor }"
                        >
                          ISOLATE
                        </button>
                      </div>
                    </td>
                  </tr>
                </tbody>
              </table>
            </div>
            
            <div class="p-4 border-t border-white/10 bg-white/[0.01] flex justify-between items-center">
              <span class="text-[9px] font-black text-gray-600 uppercase tracking-widest italic">Viewing_50_of_124_Identities</span>
              <div class="flex gap-2">
                <button class="px-3 py-1 border border-white/10 text-[9px] font-black hover:bg-white/5">PREV</button>
                <button class="px-3 py-1 border border-white/10 text-[9px] font-black hover:bg-white/5">NEXT</button>
              </div>
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
const search = ref('')
const filterRisk = ref('ALL')

const flaggedUsers = ref([
  { id: 'U-9901', name: 'Dev_Zero_Cool', email: 'zero@mesh.net', vectors: ['IP_JUMP', 'BULK_APPLY'], score: 88, active: true },
  { id: 'U-4412', name: 'HR_Ghost_Scanner', email: 'ghost@bot.io', vectors: ['SCRAPER_DETECTION'], score: 94, active: false },
  { id: 'U-3321', name: 'Sarah_Connor_29', email: 's.connor@sky.com', vectors: ['VPN_MISMATCH'], score: 42, active: true },
  { id: 'U-1109', name: 'Corporate_Shill_X', email: 'shill@phish.biz', vectors: ['LINK_INJECTION', 'BLACK_LIST_KW'], score: 76, active: true },
  { id: 'U-5561', name: 'Admin_Test_Acc', email: 'test@portal.com', vectors: ['BRUTE_FORCE_ATTEMPT'], score: 65, active: false },
])

const filteredUsers = computed(() => {
  return flaggedUsers.value.filter(u => {
    const matchesSearch = u.name.toLowerCase().includes(search.value.toLowerCase()) || u.email.toLowerCase().includes(search.value.toLowerCase());
    const riskLevel = u.score >= 80 ? 'CRITICAL' : u.score >= 50 ? 'ELEVATED' : 'LOW';
    const matchesRisk = filterRisk.value === 'ALL' || riskLevel === filterRisk.value;
    return matchesSearch && matchesRisk;
  })
})

const getRiskBg = (score: number) => {
  if (score >= 80) return 'bg-red-500';
  if (score >= 50) return 'bg-amber-500';
  return 'bg-cyan-500';
}

const getRiskText = (score: number) => {
  if (score >= 80) return 'text-red-500';
  if (score >= 50) return 'text-amber-500';
  return 'text-cyan-500';
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { height: 4px; width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.05); }

table thead th {
  backdrop-filter: blur(10px);
}
</style>