<template>
  <div class="min-h-screen bg-[#020205] text-white font-sans overflow-hidden flex selection:bg-cyan-500/30">
    
    <aside class="w-72 border-r border-white/5 bg-[#050508] flex flex-col shrink-0 z-20">
      <div class="p-8 border-b border-white/10">
        <div class="flex items-center gap-2 mb-1">
          <div class="w-1.5 h-1.5 bg-cyan-500 rounded-full animate-pulse"></div>
          <p class="text-[9px] font-mono text-cyan-500 tracking-[0.3em] uppercase italic">Admin_Protocol</p>
        </div>
        <h2 class="text-xl font-black italic uppercase tracking-tighter">System_Config</h2>
      </div>
      
      <nav class="flex-1 py-4 overflow-y-auto custom-scrollbar">
        <div v-for="group in menuGroups" :key="group.title" class="mb-6">
          <p class="px-8 text-[8px] font-mono text-gray-600 uppercase tracking-[0.4em] mb-3">{{ group.title }}</p>
          <button 
            v-for="tab in group.items" :key="tab.id"
            @click="activeTab = tab.id"
            :class="['w-full flex items-center px-8 py-3 gap-4 transition-all group', 
                     activeTab === tab.id ? 'bg-cyan-500/10 text-cyan-500 border-r-2 border-cyan-500' : 'text-gray-500 hover:text-gray-300 hover:bg-white/[0.02]']"
          >
            <span class="text-[10px] font-black uppercase tracking-[0.15em]">{{ tab.label }}</span>
          </button>
        </div>
      </nav>

      <div class="p-6 border-t border-white/5 bg-black/40">
        <button @click="triggerBackup" class="w-full py-4 border border-white/10 text-[9px] font-mono uppercase text-gray-400 hover:bg-white/5 hover:text-white transition-all">
          Master_Snapshot.bak
        </button>
      </div>
    </aside>

    <main class="flex-1 overflow-y-auto custom-scrollbar relative bg-[#020205]">
      <div class="fixed inset-0 bg-[radial-gradient(circle_at_50%_50%,rgba(0,243,255,0.01),transparent)] pointer-events-none"></div>

      <div class="relative z-10 p-16 max-w-5xl mx-auto">
        
        <section v-if="activeTab === 'theme'" class="space-y-12 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Visual_Aesthetics</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">UI & Neural Interface Sync</p>
          </header>

          <div class="bg-white/[0.02] border border-white/5 p-10">
             <h3 class="text-[10px] font-mono text-cyan-500 uppercase tracking-[0.4em] mb-8">>> Accent_Hex_Sync</h3>
             <div class="flex gap-4">
               <button v-for="color in themeColors" :key="color" @click="themeStore.setAccentColor(color)"
                 class="w-16 h-16 border border-white/10 flex items-center justify-center relative hover:scale-105 transition-all">
                 <div class="absolute inset-0 opacity-10" :style="{ backgroundColor: color }"></div>
                 <div class="w-5 h-5" :style="{ backgroundColor: color, boxShadow: `0 0 20px ${color}` }"
                      :class="themeStore.accentColor === color ? 'scale-125 border-2 border-white' : 'scale-75'"></div>
               </button>
             </div>
          </div>
        </section>

        <section v-if="activeTab === 'privacy'" class="space-y-8 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">GDPR_Vault</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">Data Privacy & Regulatory Compliance</p>
          </header>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="p-8 border border-white/5 bg-white/[0.01] space-y-6">
              <h3 class="text-[10px] font-mono text-white uppercase tracking-widest">Retention_Policy</h3>
              <div class="space-y-4">
                <p class="text-[9px] text-gray-500 uppercase leading-relaxed italic">Automatically purge candidate records after inactivity period.</p>
                <select class="w-full bg-black border border-white/10 p-4 text-[10px] font-mono uppercase text-white outline-none focus:border-cyan-500">
                  <option>6_MONTHS (EU_GDPR_STRICT)</option>
                  <option>12_MONTHS (GLOBAL_DEFAULT)</option>
                  <option>24_MONTHS (US_RECORDS_ACT)</option>
                </select>
              </div>
            </div>

            <div class="p-8 border border-white/5 bg-white/[0.01] space-y-6">
              <h3 class="text-[10px] font-mono text-white uppercase tracking-widest">DSAR_Automation</h3>
              <div class="flex items-center justify-between py-2 border-b border-white/5" v-for="opt in gdprOptions" :key="opt.id">
                <span class="text-[9px] font-mono text-gray-500 uppercase">{{ opt.label }}</span>
                <button @click="opt.state = !opt.state" :class="opt.state ? 'text-cyan-500' : 'text-gray-800'" class="text-[10px] font-mono font-bold">{{ opt.state ? '[ACTIVE]' : '[INACTIVE]' }}</button>
              </div>
            </div>
          </div>

          <div class="p-8 border border-white/5 bg-white/[0.02]">
            <h3 class="text-[10px] font-mono text-cyan-500 uppercase tracking-widest mb-4">Export_Governance</h3>
            <div class="flex flex-wrap gap-4">
              <button class="px-6 py-3 border border-white/10 text-[9px] font-mono uppercase text-gray-500 hover:text-white hover:bg-white/5 transition-all">Download_PII_Audit.json</button>
              <button class="px-6 py-3 border border-white/10 text-[9px] font-mono uppercase text-gray-500 hover:text-white hover:bg-white/5 transition-all">Generate_Compliance_Report</button>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'notifications'" class="space-y-8 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Comms_Relay</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">Themed Template Engine (Email/SMS)</p>
          </header>

          <div class="grid grid-cols-1 gap-4">
            <div v-for="template in emailTemplates" :key="template.name" class="p-8 bg-white/[0.02] border border-white/5 group hover:border-cyan-500/30 transition-all">
              <div class="flex justify-between items-start mb-6">
                <div>
                  <span class="px-2 py-0.5 border border-cyan-500/30 text-cyan-500 text-[8px] font-mono uppercase mr-3">{{ template.channel }}</span>
                  <span class="text-[8px] font-mono text-gray-600 uppercase">{{ template.trigger }}</span>
                </div>
                <div class="w-1.5 h-1.5 bg-green-500 rounded-full"></div>
              </div>
              <h4 class="text-xl font-black uppercase italic text-white mb-6 tracking-tight">{{ template.name }}</h4>
              <div class="flex gap-4">
                <button class="flex-1 py-3 bg-cyan-500 text-black text-[10px] font-black uppercase tracking-widest hover:brightness-110 active:scale-95 transition-all">
                  Edit_Protocol
                </button>
                <button class="flex-1 py-3 border border-cyan-500/50 text-cyan-500 text-[10px] font-black uppercase tracking-widest hover:bg-cyan-500/10 transition-all">
                  Test_Relay
                </button>
              </div>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'logs'" class="space-y-8 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8 flex justify-between items-end">
            <div>
              <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Audit_Logs</h1>
              <p class="text-xs font-mono text-gray-500 uppercase tracking-widest">Forensic Activity Stream</p>
            </div>
          </header>

          <div class="bg-black/80 border border-white/10 p-8 h-[500px] overflow-y-auto font-mono text-[10px] space-y-3 custom-scrollbar" ref="logContainer">
            <div v-for="(log, idx) in systemLogs" :key="idx" class="flex gap-4 group py-1 border-b border-white/[0.02]">
              <span class="text-gray-700 shrink-0">[{{ log.time }}]</span>
              <span :class="getLogColor(log.type)" class="shrink-0 w-20">[{{ log.type.toUpperCase() }}]</span>
              <span class="text-gray-400 italic">{{ log.message }}</span>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'general'" class="space-y-12 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Core_Identity</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">Global Site Configuration</p>
          </header>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="space-y-6">
              <div class="space-y-2">
                <label class="text-[9px] font-mono text-cyan-500 uppercase tracking-widest">Site_Title</label>
                <input v-model="generalSettings.siteTitle" type="text" class="w-full bg-white/5 border border-white/10 p-4 text-xs font-mono uppercase outline-none focus:border-cyan-500 transition-all">
              </div>
              <div class="space-y-2">
                <label class="text-[9px] font-mono text-cyan-500 uppercase tracking-widest">Support_Uplink_Email</label>
                <input v-model="generalSettings.supportEmail" type="email" class="w-full bg-white/5 border border-white/10 p-4 text-xs font-mono outline-none focus:border-cyan-500 transition-all">
              </div>
            </div>
            <div class="space-y-6">
              <div class="p-8 border border-white/5 bg-white/[0.01] flex flex-col items-center justify-center gap-4">
                <div class="w-20 h-20 border border-dashed border-white/20 flex items-center justify-center text-gray-600 text-[10px] font-mono uppercase">Logo_Slot</div>
                <button class="text-[9px] font-mono text-cyan-500 uppercase hover:underline">Upload_New_Vector</button>
              </div>
              <div class="flex items-center justify-between p-4 border border-white/5 bg-white/[0.01]">
                <span class="text-[9px] font-mono text-gray-500 uppercase">Maintenance_Mode</span>
                <button @click="generalSettings.maintenanceMode = !generalSettings.maintenanceMode" 
                        :class="generalSettings.maintenanceMode ? 'text-red-500' : 'text-gray-700'" class="text-[10px] font-mono font-black">
                  {{ generalSettings.maintenanceMode ? '[OFFLINE]' : '[ONLINE]' }}
                </button>
              </div>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'job-logic'" class="space-y-12 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Market_Logic</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">Job Board Algorithms & Rules</p>
          </header>

          <div class="space-y-8">
            <div class="p-8 border border-white/5 bg-white/[0.01] flex items-center justify-between">
              <div>
                <h4 class="text-sm font-black uppercase italic mb-1">Auto_Approve_Listings</h4>
                <p class="text-[9px] text-gray-600 uppercase font-mono">Bypass manual moderation for verified clients</p>
              </div>
              <button @click="jobSettings.autoApprove = !jobSettings.autoApprove" 
                      :class="jobSettings.autoApprove ? 'bg-cyan-500 text-black' : 'bg-white/5 text-gray-500'" 
                      class="px-6 py-2 text-[9px] font-black uppercase tracking-widest transition-all">
                {{ jobSettings.autoApprove ? 'ENABLED' : 'DISABLED' }}
              </button>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
              <div class="p-8 border border-white/5 bg-white/[0.01] space-y-4">
                <label class="text-[9px] font-mono text-cyan-500 uppercase tracking-widest">Default_Listing_Duration</label>
                <div class="flex items-center gap-4">
                  <input v-model="jobSettings.defaultDuration" type="number" class="w-24 bg-black border border-white/10 p-3 text-xs font-mono outline-none focus:border-cyan-500">
                  <span class="text-[9px] font-mono text-gray-600 uppercase">Solar_Days</span>
                </div>
              </div>
              <div class="p-8 border border-white/5 bg-white/[0.01] space-y-4">
                <label class="text-[9px] font-mono text-cyan-500 uppercase tracking-widest">Featured_Boost_Multiplier</label>
                <select v-model="jobSettings.boostMultiplier" class="w-full bg-black border border-white/10 p-3 text-xs font-mono uppercase outline-none focus:border-cyan-500">
                  <option value="1.5">1.5x Visibility</option>
                  <option value="2.0">2.0x Visibility</option>
                  <option value="5.0">5.0x Visibility (Extreme)</option>
                </select>
              </div>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'payments'" class="space-y-12 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Fiscal_Gateways</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">Revenue & Transaction Protocols</p>
          </header>

          <div class="space-y-6">
            <div v-for="gateway in paymentGateways" :key="gateway.id" class="p-8 border border-white/5 bg-white/[0.01] group hover:border-cyan-500/20 transition-all">
              <div class="flex justify-between items-center mb-8">
                <div class="flex items-center gap-4">
                  <div class="w-10 h-10 bg-white/5 flex items-center justify-center font-black italic text-lg">{{ gateway.name[0] }}</div>
                  <h3 class="text-xl font-black uppercase italic tracking-tight">{{ gateway.name }}</h3>
                </div>
                <button @click="gateway.enabled = !gateway.enabled" 
                        :class="gateway.enabled ? 'text-cyan-500' : 'text-gray-800'" class="text-[10px] font-mono font-black">
                  {{ gateway.enabled ? '[ACTIVE]' : '[INACTIVE]' }}
                </button>
              </div>
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6" v-if="gateway.enabled">
                <div class="space-y-2">
                  <label class="text-[8px] font-mono text-gray-600 uppercase tracking-widest">Public_Key</label>
                  <input type="password" value="pk_test_************************" class="w-full bg-black border border-white/10 p-3 text-[10px] font-mono outline-none focus:border-cyan-500">
                </div>
                <div class="space-y-2">
                  <label class="text-[8px] font-mono text-gray-600 uppercase tracking-widest">Secret_Hash</label>
                  <input type="password" value="sk_test_************************" class="w-full bg-black border border-white/10 p-3 text-[10px] font-mono outline-none focus:border-cyan-500">
                </div>
              </div>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'security'" class="space-y-12 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Defense_Grid</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">System Hardening & Access Control</p>
          </header>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
            <div class="p-8 border border-white/5 bg-white/[0.01] space-y-6">
              <h3 class="text-[10px] font-mono text-cyan-500 uppercase tracking-widest">Multi_Factor_Auth</h3>
              <p class="text-[9px] text-gray-500 uppercase font-mono leading-relaxed">Enforce hardware-level 2FA for all administrative nodes.</p>
              <button @click="securitySettings.enforce2FA = !securitySettings.enforce2FA" 
                      class="w-full py-4 border border-cyan-500/30 text-cyan-500 text-[10px] font-black uppercase tracking-widest hover:bg-cyan-500/10 transition-all">
                {{ securitySettings.enforce2FA ? 'ENFORCEMENT_ACTIVE' : 'ENABLE_ENFORCEMENT' }}
              </button>
            </div>

            <div class="p-8 border border-white/5 bg-white/[0.01] space-y-6">
              <h3 class="text-[10px] font-mono text-cyan-500 uppercase tracking-widest">Session_Persistence</h3>
              <div class="space-y-4">
                <label class="text-[8px] text-gray-600 uppercase">Auto_Logout_Threshold (Minutes)</label>
                <input v-model="securitySettings.sessionTimeout" type="number" class="w-full bg-black border border-white/10 p-4 text-xs font-mono outline-none focus:border-cyan-500">
              </div>
            </div>

            <div class="md:col-span-2 p-8 border border-white/5 bg-white/[0.01] space-y-6">
              <h3 class="text-[10px] font-mono text-cyan-500 uppercase tracking-widest">IP_Whitelisting</h3>
              <div class="flex gap-4">
                <input type="text" placeholder="0.0.0.0" class="flex-1 bg-black border border-white/10 p-4 text-xs font-mono outline-none focus:border-cyan-500">
                <button class="px-8 bg-white text-black text-[10px] font-black uppercase tracking-widest hover:bg-cyan-500 transition-all">Authorize_IP</button>
              </div>
              <div class="flex flex-wrap gap-2">
                <span v-for="ip in securitySettings.whitelistedIPs" :key="ip" class="px-3 py-1 bg-white/5 text-[9px] font-mono text-gray-500 uppercase flex items-center gap-2">
                  {{ ip }} <button class="hover:text-red-500">✕</button>
                </span>
              </div>
            </div>
          </div>
        </section>

        <section v-if="activeTab === 'webhooks'" class="space-y-12 animate-in fade-in duration-500">
          <header class="border-b border-white/5 pb-8">
            <h1 class="text-6xl font-black italic uppercase tracking-tighter mb-4 text-cyan-500">Neural_Outbound</h1>
            <p class="text-xs font-mono text-gray-500 uppercase tracking-widest italic">External System Synchronization</p>
          </header>

          <div class="space-y-6">
            <div v-for="webhook in webhooks" :key="webhook.id" class="p-8 border border-white/5 bg-white/[0.01] flex justify-between items-center group">
              <div>
                <h4 class="text-lg font-black uppercase italic mb-1">{{ webhook.target }}</h4>
                <p class="text-[9px] text-gray-600 font-mono uppercase">{{ webhook.url }}</p>
                <div class="flex gap-4 mt-4">
                  <span v-for="event in webhook.events" :key="event" class="text-[7px] font-mono text-cyan-500/50 uppercase border border-cyan-500/20 px-2 py-0.5">{{ event }}</span>
                </div>
              </div>
              <div class="flex gap-4">
                <button class="p-3 border border-white/10 hover:border-cyan-500 transition-all group-hover:bg-white/5">
                  <span class="text-[10px] font-mono text-gray-500 group-hover:text-white uppercase">Ping</span>
                </button>
                <button class="p-3 border border-white/10 hover:border-red-500 transition-all group-hover:bg-red-500/10">
                  <span class="text-[10px] font-mono text-gray-500 group-hover:text-red-500 uppercase">Delete</span>
                </button>
              </div>
            </div>
            <button class="w-full py-8 border border-dashed border-white/10 text-[10px] font-black uppercase tracking-[0.5em] text-gray-600 hover:border-cyan-500 hover:text-cyan-500 transition-all bg-white/[0.01]">
              Initialize_New_Uplink
            </button>
          </div>
        </section>

      </div>
    </main>

    <transition name="toast">
      <div v-if="showToast" class="fixed bottom-10 right-10 bg-black border border-cyan-500 px-10 py-5 z-[100] shadow-[0_0_40px_rgba(0,243,255,0.1)]">
        <p class="text-[10px] font-mono text-cyan-500 uppercase tracking-[0.2em] font-black">{{ toastMsg }}</p>
      </div>
    </transition>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, nextTick } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const activeTab = ref('theme')
const showToast = ref(false)
const toastMsg = ref('')
const logContainer = ref<HTMLElement | null>(null)

const themeColors = ['#00f3ff', '#ff0055', '#7000ff', '#00ff66', '#ffaa00']
const passFields = ref({ old: '', new: '' })

const menuGroups = [
  {
    title: 'Core_Infrastructure',
    items: [
      { id: 'general', label: 'Core Identity' },
      { id: 'theme', label: 'Neural Aesthetics' },
      { id: 'security', label: 'Defense Grid' },
      { id: 'logs', label: 'Forensic Audit' }
    ]
  },
  {
    title: 'Recruitment_Ops',
    items: [
      { id: 'job-logic', label: 'Market Logic' },
      { id: 'notifications', label: 'Comms Templates' },
      { id: 'payments', label: 'Fiscal Gateways' },
      { id: 'privacy', label: 'GDPR & Compliance' },
      { id: 'webhooks', label: 'Neural Outbound' }
    ]
  }
]

const generalSettings = ref({
  siteTitle: 'Neural_Talent_Mesh',
  supportEmail: 'uplink@neural.io',
  maintenanceMode: false
})

const jobSettings = ref({
  autoApprove: false,
  defaultDuration: 30,
  boostMultiplier: '2.0'
})

const securitySettings = ref({
  enforce2FA: true,
  sessionTimeout: 60,
  whitelistedIPs: ['127.0.0.1', '192.168.1.1']
})

const paymentGateways = ref([
  { id: 1, name: 'Stripe_Protocol', enabled: true },
  { id: 2, name: 'PayPal_Relay', enabled: false },
  { id: 3, name: 'Crypto_Uplink', enabled: true }
])

const webhooks = ref([
  { id: 1, target: 'Slack_Sync', url: 'https://hooks.slack.com/services/...', events: ['JOB_POSTED', 'NEW_TALENT'] },
  { id: 2, target: 'Discord_Relay', url: 'https://discord.com/api/webhooks/...', events: ['SYSTEM_ALERT'] }
])

const gdprOptions = ref([
  { id: 1, label: 'Right_to_be_Forgotten_Portal', state: true },
  { id: 2, label: 'Auto_Anonymize_Archived_Data', state: false },
  { id: 3, label: 'Data_Portability_Self_Service', state: true }
])

const emailTemplates = [
  { name: 'Application_Acknowledgement', trigger: 'ENTRY_SUBMIT', channel: 'EMAIL' },
  { name: 'Interview_Invitation_V4', trigger: 'STAGE_ADVANCE', channel: 'EMAIL' },
  { name: 'SMS_Interview_Reminder', trigger: 'T_MINUS_1_HOUR', channel: 'SMS' },
  { name: 'Offer_Letter_Master', trigger: 'OFFER_INITIATED', channel: 'EMAIL' }
]

const systemLogs = ref([
  { time: '10:00:00', type: 'system', message: 'Hypervisor core established.' },
  { time: '10:02:45', type: 'privacy', message: 'Retention purge: 1,204 records cleared.' }
])

const getLogColor = (type: string) => {
  switch(type) {
    case 'error': return 'text-red-500';
    case 'privacy': return 'text-amber-500';
    default: return 'text-cyan-500';
  }
}

const triggerToast = (msg: string) => {
  toastMsg.value = msg;
  showToast.value = true;
  setTimeout(() => showToast.value = false, 3000);
}

const triggerBackup = () => {
  triggerToast('SNAPSHOT_GENERATED');
}

onMounted(() => {
  setInterval(() => {
    const events = ['DDoS Filter: Passive', 'GDPR Node: Synced', 'SMTP Handshake: Success', 'API Payload: Received'];
    systemLogs.value.push({
      time: new Date().toLocaleTimeString('en-GB'),
      type: 'info',
      message: events[Math.floor(Math.random() * events.length)]
    });
    if(systemLogs.value.length > 30) systemLogs.value.shift();
    nextTick(() => { if(logContainer.value) logContainer.value.scrollTop = logContainer.value.scrollHeight });
  }, 7000);
})
</script>

<style scoped>
  @reference 'tailwindcss';
:deep(.text-cyan-500) { color: v-bind('themeStore.accentColor'); }
:deep(.bg-cyan-500) { background-color: v-bind('themeStore.accentColor'); }
:deep(.border-cyan-500) { border-color: v-bind('themeStore.accentColor'); }
:deep(.decoration-cyan-500\/30) { text-decoration-color: v-bind('themeStore.accentColor + "4D"'); }

.custom-scrollbar::-webkit-scrollbar { width: 3px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.05); }

.toast-enter-active, .toast-leave-active { transition: all 0.5s cubic-bezier(0.19, 1, 0.22, 1); }
.toast-enter-from { transform: translateY(100px); opacity: 0; }
.toast-leave-to { transform: translateX(100px); opacity: 0; }

section { animation: section-fade 0.5s ease-out; }
@keyframes section-fade {
  from { opacity: 0; transform: translateY(15px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>