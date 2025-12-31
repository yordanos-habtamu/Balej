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
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">Capability_Registry_v4</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                SKILL_NODES
              </h1>
            </div>

            <button class="px-8 py-4 text-[10px] font-black uppercase tracking-widest text-black transition-all hover:brightness-110 flex items-center gap-3" :style="{ backgroundColor: themeStore.accentColor }">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M12 4v16m8-8H4" />
              </svg>
              INITIALIZE_NEW_NODE
            </button>
          </div>

          <div class="bg-black border border-white/10 overflow-hidden">
            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-white/10 bg-white/[0.02]">
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Node_ID</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Designation</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Category</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Uplink_Count</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Complexity</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500 text-right">Operations</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="skill in skills" :key="skill.id" class="hover:bg-white/[0.03] transition-colors group">
                  <td class="p-6 font-bold text-xs text-white/40 uppercase tracking-widest">#{{ skill.id }}</td>
                  <td class="p-6">
                    <span class="text-sm font-black uppercase tracking-tight text-white group-hover:text-cyan-400 transition-colors">
                      {{ skill.name }}
                    </span>
                  </td>
                  <td class="p-6">
                    <span class="text-[10px] font-bold text-gray-400 uppercase border border-white/10 px-2 py-1">
                      {{ skill.category }}
                    </span>
                  </td>
                  <td class="p-6">
                    <div class="flex items-center gap-2">
                      <span class="text-sm font-black italic">{{ skill.users }}</span>
                      <span class="text-[10px] text-gray-600 uppercase">Active_Units</span>
                    </div>
                  </td>
                  <td class="p-6">
                    <div class="flex gap-1">
                      <div v-for="i in 5" :key="i" class="w-2 h-4 border border-white/10"
                           :style="{ backgroundColor: i <= skill.complexity ? themeStore.accentColor : 'transparent', opacity: i <= skill.complexity ? '1' : '0.2' }">
                      </div>
                    </div>
                  </td>
                  <td class="p-6 text-right">
                    <div class="flex justify-end gap-4 opacity-40 group-hover:opacity-100 transition-opacity">
                      <button class="hover:text-white text-gray-500">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                        </svg>
                      </button>
                      <button class="hover:text-red-500 text-gray-500">
                        <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <div class="mt-8 grid grid-cols-1 md:grid-cols-3 gap-6">
            <div v-for="stat in summaryStats" :key="stat.label" class="bg-white/[0.02] border border-white/5 p-6">
              <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest mb-2">{{ stat.label }}</p>
              <div class="flex items-end justify-between">
                <span class="text-3xl font-black italic tracking-tighter">{{ stat.value }}</span>
                <span class="text-[10px] font-black uppercase tracking-widest" :style="{ color: themeStore.accentColor }">
                  {{ stat.trend }}
                </span>
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

const skills = ref([
  { id: 'SK-01', name: 'Neural_Rust', category: 'Backend', users: 124, complexity: 5 },
  { id: 'SK-02', name: 'Cyber_Vue_3', category: 'Frontend', users: 89, complexity: 3 },
  { id: 'SK-03', name: 'Solidity_Contract', category: 'Blockchain', users: 42, complexity: 5 },
  { id: 'SK-04', name: 'Tailwind_Grid', category: 'Design', users: 210, complexity: 2 },
  { id: 'SK-05', name: 'Python_AI_Core', category: 'Intelligence', users: 67, complexity: 4 },
])

const summaryStats = ref([
  { label: 'Total_Skill_Nodes', value: '42', trend: '+2_THIS_WEEK' },
  { label: 'Avg_Node_Complexity', value: '3.8', trend: 'STABLE' },
  { label: 'Unlinked_Capabilities', value: '08', trend: 'ACTION_REQ' },
])
</script>