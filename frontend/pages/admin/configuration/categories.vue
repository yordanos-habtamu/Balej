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

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-6 py-8 lg:px-12">
        <div class="max-w-6xl mx-auto">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-12">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12 bg-white/20"></div>
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">Sector_Hierarchy_v1.2</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                CATEGORY_MATRIX
              </h1>
            </div>

            <button 
              @click="openCreateModal"
              class="px-8 py-4 text-[10px] font-black uppercase tracking-widest text-black transition-all hover:brightness-110 flex items-center gap-3 active:scale-95" 
              :style="{ backgroundColor: themeStore.accentColor }"
            >
              <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M12 4v16m8-8H4" />
              </svg>
              REGISTER_NEW_SECTOR
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-10">
            <div v-for="stat in matrixStats" :key="stat.label" class="bg-white/[0.02] border border-white/5 p-6 relative overflow-hidden group">
              <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest mb-2">{{ stat.label }}</p>
              <div class="flex items-end justify-between">
                <span class="text-3xl font-black italic tracking-tighter">{{ stat.value }}</span>
                <span class="text-[10px] font-black uppercase tracking-widest" :style="{ color: themeStore.accentColor }">
                  {{ stat.subtext }}
                </span>
              </div>
            </div>
          </div>

          <div class="bg-black border border-white/10 overflow-hidden shadow-2xl mb-12">
            <div class="p-6 border-b border-white/10 bg-white/[0.02] flex justify-between items-center">
              <input 
                v-model="searchQuery"
                type="text" 
                placeholder="SEARCH_SECTOR_ID..." 
                class="bg-transparent border border-white/10 px-4 py-2 text-[10px] uppercase font-bold tracking-widest focus:outline-none focus:border-cyan-500/50 w-64"
              />
              <div class="text-[9px] font-black text-gray-500 uppercase tracking-widest">
                {{ filteredCategories.length }} Active_Sectors
              </div>
            </div>

            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-white/10 bg-white/[0.01]">
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Designation</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Live_Jobs</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Sub_Nodes</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500 text-right">Operations</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="cat in filteredCategories" :key="cat.id" class="hover:bg-white/[0.03] transition-colors group">
                  <td class="p-6">
                    <div class="flex items-center gap-3">
                      <div class="w-1 h-4 transition-all group-hover:h-6" :style="{ backgroundColor: themeStore.accentColor }"></div>
                      <span class="text-sm font-black uppercase tracking-tight text-white group-hover:text-cyan-400 transition-colors">
                        {{ cat.name }}
                      </span>
                    </div>
                  </td>
                  <td class="p-6 text-sm font-black italic">{{ cat.jobCount }}</td>
                  <td class="p-6">
                    <div class="flex flex-wrap gap-2">
                      <span v-for="sub in cat.subcategories.slice(0, 2)" :key="sub" class="text-[9px] font-bold text-white/40 border border-white/10 px-2 py-0.5 uppercase">
                        {{ sub }}
                      </span>
                    </div>
                  </td>
                  <td class="p-6 text-right">
                    <div class="flex justify-end gap-3 opacity-20 group-hover:opacity-100 transition-opacity">
                      <button @click="openEditModal(cat)" class="w-8 h-8 flex items-center justify-center border border-white/10 hover:border-white/40">
                        <svg class="w-3 h-3 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" stroke-width="2" /></svg>
                      </button>
                      <button @click="confirmDelete(cat)" class="w-8 h-8 flex items-center justify-center border border-white/10 hover:border-red-500 group/del">
                        <svg class="w-3 h-3 text-red-600/50 group-hover/del:text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" stroke-width="2" /></svg>
                      </button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>
    </main>

    <transition name="modal-fade">
      <div v-if="isModalOpen" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-black/80 backdrop-blur-sm">
        <div class="w-full max-w-2xl bg-[#050508] border border-white/10 shadow-[0_0_50px_rgba(0,0,0,1)] relative">
          <div class="absolute top-0 left-0 h-[2px] bg-cyan-500" :style="{ width: '100%', backgroundColor: themeStore.accentColor }"></div>
          <div class="p-10">
            <h2 class="text-3xl font-black italic uppercase tracking-tighter mb-8">
              {{ editingId ? 'UPDATE_SECTOR_DATA' : 'INIT_NEW_SECTOR' }}
            </h2>
            <div class="space-y-6">
              <div class="grid grid-cols-2 gap-6">
                <div class="space-y-2">
                  <label class="text-[9px] font-black uppercase text-gray-500 tracking-widest">Designation</label>
                  <input v-model="sectorForm.name" type="text" placeholder="E.G. ENGINEERING" class="hud-input uppercase">
                </div>
                <div class="space-y-2">
                  <label class="text-[9px] font-black uppercase text-gray-500 tracking-widest">Node_Slug</label>
                  <input v-model="sectorForm.slug" type="text" placeholder="eng-node" class="hud-input lowercase">
                </div>
              </div>
              <div class="space-y-2">
                <label class="text-[9px] font-black uppercase text-gray-500 tracking-widest">Child_Capabilities</label>
                <textarea v-model="sectorForm.subs" rows="3" placeholder="UI, UX, BACKEND" class="hud-input resize-none uppercase"></textarea>
              </div>
            </div>
            <div class="mt-10 flex gap-4">
              <button @click="saveSector" class="flex-1 py-4 text-[10px] font-black uppercase tracking-[0.2em] text-black" :style="{ backgroundColor: themeStore.accentColor }">
                {{ editingId ? 'SYNC_CHANGES' : 'EXECUTE_REGISTRATION' }}
              </button>
              <button @click="isModalOpen = false" class="px-8 py-4 text-[10px] font-black uppercase tracking-[0.2em] border border-white/10">ABORT_SEQUENCE</button>
            </div>
          </div>
        </div>
      </div>
    </transition>

    <transition name="modal-fade">
      <div v-if="deleteConfirm.show" class="fixed inset-0 z-[200] flex items-center justify-center p-6 bg-red-950/20 backdrop-blur-md">
        <div class="w-full max-w-md bg-[#050508] border border-red-500/30 shadow-[0_0_50px_rgba(239,68,68,0.1)] relative">
          <div class="p-8 text-center">
            <div class="w-16 h-16 border-2 border-red-500 mx-auto mb-6 flex items-center justify-center animate-pulse">
              <svg class="w-8 h-8 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
              </svg>
            </div>
            <h3 class="text-xl font-black italic uppercase tracking-tighter text-red-500 mb-2">Node_Purge_Required</h3>
            <p class="text-[10px] font-bold text-gray-500 uppercase tracking-widest leading-relaxed mb-8">
              Are you sure you want to de-allocate <span class="text-white">"{{ deleteConfirm.targetName }}"</span>? This will permanently sever all child uplinks.
            </p>
            <div class="flex flex-col gap-3">
              <button @click="executeDelete" class="w-full py-4 bg-red-600 text-black text-[10px] font-black uppercase tracking-widest hover:bg-red-500 transition-colors">
                CONFIRM_PURGE
              </button>
              <button @click="deleteConfirm.show = false" class="w-full py-4 border border-white/10 text-[10px] font-black uppercase tracking-widest hover:bg-white/5 transition-colors">
                CANCEL_SAFE
              </button>
            </div>
          </div>
        </div>
      </div>
    </transition>

    <transition name="toast">
      <div v-if="toast.show" class="fixed bottom-8 right-8 z-[250] bg-black border-l-4 p-4 min-w-[280px] shadow-2xl" :style="{ borderColor: themeStore.accentColor }">
        <div class="flex items-start gap-4">
          <div class="mt-1 w-2 h-2 rounded-full animate-ping" :style="{ backgroundColor: themeStore.accentColor }"></div>
          <div>
            <p class="text-[10px] font-black uppercase tracking-widest text-gray-500">System_Event</p>
            <p class="text-xs font-bold mt-1 tracking-tight">{{ toast.message }}</p>
          </div>
        </div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()

// --- UI STATE ---
const isModalOpen = ref(false)
const editingId = ref<number | null>(null)
const searchQuery = ref('')
const toast = ref({ show: false, message: '' })

// Delete Modal State
const deleteConfirm = ref({
  show: false,
  targetId: null as number | null,
  targetName: ''
})

const sectorForm = ref({
  name: '',
  slug: '',
  subs: '',
})

// --- DATA ---
const matrixStats = ref([
  { label: 'Active_Sectors', value: '14', subtext: 'OPERATIONAL' },
  { label: 'Job_Volume', value: '1,284', subtext: '+12%_SHIFT' },
  { label: 'Network_Density', value: '08', subtext: 'OPTIMAL' },
])

const categories = ref([
  { id: 1, name: 'Neural Engineering', slug: 'neural-eng', jobCount: 423, traffic: 85, subcategories: ['Bionic_Interfaces', 'Wetware', 'Sync_Core'] },
  { id: 2, name: 'Cyber Defense', slug: 'cyber-sec', jobCount: 156, traffic: 62, subcategories: ['Ghost_Wall', 'Encryption'] },
  { id: 3, name: 'Data Visualization', slug: 'data-viz', jobCount: 89, traffic: 40, subcategories: ['HUD_Design', 'AR_Optics'] },
  { id: 4, name: 'Blockchain Logistics', slug: 'ledger-log', jobCount: 212, traffic: 78, subcategories: ['Smart_Supply', 'Minting'] },
])

// --- COMPUTED ---
const filteredCategories = computed(() => {
  if (!searchQuery.value) return categories.value
  const query = searchQuery.value.toLowerCase()
  return categories.value.filter(cat => 
    cat.name.toLowerCase().includes(query) || cat.slug.toLowerCase().includes(query)
  )
})

// --- ACTIONS ---
const showToast = (msg: string) => {
  toast.value = { show: true, message: msg }
  setTimeout(() => { toast.value.show = false }, 4000)
}

const openCreateModal = () => {
  editingId.value = null
  sectorForm.value = { name: '', slug: '', subs: '' }
  isModalOpen.value = true
}

const openEditModal = (cat: any) => {
  editingId.value = cat.id
  sectorForm.value = {
    name: cat.name,
    slug: cat.slug,
    subs: cat.subcategories.join(', ')
  }
  isModalOpen.value = true
}

const saveSector = () => {
  themeStore.initializeUplink()
  setTimeout(() => {
    if (editingId.value) {
      const idx = categories.value.findIndex(c => c.id === editingId.value)
      categories.value[idx] = { 
        ...categories.value[idx], 
        name: sectorForm.value.name, 
        slug: sectorForm.value.slug,
        subcategories: sectorForm.value.subs.split(',').map(s => s.trim())
      }
      showToast('SECTOR_MODIFICATION_COMPLETE')
    } else {
      categories.value.unshift({
        id: Date.now(),
        name: sectorForm.value.name,
        slug: sectorForm.value.slug,
        jobCount: 0,
        traffic: 0,
        subcategories: sectorForm.value.subs.split(',').map(s => s.trim())
      })
      showToast('NEW_SECTOR_UPLINK_ESTABLISHED')
    }
    isModalOpen.value = false
  }, 800)
}

// Delete Logic
const confirmDelete = (cat: any) => {
  deleteConfirm.value = {
    show: true,
    targetId: cat.id,
    targetName: cat.name
  }
}

const executeDelete = () => {
  themeStore.glitchMode = true
  const idToDelete = deleteConfirm.value.targetId
  
  setTimeout(() => {
    categories.value = categories.value.filter(c => c.id !== idToDelete)
    themeStore.glitchMode = false
    deleteConfirm.value.show = false
    showToast('NODE_DE_ALLOCATED_SUCCESSFULLY')
  }, 600)
}
</script>

<style scoped>
  @reference 'tailwindcss';
.hud-input {
  @apply w-full bg-white/[0.03] border border-white/10 p-4 text-xs font-bold tracking-widest focus:outline-none focus:border-white/30 transition-all;
}

/* Animations */
.modal-fade-enter-active, .modal-fade-leave-active { transition: all 0.3s ease; }
.modal-fade-enter-from, .modal-fade-leave-to { opacity: 0; transform: scale(0.98); }

.toast-enter-active, .toast-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-enter-from, .toast-leave-to { transform: translateX(100%) translateY(20px); opacity: 0; }

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }
</style>