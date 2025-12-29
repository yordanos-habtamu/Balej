<template>
  <div class="min-h-screen w-full bg-gray-100">
    <div class="container mx-auto px-4 py-6">
      <!-- Search Bar -->
      <div class="mb-6">
        <input 
          type="text" 
          v-model="searchTerm"
          placeholder="Search for something..." 
          class="w-full max-w-2xl mx-auto block border-2 border-black rounded-lg p-3 text-lg focus:outline-none focus:ring-2 focus:ring-black"
        >
      </div>

      <!-- Main Content -->
      <div class="flex flex-col lg:flex-row gap-6">
        <!-- Sidebar -->
        <aside class="w-full lg:w-80 bg-white border-2 border-black rounded-lg p-4 h-fit">
          <div class="flex flex-col gap-4">
            <!-- Profile Section -->
            <div class="flex items-center gap-4">
              <div class="w-16 h-16 rounded-full overflow-hidden">
                <img src="@/assets/painter.png" alt="Profile" class="object-cover w-full h-full">
              </div>
              <h1 class="text-lg font-bold">Lorem Ipsum</h1>
            </div>

            <!-- Job Type Filters -->
            <div class="border-t-2 border-black pt-4">
              <h3 class="font-semibold mb-2">Job Type</h3>
              <div class="grid grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="part-time" class="h-4 w-4">
                  <span>Part-time</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="contract" class="h-4 w-4">
                  <span>Contract</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="full-time" class="h-4 w-4">
                  <span>Full-time</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="internship" class="h-4 w-4">
                  <span>Internship</span>
                </label>
              </div>
            </div>

            <!-- Category Filters -->
            <div class="border-t-2 border-black pt-4">
              <h3 class="font-semibold mb-2">Category</h3>
              <div class="grid grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="construction" class="h-4 w-4">
                  <span>Construction</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="house_keeping" class="h-4 w-4">
                  <span>House Keeping</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="repairing" class="h-4 w-4">
                  <span>Repairing</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="management" class="h-4 w-4">
                  <span>Management</span>
                </label>
              </div>
            </div>

            <!-- Price Range -->
            <div class="border-t-2 border-black pt-4">
              <h3 class="font-semibold mb-2">Price Range</h3>
              <div class="flex gap-2">
                <input 
                  type="number" 
                  placeholder="Min" 
                  class="w-1/2 border-2 border-black rounded-lg p-2"
                >
                <input 
                  type="number" 
                  placeholder="Max" 
                  class="w-1/2 border-2 border-black rounded-lg p-2"
                >
              </div>
            </div>
          </div>
        </aside>

        <!-- Job Listings -->
        <main class="flex-1">
          <div v-if="filteredJobs.length" class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            <div 
              v-for="(job, index) in filteredJobs" 
              :key="index"
              class="bg-white border-2 border-black rounded-lg overflow-hidden hover:shadow-lg transition-shadow"
            >
              <div class="w-full h-32">
                <img 
                  src="@/assets/painter.png" 
                  alt="Job Image" 
                  class="object-cover w-full h-full"
                >
              </div>
              <div class="p-4">
                <div class="flex justify-between items-center mb-2">
                  <h2 class="text-xl font-bold">{{ job.title }}</h2>
                  <p class="text-lg">${{ job.job_type }}</p>
                </div>
                <p class="text-sm text-gray-600">{{ job.category }}</p>
              </div>
            </div>
          </div>
          <div v-else class="flex items-center justify-center h-64">
            <p class="text-lg text-black">No Jobs Available</p>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, onMounted, reactive, computed, ref } from 'vue';
import { gql } from 'graphql-tag';
import { CombinedError } from '@urql/core';
import { useUrqlClient } from '~/composables/useUrqlClient';

const USER_QUERY = gql`
  query {
    jobs {
      job_type
      title
      applicationCount
      category
    }
  }
`;

export default defineComponent({
  setup() {
    const client = useUrqlClient();
    const searchTerm = ref('');

    // Reactive object to store jobs and state
    const jobsData = reactive({
      jobs: [] as Array<any>,
      loading: false,
      error: null as CombinedError | null,
    });

    // Computed property for filtered jobs
    const filteredJobs = computed(() => {
      if (!searchTerm.value) return jobsData.jobs;
      const lowerSearch = searchTerm.value.toLowerCase();
      return jobsData.jobs.filter(job =>
        job.title?.toLowerCase().includes(lowerSearch) ||
        job.category?.toLowerCase().includes(lowerSearch)
      );
    });

    // Fetch jobs from backend
    async function loadJobs() {
      jobsData.loading = true;
      try {
        const response = await client.query(USER_QUERY, {}).toPromise();

        if (response.error) {
          jobsData.error = response.error;
        } else {
          jobsData.jobs = response.data?.jobs || [];
        }
      } catch (err) {
        jobsData.error = new CombinedError({ networkError: err as Error });
      }
      jobsData.loading = false;
    }

    // Load jobs on component mount
    onMounted(() => {
      loadJobs();
    });

    return {
      jobsData,
      searchTerm,
      filteredJobs,
    };
  },
});
</script>

<style scoped>
@media (max-width: 1024px) {
  aside {
    position: static;
    width: 100%;
  }
}
</style>