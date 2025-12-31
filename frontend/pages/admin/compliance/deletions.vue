<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-red-500/30">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)] relative',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-5" 
           style="background: radial-gradient(circle at 50% 0%, #ef4444, transparent 70%)">
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-6 py-8 lg:px-12">
        <div class="max-w-6xl mx-auto">
          
          <div class="mb-12 flex flex-col md:flex-row md:items-end justify-between gap-6">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12 bg-red-500/40"></div>
                <span class="text-xs font-bold tracking-widest text-red-500 uppercase italic">Security_Level: OMEGA</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                DELETION_VAULT
              </h1>
            </div>

            <div class="flex flex-col items-end">
              <span class="text-[10px] font-black text-gray-600 uppercase mb-1">Vault_Retention_Status</span>
              <div class="flex gap-1">
                <div v-for="i in 10" :key="i" class="w-4 h-1" :class="i < 8 ? 'bg-red-500' : 'bg-white/10'"></div>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-4 gap-4 mb-8">
            <button 
              v-for="filter in filters" :key="filter.label"
              @click="activeFilter = filter.label"
              :class="[
                'p-6 border transition-all text-left relative group',
                activeFilter === filter.label ? 'border-red-500 bg-red-500/5' : 'border-white/5 bg-white/[0.02] hover:border-white/20'
              ]"
            >
              <p class="text-[9px] font-black uppercase tracking-widest" :class="activeFilter === filter.label ? 'text-red-500' : 'text-gray-500'">{{ filter.label }}</p>
              <p class="text-2xl font-black italic mt-1">{{ filter.count }}</p>
              <div v-if="activeFilter === filter.label" class="absolute top-0 right-0 p-2">
                <div class="w-1.5 h-1.5 bg-red-500 animate-ping"></div>
              </div>
            </button>
          </div>

          <div class="bg-black border border-white/10 shadow-2xl overflow-hidden relative">
            <div class="bg-red-950/20 border-b border-red-500/20 p-4 flex items-center gap-4">
              <svg class="w-5 h-5 text-red-500" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" /></svg>
              <p class="text-[10px] font-black uppercase tracking-widest text-red-400">
                Items in the vault will be permanently purged after 30 days of inactivity.
              </p>
            </div>

            <table class="w-full text-left border-collapse">
              <thead>
                <tr class="border-b border-white/10 bg-white/[0.01]">
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">De-allocated_Entity</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Origin_Node</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-gray-500">Purge_ETA</th>
                  <th class="p-6 text-[10px] font-black uppercase tracking-[0.2em] text-right text-gray-500">Protocol</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/5">
                <tr v-for="item in filteredItems" :key="item.id" class="hover:bg-red-500/[0.02] transition-colors group">
                  <td class="p-6">
                    <div class="flex flex-col">
                      <span class="text-sm font-black uppercase tracking-tight text-white group-hover:text-red-400 transition-colors">{{ item.name }}</span>
                      <span class="text-[9px] font-bold text-gray-600 uppercase mt-1">ID: {{ item.id }}</span>
                    </div>
                  </td>
                  <td class="p-6">
                    <span class="text-[10px] font-black uppercase text-gray-400 px-2 py-1 border border-white/10 bg-white/5">
                      {{ item.type }}
                    </span>
                  </td>
                  <td class="p-6">
                    <div class="flex flex-col gap-1">
                      <span class="text-[11px] font-bold text-red-500/80 uppercase tracking-tighter">{{ item.daysLeft }} DAYS REMAINING</span>
                      <div class="w-24 h-1 bg-white/5 rounded-full overflow-hidden">
                        <div class="h-full bg-red-600" :style="{ width: (item.daysLeft / 30 * 100) + '%' }"></div>
                      </div>
                    </div>
                  </td>
                  <td class="p-6 text-right">
                    <div class="flex justify-end gap-3">
                      <button @click="restoreItem(item)" class="px-4 py-2 border border-white/10 text-[9px] font-black uppercase hover:bg-white/5 transition-all">
                        RESTORE_UPLINK
                      </button>
                      <button @click="confirmPermanentPurge(item)" class="px-4 py-2 bg-red-600/10 border border-red-500/30 text-[9px] font-black uppercase text-red-500 hover:bg-red-600 hover:text-white transition-all">
                        PURGE_NOW
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

    <transition name="fade">
      <div v-if="showDeleteModal" class="fixed inset-0 z-[100] flex items-center justify-center p-6 bg-black/90 backdrop-blur-md">
        <div class="w-full max-w-md bg-[#050508] border-2 border-red-600 p-8 text-center shadow-[0_0_50px_rgba(239,68,68,0.2)]">
          <div class="mb-6 inline-block p-4 rounded-full bg-red-500/10 text-red-500">
            <svg class="w-12 h-12" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" stroke-width="2"/></svg>
          </div>
          <h2 class="text-2xl font-black italic uppercase tracking-tighter mb-2">Irreversible_Purge</h2>
          <p class="text-[11px] font-bold text-gray-500 uppercase leading-relaxed mb-8">
            Targeting: <span class="text-white">{{ itemToPurge?.name }}</span><br/>
            This action will destroy all meta-data and file attachments associated with this entity.
          </p>
          <div class="flex flex-col gap-3">
            <button @click="executePurge" class="w-full py-4 bg-red-600 text-black text-[10px] font-black uppercase tracking-[0.2em] hover:bg-red-500">
              EXECUTE_PERMANENT_DELETION
            </button>
            <button @click="showDeleteModal = false" class="w-full py-4 border border-white/10 text-[10px] font-black uppercase tracking-[0.2em]">
              ABORT_PROTOCOL
            </button>
          </div>
        </div>
      </div>
    </transition>

    <transition name="toast">
      <div v-if="toast.show" class="fixed bottom-8 right-8 z-[200] bg-black border-l-4 border-red-500 p-4 min-w-[280px] shadow-2xl">
        <p class="text-[10px] font-black uppercase tracking-widest text-red-500">Vault_Update</p>
        <p class="text-xs font-bold mt-1 tracking-tight">{{ toast.message }}</p>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const activeFilter = ref('ALL')
const showDeleteModal = ref(false)
const itemToPurge = ref(null)
const toast = ref({ show: false, message: '' })

const filters = ref([
  { label: 'ALL', count: 24 },
  { label: 'JOBS', count: 12 },
  { label: 'USERS', count: 8 },
  { label: 'EMPLOYERS', count: 4 }
])

const trashItems = ref([
  { id: 'JOB-992', name: 'Senior Blockchain Arch', type: 'JOBS', daysLeft: 4, category: 'Engineering' },
  { id: 'USR-112', name: 'Alex_Vapor_99', type: 'USERS', daysLeft: 28, category: 'Candidate' },
  { id: 'EMP-004', name: 'Cyberdyne Systems', type: 'EMPLOYERS', daysLeft: 12, category: 'Enterprise' },
  { id: 'JOB-441', name: 'Neural Linker VR', type: 'JOBS', daysLeft: 2, category: 'Medical' },
  { id: 'USR-882', name: 'Sarah_K_01', type: 'USERS', daysLeft: 15, category: 'Candidate' },
])

const filteredItems = computed(() => {
  if (activeFilter.value === 'ALL') return trashItems.value
  return trashItems.value.filter(i => i.type === activeFilter.value)
})

const showToast = (msg) => {
  toast.value = { show: true, message: msg }
  setTimeout(() => toast.value.show = false, 3000)
}

const restoreItem = (item) => {
  themeStore.initializeUplink()
  setTimeout(() => {
    trashItems.value = trashItems.value.filter(i => i.id !== item.id)
    showToast(`${item.id}_RESTORED_TO_LIVE_DATABASE`)
  }, 600)
}

const confirmPermanentPurge = (item) => {
  itemToPurge.value = item
  showDeleteModal.value = true
}

const executePurge = () => {
  themeStore.glitchMode = true
  setTimeout(() => {
    trashItems.value = trashItems.value.filter(i => i.id !== itemToPurge.value.id)
    showDeleteModal.value = false
    themeStore.glitchMode = false
    showToast(`ENTITY_${itemToPurge.value.id}_PURGED_PERMANENTLY`)
  }, 800)
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(239, 68, 68, 0.2); }

.fade-enter-active, .fade-leave-active { transition: opacity 0.3s ease; }
.fade-enter-from, .fade-leave-to { opacity: 0; }

.toast-enter-active, .toast-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-enter-from, .toast-leave-to { transform: translateX(100%); opacity: 0; }
</style>