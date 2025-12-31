<template>
  <div class="min-h-screen w-full font-poppins bg-gray-100 dark:bg-gray-900 transition-colors duration-300">
    <div class="container mx-auto px-4 py-6 h-full">
      <!-- Search Bar -->
      <div class="mb-6 flex items-center justify-around">
        <input 
          type="text" 
          v-model="searchTerm" 
          placeholder="Search for something..." 
          class="w-full max-w-2xl mx-auto block rounded-3xl py-3 px-3 text-lg bg-white dark:bg-gray-800 text-left text-gray-900 dark:text-white focus:outline-none focus:ring-2 focus:ring-black dark:focus:ring-white border border-transparent dark:border-gray-700"
        >
        <div class="flex justify-around items-center p-2">
          <nuxt-link to="/admin/profile">
            <div class="h-10 w-10 rounded-full bg-white dark:bg-gray-800 mx-2 hover:scale-105 hover:text-gray-300 transition-all ease-out duration-75 flex items-center justify-center shadow-sm">
              <user-icon class="text-gray-500 dark:text-gray-400 h-6 w-6"/>
            </div>
          </nuxt-link>
          <div class="h-10 w-10 rounded-full bg-white dark:bg-gray-800 mx-2 hover:scale-105 hover:text-gray-300 transition-all ease-out duration-75 flex items-center justify-center shadow-sm">
            <bell-icon class="text-gray-500 dark:text-gray-400 h-6 w-6"/>
          </div>
          <button @click="themeStore.toggleDarkMode" class="h-10 w-10 rounded-full bg-white dark:bg-gray-800 mx-2 hover:scale-105 hover:text-gray-300 transition-all ease-out duration-75 flex items-center justify-center shadow-sm">
            <moon-icon v-if="!themeStore.isDarkMode" class="text-gray-500 h-6 w-6"/>
            <sun-icon v-else class="text-yellow-500 h-6 w-6"/>
          </button>
        </div>
      </div>

      <!-- Main Content -->
      <div class="flex flex-col lg:flex-row gap-6">
       
        <aside>
          <admin-sidebar></admin-sidebar>
        </aside> 

        <main class="flex-1">
          <div class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-3 gap-2">
             <div class="flex bg-white dark:bg-gray-800 rounded-lg mt-2 ml-2 mr-2 mb-4 p-4 shadow-lg hover:shadow-2xl hover:scale-105 transition-all delay-75 border dark:border-gray-700">
                <div class="w-18 h-18 rounded-full bg-blue-100 dark:bg-blue-900 p-2 flex items-center justify-center">
                  <BriefcaseIcon class="w-10 h-10 text-blue-500 dark:text-blue-300" />
                </div>
                <div class="p-2 text-gray-500 dark:text-gray-400">
                  <h1 class="text-lg">Total Jobs Posted</h1>
                  <p class="text-center text-2xl text-black dark:text-white font-semibold">2000</p>
                </div>
             </div>
               <div class="flex bg-white dark:bg-gray-800 rounded-lg mt-2 ml-2 mr-2 mb-4 p-4 shadow-lg hover:shadow-2xl hover:scale-105 transition-all delay-75 border dark:border-gray-700">
                <div class="w-18 h-18 rounded-full bg-green-200 dark:bg-green-900 p-2 flex items-center justify-center">
                  <ClipboardDocumentIcon class="w-10 h-10 text-green-500 dark:text-green-300" />
                </div>
                <div class="p-2 text-gray-500 dark:text-gray-400">
                  <h1 class="text-lg">Total Application</h1>
                  <p class="text-center text-2xl text-black dark:text-white font-semibold">2000</p>
                </div>
             </div>
               <div class="flex bg-white dark:bg-gray-800 rounded-lg mt-2 ml-2 mr-2 mb-4 p-4 shadow-lg hover:shadow-2xl hover:scale-105 transition-all delay-75 border dark:border-gray-700">
                <div class="w-18 h-18 rounded-full bg-purple-200 dark:bg-purple-900 p-2 flex items-center justify-center">
                  <UsersIcon class="w-10 h-10 text-purple-500 dark:text-purple-300" />
                </div>
                <div class="p-2 text-gray-500 dark:text-gray-400">
                  <h1 class="text-lg">Total Clients</h1>
                  <p class="text-center text-2xl text-black dark:text-white font-semibold">2000</p>
                </div>
             </div>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-2">
            <div class="lg:w-[35vw] h-[50vh] mr-4 rounded-lg border-2 border-blue-100 dark:border-gray-700 bg-white dark:bg-gray-800 p-2">
                 <Bar :data="chartData" :options="chartOptions" />
            </div>
            <div class="lg:w-[35vw] h-[50vh] ml-4 rounded-lg bg-white dark:bg-gray-800 p-2 border dark:border-gray-700">
                     <Pie :data="chartData" :options="chartOptions" />
            </div>

          </div>
      <!-- Featured Jobs Section -->
<section class="mt-8">
  <h1 class="text-2xl font-bold text-center mb-6 text-gray-900 dark:text-white">Featured Jobs</h1>
  <div v-if="filteredJobs.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
    <div 
      v-for="(job, index) in filteredJobs" 
      :key="index"
      class="bg-white dark:bg-gray-800 border border-gray-200 dark:border-gray-700 rounded-xl shadow-md hover:shadow-xl hover:scale-105 transition-all ease-in-out duration-100 overflow-hidden"
    >
      <div class="w-full h-40 overflow-hidden bg-gray-200 dark:bg-gray-700">
        <img 
          src="@/assets/painter.png" 
          alt="Job Image" 
          class="object-cover w-full h-full"
        >
      </div>
      <div class="p-4">
        <div class="flex justify-between items-center mb-2">
          <h2 class="text-lg font-semibold text-gray-800 dark:text-white">{{ job.title }}</h2>
          <span class="text-sm font-medium" :class="`text-${themeStore.primaryColor}-600 dark:text-${themeStore.primaryColor}-400`">{{ job.job_type }}</span>
        </div>
        <p class="text-sm text-gray-500 dark:text-gray-400">{{ job.category }}</p>
      </div>
    </div>
  </div>
  <div v-else class="flex items-center justify-center h-64 bg-white dark:bg-gray-800 rounded-xl shadow-inner mt-4 border dark:border-gray-700">
    <p class="text-lg text-gray-600 dark:text-gray-400">No Jobs Available</p>
  </div>
</section>

        </main>
      </div>
    </div>
  </div>
</template>


<script lang="ts">
import { defineComponent, onMounted, reactive, computed, ref } from 'vue'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { useUrqlClient } from '~/composables/useUrqlClient'
import { useThemeStore } from '~/stores/theme'

// Heroicons
import {
  BriefcaseIcon,
  ClipboardDocumentIcon,
  UsersIcon,
  UserIcon,
  BellIcon,
  MoonIcon,
  SunIcon
} from '@heroicons/vue/24/outline'

// Chart.js setup
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  ArcElement,
  CategoryScale,
  LinearScale
} from 'chart.js'
import { Bar,Pie } from 'vue-chartjs'


ChartJS.register(Title, Tooltip, Legend, BarElement, ArcElement,CategoryScale, LinearScale)

const USER_QUERY = gql`
  query {
    jobs {
      job_type
      title
      applicationCount
      category
    }
  }
`

export default defineComponent({
  components: {
    Bar, 
    Pie,
    BriefcaseIcon,
    ClipboardDocumentIcon,
    UsersIcon,
    UserIcon,
    BellIcon,
    MoonIcon,
    SunIcon
  },
  setup() {
    const themeStore = useThemeStore()
    const chartData = ref({
      labels: ['January', 'February', 'March', 'April'],
      datasets: [
        {
          label: 'Jobs',
          data: [40, 20, 12, 39],
          backgroundColor: '#4f46e5'
        }
      ]
    })

    const chartOptions = ref({
      responsive: true,
      maintainAspectRatio: false,
      scales:{
        x:{
          grid:{
            display:false,
            color: themeStore.isDarkMode ? '#374151' : '#e5e7eb'
          },
          ticks: {
             color: themeStore.isDarkMode ? '#9ca3af' : '#374151'
          }
        },
        y:{
          grid:{
            display:false,
            color: themeStore.isDarkMode ? '#374151' : '#e5e7eb'
          },
          ticks: {
             color: themeStore.isDarkMode ? '#9ca3af' : '#374151'
          }
        }
      },
      plugins: {
        legend: {
            labels: {
                color: themeStore.isDarkMode ? '#9ca3af' : '#374151'
            }
        }
      }
    })

    const client = useUrqlClient()
    const searchTerm = ref('')

    const jobsData = reactive({
      jobs: [] as Array<any>,
      loading: false,
      error: null as CombinedError | null
    })

    const filteredJobs = computed(() => {
      if (!searchTerm.value) return jobsData.jobs
      const lowerSearch = searchTerm.value.toLowerCase()
      return jobsData.jobs.filter(
        (job) =>
          job.title?.toLowerCase().includes(lowerSearch) ||
          job.category?.toLowerCase().includes(lowerSearch)
      )
    })

    async function loadJobs() {
      jobsData.loading = true
      try {
        const response = await client.query(USER_QUERY, {}).toPromise()
        if (response.error) {
          jobsData.error = response.error
        } else {
          jobsData.jobs = response.data?.jobs || []
        }
      } catch (err) {
        jobsData.error = new CombinedError({ networkError: err as Error })
      }
      jobsData.loading = false
    }

    onMounted(() => {
      loadJobs()
    })

    return {
      chartData,
      chartOptions,
      jobsData,
      searchTerm,
      filteredJobs,
      themeStore
    }
  }
})
</script>


<style scoped>
@media (max-width: 1024px) {
  aside {
    position: static;
    width: 100%;
  }
  .font-poppins {
    font-family: 'Poppins', sans-serif;
  }
}
</style>