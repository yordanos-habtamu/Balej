<template>
  <div class="h-screen w-full bg-[#020205] text-white font-mono flex overflow-hidden selection:bg-cyan-500/30">
    <admin-sidebar />

    <main 
      :class="[
        'flex-1 flex flex-col h-full transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)]',
        themeStore.terminalCollapsed ? 'pl-20' : 'pl-72'
      ]"
    >
      <div class="absolute inset-0 pointer-events-none opacity-10" 
           :style="{ background: `radial-gradient(circle at 50% 0%, ${themeStore.accentColor}, transparent 70%)` }">
      </div>

      <div class="flex-1 overflow-y-auto custom-scrollbar relative z-10 px-6 py-8 lg:px-12">
        
        <div v-if="fetching" class="flex flex-col justify-center items-center h-64 gap-4">
          <div class="animate-spin h-12 w-12 border-2 border-t-transparent" :style="{ borderColor: themeStore.accentColor, borderTopColor: 'transparent' }"></div>
          <p class="text-xs font-bold tracking-widest uppercase opacity-60">Initializing_Data_Stream...</p>
        </div>
        
        <div v-else-if="error" class="max-w-md mx-auto bg-red-500/10 border border-red-500/40 p-6 text-center">
          <p class="text-red-500 font-black uppercase tracking-tighter text-xl mb-2">ACCESS_DENIED</p>
          <p class="text-xs text-red-400/70 uppercase tracking-widest">{{ error.message }}</p>
        </div>

        <div v-else-if="user" class="max-w-6xl mx-auto">
          
          <div class="bg-black border border-white/10 p-8 lg:p-10 mb-8 relative overflow-hidden group">
            <div class="absolute inset-0 opacity-5 pointer-events-none bg-[url('https://www.transparenttextures.com/patterns/carbon-fibre.png')]"></div>
            
            <div class="flex flex-col xl:flex-row xl:items-end justify-between gap-8 relative z-10">
              <div class="flex flex-col md:flex-row items-center md:items-end gap-8">
                
                <div class="w-32 h-32 border-2 p-1 relative" :style="{ borderColor: themeStore.accentColor }">
                  <div class="absolute -top-2 -left-2 w-4 h-4 border-t-2 border-l-2" :style="{ borderColor: themeStore.accentColor }"></div>
                  <img 
                    :src="`https://ui-avatars.com/api/?name=${user.first_name}+${user.last_name}&background=0D0D0D&color=fff&size=256&bold=true`" 
                    alt="Profile" 
                    class="w-full h-full object-cover grayscale brightness-90 group-hover:grayscale-0 transition-all duration-500"
                  />
                </div>

                <div class="text-center md:text-left">
                  <div class="flex items-center justify-center md:justify-start gap-2 mb-2 text-gray-400">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                    </svg>
                    <span class="text-xs font-bold uppercase tracking-widest italic">System_Administrator</span>
                  </div>
                  <h1 class="text-4xl lg:text-6xl font-black italic uppercase tracking-tighter">
                    {{ user.first_name }}_{{ user.last_name }}
                  </h1>
                </div>
              </div>

              <div>
                <nuxt-link to="/admin/settings" class="px-8 py-4 border border-white/20 text-xs font-black uppercase tracking-widest hover:bg-white hover:text-black transition-all flex items-center gap-3">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                  </svg>
                  [ EDIT_PARAMETERS ]
                </nuxt-link>
              </div>
            </div>
          </div>

          <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <div class="lg:col-span-2 space-y-8">
                <div class="bg-white/[0.02] border border-white/5 p-8 lg:p-10 relative">
                    <div class="flex items-center gap-3 mb-8">
                      <div class="h-[1px] w-12 bg-white/20"></div>
                      <h2 class="text-xs font-bold tracking-widest uppercase text-gray-400">Core_Identity_Protocol</h2>
                    </div>

                    <div class="grid grid-cols-1 md:grid-cols-2 gap-y-10 gap-x-12">
                         <div>
                            <label class="block text-[11px] font-bold text-gray-500 uppercase tracking-widest mb-1">Comm_Uplink</label>
                            <p class="text-base text-white font-bold uppercase">{{ user.email }}</p>
                        </div>
                        <div>
                            <label class="block text-[11px] font-bold text-gray-500 uppercase tracking-widest mb-1">Direct_Freq</label>
                            <p class="text-base text-white font-bold uppercase">{{ user.phone_number }}</p>
                        </div>
                         <div>
                            <label class="block text-[11px] font-bold text-gray-500 uppercase tracking-widest mb-1">ID_Registry</label>
                            <p class="text-sm font-black bg-white/10 text-white inline-block px-3 py-1 border border-white/10">
                              ADMIN-{{ user.id }}
                            </p>
                        </div>
                        <div>
                            <label class="block text-[11px] font-bold text-gray-500 uppercase tracking-widest mb-1">Uplink_Status</label>
                            <div class="flex items-center gap-3 mt-1">
                              <span class="w-2.5 h-2.5 rounded-full animate-pulse bg-green-500"></span>
                              <span class="text-xs font-black text-green-500 uppercase">ACTIVE_CONNECTED</span>
                            </div>
                        </div>
                         <div class="md:col-span-2">
                            <label class="block text-[11px] font-bold text-gray-500 uppercase tracking-widest mb-1">Station_Coordinates</label>
                            <p class="text-base text-gray-300 uppercase leading-relaxed">{{ user.address }}</p>
                        </div>
                    </div>
                </div>
            </div>

            <div class="space-y-8">
                <div class="bg-black border border-white/10 p-8">
                    <h3 class="text-xs font-black tracking-widest uppercase text-white mb-8 border-b border-white/10 pb-4">SYSTEM_METRICS</h3>
                    <div class="space-y-6">
                        <div v-for="(val, label, index) in { 'Registry': '1,234', 'Active_Jobs': '56', 'Pending': '12' }" :key="label" class="flex justify-between items-center border-b border-white/5 pb-4">
                            <span class="text-[11px] text-gray-500 font-bold uppercase tracking-wider">{{ label }}</span>
                            <span class="text-xl font-black italic tracking-tighter" :class="index === 1 ? 'text-green-400' : index === 2 ? 'text-yellow-400' : 'text-white'">{{ val }}</span>
                        </div>
                    </div>
                </div>

                <div class="bg-black border border-white/10 p-8">
                     <h3 class="text-xs font-black tracking-widest uppercase text-white mb-8">MANAGEMENT_CONSOLE</h3>
                     <div class="space-y-4">
                         <button class="w-full text-left p-4 border border-white/5 bg-white/[0.02] hover:bg-white/10 hover:border-white/20 transition-all flex items-center group">
                            <div class="p-2 mr-4 border border-white/10 opacity-60 group-hover:opacity-100 text-blue-400">
                               <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
                               </svg>
                            </div>
                            <span class="text-xs font-black text-gray-400 group-hover:text-white uppercase tracking-widest">MANAGE_REGISTRY</span>
                         </button>
                         
                         <button class="w-full text-left p-4 border border-white/5 bg-white/[0.02] hover:bg-white/10 hover:border-white/20 transition-all flex items-center group">
                            <div class="p-2 mr-4 border border-white/10 opacity-60 group-hover:opacity-100 text-purple-400">
                               <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
                               </svg>
                            </div>
                            <span class="text-xs font-black text-gray-400 group-hover:text-white uppercase tracking-widest">FETCH_LOG_DATA</span>
                         </button>

                         <button class="w-full text-left p-4 border border-white/5 bg-white/[0.02] hover:bg-white/10 hover:border-white/20 transition-all flex items-center group">
                            <div class="p-2 mr-4 border border-white/10 opacity-60 group-hover:opacity-100 text-red-400">
                               <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                               </svg>
                            </div>
                            <span class="text-xs font-black text-gray-400 group-hover:text-white uppercase tracking-widest">SYSTEM_ALERTS</span>
                         </button>
                     </div>
                </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useCookie } from '#app'
import { jwtDecode } from 'jwt-decode'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const tokenCookie = useCookie('token')
const client = useUrqlClient()
const fetching = ref(true)
const error = ref<CombinedError | null>(null)
const user = ref<any>(null)
const userId = ref<string | null>(null)

// Decode token to get user ID
if (tokenCookie.value) {
  try {
    const decoded: any = jwtDecode(tokenCookie.value)
    userId.value = decoded.user_id || decoded.id || decoded.sub 
  } catch (e) {
    error.value = { message: "Invalid token" } as any
    fetching.value = false
  }
} else {
    fetching.value = false
    error.value = { message: "No token found" } as any
}

const UserQuery = gql`
  query GetUser($id: ID!) {
    user(id: $id) {
      id
      first_name
      last_name
      email
      phone_number
      address
      role
      is_active
    }
  }
`

const fetchUser = async () => {
    if (!userId.value) return
    fetching.value = true
    error.value = null
    
    try {
        const result = await client.query(UserQuery, { id: userId.value }).toPromise()
        if (result.error) {
            error.value = result.error
        } else {
            user.value = result.data?.user
        }
    } catch (err: any) {
        error.value = err
    } finally {
        fetching.value = false
    }
}

onMounted(() => {
    fetchUser()
})
</script>

<style scoped>
@reference 'tailwindcss';

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-track { background: transparent; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: #1f2937; }
.custom-scrollbar::-webkit-scrollbar-thumb:hover { background: v-bind('themeStore.accentColor'); }

/* Animation for theme coherence */
.bg-black, .bg-white\/\[0\.02\] {
  animation: nodeEnter 0.5s cubic-bezier(0.2, 1, 0.3, 1) forwards;
  opacity: 0;
}

@keyframes nodeEnter {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
</style>