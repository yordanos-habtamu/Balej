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

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-4 py-6 lg:px-10">
        <div class="max-w-6xl mx-auto">
          
          <div class="mb-6 group">
            <NuxtLink to="/admin/clients" class="inline-flex items-center gap-3 text-gray-500 hover:text-white transition-all">
              <span class="text-lg group-hover:-translate-x-1 transition-transform">«</span>
              <span class="text-[9px] tracking-[0.4em] uppercase font-bold">RETURN_TO_REGISTRY</span>
            </NuxtLink>
          </div>

          <div class="bg-black border border-white/10 p-6 lg:p-10 mb-6 relative overflow-hidden group">
            <div class="absolute top-0 right-0 w-32 h-32 bg-white/[0.02] -rotate-45 translate-x-16 -translate-y-16"></div>
            
            <div class="flex flex-col xl:flex-row xl:items-center justify-between gap-6 relative z-10">
              <div>
                <div class="flex items-center gap-3 mb-4">
                  <div class="w-1.5 h-1.5" :style="{ backgroundColor: themeStore.accentColor }"></div>
                  <span class="text-[9px] tracking-[0.5em] text-gray-500 uppercase italic">Active_Mission_Control</span>
                </div>
                <h1 class="text-3xl lg:text-4xl xl:text-5xl font-black italic uppercase tracking-tighter mb-4">
                  {{ project.name.replace(/ /g, '_') }}
                </h1>
                
                <div class="flex flex-wrap items-center gap-6">
                  <div class="flex items-center gap-2 text-[9px] tracking-widest text-gray-400">
                    <span :style="{ color: themeStore.accentColor }">ID_REF:</span>
                    <span>#PROJ-1248</span>
                  </div>
                  <div class="flex items-center gap-2 text-[9px] tracking-widest text-gray-400">
                    <span :style="{ color: themeStore.accentColor }">DEADLINE:</span>
                    <span>{{ project.deadline.toUpperCase() }}</span>
                  </div>
                  <span :class="[
                    'text-[8px] px-3 py-0.5 border font-black tracking-widest uppercase',
                    project.status === 'active' ? 'border-green-500/50 text-green-500' : 'border-blue-500/50 text-blue-500'
                  ]">
                    {{ project.statusLabel }}
                  </span>
                </div>
              </div>

              <div class="flex flex-wrap gap-3">
                <button class="px-6 py-3 border border-white/10 text-[9px] font-black tracking-[0.3em] uppercase hover:bg-white hover:text-black transition-all">
                  [ EDIT_PARAMETERS ]
                </button>
                <button 
                  @click="RouteChat()"
                  class="px-6 py-3 text-[9px] font-black tracking-[0.3em] uppercase text-black transition-all hover:scale-105"
                  :style="{ backgroundColor: themeStore.accentColor }"
                >
                  INITIALIZE_COMMS
                </button>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 xl:grid-cols-3 gap-6 mb-8">
            
            <div class="xl:col-span-2 space-y-6">
              <div class="bg-white/[0.02] border border-white/5 p-6 lg:p-8">
                <div class="flex items-center gap-2 mb-6">
                  <div class="h-[1px] w-8 bg-white/20"></div>
                  <h2 class="text-[9px] font-black tracking-[0.4em] uppercase text-gray-500">Synchronization_Status</h2>
                </div>
                
                <div class="flex flex-col md:flex-row items-end justify-between gap-6 mb-8">
                  <div class="flex items-baseline gap-4">
                    <p class="text-5xl lg:text-6xl font-black italic tracking-tighter" :style="{ color: themeStore.accentColor }">{{ project.progress }}%</p>
                    <p class="text-[9px] text-gray-500 tracking-widest uppercase mb-1">Complete</p>
                  </div>
                  <div class="text-right border-r-2 pr-5" :style="{ borderColor: themeStore.accentColor }">
                    <p class="text-2xl font-black tracking-tighter text-white">${{ project.spent }}</p>
                    <p class="text-[8px] text-gray-600 uppercase tracking-widest">Consumed / ${{ project.budget }}</p>
                  </div>
                </div>

                <div class="w-full bg-white/5 h-1 relative">
                  <div
                    class="h-full transition-all duration-1000 ease-out"
                    :style="{ width: project.progress + '%', backgroundColor: themeStore.accentColor }"
                  ></div>
                </div>
              </div>

              <div class="bg-white/[0.02] border border-white/5 p-6 lg:p-8">
                <div class="flex items-center gap-2 mb-4">
                  <div class="h-[1px] w-8 bg-white/20"></div>
                  <h2 class="text-[9px] font-black tracking-[0.4em] uppercase text-gray-500">Mission_Directives</h2>
                </div>
                <p class="text-[11px] text-gray-400 leading-relaxed uppercase tracking-widest">
                  {{ project.description }}
                </p>
              </div>
            </div>

            <div class="space-y-6">
              <div class="bg-black border border-white/10 p-6">
                <h2 class="text-[9px] font-black tracking-[0.4em] uppercase text-white mb-6">CORE_METRICS</h2>
                <div class="space-y-4">
                  <div v-for="(val, label) in { 'Client': 'TechCorp', 'Uplink': 'Oct 1, 2025', 'Load': project.talents + ' Units' }" :key="label" class="flex justify-between border-b border-white/5 pb-2">
                    <span class="text-[8px] text-gray-600 uppercase tracking-widest">{{ label }}</span>
                    <span class="text-[9px] font-bold text-white uppercase">{{ val }}</span>
                  </div>
                </div>
              </div>

              <div class="bg-black border border-white/10 p-6 relative overflow-hidden">
                <h2 class="text-[9px] font-black tracking-[0.4em] uppercase text-white mb-6">EVENT_LOG</h2>
                <div class="space-y-6 relative">
                  <div class="absolute left-[7px] top-2 bottom-2 w-[1px] bg-white/10"></div>
                  
                  <div class="flex items-center gap-4 relative">
                    <div class="w-2.5 h-2.5 rounded-full bg-green-500 z-10 border-2 border-black"></div>
                    <div>
                      <p class="text-[9px] font-bold text-white uppercase">Kickoff_Meeting</p>
                      <p class="text-[7px] text-gray-600 tracking-widest uppercase">Oct 05</p>
                    </div>
                  </div>
                  <div class="flex items-center gap-4 relative">
                    <div class="w-2.5 h-2.5 rounded-full z-10 border-2 border-black" :style="{ backgroundColor: themeStore.accentColor }"></div>
                    <div>
                      <p class="text-[9px] font-bold text-white uppercase">Design_Finalized</p>
                      <p class="text-[7px] text-gray-600 tracking-widest uppercase">Nov 20</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>

          <div class="bg-white/[0.02] border border-white/5 p-6 lg:p-8 mb-10">
            <div class="flex items-center justify-between mb-8">
              <div class="flex items-center gap-2">
                <div class="h-[1px] w-8 bg-white/20"></div>
                <h2 class="text-[9px] font-black tracking-[0.4em] uppercase text-white">Assigned_Talent_Nodes</h2>
              </div>
              <span class="text-[8px] text-gray-600 uppercase tracking-widest">Active_Units: {{ project.assignedTalents.length }}</span>
            </div>

            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
              <div v-for="talent in project.assignedTalents" :key="talent.id" class="group flex items-center gap-4 bg-black border border-white/5 p-4 hover:border-white/20 transition-all">
                <div class="w-10 h-10 border border-white/10 p-0.5 grayscale group-hover:grayscale-0 transition-all duration-300">
                  <img src="~/assets/painter.png" alt="Talent" class="w-full h-full object-cover" />
                </div>
                <div>
                  <h4 class="text-[10px] font-black text-white uppercase tracking-tighter">{{ talent.name }}</h4>
                  <p class="text-[7px] text-gray-500 uppercase tracking-widest">{{ talent.role }}</p>
                  <div class="flex items-center gap-2 mt-2">
                    <div class="h-0.5 w-10 bg-white/5 overflow-hidden">
                      <div class="h-full bg-white/20" :style="{ width: (talent.hours / 40 * 100) + '%' }"></div>
                    </div>
                    <span class="text-[6px] text-gray-600 uppercase">{{ talent.hours }}H</span>
                  </div>
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
import { useRouter } from 'vue-router'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const router = useRouter()

const RouteChat = () => {
  router.push('/admin/chats')
}

const project = ref({
  name: 'Enterprise Dashboard Redesign',
  description: 'Complete overhaul of the internal analytics platform with modern UI/UX, real-time data visualization, and improved performance. The new dashboard will support multiple user roles and advanced filtering capabilities.',
  status: 'active',
  statusLabel: 'Active_Node',
  budget: '45,000',
  spent: '30,600',
  progress: 68,
  talents: 5,
  deadline: 'Feb 15, 2026',
  assignedTalents: [
    { id: 1, name: 'Maria Gonzalez', role: 'Senior UI Designer', hours: 32 },
    { id: 2, name: 'James Wilson', role: 'Full-Stack Developer', hours: 40 },
    { id: 3, name: 'Aisha Khan', role: 'Frontend Engineer', hours: 28 },
    { id: 4, name: 'Carlos Rodriguez', role: 'Backend Developer', hours: 35 },
    { id: 5, name: 'Sophie Dubois', role: 'Project Manager', hours: 20 }
  ]
})
</script>

<style scoped>
@reference 'tailwindcss';

/* Slimmer scrollbar for a cleaner fit */
.custom-scrollbar::-webkit-scrollbar { width: 3px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #1f2937; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: v-bind('themeStore.accentColor'); }

/* Refined animations with less travel distance */
.bg-white\/\[0\.02\], .bg-black {
  animation: nodeEnter 0.5s cubic-bezier(0.2, 1, 0.3, 1) forwards;
  opacity: 0;
}

@keyframes nodeEnter {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}

.bg-white\/\[0\.02\]:nth-child(1) { animation-delay: 0.1s; }
.bg-white\/\[0\.02\]:nth-child(2) { animation-delay: 0.15s; }
</style>