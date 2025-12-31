<template>
  <div class="min-h-screen flex">
    <!-- Sidebar -->
    <aside class="flex-shrink-0 pt-8">
      <admin-sidebar />
    </aside>

    <!-- Main Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="p-6 md:p-8 lg:p-12 max-w-7xl mx-auto">
        <!-- Header -->
        <div class="mb-12">
          <h1 class="text-4xl lg:text-5xl font-bold mb-4">User Management</h1>
          <p class="text-xl text-gray-400">Manage administrators, clients, and talents across the platform</p>
        </div>

        <!-- Stats Overview -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Total Users</p>
              <svg class="w-8 h-8 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
              </svg>
            </div>
            <p class="text-4xl font-bold">1,248</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Admins</p>
              <svg class="w-8 h-8 text-[color:var(--primary-color)]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-[color:var(--primary-color)]">12</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Clients</p>
              <svg class="w-8 h-8 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-green-400">78</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Talents</p>
              <svg class="w-8 h-8 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-blue-400">1,158</p>
          </div>
        </div>

        <!-- Filters & Search -->
        <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6 mb-10">
          <div class="flex flex-col lg:flex-row gap-6 items-center justify-between">
            <div class="flex flex-wrap gap-4">
              <select v-model="roleFilter" class="px-5 py-3 bg-gray-700 border border-gray-600 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]">
                <option value="all">All Roles</option>
                <option value="admin">Admin</option>
                <option value="client">Client</option>
                <option value="talent">Talent</option>
              </select>
              <select v-model="statusFilter" class="px-5 py-3 bg-gray-700 border border-gray-600 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]">
                <option value="all">All Status</option>
                <option value="active">Active</option>
                <option value="inactive">Inactive</option>
                <option value="suspended">Suspended</option>
              </select>
            </div>
            <div class="relative w-full lg:w-auto">
              <input
                type="text"
                v-model="searchTerm"
                placeholder="Search users..."
                class="w-full lg:w-96 px-6 py-3 pl-12 bg-gray-700 border border-gray-600 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)] placeholder-gray-500"
              />
              <svg class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>
        </div>

        <!-- Users Table -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 overflow-hidden">
          <div class="p-6 border-b border-gray-700 flex items-center justify-between">
            <h2 class="text-2xl font-bold">All Users</h2>
            <button class="px-6 py-3 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-xl font-medium transition flex items-center gap-3">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              Add User
            </button>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-700/50">
                <tr>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">User</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Role</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Status</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Joined</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Last Active</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Actions</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-700">
                <tr v-for="user in filteredUsers" :key="user.id" class="hover:bg-gray-700/30 transition">
                  <td class="px-6 py-5">
                    <div class="flex items-center gap-4">
                      <div class="w-12 h-12 rounded-full overflow-hidden ring-2 ring-gray-600">
                        <img :src="user.avatar" alt="User" class="w-full h-full object-cover" />
                      </div>
                      <div>
                        <p class="font-medium text-white">{{ user.name }}</p>
                        <p class="text-sm text-gray-400">{{ user.email }}</p>
                      </div>
                    </div>
                  </td>
                  <td class="px-6 py-5">
                    <span :class="[
                      'px-4 py-2 rounded-full text-xs font-medium',
                      user.role === 'admin' ? 'bg-purple-600/20 text-purple-400' :
                      user.role === 'client' ? 'bg-green-600/20 text-green-400' :
                      'bg-blue-600/20 text-blue-400'
                    ]">
                      {{ user.roleLabel }}
                    </span>
                  </td>
                  <td class="px-6 py-5">
                    <span :class="[
                      'px-4 py-2 rounded-full text-xs font-medium',
                      user.status === 'active' ? 'bg-green-600/20 text-green-400' :
                      user.status === 'inactive' ? 'bg-gray-600/20 text-gray-400' :
                      'bg-red-600/20 text-red-400'
                    ]">
                      {{ user.statusLabel }}
                    </span>
                  </td>
                  <td class="px-6 py-5 text-sm text-gray-300">{{ user.joined }}</td>
                  <td class="px-6 py-5 text-sm text-gray-300">{{ user.lastActive }}</td>
                  <td class="px-6 py-5">
                    <div class="flex items-center gap-3">
                      <button class="text-gray-400 hover:text-white transition">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                        </svg>
                      </button>
                      <button class="text-gray-400 hover:text-red-400 transition">
                        <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                        </svg>
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
import { ref, computed, onMounted, watch } from 'vue'
import { useThemeStore } from '~/stores/theme'

const theme = useThemeStore()

const roleFilter = ref('all')
const statusFilter = ref('all')
const searchTerm = ref('')

// Mock user data
const users = ref([
  { id: 1, name: 'John Doe', email: 'john@admin.com', role: 'admin', roleLabel: 'Administrator', status: 'active', statusLabel: 'Active', joined: 'Jan 15, 2024', lastActive: '2 hours ago', avatar: '~/assets/painter.png' },
  { id: 2, name: 'Sarah Chen', email: 'sarah@techcorp.com', role: 'client', roleLabel: 'Client', status: 'active', statusLabel: 'Active', joined: 'Mar 8, 2024', lastActive: '1 day ago', avatar: '~/assets/painter.png' },
  { id: 3, name: 'Mike Rivera', email: 'mike.rivera@example.com', role: 'talent', roleLabel: 'Talent', status: 'active', statusLabel: 'Active', joined: 'Jun 22, 2025', lastActive: '5 minutes ago', avatar: '~/assets/painter.png' },
  { id: 4, name: 'Emma Wilson', email: 'emma@globalinc.com', role: 'client', roleLabel: 'Client', status: 'inactive', statusLabel: 'Inactive', joined: 'Feb 12, 2024', lastActive: '3 weeks ago', avatar: '~/assets/painter.png' },
  { id: 5, name: 'Alex Kumar', email: 'alex.k@example.com', role: 'talent', roleLabel: 'Talent', status: 'active', statusLabel: 'Active', joined: 'Sep 1, 2025', lastActive: '30 minutes ago', avatar: '~/assets/painter.png' },
  { id: 6, name: 'Lisa Park', email: 'lisa@enterprise.com', role: 'client', roleLabel: 'Client', status: 'active', statusLabel: 'Active', joined: 'Apr 30, 2024', lastActive: '4 hours ago', avatar: '~/assets/painter.png' }
])

const filteredUsers = computed(() => {
  let filtered = users.value

  if (searchTerm.value) {
    const term = searchTerm.value.toLowerCase()
    filtered = filtered.filter(u => 
      u.name.toLowerCase().includes(term) || 
      u.email.toLowerCase().includes(term)
    )
  }

  if (roleFilter.value !== 'all') {
    filtered = filtered.filter(u => u.role === roleFilter.value)
  }

  if (statusFilter.value !== 'all') {
    filtered = filtered.filter(u => u.status === statusFilter.value)
  }

  return filtered
})

// Theme sync
onMounted(() => {
  const updateColor = () => {
    const colorMap: Record<string, string> = {
      blue: '#3b82f6',
      indigo: '#6366f1',
      purple: '#a78bfa',
      pink: '#ec4899',
      emerald: '#10b981',
      red: '#ef4444'
    }
    const color = colorMap[theme.primaryColor] || '#3b82f6'
    document.documentElement.style.setProperty('--primary-color', color)
  }
  updateColor()
  watch(() => theme.primaryColor, updateColor)
})
</script>