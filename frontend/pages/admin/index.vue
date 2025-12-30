<template>
  <div class="min-h-screen w-full font-poppins bg-gray-100">
    <div class="container mx-auto px-4 py-6 h-full">
      <!-- Search Bar -->
      <div class="mb-6 flex items-center justify-around">
        <input 
          type="text" 
          v-model="searchTerm" 
          placeholder="Search for something..." 
          class="w-full max-w-2xl mx-auto block  rounded-3xl py-3  px-3 text-lg bg-white text-left focus:outline-none focus:ring-2 focus:ring-black"
        >
        <div class="flex justify-around items-center p-2">
          <nuxt-link to="/admin/profile"><div class="h-10 w-10 rounded-full bg-white mx-2 hover:scale-105 hover:text-gray-300  transition-all ease-out duration-75">
            <user-icon class="text-gray-500 h-7 w-7 translate-1"/>
          </div></nuxt-link>
          <div class="h-10 w-10 rounded-full bg-white  mx-2 hover:scale-105 hover:text-gray-300 transition-all ease-out duration-75">
            <bell-icon class="text-gray-500 h-7 w-7 translate-1"/>
          </div>
         <div class="h-10 w-10 rounded-full bg-white  mx-2 hover:scale-105 hover:text-gray-300 transition-all ease-out duration-75">
          <moon-icon class="text-gray-500 h-7 w-7 translate-1"/>
         </div>

        </div>
      </div>

      <!-- Main Content -->
      <div class="flex flex-col lg:flex-row gap-6">
       
        <aside >
          <admin-sidebar></admin-sidebar>
           
        </aside> 

      
        <main class="flex-1">
          <div class="grid grid-cols-1 lg:grid-cols-3 md:grid-cols-3 gap-2">
             <div class="flex  bg-gray rounded-lg mt-2 ml-2 mr-2 mb-4 p-4 shadow-black/20 shadow-lg bg-white  hover:shadow-2xl hover:scale-105 transition-all delay-75">
                <div class="w-18 h-18 rounded-full bg-blue-100  p-2">
                  <BriefcaseIcon class="w-10 h-10 text-blue-500 translate-x-1 translate-y-1" />
                </div>
                <div class="p-2 text-black/50">
                  <h1 class=" text-lg">Total Jobs Posted</h1>
                  <p class="text-center text-2xl text-black font-semibold">2000</p>
                </div>
             </div>
               <div class="flex bg-gray rounded-lg mt-2 ml-2 mr-2 mb-4 p-4 shadow-black/20 shadow-lg bg-white hover:shadow-2xl hover:scale-105 transition-all delay-75">
                <div class="w-18 h-18 rounded-full bg-green-200 p-2">
                  <ClipboardDocumentIcon class="w-10 h-10 text-green-500 translate-1" />
                </div>
                <div class="p-2 text-black/50">
                  <h1 class=" text-lg">Total Application</h1>
                  <p class="text-center text-2xl text-black font-semibold">2000</p>
                </div>
             </div>
               <div class="flex bg-gray rounded-lg mt-2 ml-2 mr-2 mb-4 p-4 shadow-black/20 shadow-lg bg-white hover:shadow-2xl hover:scale-105 transition-all delay-75">
                <div class="w-18 h-18 rounded-full bg-purple-200 p-2">
                  <UsersIcon class="w-10 h-10 text-purple-500 translate-1" />
                </div>
                <div class="p-2 text-black/50">
                  <h1 class=" text-lg">Total Clients</h1>
                  <p class="text-center text-2xl text-black font-semibold">2000</p>
                </div>
             </div>
          </div>
          <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-2">
            <div class="lg:w-[35vw] h-[50vh] mr-4 rounded-lg border-2 border-blue-100">
                 <Bar :data="chartData" :options="chartOptions" class="bg-white rounded-lg"/>
            </div>
            <div class="lg:w-[35vw] h-[50vh] ml-4 rounded-lg">
                     <Pie :data="chartData" :options="chartOptions" class="bg-white"/>
            </div>

          </div>
      <!-- Featured Jobs Section -->
<section class="mt-8">
  <h1 class="text-2xl font-bold text-center mb-6">Featured Jobs</h1>
  <div v-if="filteredJobs.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
    <div 
      v-for="(job, index) in filteredJobs" 
      :key="index"
      class="bg-white border border-gray-200 rounded-xl shadow-md hover:shadow-xl  hover:scale-105 transition-all ease-in-out duration-100 overflow-hidden"
    >
      <div class="w-full h-40 overflow-hidden">
        <img 
          src="@/assets/painter.png" 
          alt="Job Image" 
          class="object-cover w-full h-full"
        >
      </div>
      <div class="p-4">
        <div class="flex justify-between items-center mb-2">
          <h2 class="text-lg font-semibold text-gray-800">{{ job.title }}</h2>
          <span class="text-sm font-medium text-blue-600">{{ job.job_type }}</span>
        </div>
        <p class="text-sm text-gray-500">{{ job.category }}</p>
      </div>
    </div>
  </div>
  <div v-else class="flex items-center justify-center h-64 bg-white rounded-xl shadow-inner mt-4">
    <p class="text-lg text-gray-600">No Jobs Available</p>
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

// Heroicons
import {
  BriefcaseIcon,
  ClipboardDocumentIcon,
  UsersIcon
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
    Pie,// ⬅️ Make Bar component available in template
    BriefcaseIcon,
    ClipboardDocumentIcon,
    UsersIcon
  },
  setup() {
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
            display:false
          }
        },
        y:{
          grid:{
            display:false
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
      filteredJobs
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