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

      <div class="flex-1 flex flex-col relative z-10 px-6 py-8 lg:px-12 overflow-hidden">
        <div class="max-w-7xl mx-auto w-full h-full flex flex-col">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-8 shrink-0">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12 bg-white/20"></div>
                <span class="text-xs font-bold tracking-widest text-gray-400 uppercase">Compliance_Protocol_v9.0</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                LEGAL_CORE
              </h1>
            </div>

            <div class="flex gap-4">
              <button 
                @click="saveCurrentPolicy"
                class="px-8 py-4 text-[10px] font-black uppercase tracking-widest text-black transition-all hover:brightness-110 flex items-center gap-3" 
                :style="{ backgroundColor: themeStore.accentColor }"
              >
                COMMIT_REVISIONS
              </button>
            </div>
          </div>

          <div class="flex-1 flex gap-6 overflow-hidden min-h-0 mb-6">
            
            <div class="w-72 shrink-0 border border-white/10 bg-black/40 overflow-y-auto custom-scrollbar">
              <div class="p-4 border-b border-white/10 bg-white/[0.02]">
                <span class="text-[9px] font-black text-gray-500 uppercase tracking-widest">Available_Protocols</span>
              </div>
              <div class="divide-y divide-white/5">
                <button 
                  v-for="policy in policies" 
                  :key="policy.id"
                  @click="activePolicyId = policy.id"
                  :class="[
                    'w-full p-5 text-left transition-all group relative',
                    activePolicyId === policy.id ? 'bg-white/[0.05]' : 'hover:bg-white/[0.02]'
                  ]"
                >
                  <div v-if="activePolicyId === policy.id" 
                       class="absolute left-0 top-0 bottom-0 w-1" 
                       :style="{ backgroundColor: themeStore.accentColor }"></div>
                  
                  <p class="text-[10px] font-bold uppercase tracking-tighter mb-1"
                     :class="activePolicyId === policy.id ? 'text-white' : 'text-gray-500'">
                    {{ policy.code }}
                  </p>
                  <p class="text-xs font-black uppercase tracking-tight"
                     :class="activePolicyId === policy.id ? 'text-cyan-400' : 'text-gray-300'">
                    {{ policy.title }}
                  </p>
                  <div class="mt-3 flex items-center justify-between">
                    <span class="text-[8px] font-bold text-gray-600 uppercase">Rev: {{ policy.version }}</span>
                    <span v-if="policy.status === 'Draft'" class="text-[8px] px-1 bg-amber-500/10 text-amber-500 border border-amber-500/20">DRAFT</span>
                  </div>
                </button>
              </div>
            </div>

            <div class="flex-1 border border-white/10 bg-black flex flex-col overflow-hidden">
              <div class="p-4 border-b border-white/10 bg-white/[0.02] flex justify-between items-center shrink-0">
                <div class="flex items-center gap-4">
                  <span class="text-[10px] font-black text-gray-500 uppercase">Buffer_Mode: EDIT_MARKDOWN</span>
                  <div class="h-4 w-[1px] bg-white/10"></div>
                  <span class="text-[10px] font-black text-cyan-500 uppercase">Active: {{ currentPolicy?.title }}</span>
                </div>
                <div class="flex gap-2">
                  <div class="w-2 h-2 rounded-full bg-green-500 animate-pulse"></div>
                  <span class="text-[9px] font-bold text-gray-600 uppercase">Encrypted_Uplink</span>
                </div>
              </div>

              <div class="flex-1 p-8 overflow-y-auto custom-scrollbar bg-[url('https://www.transparenttextures.com/patterns/carbon-fibre.png')]">
                <textarea 
                  v-model="currentPolicy.content"
                  class="w-full h-full bg-transparent text-gray-300 font-mono text-sm leading-relaxed focus:outline-none resize-none scroll-smooth"
                  spellcheck="false"
                ></textarea>
              </div>

              <div class="p-4 border-t border-white/10 bg-white/[0.01] flex justify-between items-center shrink-0">
                <div class="flex gap-6 text-[9px] font-bold text-gray-600 uppercase tracking-widest">
                  <span>Lines: {{ currentPolicy?.content.split('\n').length }}</span>
                  <span>Chars: {{ currentPolicy?.content.length }}</span>
                </div>
                <span class="text-[9px] font-bold text-gray-600 uppercase tracking-widest">
                  Last_Backup: {{ lastSaved }}
                </span>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-4 gap-6 shrink-0">
            <div v-for="item in complianceMetrics" :key="item.label" class="bg-white/[0.02] border border-white/5 p-4 flex justify-between items-center">
              <div>
                <p class="text-[9px] font-black text-gray-500 uppercase tracking-widest">{{ item.label }}</p>
                <p class="text-xl font-black italic tracking-tighter mt-1">{{ item.value }}</p>
              </div>
              <div :class="item.alert ? 'text-red-500' : 'text-green-500'">
                <svg class="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><path :d="item.icon" /></svg>
              </div>
            </div>
          </div>

        </div>
      </div>
    </main>

    <transition name="toast">
      <div v-if="showToast" class="fixed bottom-8 right-8 z-[200] bg-black border-l-4 p-4 min-w-[300px] shadow-2xl" :style="{ borderColor: themeStore.accentColor }">
        <p class="text-[10px] font-black uppercase tracking-widest text-gray-500">System_Update</p>
        <p class="text-xs font-bold mt-1 tracking-tight">POLICIES_REWRITTEN_TO_BLOCKCHAIN</p>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const activePolicyId = ref(1)
const showToast = ref(false)
const lastSaved = ref('14:20:03')

const policies = ref([
  { 
    id: 1, 
    code: 'POL-TOS', 
    title: 'Terms of Service', 
    version: '4.2.1', 
    status: 'Live',
    content: '# TERMS_OF_SERVICE_V4\n\n1. NEURAL_INTERFACE_USAGE\nUsers must ensure all neural uplinks are calibrated to system standards.\n\n2. DATA_EXTRACTION_LIMITS\nJob seekers may only pull 500 records per cycle.\n\n3. LIABILITY_SHIELD\nThe portal is not responsible for bio-data leaks caused by third-party hacking units.'
  },
  { 
    id: 2, 
    code: 'POL-PRIV', 
    title: 'Privacy Protocol', 
    version: '2.0.0', 
    status: 'Draft',
    content: '# PRIVACY_PROTOCOL_V2\n\nYour data is hashed using SHA-512 and stored in off-shore server nodes. We do not sell your biological signatures to recruiters without express consent.'
  },
  { 
    id: 3, 
    code: 'POL-GDPR', 
    title: 'GDPR Compliance', 
    version: '1.5.0', 
    status: 'Live',
    content: '# GDPR_COMPLIANCE\n\nRight to be forgotten: Users can initiate a "NODE_PURGE" at any time to remove all traces of their identity from the central mainframe.'
  },
  { 
    id: 4, 
    code: 'POL-REFUND', 
    title: 'Refund Policy', 
    version: '1.1.0', 
    status: 'Live',
    content: '# REFUND_POLICY\n\nCredits used for featured job placement are non-refundable once the data stream has hit the public index.'
  }
])

const complianceMetrics = ref([
  { label: 'System_Audit', value: 'PASS', icon: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z', alert: false },
  { label: 'Pending_Drafts', value: '01', icon: 'M11 5H6a2 2 0 00-2 2v11a2 2 0 002 2h11a2 2 0 002-2v-5m-1.414-9.414a2 2 0 112.828 2.828L11.828 15H9v-2.828l8.586-8.586z', alert: true },
  { label: 'Data_Retention', value: '365_DAYS', icon: 'M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z', alert: false },
  { label: 'Legal_Last_Sync', value: 'STABLE', icon: 'M21 12a9 9 0 01-9 9m9-9a9 9 0 00-9-9m9 9H3m9 9a9 9 0 01-9-9m9 9c1.657 0 3-4.03 3-9s-1.343-9-3-9m0 18c-1.657 0-3-4.03-3-9s1.343-9 3-9m-9 9a9 9 0 019-9', alert: false }
])

const currentPolicy = computed(() => {
  return policies.value.find(p => p.id === activePolicyId.value) || policies.value[0]
})

const saveCurrentPolicy = () => {
  themeStore.initializeUplink()
  showToast.value = true
  lastSaved.value = new Date().toLocaleTimeString('en-GB', { hour12: false }).slice(0, 5)
  setTimeout(() => { showToast.value = false }, 3000)
}
</script>

<style scoped>
  @reference 'tailwindcss';
.hud-input {
  @apply w-full bg-white/[0.03] border border-white/10 p-4 text-xs font-bold tracking-widest focus:outline-none focus:border-white/30 transition-all;
}

textarea::selection {
  background: rgba(0, 243, 255, 0.3);
}

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }

.toast-enter-active, .toast-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-enter-from, .toast-leave-to { transform: translateX(100%); opacity: 0; }
</style>