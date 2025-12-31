<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <div :class="['flex-1 flex flex-col h-full relative transition-all duration-500', themeStore.terminalCollapsed ? 'ml-20' : 'ml-72']">
      
      <header class="h-32 flex flex-col justify-center px-12 shrink-0 border-b border-white/5 bg-black/60 relative overflow-hidden z-20">
        <div class="flex items-center justify-between">
          <div>
            <span class="text-[10px] font-black tracking-[0.4em] uppercase opacity-40 text-cyan-400">Marketplace // Inventory_v9</span>
            <h1 class="text-5xl font-black italic uppercase tracking-tighter">JOB_LISTINGS</h1>
          </div>
          <div class="flex gap-8">
            <div class="text-right">
              <p class="text-[9px] font-black text-gray-500 uppercase">Active_Postings</p>
              <p class="text-2xl font-black italic text-white">1,284</p>
            </div>
            <div class="text-right">
              <p class="text-[9px] font-black text-gray-500 uppercase">Avg_AI_Match</p>
              <p class="text-2xl font-black italic text-green-500">84%</p>
            </div>
          </div>
        </div>
      </header>

      <div class="px-12 py-6 bg-white/[0.02] border-b border-white/5 flex gap-4 items-center shrink-0 z-20">
        <div class="relative flex-1 group">
          <span class="absolute left-4 top-1/2 -translate-y-1/2 text-gray-600 group-focus-within:text-cyan-500 transition-colors">🔍</span>
          <input type="text" placeholder="QUERY_BY_SKILLS_ROLE_OR_COMPANY..." class="w-full bg-black border border-white/10 text-[10px] py-3 pl-12 pr-4 uppercase outline-none focus:border-cyan-500 transition-all" />
        </div>
        <select class="bg-black border border-white/10 text-[10px] p-3 uppercase outline-none hover:border-white/30 cursor-pointer">
          <option>Department: Engineering</option>
          <option>Department: Design</option>
          <option>Department: Marketing</option>
        </select>
        <button class="px-8 py-3 bg-cyan-600 text-black text-[10px] font-black uppercase hover:bg-cyan-400 transition-all flex items-center gap-2">
          <span>+</span> Create_New_Posting
        </button>
      </div>

      <main class="flex-1 overflow-y-auto p-12 bg-[#020205] custom-scrollbar z-10">
        <div class="grid grid-cols-1 xl:grid-cols-2 gap-6">
          
          <div v-for="job in jobs" :key="job.id" 
               class="group relative bg-white/[0.02] border border-white/10 p-8 hover:border-cyan-500/50 transition-all duration-300">
            
            <div class="absolute top-6 right-8 text-right">
              <div class="text-[9px] font-black text-gray-500 uppercase mb-1">AI_Sentiment</div>
              <div class="text-xl font-black italic" :class="job.health > 80 ? 'text-green-500' : 'text-amber-500'">{{ job.health }}%</div>
            </div>

            <div class="flex items-start gap-6">
              <div class="w-14 h-14 bg-white/5 border border-white/10 flex items-center justify-center text-xl font-black text-gray-400 group-hover:border-cyan-500/50 group-hover:text-cyan-400 transition-all">
                {{ job.company[0] }}
              </div>

              <div class="flex-1">
                <h3 class="text-xl font-black italic uppercase tracking-tight group-hover:text-cyan-400 transition-colors">{{ job.title }}</h3>
                <div class="flex items-center gap-4 mt-2">
                  <span class="text-[10px] font-black text-gray-400 uppercase tracking-widest">{{ job.company }}</span>
                  <span class="w-1 h-1 rounded-full bg-gray-700"></span>
                  <span class="text-[10px] font-black text-cyan-500 uppercase">{{ job.location }}</span>
                </div>
              </div>
            </div>

            <div class="mt-8 grid grid-cols-3 gap-4 border-y border-white/5 py-6">
              <div>
                <p class="text-[8px] font-black text-gray-600 uppercase mb-1">Salary_Range</p>
                <p class="text-[11px] font-bold text-white tracking-wide">{{ job.salary }}</p>
              </div>
              <div>
                <p class="text-[8px] font-black text-gray-600 uppercase mb-1">Applicants</p>
                <div class="flex items-center gap-2">
                  <p class="text-[11px] font-bold text-white">{{ job.applicants }}</p>
                  <span class="text-[8px] font-black px-1.5 py-0.5 bg-green-500/10 text-green-500">+{{ job.newApplicants }} NEW</span>
                </div>
              </div>
              <div>
                <p class="text-[8px] font-black text-gray-600 uppercase mb-1">Experience</p>
                <p class="text-[11px] font-bold text-white">{{ job.experience }}</p>
              </div>
            </div>

            <div class="mt-6 flex flex-wrap gap-2">
              <span v-for="tag in job.tags" :key="tag" 
                    class="text-[8px] font-black px-3 py-1 bg-white/5 border border-white/5 text-gray-400 uppercase group-hover:border-cyan-500/20 group-hover:text-gray-300">
                #{{ tag }}
              </span>
            </div>

            <div class="mt-8 flex justify-between items-center">
              <div class="flex gap-4">
                <button class="text-[9px] font-black uppercase text-gray-500 hover:text-white transition-colors">Edit_Details</button>
                <button class="text-[9px] font-black uppercase text-gray-500 hover:text-white transition-colors">Promote</button>
              </div>
              <button @click="viewApplications(job)" class="px-6 py-2 bg-white text-black text-[9px] font-black uppercase hover:bg-cyan-500 hover:text-black transition-all">
                Manage_Candidates
              </button>
            </div>

            <div class="absolute bottom-0 right-0 w-8 h-8 opacity-0 group-hover:opacity-100 transition-opacity">
              <div class="absolute bottom-2 right-2 w-4 h-[1px] bg-cyan-500"></div>
              <div class="absolute bottom-2 right-2 w-[1px] h-4 bg-cyan-500"></div>
            </div>
          </div>

        </div>
      </main>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()

const jobs = ref([
  { 
    id: 1, 
    title: 'Sr. Backend Engineer', 
    company: 'NeuralSync Systems', 
    location: 'Remote // Berlin', 
    salary: '€140k - €175k', 
    applicants: 42, 
    newApplicants: 12, 
    experience: '6+ Years',
    health: 94,
    tags: ['RUST', 'DISTRIBUTED_SYSTEMS', 'KAFKA']
  },
  { 
    id: 2, 
    title: 'Lead Product Designer', 
    company: 'Aether_Flow', 
    location: 'Hybrid // Tokyo', 
    salary: '¥12M - ¥18M', 
    applicants: 89, 
    newApplicants: 5, 
    experience: 'Senior',
    health: 78,
    tags: ['FIGMA', 'MOTION_DESIGN', 'UX_RESEARCH']
  },
  { 
    id: 3, 
    title: 'Cybersecurity Analyst', 
    company: 'Void_Security', 
    location: 'On-Site // London', 
    salary: '£90k - £120k', 
    applicants: 15, 
    newApplicants: 2, 
    experience: '4+ Years',
    health: 98,
    tags: ['ETHICAL_HACKING', 'SIEM', 'PEN_TESTING']
  },
  { 
    id: 4, 
    title: 'AI Researcher', 
    company: 'DeepMind_Labs', 
    location: 'Remote // US', 
    salary: '$200k - $250k', 
    applicants: 231, 
    newApplicants: 44, 
    experience: 'PhD Preferred',
    health: 82,
    tags: ['PYTORCH', 'LLM', 'TRANSFORMERS']
  }
])

const viewApplications = (job) => {
  console.log('Routing to candidates for:', job.title)
  navigateTo(`/admin/jobs/${job.id}/candidates`)
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }
</style>