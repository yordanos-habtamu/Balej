<template>
  <div class="min-h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col min-h-screen transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)]',
        themeStore.terminalCollapsed ? 'ml-20' : 'ml-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-10" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 p-6 lg:p-12">
        <div class="max-w-6xl mx-auto">
          
          <div class="mb-12 border-b border-white/5 pb-10">
            <div class="flex items-center gap-3 mb-4">
              <div class="w-2 h-2 animate-pulse" :style="{ backgroundColor: themeStore.accentColor }"></div>
              <span class="text-[10px] tracking-[0.5em] text-gray-500 uppercase italic">Central_Database // Client_Registry</span>
            </div>
            <h1 class="text-5xl font-black italic uppercase tracking-tighter">
              CLIENT_<span :style="{ color: themeStore.accentColor }">DIRECTORY</span>
            </h1>
          </div>

          <div class="mb-12 relative group">
            <div class="absolute -inset-0.5 bg-gradient-to-r from-transparent via-white/10 to-transparent opacity-30"></div>
            <div class="relative bg-black border border-white/10 flex items-center">
              <div class="pl-6 pr-4">
                <svg class="w-5 h-5" :style="{ color: themeStore.accentColor }" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
                </svg>
              </div>
              <input
                type="text"
                v-model="searchTerm"
                placeholder="EXECUTE_SEARCH: [NAME / EMAIL / ID]..."
                class="w-full bg-transparent py-6 text-xs tracking-widest uppercase focus:outline-none placeholder:text-gray-800"
              />
              <div class="absolute bottom-0 left-0 h-[1px] w-0 transition-all duration-500 group-focus-within:w-full" :style="{ backgroundColor: themeStore.accentColor }"></div>
            </div>
          </div>

          <div v-if="clientData.loading" class="flex flex-col items-center justify-center py-32 space-y-4">
            <div class="w-12 h-12 border-2 border-t-transparent animate-spin rounded-full" :style="{ borderColor: `${themeStore.accentColor} transparent ${themeStore.accentColor} ${themeStore.accentColor}` }"></div>
            <p class="text-[10px] tracking-[0.5em] text-gray-500 animate-pulse">DECRYPTING_CLIENT_DATA...</p>
          </div>

          <div v-else-if="clientData.error" class="bg-red-500/5 border border-red-500/20 p-10 text-center">
            <p class="text-red-500 text-xs font-bold tracking-widest uppercase">CRITICAL_ERROR: FAILED_TO_ACCESS_UPLINK</p>
            <button @click="loadUsers" class="mt-4 text-[10px] underline hover:text-white transition-colors">RETRY_CONNECTION</button>
          </div>

          <div v-else class="space-y-4">
            <div class="flex items-center justify-between mb-8 px-2">
              <span class="text-[9px] text-gray-500 tracking-[0.3em] uppercase italic">>> {{ filteredUsers.length }} Nodes_Identified</span>
              <div class="h-[1px] flex-1 mx-6 bg-white/5"></div>
            </div>

            <div v-if="filteredUsers.length === 0" class="text-center py-20 border border-dashed border-white/10">
              <p class="text-xs text-gray-600 tracking-widest">ZERO_MATCHES_FOUND_IN_NODE_ARRAY</p>
            </div>

            <div
              v-for="user in filteredUsers"
              :key="user.email"
              class="client-node group relative bg-white/[0.02] border border-white/5 p-6 flex items-center gap-8 transition-all hover:bg-white/[0.05] hover:border-white/20"
            >
              <div class="relative flex-shrink-0">
                <div class="w-16 h-16 grayscale group-hover:grayscale-0 transition-all duration-500 border border-white/10 p-1">
                  <img src="~/assets/painter.png" alt="Node_Img" class="w-full h-full object-cover" />
                </div>
                <div class="absolute -bottom-1 -right-1 w-3 h-3 border border-black" :class="user.role === 'admin' ? 'bg-cyan-500' : 'bg-green-500'"></div>
              </div>

              <div class="flex-1 min-w-0">
                <div class="flex items-center gap-4">
                  <h3 class="text-lg font-black italic tracking-tighter uppercase group-hover:text-cyan-400 transition-colors" 
                      :style="selectedUser === user.email ? { color: themeStore.accentColor } : {}">
                    {{ user.first_name }} {{ user.last_name }}
                  </h3>
                  <div class="h-[1px] w-8 bg-white/10"></div>
                  <span class="text-[8px] px-2 py-0.5 border border-white/10 text-gray-500 tracking-widest uppercase group-hover:border-cyan-500/50 group-hover:text-white">
                    {{ user.role || 'CLIENT' }}
                  </span>
                </div>
                <div class="flex items-center gap-6 mt-2">
                  <p class="text-[10px] text-gray-500 truncate lowercase font-mono tracking-wider">{{ user.email }}</p>
                  <p class="text-[10px] text-gray-400 uppercase hidden md:block tracking-widest opacity-40">LOC: {{ user.address || 'UNDEFINED' }}</p>
                </div>
              </div>

              <div class="flex gap-4 opacity-40 group-hover:opacity-100 transition-opacity">
                <NuxtLink to="/admin/chats">
                  <button class="px-6 py-2 border border-white/10 text-[10px] font-bold tracking-[0.2em] hover:bg-white hover:text-black transition-all">
                    OPEN_COMMS
                  </button>
                </NuxtLink>
                <NuxtLink :to="`/admin/clients/${user.first_name}`">
                  <button class="px-6 py-2 bg-white/5 border border-white/10 text-[10px] font-bold tracking-[0.2em] hover:border-cyan-500 transition-all"
                          :style="{ ':hover': { borderColor: themeStore.accentColor } }">
                    VIEW_PROJ
                  </button>
                </NuxtLink>
              </div>

              <div class="absolute bottom-0 left-0 h-[1px] w-0 group-hover:w-full transition-all duration-700" :style="{ backgroundColor: themeStore.accentColor }"></div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <div v-if="isOpen" class="fixed inset-0 bg-[#020205]/90 backdrop-blur-sm flex items-center justify-center z-50 p-4">
      <div class="bg-black border-2 p-12 max-w-md w-full relative overflow-hidden" :style="{ borderColor: themeStore.accentColor }">
        <div class="absolute top-0 right-0 p-2 text-[8px] text-gray-600">ID_992-X</div>
        <div class="w-16 h-16 mx-auto mb-8 border-2 flex items-center justify-center" :style="{ borderColor: themeStore.accentColor, color: themeStore.accentColor }">
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h2 class="text-3xl font-black italic uppercase tracking-tighter text-center mb-2">Node_Integrated</h2>
        <p class="text-[10px] text-gray-500 tracking-[0.4em] text-center uppercase">Successfully added to global registry.</p>
        <button @click="isOpen = false" 
                class="w-full mt-10 py-4 text-[10px] font-black tracking-[0.4em] uppercase transition-all"
                :style="{ backgroundColor: themeStore.accentColor, color: '#000' }">
          CLOSE_TERMINAL
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref, computed, onMounted } from 'vue'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { useThemeStore } from '@/stores/theme'
import { useUrqlClient } from '@/composables/useUrqlClient'

const themeStore = useThemeStore()
const client = useUrqlClient()
const searchTerm = ref('')
const isOpen = ref(false)
const selectedUser = ref('')

const USER_QUERY = gql`
  query {
    users {
      first_name
      last_name
      email
      phone_number
      address
      role
    }
  }
`

const clientData = reactive({
  users: [] as Array<any>,
  loading: false,
  error: null as CombinedError | null
})

const filteredUsers = computed(() => {
  if (!searchTerm.value) return clientData.users
  const lowerSearch = searchTerm.value.toLowerCase()
  return clientData.users.filter(user =>
    user.first_name?.toLowerCase().includes(lowerSearch) ||
    user.last_name?.toLowerCase().includes(lowerSearch) ||
    user.email?.toLowerCase().includes(lowerSearch)
  )
})

async function loadUsers() {
  clientData.loading = true
  clientData.error = null
  try {
    const response = await client.query(USER_QUERY, {}).toPromise()
    if (response.error) {
      clientData.error = response.error
    } else {
      clientData.users = response.data.users
    }
  } catch (err: any) {
    clientData.error = err
  } finally {
    clientData.loading = false
  }
}

onMounted(() => {
  loadUsers()
})
</script>

<style scoped>
@reference 'tailwindcss';

/* Custom Scrollbar */
.custom-scrollbar::-webkit-scrollbar {
  width: 4px;
}
.custom-scrollbar::-webkit-scrollbar-track {
  background: transparent;
}
.custom-scrollbar::-webkit-scrollbar-thumb {
  background: #1f2937;
}
.custom-scrollbar::-webkit-scrollbar-thumb:hover {
  background: v-bind('themeStore.accentColor');
}

.client-node {
  animation: slideIn 0.4s ease-out forwards;
}

@keyframes slideIn {
  from { opacity: 0; transform: translateX(-10px); }
  to { opacity: 1; transform: translateX(0); }
}

/* Staggered animation for list items */
.client-node:nth-child(1) { animation-delay: 0.05s; }
.client-node:nth-child(2) { animation-delay: 0.1s; }
.client-node:nth-child(3) { animation-delay: 0.15s; }
</style>