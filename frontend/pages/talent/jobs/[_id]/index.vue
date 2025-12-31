<template>
  <div class="min-h-screen w-full bg-gray-50">
    <!-- Back button -->
    <div class="p-4">
      <nuxt-link to="/admin/jobs">
        <ArrowLeftCircleIcon
          class="h-8 w-8 text-gray-700 hover:text-gray-900 transition-colors"
        />
      </nuxt-link>
    </div>

    <!-- Loading state -->
    <div v-if="jobsData.loading" class="flex items-center justify-center h-[70vh]">
      <p class="text-gray-600 text-lg animate-pulse">Loading job data...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="jobsData.error" class="flex items-center justify-center h-[70vh]">
      <p class="text-red-500 font-medium">Error: {{ jobsData.error.message }}</p>
    </div>

    <!-- User profile -->
    <div v-else-if="job" class="flex flex-col items-center px-4 md:px-12">
      <!-- Profile image -->
      <div class="w-40 h-40 md:w-60 md:h-60 rounded-full overflow-hidden shadow-md">
        <BriefcaseIcon class="text-3xl text-blue"/> 
      </div>

      <!-- User info card -->
      <div
        class="bg-white rounded-2xl shadow-md w-full max-w-4xl mt-6 p-6 md:p-10 grid grid-cols-1 md:grid-cols-2 gap-6"
      >
        <!-- Left column -->
        <div class="space-y-4">
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Title</h2>
            <p class="text-lg font-medium text-gray-900">{{ job.title }}</p>
          </div>
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Type</h2>
            <p class="text-lg font-medium text-gray-900">{{ job.job_type}}</p>
          </div>
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Application Count</h2>
            <p class="text-lg font-medium text-gray-900">{{ job.applicationCount}}</p>
          </div>
        </div>

        <!-- Right column -->
        <div class="space-y-4">
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Category</h2>
            <p class="text-lg font-medium text-gray-900">{{ job.category }}</p>
          </div>

        </div>
      </div>

      <!-- Details Section -->
      <div class="bg-white rounded-2xl shadow-md w-full max-w-4xl mt-6 p-6 md:p-10">
        <h2 class="text-xl font-semibold mb-3 text-gray-800">Details</h2>
        <p class="text-gray-600 leading-relaxed">
           {{ job.description }}
        </p>

        <!-- Actions -->
        <div class="flex flex-col sm:flex-row gap-4 mt-6">
          <button
            class="flex-1 px-6 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-lg shadow transition"
           @click="">
            Apply
          </button>
          <button
            class="flex-1 px-6 py-3 bg-green-500 hover:bg-green-600 text-white rounded-lg shadow transition"
          >
            Add to Catalogue
          </button>
        </div>
      </div>
    </div>

    <!-- User not found -->
    <div v-else class="flex items-center justify-center h-[70vh]">
      <p class="text-gray-600 text-lg">Job not found</p>
    </div>
  </div>
</template>

<script lang="ts">
import { ArrowLeftCircleIcon,BriefcaseIcon } from "@heroicons/vue/24/outline";
import { defineComponent, onMounted, reactive, ref, computed } from "vue";
import { gql } from "graphql-tag";
import { CombinedError } from "@urql/core";

const USER_QUERY = gql`
  query {
    jobs {
    job_type
    title
    description
    applicationCount
    category
    }
  }
`;

export default defineComponent({
  setup() {
    const searchTerm = ref("");
    const client = useUrqlClient();
    const route = useRoute();

    const jobsData = reactive({
      jobs: [] as Array<any>,
      loading: false,
      error: null as CombinedError | null,
    });

    const job = computed(() => {
      return (
        jobsData.jobs.find(
          (item) => item.title === route.params.title
        ) || null
      );
    });

    async function loadJobs() {
      jobsData.loading = true;
      const response = await client.query(USER_QUERY, {}).toPromise();

      if (response.error) {
        jobsData.error = response.error;
      } else {
        jobsData.jobs = response.data.jobs;
      }

      jobsData.loading = false;
    }

    onMounted(() => {
      loadJobs();
    });

    return {
      jobsData,
      job,
      searchTerm,
    };
  },
});
</script>
