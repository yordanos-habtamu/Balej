<template>
  <div class="min-h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col min-h-screen transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)]',
        themeStore.terminalCollapsed ? 'ml-20' : 'ml-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-10" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 p-6 lg:p-12">
        <div class="max-w-7xl mx-auto">
          
          <div class="mb-12 border-b border-white/5 pb-10 flex flex-col md:flex-row md:items-end justify-between gap-6">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="w-2 h-2 animate-pulse" :style="{ backgroundColor: themeStore.accentColor }"></div>
                <span class="text-[10px] tracking-[0.5em] text-gray-500 uppercase italic">Project_Archive // Client_View</span>
              </div>
              <h1 class="text-5xl font-black italic uppercase tracking-tighter">
                CLIENT_<span :style="{ color: themeStore.accentColor }">PROJECTS</span>
              </h1>
            </div>
            <p class="text-[10px] text-gray-500 tracking-[0.4em] uppercase">Status: Connected // Archive_ID: TC-882</p>
          </div>

          <div class="bg-black border border-white/10 p-8 mb-12 relative overflow-hidden shadow-2xl group">
             <div class="absolute top-0 left-0 w-full h-[1px] bg-gradient-to-r from-transparent via-white/20 to-transparent group-hover:via-white/50 transition-all duration-700"></div>
             
             <div class="flex flex-col lg:flex-row items-center gap-10 relative z-10">
                <div class="relative">
                  <div class="w-32 h-32 border border-white/10 p-2 grayscale group-hover:grayscale-0 transition-all duration-700">
                    <img src="~/assets/painter.png" alt="Client" class="w-full h-full object-cover shadow-2xl" />
                  </div>
                  <div class="absolute -bottom-2 -right-2 w-6 h-6 border-4 border-black" :style="{ backgroundColor: themeStore.accentColor }"></div>
                </div>

                <div class="flex-1 text-center lg:text-left">
                  <h2 class="text-4xl font-black italic uppercase tracking-tighter mb-2">TechCorp Solutions</h2>
                  <p class="text-[10px] text-gray-500 tracking-[0.3em] uppercase mb-8 italic">>> Premium_Tier_Client // Verified_Since_2024</p>
                  
                  <div class="grid grid-cols-1 sm:grid-cols-3 gap-8">
                    <div class="border-l border-white/10 pl-4">
                      <p class="text-3xl font-black tracking-tighter" :style="{ color: themeStore.accentColor }">08</p>
                      <p class="text-[9px] text-gray-600 uppercase tracking-widest">Active_Nodes</p>
                    </div>
                    <div class="border-l border-white/10 pl-4">
                      <p class="text-3xl font-black tracking-tighter text-green-500">24</p>
                      <p class="text-[9px] text-gray-600 uppercase tracking-widest">Successful_Link</p>
                    </div>
                    <div class="border-l border-white/10 pl-4">
                      <p class="text-3xl font-black tracking-tighter text-white">$184K</p>
                      <p class="text-[9px] text-gray-600 uppercase tracking-widest">Resource_Value</p>
                    </div>
                  </div>
                </div>
             </div>
          </div>

          <div class="flex flex-col lg:flex-row gap-6 mb-12">
            <div class="flex-1 relative group">
              <input
                type="text"
                v-model="searchTerm"
                placeholder="EXECUTE_PROJECT_SEARCH..."
                class="w-full bg-white/[0.03] border border-white/10 px-6 py-4 text-xs font-mono focus:outline-none transition-all uppercase placeholder:text-gray-800"
                :style="{ borderLeft: `4px solid ${themeStore.accentColor}` }"
              />
            </div>
            <select 
              v-model="statusFilter" 
              class="bg-black border border-white/10 px-8 py-4 text-[10px] uppercase tracking-widest font-bold focus:outline-none focus:border-white/40 appearance-none cursor-pointer"
            >
              <option value="all">ALL_STATUS</option>
              <option value="active">ACTIVE_NODES</option>
              <option value="completed">ARCHIVED_LINKS</option>
              <option value="on-hold">ON_STANDBY</option>
            </select>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-6">
            <div
              v-for="project in filteredProjects"
              :key="project.id"
              class="group relative bg-white/[0.02] border border-white/5 flex flex-col transition-all duration-500 hover:bg-white/[0.05] hover:border-white/20"
            >
              <div class="p-6 border-b border-white/5">
                <div class="flex items-start justify-between mb-4">
                  <h3 class="text-lg font-black italic uppercase tracking-tighter leading-tight group-hover:text-cyan-400 transition-colors">
                    {{ project.name }}
                  </h3>
                  <span :class="[
                    'text-[8px] px-2 py-0.5 border font-bold tracking-widest uppercase',
                    project.status === 'active' ? 'border-green-500/50 text-green-500' :
                    project.status === 'completed' ? 'border-blue-500/50 text-blue-500' :
                    'border-amber-500/50 text-amber-500'
                  ]">
                    {{ project.statusLabel }}
                  </span>
                </div>
                <p class="text-[10px] text-gray-500 leading-relaxed uppercase">{{ project.description }}</p>
              </div>

              <div class="p-6 flex-1 space-y-4">
                <div class="flex justify-between items-center text-[9px] tracking-widest">
                  <span class="text-gray-600 uppercase">>> RESOURCE_BUDGET</span>
                  <span class="font-bold text-white">{{ project.budget }}</span>
                </div>
                <div class="flex justify-between items-center text-[9px] tracking-widest">
                  <span class="text-gray-600 uppercase">>> UPLINK_DEADLINE</span>
                  <span class="font-bold text-white italic">{{ project.deadline }}</span>
                </div>
                <div class="flex justify-between items-center text-[9px] tracking-widest">
                  <span class="text-gray-600 uppercase">>> TALENT_LOAD</span>
                  <span class="font-bold text-white">{{ project.talents }} UNITS</span>
                </div>
              </div>

              <div class="px-6 pb-4">
                <div class="flex justify-between items-center mb-2">
                  <span class="text-[8px] text-gray-600 uppercase tracking-[0.2em]">Synchronization</span>
                  <span class="text-[10px] font-black" :style="{ color: themeStore.accentColor }">{{ project.progress }}%</span>
                </div>
                <div class="w-full bg-white/5 h-[2px] relative overflow-hidden">
                  <div
                    class="h-full transition-all duration-1000 ease-out"
                    :style="{ 
                      width: project.progress + '%', 
                      backgroundColor: project.status === 'active' ? themeStore.accentColor : '#22c55e',
                      boxShadow: `0 0 10px ${themeStore.accentColor}`
                    }"
                  ></div>
                </div>
              </div>

              <div class="p-6 pt-2">
                <nuxt-link :to="`/admin/clients/projects/${project.name}`" class="block w-full">
                  <button class="w-full py-3 border border-white/10 text-[10px] font-black tracking-[0.4em] uppercase hover:bg-white hover:text-black transition-all">
                    VIEW_DETAILS_LOG
                  </button>
                </nuxt-link>
              </div>

              <div class="absolute bottom-0 left-0 h-[1px] w-0 group-hover:w-full transition-all duration-700" :style="{ backgroundColor: themeStore.accentColor }"></div>
            </div>
          </div>

          <div v-if="filteredProjects.length === 0" class="text-center py-32 border border-dashed border-white/10">
            <p class="text-xs text-gray-600 tracking-[0.5em] uppercase">No_Project_Nodes_Found_In_Current_Filter</p>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const searchTerm = ref('')
const statusFilter = ref('all')

// Mock project data
const projects = ref([
  { id: 1, name: 'Enterprise Dashboard Redesign', description: 'Complete overhaul of internal analytics platform', status: 'active', statusLabel: 'Active', budget: '$45,000', deadline: 'Feb 15, 2026', talents: 5, progress: 68 },
  { id: 2, name: 'Mobile App Development', description: 'Native iOS & Android apps for customer portal', status: 'active', statusLabel: 'Active', budget: '$72,000', deadline: 'Mar 30, 2026', talents: 8, progress: 42 },
  { id: 3, name: 'Cloud Migration Phase 2', description: 'Migrating legacy systems to AWS', status: 'active', statusLabel: 'Active', budget: '$38,000', deadline: 'Jan 20, 2026', talents: 4, progress: 85 },
  { id: 4, name: 'API Integration Suite', description: 'Connecting third-party services', status: 'completed', statusLabel: 'Completed', budget: '$28,000', deadline: 'Nov 10, 2025', talents: 3, progress: 100 },
  { id: 5, name: 'Security Audit & Compliance', description: 'GDPR and SOC 2 compliance preparation', status: 'on-hold', statusLabel: 'On Hold', budget: '$15,000', deadline: 'Pending', talents: 2, progress: 30 },
  { id: 6, name: 'Marketing Website Revamp', description: 'New corporate website with CMS', status: 'active', statusLabel: 'Active', budget: '$22,000', deadline: 'Feb 28, 2026', talents: 4, progress: 55 }
])

const filteredProjects = computed(() => {
  let filtered = projects.value
  if (searchTerm.value) {
    const term = searchTerm.value.toLowerCase()
    filtered = filtered.filter(p => 
      p.name.toLowerCase().includes(term) ||
      p.description.toLowerCase().includes(term)
    )
  }
  if (statusFilter.value !== 'all') {
    filtered = filtered.filter(p => p.status === statusFilter.value)
  }
  return filtered
})
</script>

<style scoped>
@reference 'tailwindcss';

.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #1f2937;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: v-bind('themeStore.accentColor');
}

/* Entry Animation */
.group {
  animation: nodeFadeIn 0.5s ease-out forwards;
}

@keyframes nodeFadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

/* Staggered load for cards */
.group:nth-child(2) { animation-delay: 0.1s; }
.group:nth-child(3) { animation-delay: 0.2s; }
</style>