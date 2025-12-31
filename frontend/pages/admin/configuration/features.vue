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
          <h1 class="text-4xl lg:text-5xl font-bold mb-4">System Alerts & Logs</h1>
          <p class="text-xl text-gray-400">Monitor real-time system activity, alerts, and audit logs</p>
        </div>

        <!-- Stats Overview -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Critical Alerts</p>
              <div class="w-4 h-4 rounded-full bg-red-500"></div>
            </div>
            <p class="text-4xl font-bold text-red-400">3</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Warnings</p>
              <div class="w-4 h-4 rounded-full bg-amber-500"></div>
            </div>
            <p class="text-4xl font-bold text-amber-400">12</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Info Events</p>
              <div class="w-4 h-4 rounded-full bg-blue-500"></div>
            </div>
            <p class="text-4xl font-bold text-blue-400">48</p>
          </div>
          <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6">
            <div class="flex items-center justify-between mb-4">
              <p class="text-gray-400">Total Today</p>
              <div class="w-4 h-4 rounded-full bg-green-500"></div>
            </div>
            <p class="text-4xl font-bold text-green-400">156</p>
          </div>
        </div>

        <!-- Filters -->
        <div class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 p-6 mb-10">
          <div class="flex flex-col lg:flex-row gap-6 items-center justify-between">
            <div class="flex flex-wrap gap-4">
              <select v-model="severityFilter" class="px-5 py-3 bg-gray-700 border border-gray-600 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]">
                <option value="all">All Severity</option>
                <option value="critical">Critical</option>
                <option value="warning">Warning</option>
                <option value="info">Info</option>
              </select>
              <select v-model="typeFilter" class="px-5 py-3 bg-gray-700 border border-gray-600 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]">
                <option value="all">All Types</option>
                <option value="auth">Authentication</option>
                <option value="system">System</option>
                <option value="user">User Activity</option>
                <option value="job">Job Posting</option>
              </select>
            </div>
            <button class="px-6 py-3 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-xl font-medium transition flex items-center gap-3">
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15" />
              </svg>
              Refresh Logs
            </button>
          </div>
        </div>

        <!-- Logs Table -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 overflow-hidden">
          <div class="p-6 border-b border-gray-700">
            <h2 class="text-2xl font-bold">Activity Log</h2>
          </div>
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead class="bg-gray-700/50">
                <tr>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Timestamp</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Severity</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Type</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">Message</th>
                  <th class="px-6 py-4 text-left text-sm font-medium text-gray-300">User / Source</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-700">
                <tr v-for="log in filteredLogs" :key="log.id" class="hover:bg-gray-700/30 transition">
                  <td class="px-6 py-5 text-sm text-gray-300">{{ log.timestamp }}</td>
                  <td class="px-6 py-5">
                    <span :class="[
                      'px-4 py-2 rounded-full text-xs font-medium',
                      log.severity === 'critical' ? 'bg-red-600/20 text-red-400' :
                      log.severity === 'warning' ? 'bg-amber-600/20 text-amber-400' :
                      'bg-blue-600/20 text-blue-400'
                    ]">
                      {{ log.severityLabel }}
                    </span>
                  </td>
                  <td class="px-6 py-5 text-sm text-gray-300 capitalize">{{ log.type }}</td>
                  <td class="px-6 py-5 text-sm text-gray-200 max-w-md truncate">{{ log.message }}</td>
                  <td class="px-6 py-5 text-sm text-gray-400">{{ log.source }}</td>
                </tr>
              </tbody>
            </table>
          </div>
          <div class="p-6 border-t border-gray-700 text-center text-gray-400">
            Showing latest 50 entries • Auto-refresh every 30s
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

const severityFilter = ref('all')
const typeFilter = ref('all')

// Mock log data
const logs = ref([
  { id: 1, timestamp: '2025-12-31 14:32:18', severity: 'critical', severityLabel: 'Critical', type: 'auth', message: 'Failed login attempt from IP 192.168.1.105 (5th attempt)', source: 'System' },
  { id: 2, timestamp: '2025-12-31 14:28:45', severity: 'warning', severityLabel: 'Warning', type: 'job', message: 'Job posting #JOB-892 reached 95% of application limit', source: 'Job Engine' },
  { id: 3, timestamp: '2025-12-31 14:25:12', severity: 'info', severityLabel: 'Info', type: 'user', message: 'Admin user "john.doe" updated client profile for TechCorp Solutions', source: 'john.doe@admin.com' },
  { id: 4, timestamp: '2025-12-31 14:20:03', severity: 'info', severityLabel: 'Info', type: 'system', message: 'Daily backup completed successfully', source: 'Backup Service' },
  { id: 5, timestamp: '2025-12-31 14:15:57', severity: 'warning', severityLabel: 'Warning', type: 'auth', message: 'Unusual login location detected for user maria.gonzalez', source: 'Security Module' },
  { id: 6, timestamp: '2025-12-31 14:10:22', severity: 'info', severityLabel: 'Info', type: 'job', message: 'New application received for Senior Developer position', source: 'Talent Portal' },
  { id: 7, timestamp: '2025-12-31 14:05:41', severity: 'critical', severityLabel: 'Critical', type: 'system', message: 'Database connection pool exhausted - high traffic detected', source: 'Database Monitor' },
  { id: 8, timestamp: '2025-12-31 13:58:19', severity: 'info', severityLabel: 'Info', type: 'user', message: 'Client "TechCorp" started chat with talent "James Wilson"', source: 'Messaging Service' }
])

const filteredLogs = computed(() => {
  let filtered = logs.value

  if (severityFilter.value !== 'all') {
    filtered = filtered.filter(log => log.severity === severityFilter.value)
  }

  if (typeFilter.value !== 'all') {
    filtered = filtered.filter(log => log.type === typeFilter.value)
  }

  return filtered
})

// Theme primary color sync
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