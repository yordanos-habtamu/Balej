<template>
  <div class="min-h-screen flex">
    <!-- Sidebar -->
    <aside class="flex-shrink-0 pt-8">
      <admin-sidebar />
    </aside>

    <!-- Main Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="max-w-7xl mx-auto px-6 py-10 lg:px-12 lg:py-14">
        <!-- Back Button -->
        <div class="mb-10">
          <NuxtLink to="/admin/clients" class="inline-flex items-center gap-3 text-gray-400 hover:text-[color:var(--primary-color)] transition">
            <ArrowLeftCircleIcon class="w-8 h-8" />
            <span class="text-lg font-medium">Back to Clients</span>
          </NuxtLink>
        </div>

        <!-- Project Header -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-10 mb-10">
          <div class="flex flex-col lg:flex-row lg:items-center justify-between gap-8">
            <div>
              <h1 class="text-4xl lg:text-5xl font-bold mb-4">Enterprise Dashboard Redesign</h1>
              <div class="flex flex-wrap items-center gap-6 text-gray-400">
                <div class="flex items-center gap-3">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13.255A23.931 23.931 0 0112 15c-3.183 0-6.22-.62-9-1.745M16 6V4a2 2 0 00-2-2h-4a2 2 0 00-2 2v2m4 6h.01M5 20h14a2 2 0 002-2V8a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                  </svg>
                  <span>Project ID: #PROJ-1248</span>
                </div>
                <div class="flex items-center gap-3">
                  <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  <span>Deadline: Feb 15, 2026</span>
                </div>
                <span :class="[
                  'px-5 py-2 rounded-full font-medium',
                  project.status === 'active' ? 'bg-green-600/20 text-green-400' :
                  project.status === 'completed' ? 'bg-blue-600/20 text-blue-400' :
                  'bg-amber-600/20 text-amber-400'
                ]">
                  {{ project.statusLabel }}
                </span>
              </div>
            </div>

            <div class="flex flex-col sm:flex-row gap-4">
              <button class="px-8 py-4 bg-gray-700 hover:bg-gray-600 text-white rounded-2xl font-semibold transition flex items-center justify-center gap-3">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.232 5.232l3.536 3.536m-2.036-5.036a2.5 2.5 0 113.536 3.536L6.5 21.036H3v-3.572L16.732 3.732z" />
                </svg>
                Edit Project
              </button>
              
                <button class="px-8 py-4 bg-[color:var(--primary-color)] hover:bg-[color:var(--primary-color)/0.9] text-white rounded-2xl font-semibold transition flex items-center justify-center gap-3 shadow-lg" @click="RouteChat()">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                </svg>
                Message Client
              </button>
              
              
            </div>
          </div>
        </div>

        <!-- Project Overview Grid -->
        <div class="grid grid-cols-1 md:grid-cols-3 gap-8 mb-12">
          <!-- Budget & Progress -->
          <div class="md:col-span-2 space-y-8">
            <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-8">
              <h2 class="text-2xl font-bold mb-6">Project Progress</h2>
              <div class="flex items-end justify-between mb-6">
                <div>
                  <p class="text-5xl font-bold text-white">{{ project.progress }}%</p>
                  <p class="text-gray-400 mt-2">Complete</p>
                </div>
                <div class="text-right">
                  <p class="text-3xl font-bold text-[color:var(--primary-color)]">${{ project.spent }}</p>
                  <p class="text-gray-400">of ${{ project.budget }} spent</p>
                </div>
              </div>
              <div class="w-full bg-gray-700 rounded-full h-6">
                <div
                  class="h-6 rounded-full transition-all duration-700"
                  :style="{ width: project.progress + '%', backgroundColor: 'var(--primary-color)' }"
                ></div>
              </div>
            </div>

            <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-8">
              <h2 class="text-2xl font-bold mb-6">Project Description</h2>
              <p class="text-lg text-gray-300 leading-relaxed">
                {{ project.description }}
              </p>
            </div>
          </div>

          <!-- Quick Stats -->
          <div class="space-y-8">
            <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-8">
              <h2 class="text-2xl font-bold mb-6">Quick Stats</h2>
              <div class="space-y-6">
                <div class="flex justify-between">
                  <span class="text-gray-400">Client</span>
                  <span class="font-semibold">TechCorp Solutions</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-400">Start Date</span>
                  <span class="font-semibold">Oct 1, 2025</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-400">Talents Assigned</span>
                  <span class="font-semibold text-[color:var(--primary-color)]">{{ project.talents }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-gray-400">Tasks Completed</span>
                  <span class="font-semibold">42 / 68</span>
                </div>
              </div>
            </div>

            <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-8">
              <h2 class="text-2xl font-bold mb-6">Timeline</h2>
              <div class="space-y-4">
                <div class="flex items-center gap-4">
                  <div class="w-4 h-4 rounded-full bg-green-500"></div>
                  <span class="text-gray-300">Kickoff Meeting</span>
                  <span class="text-gray-500 ml-auto">Oct 5</span>
                </div>
                <div class="flex items-center gap-4">
                  <div class="w-4 h-4 rounded-full bg-[color:var(--primary-color)]"></div>
                  <span class="text-gray-300">Design Phase Complete</span>
                  <span class="text-gray-500 ml-auto">Nov 20</span>
                </div>
                <div class="flex items-center gap-4">
                  <div class="w-4 h-4 rounded-full bg-gray-600"></div>
                  <span class="text-gray-300">Development Phase</span>
                  <span class="text-gray-500 ml-auto">Ongoing</span>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Assigned Talents -->
        <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-10">
          <h2 class="text-2xl font-bold mb-8">Assigned Talents</h2>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
            <div v-for="talent in project.assignedTalents" :key="talent.id" class="flex items-center gap-6 bg-gray-700/50 rounded-2xl p-6 hover:bg-gray-700/70 transition">
              <div class="w-16 h-16 rounded-full overflow-hidden ring-4 ring-gray-600">
                <img :src="talent.avatar" alt="Talent" class="w-full h-full object-cover" />
              </div>
              <div>
                <h4 class="text-lg font-semibold text-white">{{ talent.name }}</h4>
                <p class="text-gray-400">{{ talent.role }}</p>
                <p class="text-sm text-gray-500 mt-1">{{ talent.hours }} hours this week</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import {useRouter} from 'vue-router'
import { useThemeStore } from '~/stores/theme'
import { ArrowLeftCircleIcon } from '@heroicons/vue/24/outline'

const theme = useThemeStore()
const router = useRouter()

const RouteChat = () => {
  router.push('/admin/chats')
}
// Mock project data
const project = ref({
  name: 'Enterprise Dashboard Redesign',
  description: 'Complete overhaul of the internal analytics platform with modern UI/UX, real-time data visualization, and improved performance. The new dashboard will support multiple user roles and advanced filtering capabilities.',
  status: 'active',
  statusLabel: 'Active',
  budget: '45,000',
  spent: '30,600',
  progress: 68,
  talents: 5,
  assignedTalents: [
    { id: 1, name: 'Maria Gonzalez', role: 'Senior UI Designer', hours: 32, avatar: '~/assets/painter.png' },
    { id: 2, name: 'James Wilson', role: 'Full-Stack Developer', hours: 40, avatar: '~/assets/painter.png' },
    { id: 3, name: 'Aisha Khan', role: 'Frontend Engineer', hours: 28, avatar: '~/assets/painter.png' },
    { id: 4, name: 'Carlos Rodriguez', role: 'Backend Developer', hours: 35, avatar: '~/assets/painter.png' },
    { id: 5, name: 'Sophie Dubois', role: 'Project Manager', hours: 20, avatar: '~/assets/painter.png' }
  ]
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