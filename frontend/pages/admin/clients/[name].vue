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
          <h1 class="text-4xl font-bold mb-2">Client Projects</h1>
          <p class="text-xl text-gray-400">Overview of all active and completed projects for this client</p>
        </div>

        <!-- Client Summary Card -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-8 mb-10 flex flex-col md:flex-row items-center gap-8">
          <div class="w-32 h-32 rounded-full overflow-hidden ring-6 ring-[color:var(--primary-color)/0.2] shadow-xl">
            <img src="~/assets/painter.png" alt="Client" class="w-full h-full object-cover" />
          </div>
          <div class="text-center md:text-left">
            <h2 class="text-3xl font-bold mb-2">TechCorp Solutions</h2>
            <p class="text-xl text-gray-400 mb-4">Premium Client • Since March 2024</p>
            <div class="flex flex-wrap gap-4 justify-center md:justify-start">
              <div class="text-center">
                <p class="text-3xl font-bold text-[color:var(--primary-color)]">8</p>
                <p class="text-gray-400">Active Projects</p>
              </div>
              <div class="text-center">
                <p class="text-3xl font-bold text-green-400">24</p>
                <p class="text-gray-400">Completed</p>
              </div>
              <div class="text-center">
                <p class="text-3xl font-bold text-white">$184K</p>
                <p class="text-gray-400">Total Value</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Filters & Search -->
        <div class="flex flex-col lg:flex-row gap-6 mb-10">
          <div class="flex-1">
            <div class="relative">
              <input
                type="text"
                v-model="searchTerm"
                placeholder="Search projects..."
                class="w-full px-6 py-4 pl-14 bg-gray-800 border border-gray-700 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.5] placeholder-gray-500"
              />
              <svg class="absolute left-5 top-1/2 -translate-y-1/2 w-6 h-6 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
              </svg>
            </div>
          </div>
          <select v-model="statusFilter" class="px-6 py-4 bg-gray-800 border border-gray-700 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.5]">
            <option value="all">All Status</option>
            <option value="active">Active</option>
            <option value="completed">Completed</option>
            <option value="on-hold">On Hold</option>
          </select>
        </div>

        <!-- Projects Grid -->
        <div class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-8">
          <div
            v-for="project in filteredProjects"
            :key="project.id"
            class="bg-gray-800 rounded-2xl shadow-xl border border-gray-700 overflow-hidden transition-all hover:shadow-[0_0_30px_rgba(var(--primary-color-rgb),0.2)] hover:border-[color:var(--primary-color)/0.4]"
          >
            <!-- Project Header -->
            <div class="p-6 border-b border-gray-700">
              <div class="flex items-center justify-between mb-4">
                <h3 class="text-xl font-bold text-white">{{ project.name }}</h3>
                <span :class="[
                  'px-4 py-2 rounded-full text-sm font-medium',
                  project.status === 'active' ? 'bg-green-600/20 text-green-400' :
                  project.status === 'completed' ? 'bg-blue-600/20 text-blue-400' :
                  'bg-amber-600/20 text-amber-400'
                ]">
                  {{ project.statusLabel }}
                </span>
              </div>
              <p class="text-gray-400">{{ project.description }}</p>
            </div>

            <!-- Project Stats -->
            <div class="p-6 space-y-4">
              <div class="flex justify-between">
                <span class="text-gray-400">Budget</span>
                <span class="font-semibold">{{ project.budget }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-400">Deadline</span>
                <span class="font-semibold">{{ project.deadline }}</span>
              </div>
              <div class="flex justify-between">
                <span class="text-gray-400">Talents Assigned</span>
                <span class="font-semibold">{{ project.talents }}</span>
              </div>
            </div>

            <!-- Progress Bar -->
            <div class="px-6 pb-6">
              <div class="flex justify-between mb-2">
                <span class="text-sm text-gray-400">Progress</span>
                <span class="text-sm font-medium">{{ project.progress }}%</span>
              </div>
              <div class="w-full bg-gray-700 rounded-full h-3">
                <div
                  class="h-3 rounded-full transition-all duration-500"
                  :class="project.status === 'active' ? 'bg-[color:var(--primary-color)]' : 'bg-green-500'"
                  :style="{ width: project.progress + '%' }"
                ></div>
              </div>
            </div>

            <!-- Actions -->
                <nuxt-link :to="`/admin/clients/projects/${project.name}`">
            <div class="px-6 pb-6 flex gap-4">
           <button class="flex-1 py-3 bg-gray-700 hover:bg-gray-600 text-white rounded-xl font-medium transition">
                View Details
              </button>
            </div>
          </nuxt-link>

          </div>
        </div>

        <!-- Empty State -->
        <div v-if="filteredProjects.length === 0" class="text-center py-32">
          <p class="text-2xl text-gray-500">No projects found.</p>
          <p class="text-gray-600 mt-4">Try adjusting your search or filters.</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useThemeStore } from '~/stores/theme'

const theme = useThemeStore()

const searchTerm = ref('')
const statusFilter = ref('all')

// Mock project data
const projects = ref([
  { id: 1, name: 'Enterprise Dashboard Redesign', description: 'Complete overhaul of internal analytics platform', status: 'active', statusLabel: 'Active', budget: '$45,000', deadline: 'Feb 15, 2026', talents: 5, progress: 68 },
  { id: 2, name: 'Mobile App Development', description: 'Native iOS & Android apps for customer portal', status: 'active', statusLabel: 'Active', budget: '$72,000', deadline: 'Mar 30, 2026', talents: 8, progress: 42 },
  { id: 3, name: 'Cloud Migration Phase 2', description: 'Migrating legacy systems to AWS', status: 'active', statusLabel: 'Active', budget: '$38,000', deadline: 'Jan 20, 2026', talents: 4, progress: 85 },
  { id: 4, name: 'API Integration Suite', description: 'Connecting third-party services', status: 'completed', statusLabel: 'Completed', budget: '$28,000', deadline: 'Nov 10, 2025', talents: 3, progress: 100 },
  { id: 5, name: 'Security Audit & Compliance', description: 'GDPR and SOC 2 compliance preparation', status: 'on-hold', statusLabel: 'On Hold', budget: '$15,000', deadline: 'Pending', talents: 2, progress: 30 },
  { id: 6, name: 'Marketing Website Revamp', description: 'New corporate website with CMS', status: 'active', statusLabel: 'Active', budget: '$22,000', deadline: 'Feb 28, 2026', talents: 4, progress: 55 }
])

const filteredProjects = computed(() => {
  let filtered = projects.value

  if (searchTerm.value) {
    const term = searchTerm.value.toLowerCase()
    filtered = filtered.filter(p => 
      p.name.toLowerCase().includes(term) ||
      p.description.toLowerCase().includes(term)
    )
  }

  if (statusFilter.value !== 'all') {
    filtered = filtered.filter(p => p.status === statusFilter.value)
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