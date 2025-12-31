<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden">
    <admin-sidebar />

    <div :class="['flex-1 flex flex-col h-full relative transition-all duration-500', themeStore.terminalCollapsed ? 'ml-20' : 'ml-72']">
      
      <header class="h-32 flex flex-col justify-center px-12 shrink-0 border-b border-white/5 bg-black/60 z-20">
        <div class="flex items-center justify-between">
          <div>
            <p class="text-[9px] font-black text-cyan-500 uppercase tracking-[0.3em] mb-1">Recruitment_Ops // ATS_v15</p>
            <h1 class="text-4xl font-black italic uppercase tracking-tighter">PIPELINE_TERMINAL</h1>
          </div>
        </div>
      </header>

      <nav class="px-12 bg-white/[0.02] border-b border-white/5 flex gap-1 shrink-0 z-20">
        <button v-for="stage in stageList" :key="stage" @click="activeStage = stage"
          :class="['px-8 py-5 text-[10px] font-black uppercase tracking-widest transition-all relative', activeStage === stage ? 'text-cyan-400 bg-white/[0.05]' : 'text-gray-500 hover:text-gray-300']">
          {{ stage }} <span class="ml-2 opacity-30">[{{ getCount(stage) }}]</span>
          <div v-if="activeStage === stage" class="absolute bottom-0 left-0 w-full h-0.5 bg-cyan-500 shadow-[0_0_15px_#06b6d4]"></div>
        </button>
      </nav>

      <main class="flex-1 p-8 flex gap-8 bg-[#020205] z-10 overflow-hidden">
        <div class="flex-1 border border-white/10 bg-white/[0.01] flex flex-col overflow-hidden">
          <div class="flex-1 overflow-y-auto custom-scrollbar">
            <div v-for="candidate in filteredCandidates" :key="candidate.id"
              class="grid grid-cols-12 items-center p-6 border-b border-white/5 group transition-all">
              <div class="col-span-6 flex items-center gap-6">
                <div class="w-10 h-10 bg-white/5 flex items-center justify-center text-[10px] font-black border border-white/10">{{ candidate.name[0] }}</div>
                <h4 class="text-sm font-black uppercase">{{ candidate.name }}</h4>
              </div>

              <div class="col-span-6 flex justify-end gap-3">
                <button v-if="candidate.currentStage === 'OFFER'" @click.stop="openDeclineModal(candidate)" 
                  class="px-6 py-2 border border-red-500/50 text-red-500 text-[9px] font-black uppercase hover:bg-red-500 hover:text-white transition-all">
                  Decline_Offer
                </button>
                <button @click.stop="promoteCandidate(candidate)" 
                  class="px-6 py-2 bg-white text-black text-[9px] font-black uppercase hover:bg-cyan-500 transition-all">
                  {{ candidate.currentStage === 'OFFER' ? 'Finalize Hire' : 'Promote' }}
                </button>
              </div>
            </div>
          </div>
        </div>
      </main>

      <transition name="modal-zoom">
        <div v-if="hiredModal.show" class="fixed inset-0 z-[100] flex items-center justify-center bg-black/95 backdrop-blur-xl">
          <div v-for="n in 50" :key="n" class="confetti" :style="generateConfettiStyle()"></div>
          <div class="max-w-md w-full p-1 border border-cyan-500/50 bg-[#020205] text-center relative z-10">
            <div class="p-10 border border-cyan-500/10">
              <h2 class="text-3xl font-black italic uppercase text-white mb-2">HIRE_SUCCESS</h2>
              <p class="text-[10px] font-black text-cyan-500 uppercase mb-8">{{ hiredModal.name }}</p>
              <div class="space-y-3">
                <button @click="exportToPDF" class="w-full py-4 bg-cyan-500 text-black text-[10px] font-black uppercase hover:bg-cyan-400">Export_Record</button>
                <button @click="hiredModal.show = false" class="w-full py-3 border border-white/20 text-white text-[9px] font-black uppercase">Dismiss</button>
              </div>
            </div>
          </div>
        </div>
      </transition>

      <transition name="modal-zoom">
        <div v-if="declineModal.show" class="fixed inset-0 z-[100] flex items-center justify-center bg-black/90 backdrop-blur-md">
          <div class="max-w-md w-full p-1 border border-red-500/50 bg-[#020205]">
            <div class="p-8 border border-red-500/10 text-center">
              <h3 class="text-xl font-black italic uppercase text-red-500 mb-6">Archive_Candidate</h3>
              
              <div class="grid grid-cols-2 gap-2 mb-4">
                <button v-for="reason in declineReasons" :key="reason" 
                  @click="selectedReason = reason"
                  :class="['py-3 border text-[9px] font-black uppercase transition-all', 
                           selectedReason === reason ? 'bg-red-500 text-white border-red-400' : 'border-white/5 text-gray-500 hover:text-white']">
                  {{ reason }}
                </button>
              </div>

              <transition name="fade">
                <div v-if="selectedReason === 'Other'" class="mb-6">
                  <p class="text-[8px] font-black text-red-500/60 uppercase mb-2 text-left">Internal_Comments</p>
                  <textarea 
                    v-model="customReason" 
                    placeholder="Enter specific reason for archival..."
                    class="w-full bg-black border border-red-500/30 p-4 text-[10px] text-white font-mono outline-none focus:border-red-500 min-h-[100px] resize-none"
                  ></textarea>
                </div>
              </transition>

              <div class="flex flex-col gap-3">
                <button @click="confirmDecline" 
                  :disabled="!selectedReason || (selectedReason === 'Other' && !customReason)"
                  class="w-full py-4 bg-red-600 text-white text-[10px] font-black uppercase hover:bg-red-500 disabled:opacity-20 disabled:cursor-not-allowed transition-all">
                  Confirm_Rejection
                </button>
                <button @click="closeDeclineModal" class="text-[8px] font-black uppercase text-gray-600 underline hover:text-gray-400">Cancel_Action</button>
              </div>
            </div>
          </div>
        </div>
      </transition>

      <transition name="toast-slide">
        <div v-if="statusToast.show" 
          :class="['fixed bottom-10 right-10 px-8 py-4 z-[110] border font-black text-[10px] uppercase shadow-2xl', 
                  statusToast.type === 'error' ? 'bg-red-950 border-red-500 text-red-500' : 'bg-cyan-950 border-cyan-500 text-cyan-500']">
          {{ statusToast.message }}
        </div>
      </transition>

    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const activeStage = ref('OFFER')
const hiredModal = ref({ show: false, name: '' })
const statusToast = ref({ show: false, message: '', type: 'success' })

// Decline Modal State
const declineModal = ref({ show: false, candidate: null })
const selectedReason = ref('')
const customReason = ref('')
const declineReasons = ['Salary', 'Competing Offer', 'Culture Fit', 'Relocation', 'Role Rescinded', 'Other']

const stageList = ['INBOX', 'SCREENING', 'INTERVIEW', 'OFFER']
const candidates = ref([
  { id: 1, name: 'Marcus Holloway', currentStage: 'OFFER' },
  { id: 2, name: 'Elena Rodriguez', currentStage: 'OFFER' }
])

const filteredCandidates = computed(() => candidates.value.filter(c => c.currentStage === activeStage.value))
const getCount = (stage) => candidates.value.filter(c => c.currentStage === stage).length

const triggerToast = (msg, type = 'success') => {
  statusToast.value = { show: true, message: msg, type }
  setTimeout(() => statusToast.value.show = false, 4000)
}

const promoteCandidate = (candidate) => {
  const idx = stageList.indexOf(candidate.currentStage)
  if (idx === stageList.length - 1) {
    hiredModal.value = { show: true, name: candidate.name }
    candidates.value = candidates.value.filter(c => c.id !== candidate.id)
  } else {
    candidate.currentStage = stageList[idx + 1]
    triggerToast(`${candidate.name} PROMOTED`)
  }
}

const openDeclineModal = (candidate) => {
  declineModal.value = { show: true, candidate }
  selectedReason.value = ''
  customReason.value = ''
}

const closeDeclineModal = () => {
  declineModal.value.show = false
  selectedReason.value = ''
}

const confirmDecline = () => {
  const finalReason = selectedReason.value === 'Other' ? customReason.value : selectedReason.value
  const candidateName = declineModal.value.candidate.name
  
  // Update UI
  candidates.value = candidates.value.filter(c => c.id !== declineModal.value.candidate.id)
  declineModal.value.show = false
  
  triggerToast(`${candidateName} ARCHIVED // REASON: ${finalReason}`, 'error')
}

const generateConfettiStyle = () => ({
  left: Math.random() * 100 + 'vw',
  backgroundColor: ['#06b6d4', '#ffffff', '#22d3ee'][Math.floor(Math.random() * 3)],
  animationDuration: (Math.random() * 3 + 2) + 's',
  animationDelay: Math.random() * 2 + 's'
})

// PDF Mock
const exportToPDF = () => {
  triggerToast('GENERATING_DOCUMENT...')
}
</script>

<style scoped>
   @reference 'tailwindcss'; 
.confetti { position: absolute; top: -20px; width: 8px; height: 8px; animation: fall linear infinite; }
@keyframes fall { 0% { transform: translateY(0) rotate(0deg); opacity: 1; } 100% { transform: translateY(110vh) rotate(720deg); opacity: 0; } }

.modal-zoom-enter-active { transition: all 0.4s cubic-bezier(0.34, 1.56, 0.64, 1); }
.modal-zoom-enter-from { opacity: 0; transform: scale(0.9); }

.fade-enter-active { transition: all 0.3s ease; }
.fade-enter-from { opacity: 0; transform: translateY(-10px); }

.toast-slide-enter-active { transition: transform 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-slide-enter-from { transform: translateX(120%); }

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }
</style>