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
          <h1 class="text-4xl lg:text-5xl font-bold mb-4">Role & Permission Management</h1>
          <p class="text-xl text-gray-400">Define access levels and permissions for different user roles in the platform</p>
        </div>

        <!-- Role Stats -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Total Roles</p>
              <svg class="w-8 h-8 text-purple-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-purple-400">5</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Active Admins</p>
              <svg class="w-8 h-8 text-[color:var(--primary-color)]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-[color:var(--primary-color)]">12</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Custom Roles</p>
              <svg class="w-8 h-8 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-green-400">2</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Last Updated</p>
              <svg class="w-8 h-8 text-blue-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </div>
            <p class="text-4xl font-bold text-blue-400">2d ago</p>
          </div>
        </div>

        <!-- Role Tabs -->
        <div class="mb-10">
          <div class="flex flex-wrap gap-4 border-b border-gray-700">
            <button
              v-for="role in roles"
              :key="role.id"
              @click="selectedRole = role"
              :class="[
                'px-8 py-4 font-semibold text-lg border-b-4 transition',
                selectedRole.id === role.id
                  ? 'border-[color:var(--primary-color)] text-[color:var(--primary-color)]'
                  : 'border-transparent text-gray-500 hover:text-gray-300'
              ]"
            >
              {{ role.name }} ({{ role.userCount }})
            </button>
            <button class="ml-auto px-8 py-4 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-xl font-semibold transition flex items-center gap-3">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4v16m8-8H4" />
              </svg>
              Create New Role
            </button>
          </div>
        </div>

        <!-- Permissions Grid -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-10">
          <div class="flex items-center justify-between mb-8">
            <h2 class="text-2xl font-bold">Permissions for {{ selectedRole.name }}</h2>
            <button class="px-6 py-3 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-xl font-medium transition">
              Save Changes
            </button>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <div v-for="category in permissionCategories" :key="category.id" class="space-y-6">
              <h3 class="text-xl font-semibold text-gray-300 mb-4 flex items-center gap-3">
                <svg class="w-6 h-6 text-[color:var(--primary-color)]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="{{ category.icon }}" />
                </svg>
                {{ category.name }}
              </h3>
              <div class="space-y-4">
                <label v-for="perm in category.permissions" :key="perm.id" class="flex items-center justify-between cursor-pointer hover:bg-gray-700/50 rounded-lg p-3 transition">
                  <span class="text-gray-300">{{ perm.label }}</span>
                  <input
                    type="checkbox"
                    :checked="selectedRole.permissions.includes(perm.id)"
                    class="w-6 h-6 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]"
                  />
                </label>
              </div>
            </div>
          </div>
        </div>

        <!-- Danger Zone -->
        <div class="mt-12 bg-red-600/10 rounded-3xl border border-red-600/30 p-10">
          <h2 class="text-2xl font-bold text-red-400 mb-6">Danger Zone</h2>
          <p class="text-gray-300 mb-8">Once you delete a role, there is no going back. Please be certain.</p>
          <button class="px-8 py-4 bg-red-600 hover:bg-red-700 text-white rounded-xl font-semibold transition">
            Delete {{ selectedRole.name }} Role
          </button>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { useThemeStore } from '~/stores/theme'

const theme = useThemeStore()

// Mock roles data
const roles = ref([
  { id: 1, name: 'Super Administrator', userCount: 3, permissions: ['all'] },
  { id: 2, name: 'Administrator', userCount: 9, permissions: ['view_users', 'edit_users', 'manage_jobs', 'view_reports'] },
  { id: 3, name: 'Client Manager', userCount: 15, permissions: ['view_clients', 'manage_projects', 'chat_clients'] },
  { id: 4, name: 'Talent Coordinator', userCount: 8, permissions: ['view_talents', 'manage_applications'] },
  { id: 5, name: 'Viewer', userCount: 5, permissions: ['view_reports', 'view_logs'] }
])

const selectedRole = ref(roles.value[1])

const permissionCategories = ref([
  {
    id: 1,
    name: 'User Management',
    icon: 'M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z',
    permissions: [
      { id: 'view_users', label: 'View Users' },
      { id: 'edit_users', label: 'Edit Users' },
      { id: 'delete_users', label: 'Delete Users' },
      { id: 'manage_roles', label: 'Manage Roles & Permissions' }
    ]
  },
  {
    id: 2,
    name: 'Job & Talent',
    icon: 'M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z',
    permissions: [
      { id: 'manage_jobs', label: 'Create/Edit Jobs' },
      { id: 'view_applications', label: 'View Applications' },
      { id: 'manage_talents', label: 'Manage Talent Profiles' },
      { id: 'export_data', label: 'Export Talent Data' }
    ]
  },
  {
    id: 3,
    name: 'Client & Projects',
    icon: 'M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4',
    permissions: [
      { id: 'view_clients', label: 'View Clients' },
      { id: 'manage_projects', label: 'Create/Edit Projects' },
      { id: 'chat_clients', label: 'Message Clients' },
      { id: 'view_invoices', label: 'View Invoices' }
    ]
  },
  {
    id: 4,
    name: 'System & Reports',
    icon: 'M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z',
    permissions: [
      { id: 'view_reports', label: 'View Analytics & Reports' },
      { id: 'view_logs', label: 'View System Logs' },
      { id: 'manage_settings', label: 'Platform Settings' },
      { id: 'all', label: 'Full System Access' }
    ]
  }
])

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