<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-white/10">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)] relative',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-[0.05]" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 flex flex-col relative z-10 px-6 py-8 lg:px-12 overflow-hidden">
        <div class="max-w-7xl mx-auto w-full h-full flex flex-col">
          
          <div class="flex flex-col md:flex-row md:items-end justify-between gap-6 mb-10 shrink-0">
            <div>
              <div class="flex items-center gap-3 mb-4">
                <div class="h-[1px] w-12" :style="{ backgroundColor: themeStore.accentColor + '66' }"></div>
                <span class="text-xs font-black tracking-widest uppercase opacity-60">Communication_Audit // Neural_Filter</span>
              </div>
              <h1 class="text-5xl lg:text-6xl font-black italic uppercase tracking-tighter">
                INTERCEPT_LOGS
              </h1>
            </div>

            <div class="flex gap-4">
              <div class="bg-white/[0.02] border border-white/10 p-4 min-w-[150px]">
                <p class="text-[9px] font-black text-gray-500 uppercase">Detection_Accuracy</p>
                <p class="text-xl font-black text-green-500 mt-1">99.2%</p>
              </div>
            </div>
          </div>

          <div class="flex-1 grid grid-cols-1 lg:grid-cols-12 gap-6 overflow-hidden min-h-0 mb-6">
            
            <div class="lg:col-span-4 border border-white/10 bg-black/40 flex flex-col overflow-hidden">
              <div class="p-4 border-b border-white/10 bg-white/[0.02] flex justify-between items-center">
                <span class="text-[10px] font-black uppercase tracking-widest text-gray-400">Flagged_Streams</span>
              </div>
              <div class="flex-1 overflow-y-auto custom-scrollbar divide-y divide-white/5">
                <button 
                  v-for="msg in flaggedMessages" :key="msg.id"
                  @click="selectedMessage = msg"
                  :class="[
                    'w-full p-5 text-left transition-all relative group',
                    selectedMessage?.id === msg.id ? 'bg-white/[0.05]' : 'hover:bg-white/[0.02]'
                  ]"
                >
                  <div v-if="selectedMessage?.id === msg.id" class="absolute left-0 top-0 bottom-0 w-1" :style="{ backgroundColor: themeStore.accentColor }"></div>
                  <div class="flex justify-between items-center mb-2">
                    <span class="text-[9px] font-black uppercase px-2 py-0.5" 
                          :style="{ backgroundColor: getRiskColor(msg.risk) + '22', color: getRiskColor(msg.risk) }">
                      {{ msg.category }}
                    </span>
                    <span class="text-[8px] font-bold text-gray-600 uppercase">{{ msg.time }}</span>
                  </div>
                  <p class="text-xs font-black uppercase text-white truncate mb-1">{{ msg.sender }} → {{ msg.receiver }}</p>
                  <p class="text-[10px] text-gray-500 italic truncate uppercase">"{{ msg.preview }}"</p>
                </button>
              </div>
            </div>

            <div class="lg:col-span-8 flex flex-col border border-white/10 bg-black overflow-hidden relative">
              <div v-if="selectedMessage" class="flex flex-col h-full">
                <div class="p-6 border-b border-white/10 bg-white/[0.01] flex justify-between items-center shrink-0">
                  <div>
                    <h3 class="text-sm font-black uppercase tracking-widest text-white">Full_Context_Audit</h3>
                    <p class="text-[9px] font-bold text-gray-600 mt-1 uppercase">Channel_ID: {{ selectedMessage.channelId }}</p>
                  </div>
                  <div class="flex gap-2">
                    <button 
                      @click="handleAction('SHADOWBAN')"
                      class="px-4 py-2 border border-white/10 text-[9px] font-black uppercase hover:bg-white/5 transition-all active:scale-95"
                    >
                      Shadowban_User
                    </button>
                    <button 
                      @click="handleAction('ALLOW')"
                      class="px-4 py-2 text-black text-[9px] font-black uppercase transition-all active:scale-95 hover:brightness-110" 
                      :style="{ backgroundColor: themeStore.accentColor }"
                    >
                      Allow_Message
                    </button>
                  </div>
                </div>

                <div class="flex-1 p-8 overflow-y-auto custom-scrollbar space-y-6 bg-[url('https://www.transparenttextures.com/patterns/stardust.png')] opacity-80">
                  <div v-for="(chat, i) in selectedMessage.thread" :key="i" 
                       :class="['max-w-[80%] p-4 border', chat.sender === selectedMessage.sender ? 'border-red-500/30 bg-red-500/5 ml-auto text-right' : 'border-white/10 bg-white/[0.02] mr-auto']">
                    <p class="text-[8px] font-black uppercase mb-2" :class="chat.sender === selectedMessage.sender ? 'text-red-500' : 'text-gray-500'">{{ chat.sender }}</p>
                    <p class="text-sm font-bold text-gray-200 leading-relaxed uppercase">{{ chat.text }}</p>
                    
                    <div v-if="chat.triggered" class="mt-3 pt-3 border-t border-red-500/20">
                      <span class="text-[9px] font-black text-red-500 uppercase tracking-widest italic">!! TRIGGER: {{ chat.triggerType }} !!</span>
                    </div>
                  </div>
                </div>

                <div class="p-6 border-t border-white/10 bg-white/[0.02] shrink-0">
                  <p class="text-[10px] font-black text-gray-500 uppercase tracking-widest mb-4">AI_Detection_Score</p>
                  <div class="flex items-center gap-6">
                    <div class="flex-1">
                      <div class="flex justify-between text-[10px] font-bold mb-2 uppercase">
                        <span>Fraud_Probability</span>
                        <span :style="{ color: themeStore.accentColor }">{{ selectedMessage.fraudScore }}%</span>
                      </div>
                      <div class="h-1 w-full bg-white/5">
                        <div class="h-full" :style="{ width: selectedMessage.fraudScore + '%', backgroundColor: themeStore.accentColor }"></div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>

              <div v-else class="flex-1 flex flex-col items-center justify-center opacity-20">
                <div class="w-12 h-12 border border-dashed border-white/40 animate-spin mb-4" style="animation-duration: 8s;"></div>
                <p class="text-[10px] font-black uppercase tracking-[0.5em]">Awaiting_Signal_Select</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <transition name="toast">
      <div v-if="toast.show" 
           class="fixed bottom-8 right-8 z-[200] bg-black border-l-4 p-5 min-w-[340px] shadow-2xl flex items-center justify-between"
           :style="{ borderColor: toast.type === 'danger' ? '#ef4444' : themeStore.accentColor }">
        <div>
          <p class="text-[10px] font-black uppercase tracking-widest opacity-60">{{ toast.label }}</p>
          <p class="text-xs font-black mt-1 tracking-tight uppercase">{{ toast.message }}</p>
        </div>
        <div class="w-2 h-2 rounded-full" :class="toast.type === 'danger' ? 'bg-red-500' : 'bg-green-500'" :style="toast.type !== 'danger' ? { backgroundColor: themeStore.accentColor } : {}"></div>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const selectedMessage = ref(null)
const toast = ref({ show: false, message: '', label: '', type: 'success' })

const flaggedMessages = ref([
  { 
    id: 1, sender: 'Global_Talent_Bot', receiver: 'User_X99', time: '2m ago', 
    category: 'SCAM_RISK', risk: 'HIGH', preview: 'Please message us on WhatsApp +44...', 
    channelId: 'CH_882', fraudScore: 94,
    thread: [
      { sender: 'User_X99', text: 'Hello, I saw your job posting for the React role.' },
      { sender: 'Global_Talent_Bot', text: 'Hi! We are moving fast. Please message our manager on WhatsApp +44 7700 900123 to finalize your payment.', triggered: true, triggerType: 'OFF_PLATFORM_REDIRECT' }
    ]
  },
  { 
    id: 2, sender: 'Alpha_Tech_HR', receiver: 'John_Doe', time: '14m ago', 
    category: 'TOS_VIOLATION', risk: 'MED', preview: 'We can pay you under the table...', 
    channelId: 'CH_441', fraudScore: 68,
    thread: [
      { sender: 'Alpha_Tech_HR', text: 'We like your CV. Instead of platform billing, we can pay you under the table via crypto to avoid fees.', triggered: true, triggerType: 'FEE_AVOIDANCE' }
    ]
  }
])

const handleAction = (type: string) => {
  themeStore.initializeUplink()
  
  if (type === 'SHADOWBAN') {
    toast.value = {
      show: true,
      label: 'Security_Protocol',
      message: `USER_${selectedMessage.value.sender}_SHADOW_RESTRICTED`,
      type: 'danger'
    }
  } else {
    toast.value = {
      show: true,
      label: 'Manual_Override',
      message: 'MESSAGE_STREAM_RE_INJECTED_SUCCESSFULLY',
      type: 'success'
    }
  }

  // Simulate removal from queue
  const currentId = selectedMessage.value.id
  setTimeout(() => {
    flaggedMessages.value = flaggedMessages.value.filter(m => m.id !== currentId)
    selectedMessage.value = null
  }, 1000)

  setTimeout(() => {
    toast.value.show = false
  }, 4000)
}

const getRiskColor = (risk: string) => {
  if (risk === 'HIGH') return '#ef4444'
  if (risk === 'MED') return '#f59e0b'
  return themeStore.accentColor
}
</script>

<style scoped>
  @reference 'tailwindcss';
.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.1); }

.toast-enter-active, .toast-leave-active { transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275); }
.toast-enter-from, .toast-leave-to { transform: translateX(100%) scale(0.9); opacity: 0; }
</style>