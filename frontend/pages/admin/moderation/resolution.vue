<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-white/10">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)] relative',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-[0.07]" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 flex flex-col relative z-10 px-6 py-8 lg:px-12 overflow-hidden">
        <div class="max-w-7xl mx-auto w-full h-full flex flex-col">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-8 shrink-0">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12" :style="{ backgroundColor: themeStore.accentColor + '66' }"></div>
                <span class="text-xs font-black tracking-widest uppercase opacity-60">Moderation_Layer // Resolution</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                TRIBUNAL_CORE
              </h1>
            </div>
            
            <div class="flex gap-4">
              <div class="text-right border-r border-white/10 pr-6">
                <p class="text-[9px] font-black text-gray-500 uppercase">Pending_Cases</p>
                <p class="text-2xl font-black" :style="{ color: themeStore.accentColor }">14</p>
              </div>
              <div class="text-right">
                <p class="text-[9px] font-black text-gray-500 uppercase">Avg_Resolve_Time</p>
                <p class="text-2xl font-black text-white">2.4m</p>
              </div>
            </div>
          </div>

          <div class="flex-1 flex gap-6 overflow-hidden min-h-0 mb-6">
            
            <div class="w-80 shrink-0 border border-white/10 bg-black/40 flex flex-col overflow-hidden">
              <div class="p-4 bg-white/[0.02] border-b border-white/10 flex justify-between items-center">
                <span class="text-[10px] font-black uppercase tracking-widest text-gray-400">Case_Buffer</span>
                <span class="text-[8px] px-2 py-0.5 font-bold" 
                      :style="{ backgroundColor: themeStore.accentColor + '33', color: themeStore.accentColor }">
                  HIGH_PRIORITY
                </span>
              </div>
              <div class="flex-1 overflow-y-auto custom-scrollbar divide-y divide-white/5">
                <button 
                  v-for="caseItem in cases" :key="caseItem.id"
                  @click="selectedCase = caseItem"
                  :class="[
                    'w-full p-5 text-left transition-all relative group',
                    selectedCase?.id === caseItem.id ? 'bg-white/[0.05]' : 'hover:bg-white/[0.02]'
                  ]"
                >
                  <div v-if="selectedCase?.id === caseItem.id" 
                       class="absolute left-0 top-0 bottom-0 w-1" 
                       :style="{ backgroundColor: themeStore.accentColor }"></div>
                  
                  <div class="flex justify-between items-start mb-2">
                    <span class="text-[9px] font-black text-gray-500 uppercase">#{{ caseItem.id }}</span>
                    <span class="text-[8px] font-bold text-gray-400 italic">{{ caseItem.timeAgo }}</span>
                  </div>
                  <p class="text-xs font-black uppercase tracking-tight mb-1 text-white truncate">{{ caseItem.title }}</p>
                  <p class="text-[9px] font-bold uppercase opacity-80" :style="{ color: themeStore.accentColor }">
                    {{ caseItem.reason }}
                  </p>
                </button>
              </div>
            </div>

            <div v-if="selectedCase" class="flex-1 flex gap-6 overflow-hidden">
              <div class="flex-1 border border-white/10 bg-black flex flex-col overflow-hidden">
                <div class="p-4 border-b border-white/10 bg-white/[0.02] flex justify-between items-center">
                  <span class="text-[10px] font-black text-gray-500 uppercase tracking-widest">Evidence_Log // {{ selectedCase.type }}</span>
                  <button class="text-[9px] font-black uppercase hover:text-white transition-colors" :style="{ color: themeStore.accentColor }">
                    View_Live_Profile
                  </button>
                </div>
                <div class="flex-1 p-8 overflow-y-auto custom-scrollbar space-y-8">
                  <div class="space-y-4">
                    <h3 class="text-lg font-black uppercase tracking-tighter text-white">{{ selectedCase.title }}</h3>
                    <div class="p-4 bg-white/[0.03] border-l-2 border-white/10 italic text-sm text-gray-400 leading-relaxed">
                      "{{ selectedCase.reportedContent }}"
                    </div>
                  </div>
                  <div class="grid grid-cols-2 gap-6">
                    <div class="space-y-2">
                      <p class="text-[9px] font-black text-gray-500 uppercase">Reporter_ID</p>
                      <p class="text-xs font-bold text-white">{{ selectedCase.reporter }}</p>
                    </div>
                    <div class="space-y-2">
                      <p class="text-[9px] font-black text-gray-500 uppercase">Trust_Score</p>
                      <p class="text-xs font-bold text-green-500">9.8 / 10</p>
                    </div>
                  </div>
                </div>
              </div>

              <div class="w-96 border border-white/10 bg-black/40 p-6 flex flex-col">
                <div class="flex-1 space-y-8">
                  <div>
                    <h4 class="text-xs font-black uppercase tracking-[0.2em] mb-6">Verdict_Protocol</h4>
                    <div class="space-y-3">
                      <label class="text-[9px] font-black text-gray-500 uppercase">Action_Selection</label>
                      <div class="flex flex-col gap-2">
                        <button 
                          v-for="action in verdictActions" :key="action.id"
                          @click="selectedVerdict = action.id"
                          :class="[
                            'p-4 border text-left transition-all uppercase text-[10px] font-black tracking-widest',
                            selectedVerdict === action.id ? 'bg-white text-black border-white' : 'border-white/10 text-gray-500 hover:border-white/30'
                          ]"
                          :style="selectedVerdict === action.id ? { backgroundColor: themeStore.accentColor, borderColor: themeStore.accentColor } : {}"
                        >
                          {{ action.label }}
                        </button>
                      </div>
                    </div>
                  </div>

                  <div class="space-y-3">
                    <label class="text-[9px] font-black text-gray-500 uppercase">Internal_Notes</label>
                    <textarea 
                      placeholder="JUSTIFY_VERDICT..."
                      class="w-full h-32 bg-white/[0.03] border border-white/10 p-4 text-xs font-bold focus:outline-none uppercase"
                      :style="{ focusBorderColor: themeStore.accentColor }"
                    ></textarea>
                  </div>
                </div>

                <button 
                  @click="finalizeVerdict"
                  class="w-full py-5 text-black text-xs font-black uppercase tracking-[0.3em] hover:brightness-110 active:scale-95 transition-all mt-6"
                  :style="{ backgroundColor: themeStore.accentColor }"
                >
                  EXECUTE_VERDICT
                </button>
              </div>
            </div>

            <div v-else class="flex-1 border border-dashed border-white/10 flex items-center justify-center bg-white/[0.01]">
              <div class="text-center">
                <div class="w-12 h-12 border border-white/10 flex items-center justify-center mx-auto mb-4">
                  <div class="w-2 h-2 bg-white/20 animate-ping"></div>
                </div>
                <p class="text-[10px] font-black text-gray-600 uppercase tracking-[0.3em]">Select_Case_From_Buffer</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <transition name="toast">
      <div v-if="toast.show" 
           class="fixed bottom-8 right-8 z-[200] bg-black border-l-4 p-4 min-w-[300px] shadow-2xl"
           :style="{ borderColor: themeStore.accentColor }">
        <p class="text-[10px] font-black uppercase tracking-widest opacity-60">Protocol_Finalized</p>
        <p class="text-xs font-bold mt-1 tracking-tight">CASE_CLOSED_AND_ARCHIVED</p>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const selectedCase = ref(null)
const selectedVerdict = ref(null)
const toast = ref({ show: false })

const cases = ref([
  { 
    id: 'CASE-7712', 
    title: 'Phishing_Attempt: Neo-Tech Corp', 
    type: 'JOB_LISTING', 
    reason: 'Suspicious_External_Link', 
    timeAgo: '4m ago',
    reportedContent: 'We are looking for a Senior Architect. Please click this link to verify your identity: http://bit.ly/fake-verify-job',
    reporter: 'USR_8821'
  },
  { 
    id: 'CASE-7715', 
    title: 'Hate_Speech: Candidate_X', 
    type: 'USER_PROFILE', 
    reason: 'TOS_Violation', 
    timeAgo: '12m ago',
    reportedContent: 'Profile bio contains inflammatory language toward specific regional nodes.',
    reporter: 'SYS_DAEMON'
  },
  { 
    id: 'CASE-7719', 
    title: 'Scam: Global Logistics', 
    type: 'EMPLOYER', 
    reason: 'Payment_Bypass_Attempt', 
    timeAgo: '22m ago',
    reportedContent: 'Employer requested candidate pay for "Equipment Training" via crypto before interview.',
    reporter: 'USR_0012'
  }
])

const verdictActions = [
  { id: 'dismiss', label: 'DISMISS_CASE' },
  { id: 'suspend', label: 'SUSPEND_ENTITY' },
  { id: 'shadowban', label: 'SHADOW_RESTRICT' },
  { id: 'purge', label: 'PERMANENT_PURGE' }
]

const finalizeVerdict = () => {
  themeStore.initializeUplink()
  setTimeout(() => {
    if (selectedCase.value) {
      cases.value = cases.value.filter(c => c.id !== selectedCase.value.id)
      selectedCase.value = null
      selectedVerdict.value = null
      toast.value.show = true
      setTimeout(() => toast.value.show = false, 3000)
    }
  }, 800)
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }

.toast-enter-active, .toast-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-enter-from, .toast-leave-to { transform: translateX(100%); opacity: 0; }
</style>