<template>
  <div class="h-screen w-full bg-[#020205] text-gray-100 font-sans flex overflow-hidden selection:bg-cyan-500/30">
    
    <admin-sidebar />

    <div 
      :class="[
        'flex-1 flex flex-col min-h-screen transition-all duration-500 ease-in-out min-w-0',
        themeStore.terminalCollapsed ? 'ml-20' : 'ml-72'
      ]"
    >
      <header class="sticky top-0 z-40 border-b border-white/10 bg-[#050508]/90 backdrop-blur-xl px-8 py-4 flex items-center justify-between">
        <div class="flex items-center gap-6 flex-1">
          <div class="hidden sm:block">
            <h1 class="text-lg font-black tracking-tighter text-white uppercase italic">Command_<span :style="{ color: themeStore.accentColor }">Center</span></h1>
            <div class="flex items-center gap-2 mt-0.5">
              <div class="w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse"></div>
              <span class="text-[9px] font-mono text-gray-500 uppercase tracking-widest">System_Operational</span>
            </div>
          </div>
          
          <div class="relative max-w-md w-full group ml-4">
            <input 
              type="text" 
              v-model="searchTerm" 
              placeholder="SCAN_FOR_NODES..." 
              class="w-full bg-white/[0.03] border border-white/10 py-2.5 pl-10 pr-4 text-[10px] font-mono focus:border-cyan-500 focus:bg-white/[0.05] outline-none transition-all rounded-sm uppercase italic"
            >
            <span class="absolute left-3 top-1/2 -translate-y-1/2 text-[10px]" :style="{ color: themeStore.accentColor }">>></span>
          </div>
        </div>

        <div class="flex items-center gap-4">
          <button @click="themeStore.toggleHUD()" class="px-3 py-2 border border-white/10 text-[9px] font-mono uppercase tracking-tighter hover:bg-white hover:text-black transition-all">
            HUD_{{ themeStore.hudVisibility ? 'ACTIVE' : 'STBY' }}
          </button>
          <div class="h-9 w-9 bg-white/5 border border-white/10 rounded-sm flex items-center justify-center cursor-pointer hover:border-cyan-500 transition-colors">
            <UserIcon class="h-5 w-5 text-gray-400" />
          </div>
        </div>
      </header>

      <main class="flex-1 overflow-y-auto p-6 lg:p-10 custom-scrollbar">
        <div class="max-w-[1400px] mx-auto space-y-10">
          
          <div class="grid grid-cols-1 sm:grid-cols-2 xl:grid-cols-4 gap-4">
            <div v-for="(stat, i) in stats" :key="i" class="bg-black border border-white/5 p-6 relative group overflow-hidden transition-all hover:bg-white/[0.01]">
              <div class="absolute top-0 left-0 w-[2px] h-0 group-hover:h-full transition-all duration-500" :style="{ backgroundColor: themeStore.accentColor }"></div>
              <div class="flex justify-between items-start mb-4">
                <span class="text-[9px] font-mono text-gray-500 uppercase tracking-[0.3em]">{{ stat.label }}</span>
                <component :is="stat.icon" class="w-4 h-4 opacity-30" :style="{ color: themeStore.accentColor }" />
              </div>
              <div class="flex items-end gap-3">
                <p class="text-4xl font-black text-white italic tracking-tighter">{{ stat.value }}</p>
                <span class="text-[10px] text-green-500 font-bold mb-1.5">+{{ stat.trend }}%</span>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 xl:grid-cols-3 gap-6">
            <div class="xl:col-span-2 bg-[#050508] border border-white/10 p-8 relative overflow-hidden chart-scanlines">
              <div class="flex justify-between items-end mb-10 border-l-2 pl-4" :style="{ borderColor: themeStore.accentColor }">
                <div>
                  <h3 class="text-xs font-black uppercase tracking-[0.4em] text-white italic">Neural_Pipeline_Load</h3>
                  <p class="text-[9px] font-mono text-gray-600 mt-1 uppercase">Candidate throughput analysis</p>
                </div>
                <div class="text-right hidden md:block">
                  <p class="text-[8px] font-mono text-gray-700 uppercase">Load_Index</p>
                  <p class="text-xl font-black italic text-white leading-none">0.892</p>
                </div>
              </div>
              <div class="h-[320px] relative">
                <Bar :data="barData" :options="chartOptions" />
              </div>
            </div>

            <div class="bg-[#050508] border border-white/10 p-8 flex flex-col relative overflow-hidden">
              <div class="mb-8 border-b border-white/5 pb-4">
                <h3 class="text-xs font-black uppercase tracking-[0.4em] text-white italic">Entity_Cluster</h3>
                <p class="text-[8px] font-mono text-gray-600 uppercase">Global_Distribution</p>
              </div>
              <div class="flex-1 flex items-center justify-center min-h-[250px]">
                <Pie :data="pieData" :options="pieOptions" />
              </div>
              <div class="mt-8 grid grid-cols-2 gap-y-3 gap-x-4 border-t border-white/5 pt-6">
                <div v-for="(label, i) in pieData.labels" :key="label" class="flex items-center gap-2">
                  <div class="w-2 h-2 rounded-full" :style="{ backgroundColor: pieData.datasets[0].backgroundColor[i] }"></div>
                  <span class="text-[9px] font-mono text-gray-500 uppercase truncate">{{ label }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="space-y-6">
            <div class="flex items-center gap-4">
              <h2 class="text-xs font-black tracking-[0.5em] uppercase text-white italic">Active_Postings_Stream</h2>
              <div class="h-px flex-1 bg-gradient-to-r from-white/10 to-transparent"></div>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
              <div v-for="job in jobs" :key="job.title" class="bg-black border border-white/5 p-6 group hover:border-white/20 transition-all">
                <div class="flex justify-between items-start mb-6">
                  <span class="text-[9px] font-mono text-gray-600 tracking-[0.2em] uppercase">Ref_{{ Math.random().toString(36).substr(2, 4).toUpperCase() }}</span>
                  <div class="flex items-center gap-2">
                    <span class="text-[8px] font-bold text-gray-500 uppercase">{{ job.job_type }}</span>
                    <div class="w-1.5 h-1.5 rounded-full" :style="{ backgroundColor: themeStore.accentColor, boxShadow: `0 0 8px ${themeStore.accentColor}` }"></div>
                  </div>
                </div>
                <h4 class="text-sm font-black uppercase italic tracking-wider text-gray-300 group-hover:text-white transition-colors mb-2">{{ job.title }}</h4>
                <p class="text-[10px] text-gray-600 uppercase mb-8">{{ job.category }}</p>
                <div class="flex justify-between items-center pt-4 border-t border-white/5">
                  <div class="flex flex-col">
                    <span class="text-lg font-black italic leading-none">{{ job.appCount }}</span>
                    <span class="text-[7px] text-gray-700 uppercase font-mono">Applicants</span>
                  </div>
                  <button class="text-[9px] font-black uppercase tracking-widest px-4 py-2 border border-white/10 group-hover:bg-white group-hover:text-black transition-all">
                    Process_Node
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, reactive } from 'vue'
import { useThemeStore } from '~/stores/theme'
import { BriefcaseIcon, UsersIcon, ArrowTrendingUpIcon, UserIcon } from '@heroicons/vue/24/outline'

// Chart.js Core
import { 
  Chart as ChartJS, Title, Tooltip, Legend, 
  BarElement, ArcElement, CategoryScale, LinearScale 
} from 'chart.js'
import { Bar, Pie } from 'vue-chartjs'

ChartJS.register(Title, Tooltip, Legend, BarElement, ArcElement, CategoryScale, LinearScale)

const themeStore = useThemeStore()
const searchTerm = ref('')

// SYSTEM STATS
const stats = [
  { label: 'Active_Jobs', value: '412', icon: BriefcaseIcon, trend: '8.1' },
  { label: 'New_Talent', value: '1.8K', icon: UsersIcon, trend: '12.4' },
  { label: 'Avg_Fill_Rate', value: '74%', icon: ArrowTrendingUpIcon, trend: '2.5' },
  { label: 'System_Load', value: '14%', icon: UserIcon, trend: '-1.2' },
]

// MOCK DATA
const jobs = [
  { title: 'Neural Architect', category: 'Technology', job_type: 'Full-time', appCount: 84 },
  { title: 'Senior Security Analyst', category: 'Security', job_type: 'Full-time', appCount: 156 },
  { title: 'Blockchain Engineer', category: 'Technology', job_type: 'Remote', appCount: 92 },
  { title: 'UX Visionary', category: 'Design', job_type: 'Contract', appCount: 31 },
  { title: 'Ops Commander', category: 'Management', job_type: 'Hybrid', appCount: 45 },
  { title: 'Protocol Auditor', category: 'Legal', job_type: 'Remote', appCount: 12 },
]

// COOL BAR DATA (Vertical Gradients)
const barData = computed(() => {
  const accent = themeStore.accentColor || '#00f3ff'
  return {
    labels: ['DISCOVERY', 'APPLIED', 'SCREENING', 'INTERVIEW', 'OFFER', 'HIRED'],
    datasets: [{
      label: 'Flow_Volume',
      data: [95, 80, 45, 30, 18, 12],
      backgroundColor: (context: any) => {
        const {ctx, chartArea} = context.chart
        if (!chartArea) return null
        const gradient = ctx.createLinearGradient(0, chartArea.bottom, 0, chartArea.top)
        gradient.addColorStop(0, 'rgba(0,0,0,0)')
        gradient.addColorStop(1, accent)
        return gradient
      },
      borderColor: accent,
      borderWidth: { top: 2, left: 0, right: 0, bottom: 0 },
      borderRadius: 1,
      barPercentage: 0.5
    }]
  }
})

// HIGH CONTRAST PIE DATA
const pieData = computed(() => {
  const primary = themeStore.accentColor || '#00f3ff'
  return {
    labels: ['TECH_NODES', 'CREATIVE', 'OPERATIONS', 'LEGAL'],
    datasets: [{
      data: [45, 25, 20, 10],
      backgroundColor: [primary, '#a855f7', '#22c55e', '#1e293b'],
      borderColor: '#050508',
      borderWidth: 4,
      spacing: 5,
      hoverOffset: 15
    }]
  }
})

// CHART OPTIONS
const chartOptions = computed(() => ({
  responsive: true,
  maintainAspectRatio: false,
  plugins: {
    legend: { display: false },
    tooltip: {
      backgroundColor: '#000',
      titleFont: { family: 'monospace', size: 10 },
      bodyFont: { family: 'monospace', size: 12 },
      borderColor: themeStore.accentColor,
      borderWidth: 1,
      padding: 12,
      displayColors: false,
      cornerRadius: 0
    }
  },
  scales: {
    x: { grid: { display: false }, ticks: { color: '#444', font: { family: 'monospace', size: 8 } } },
    y: { 
      grid: { color: 'rgba(255,255,255,0.02)', borderDash: [5, 5] }, 
      ticks: { color: '#333', font: { family: 'monospace', size: 8 } } 
    }
  }
}))

const pieOptions = {
  responsive: true,
  maintainAspectRatio: false,
  plugins: { legend: { display: false } },
  cutout: '15%'
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 3px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.05); }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: v-bind('themeStore.accentColor'); }

/* Subtle Scanline Effect for Charts */
.chart-scanlines::after {
  content: "";
  position: absolute;
  top: 0; left: 0; right: 0; bottom: 0;
  background: linear-gradient(rgba(18, 16, 16, 0) 50%, rgba(0, 0, 0, 0.1) 50%);
  background-size: 100% 4px;
  pointer-events: none;
  opacity: 0.2;
}

input::placeholder { color: #333; }
</style>