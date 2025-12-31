<template>
  <div class="min-h-screen bg-[#020205] text-white font-sans flex overflow-hidden selection:bg-cyan-500/30">
    
    <aside 
      :class="[
        'flex-shrink-0 z-50 transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)]',
        themeStore.terminalCollapsed ? 'w-20' : 'w-72'
      ]"
    >
      <admin-sidebar />
    </aside>

    <main class="flex-1 overflow-y-auto custom-scrollbar relative">
      <div class="fixed inset-0 bg-[radial-gradient(circle_at_50%_50%,rgba(0,243,255,0.015),transparent)] pointer-events-none"></div>

      <div class="relative z-10 p-8 lg:p-16 max-w-[1600px] mx-auto">
        
        <div class="mb-16 flex flex-col md:flex-row justify-between items-start md:items-end gap-6 border-b border-white/5 pb-12">
          <div>
            <div class="flex items-center gap-3 mb-4">
              <div class="w-2 h-2 bg-cyan-500 animate-pulse"></div>
              <span class="text-[10px] font-mono tracking-[0.5em] text-cyan-500 uppercase italic">Core_Identity // Auth_Registry</span>
            </div>
            <h1 class="text-6xl font-black italic uppercase tracking-tighter">User_<span class="text-cyan-500">Management</span></h1>
          </div>
          <div class="flex flex-col items-end gap-2">
             <span class="text-[9px] font-mono text-gray-600 uppercase">Active_Nodes: <span class="text-white">1,248</span></span>
             <button class="px-8 py-3 bg-white text-black font-black uppercase text-[10px] tracking-[0.2em] hover:bg-cyan-500 transition-all flex items-center gap-3">
              <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M12 4v16m8-8H4" stroke-width="3"/></svg>
              Initialize_New_User
            </button>
          </div>
        </div>

        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-1px bg-white/5 border border-white/10 mb-16 shadow-[0_0_50px_rgba(0,0,0,0.5)]">
          <div v-for="stat in stats" :key="stat.label" class="bg-black p-8 group hover:bg-white/[0.02] transition-colors">
            <p class="text-[9px] font-mono text-gray-500 uppercase tracking-[0.3em] mb-4">{{ stat.label }}</p>
            <div class="flex items-baseline gap-4">
              <span class="text-4xl font-black italic" :class="stat.colorClass">{{ stat.value }}</span>
              <span class="text-[8px] font-mono text-green-500/50">+{{ stat.growth }}%</span>
            </div>
          </div>
        </div>

        <div class="bg-black border border-white/10 p-2 mb-1px flex flex-col lg:flex-row gap-2">
          <div class="flex flex-1 gap-2">
            <div class="relative flex-1">
              <input 
                type="text" 
                v-model="searchTerm" 
                placeholder="Search_Identity_Index..." 
                class="w-full bg-white/[0.03] border border-white/5 p-4 text-[10px] font-mono uppercase tracking-widest focus:border-cyan-500/50 outline-none transition-all placeholder:text-gray-700"
              />
              <span class="absolute right-4 top-1/2 -translate-y-1/2 text-[8px] font-mono text-gray-700 italic">BASH_CMD</span>
            </div>
            <select v-model="roleFilter" class="bg-black border border-white/10 px-6 text-[9px] font-mono uppercase tracking-widest text-gray-400 outline-none focus:text-cyan-500">
              <option value="all">Role_All</option>
              <option value="admin">Admin</option>
              <option value="client">Client</option>
              <option value="talent">Talent</option>
            </select>
          </div>
          <div class="flex gap-2">
            <select v-model="statusFilter" class="bg-black border border-white/10 px-6 py-4 lg:py-0 text-[9px] font-mono uppercase tracking-widest text-gray-400 outline-none focus:text-cyan-500">
              <option value="all">Status_All</option>
              <option value="active">Operational</option>
              <option value="suspended">Locked</option>
            </select>
            <button class="px-6 bg-white/[0.03] border border-white/5 text-gray-500 hover:text-white transition-colors">
              <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M4 4h16v2H4V4zm2 5h12v2H6V9zm3 5h6v2H9v-2z" stroke-width="2"/></svg>
            </button>
          </div>
        </div>

        <div class="bg-black border border-white/10 overflow-hidden shadow-2xl">
          <div class="overflow-x-auto">
            <table class="w-full border-collapse">
              <thead>
                <tr class="border-b border-white/10 bg-white/[0.01]">
                  <th class="px-8 py-5 text-left text-[9px] font-mono text-cyan-500 uppercase tracking-[0.3em]">Identity_Node</th>
                  <th class="px-8 py-5 text-left text-[9px] font-mono text-cyan-500 uppercase tracking-[0.3em]">Access_Level</th>
                  <th class="px-8 py-5 text-left text-[9px] font-mono text-cyan-500 uppercase tracking-[0.3em]">Operational_State</th>
                  <th class="px-8 py-5 text-left text-[9px] font-mono text-cyan-500 uppercase tracking-[0.3em]">Node_History</th>
                  <th class="px-8 py-5 text-right text-[9px] font-mono text-cyan-500 uppercase tracking-[0.3em]">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-white/[0.05]">
                <tr v-for="user in filteredUsers" :key="user.id" class="group hover:bg-cyan-500/[0.02] transition-colors">
                  <td class="px-8 py-6">
                    <div class="flex items-center gap-4">
                      <div class="w-10 h-10 border border-white/10 p-0.5 group-hover:border-cyan-500/50 transition-colors">
                        <img :src="user.avatar" class="w-full h-full object-cover grayscale group-hover:grayscale-0 transition-all" />
                      </div>
                      <div>
                        <p class="text-xs font-black uppercase italic tracking-wider">{{ user.name }}</p>
                        <p class="text-[9px] font-mono text-gray-600 lowercase tracking-tighter">{{ user.email }}</p>
                      </div>
                    </div>
                  </td>
                  <td class="px-8 py-6">
                    <span :class="getRoleStyles(user.role)" class="text-[9px] font-mono font-bold px-3 py-1 border">
                      {{ user.roleLabel }}
                    </span>
                  </td>
                  <td class="px-8 py-6">
                    <div class="flex items-center gap-2">
                      <div class="w-1 h-1 rounded-full" :class="user.status === 'active' ? 'bg-green-500 shadow-[0_0_8px_#22c55e]' : 'bg-red-500'"></div>
                      <span class="text-[9px] font-mono uppercase tracking-widest" :class="user.status === 'active' ? 'text-green-500' : 'text-red-500'">
                        {{ user.status === 'active' ? 'Operational' : 'Access_Locked' }}
                      </span>
                    </div>
                  </td>
                  <td class="px-8 py-6">
                    <div class="text-[9px] font-mono space-y-1">
                      <p class="text-gray-500">JOINED: {{ user.joined }}</p>
                      <p class="text-cyan-500/50">LAST_PING: {{ user.lastActive }}</p>
                    </div>
                  </td>
                  <td class="px-8 py-6 text-right">
                    <div class="flex justify-end gap-2 opacity-30 group-hover:opacity-100 transition-opacity">
                      <button class="p-2 border border-white/10 hover:border-cyan-500 hover:text-cyan-500 transition-all">
                        <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" stroke-width="2"/></svg>
                      </button>
                      <button @click="toggleLock(user)" class="p-2 border border-white/10 hover:border-red-500 hover:text-red-500 transition-all">
                        <svg class="w-3 h-3" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" stroke-width="2"/></svg>
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
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()

const roleFilter = ref('all')
const statusFilter = ref('all')
const searchTerm = ref('')

const stats = [
  { label: 'Total_Identities', value: '1,248', growth: '12', colorClass: 'text-white' },
  { label: 'Admin_Authorities', value: '12', growth: '2', colorClass: 'text-cyan-500' },
  { label: 'Corporate_Clients', value: '78', growth: '5', colorClass: 'text-green-500' },
  { label: 'Talent_Nodes', value: '1,158', growth: '24', colorClass: 'text-purple-500' }
]

const users = ref([
  { id: 1, name: 'John Wick', email: 'john@continental.com', role: 'admin', roleLabel: 'ROOT_ADMIN', status: 'active', joined: 'JAN_2024', lastActive: '2M_AGO', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=John' },
  { id: 2, name: 'Sarah Connor', email: 'sarah@skynet.com', role: 'client', roleLabel: 'CLIENT_PRIME', status: 'active', joined: 'MAR_2024', lastActive: '1D_AGO', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Sarah' },
  { id: 3, name: 'Neo Anderson', email: 'neo@matrix.io', role: 'talent', roleLabel: 'TALENT_DEV', status: 'active', joined: 'JUN_2025', lastActive: 'SYNC_NOW', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Neo' },
  { id: 4, name: 'Major Motoko', email: 'kusanagi@section9.jp', role: 'talent', roleLabel: 'TALENT_SEC', status: 'suspended', joined: 'FEB_2024', lastActive: '3W_AGO', avatar: 'https://api.dicebear.com/7.x/avataaars/svg?seed=Motoko' }
])

const filteredUsers = computed(() => {
  return users.value.filter(u => {
    const searchMatch = u.name.toLowerCase().includes(searchTerm.value.toLowerCase()) || 
                        u.email.toLowerCase().includes(searchTerm.value.toLowerCase())
    const roleMatch = roleFilter.value === 'all' || u.role === roleFilter.value
    const statusMatch = statusFilter.value === 'all' || u.status === statusFilter.value
    return searchMatch && roleMatch && statusMatch
  })
})

const getRoleStyles = (role: string) => {
  switch(role) {
    case 'admin': return 'border-cyan-500/30 text-cyan-500 bg-cyan-500/5'
    case 'client': return 'border-green-500/30 text-green-500 bg-green-500/5'
    default: return 'border-purple-500/30 text-purple-500 bg-purple-500/5'
  }
}

const toggleLock = (user: any) => {
  user.status = user.status === 'active' ? 'suspended' : 'active'
}
</script>

<style scoped>
  @reference 'tailwindcss';
/* Theme Engine Overrides */
:deep(.text-cyan-500) { color: v-bind('themeStore.accentColor'); }
:deep(.bg-cyan-500) { background-color: v-bind('themeStore.accentColor'); }
:deep(.border-cyan-500) { border-color: v-bind('themeStore.accentColor'); }

.custom-scrollbar::-webkit-scrollbar { width: 4px; }
.custom-scrollbar::-webkit-scrollbar-thumb { background: rgba(255, 255, 255, 0.05); border-radius: 10px; }

input, select {
  -webkit-appearance: none;
  border-radius: 0;
}
</style>