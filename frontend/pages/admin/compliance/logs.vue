<template>
  <div class="min-h-screen flex">
    <!-- Sidebar - Added top margin to lower it slightly -->
    <aside class="flex-shrink-0 pt-12 mt-4"> <!-- Key change: pt-8 (padding-top: 2rem) -->
      <admin-sidebar />
    </aside>

    <!-- Main Reports Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="p-6 md:p-8 lg:p-12 max-w-7xl mx-auto">
        <!-- Page Header -->
        <div class="flex flex-col md:flex-row md:items-center md:justify-between mb-10">
          <h1 class="text-3xl md:text-4xl font-bold mb-6 md:mb-0">Reports & Analytics</h1>
          <div class="flex flex-col sm:flex-row gap-4">
            <select
              v-model="timeRange"
              class="px-5 py-3 rounded-xl bg-gray-800 border border-gray-700 focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)] transition text-sm"
            >
              <option value="week">Last 7 Days</option>
              <option value="month">Last 30 Days</option>
              <option value="quarter">Last Quarter</option>
              <option value="year">Last Year</option>
            </select>
            <button
              @click="exportReport"
              class="px-6 py-3 rounded-xl bg-[color:var(--primary-color)] text-white hover:opacity-90 transition flex items-center justify-center gap-2 text-sm font-medium"
            >
              <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
              </svg>
              Export Report
            </button>
          </div>
        </div>

        <!-- Key Metrics Grid -->
        <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6 mb-12">
          <div class="p-6 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
            <h3 class="text-lg font-medium text-gray-400 mb-3">Total Jobs Posted</h3>
            <p class="text-4xl font-bold text-white">{{ metrics.totalJobs }}</p>
            <p class="text-sm text-green-400 mt-3">↑ 18% from last period</p>
          </div>
          <div class="p-6 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
            <h3 class="text-lg font-medium text-gray-400 mb-3">Total Applications</h3>
            <p class="text-4xl font-bold text-white">{{ metrics.totalApplications }}</p>
            <p class="text-sm text-green-400 mt-3">↑ 12% from last period</p>
          </div>
          <div class="p-6 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
            <h3 class="text-lg font-medium text-gray-400 mb-3">Active Talents</h3>
            <p class="text-4xl font-bold text-white">{{ metrics.activeTalents }}</p>
            <p class="text-sm text-green-400 mt-3">↑ 22% from last period</p>
          </div>
          <div class="p-6 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
            <h3 class="text-lg font-medium text-gray-400 mb-3">Active Clients</h3>
            <p class="text-4xl font-bold text-white">{{ metrics.activeClients }}</p>
            <p class="text-sm text-amber-400 mt-3">→ 3% from last period</p>
          </div>
        </div>

        <!-- Charts Section - Fixed height containers -->
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8 mb-12">
          <!-- Jobs Posted Over Time -->
          <div class="p-8 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
            <h3 class="text-xl font-semibold mb-6">Jobs Posted Over Time</h3>
            <div class="h-96">
              <canvas ref="jobsChart"></canvas>
            </div>
          </div>

          <!-- Applications by Category -->
          <div class="p-8 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
            <h3 class="text-xl font-semibold mb-6">Applications by Job Category</h3>
            <div class="h-96">
              <canvas ref="categoryChart"></canvas>
            </div>
          </div>
        </div>

        <!-- Top Performing Jobs Table -->
        <div class="p-8 rounded-2xl bg-gray-800 shadow-xl border border-gray-700">
          <h3 class="text-xl font-semibold mb-6">Top Performing Jobs</h3>
          <div class="overflow-x-auto">
            <table class="w-full text-left">
              <thead>
                <tr class="border-b border-gray-700 text-gray-400 text-sm">
                  <th class="py-4 px-6">Job Title</th>
                  <th class="py-4 px-6 text-center">Applications</th>
                  <th class="py-4 px-6 text-center">Views</th>
                  <th class="py-4 px-6 text-center">Conversion Rate</th>
                  <th class="py-4 px-6 text-center">Status</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="job in topJobs" :key="job.id" class="border-b border-gray-700 hover:bg-gray-700/30 transition">
                  <td class="py-5 px-6 font-medium text-white">{{ job.title }}</td>
                  <td class="py-5 px-6 text-center">{{ job.applications }}</td>
                  <td class="py-5 px-6 text-center">{{ job.views }}</td>
                  <td class="py-5 px-6 text-center">
                    <span class="text-green-400 font-medium">{{ job.conversion }}%</span>
                  </td>
                  <td class="py-5 px-6 text-center">
                    <span
                      :class="[
                        'px-4 py-2 rounded-full text-xs font-medium',
                        job.status === 'Active' ? 'bg-green-600/30 text-green-400' : 'bg-red-600/30 text-red-400'
                      ]"
                    >
                      {{ job.status }}
                    </span>
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
import { ref, onMounted, watch, nextTick } from 'vue'
import { useThemeStore } from '~/stores/theme'
import Chart from 'chart.js/auto'

const theme = useThemeStore()

// Reactive state
const timeRange = ref('month')

const metrics = ref({
  totalJobs: 248,
  totalApplications: 1942,
  activeTalents: 512,
  activeClients: 78
})

const topJobs = ref([
  { id: 1, title: 'Senior Full-Stack Engineer', applications: 92, views: 2450, conversion: 3.8, status: 'Active' },
  { id: 2, title: 'Product Designer (UI/UX)', applications: 78, views: 1890, conversion: 4.1, status: 'Active' },
  { id: 3, title: 'DevOps Engineer', applications: 65, views: 1520, conversion: 4.3, status: 'Active' },
  { id: 4, title: 'Marketing Lead', applications: 48, views: 980, conversion: 4.9, status: 'Closed' },
  { id: 5, title: 'Data Analyst', applications: 55, views: 1210, conversion: 4.5, status: 'Active' }
])

// Chart refs
const jobsChart = ref<HTMLCanvasElement | null>(null)
const categoryChart = ref<HTMLCanvasElement | null>(null)

let jobsChartInstance: Chart | null = null
let categoryChartInstance: Chart | null = null

// Create or update charts
const createCharts = () => {
  const primaryColor = getComputedStyle(document.documentElement).getPropertyValue('--primary-color').trim() || '#3b82f6'

  // Destroy previous instances
  if (jobsChartInstance) jobsChartInstance.destroy()
  if (categoryChartInstance) categoryChartInstance.destroy()

  // Line Chart: Jobs Posted Over Time
  if (jobsChart.value) {
    jobsChartInstance = new Chart(jobsChart.value, {
      type: 'line',
      data: {
        labels: ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'],
        datasets: [{
          label: 'Jobs Posted',
          data: [18, 24, 30, 28, 35, 42, 38, 45, 52, 48, 55, 62],
          borderColor: primaryColor,
          backgroundColor: `${primaryColor}30`,
          fill: true,
          tension: 0.4,
          pointBackgroundColor: primaryColor,
          pointBorderColor: '#fff',
          pointBorderWidth: 2,
          pointRadius: 5
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: { display: false },
          tooltip: { backgroundColor: '#1f2937' }
        },
        scales: {
          y: {
            beginAtZero: true,
            grid: { color: '#374151' },
            ticks: { color: '#9ca3af' }
          },
          x: {
            grid: { color: '#374151' },
            ticks: { color: '#9ca3af' }
          }
        }
      }
    })
  }

  // Doughnut Chart: Applications by Category
  if (categoryChart.value) {
    categoryChartInstance = new Chart(categoryChart.value, {
      type: 'doughnut',
      data: {
        labels: ['Technology', 'Design', 'Marketing', 'Sales', 'Management', 'Other'],
        datasets: [{
          data: [620, 420, 310, 240, 180, 172],
          backgroundColor: [
            primaryColor,
            '#a78bfa',
            '#f472b6',
            '#fb923c',
            '#34d399',
            '#60a5fa'
          ],
          borderWidth: 0,
          hoverOffset: 10
        }]
      },
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: {
            position: 'right',
            labels: {
              color: '#e5e7eb',
              padding: 20,
              font: { size: 14 }
            }
          }
        }
      }
    })
  }
}

// Export placeholder
const exportReport = () => {
  alert('Export to PDF/CSV feature ready for implementation (e.g., jsPDF + html2canvas)')
}

// Update theme color and refresh charts
const updateThemeColor = () => {
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
  nextTick(() => createCharts())
}

// Initialize
onMounted(() => {
  updateThemeColor()
  createCharts()
})

watch(() => theme.primaryColor, updateThemeColor)
watch(timeRange, () => {
  console.log('Time range changed:', timeRange.value)
  // In real app: refetch data here
})
</script>

<style scoped>
main {
  min-height: 100vh;
}
</style>