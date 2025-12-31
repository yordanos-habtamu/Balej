<template>
  <div class="min-h-screen flex">
   

    <!-- Main Profile Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="max-w-5xl mx-auto px-6 py-10 lg:px-12 lg:py-14">
        <!-- Back Button -->
        <div class="mb-10 ml-0 w-100">
          <NuxtLink to="/admin/talents" class="inline-flex items-center gap-3 text-gray-400 hover:text-[color:var(--primary-color)] transition">
            <ArrowLeftCircleIcon class="w-8 h-8" />
            <span class="text-lg font-medium">Back to Talents</span>
          </NuxtLink>
        </div>

        <!-- Loading State -->
        <div v-if="usersData.loading" class="flex flex-col items-center justify-center py-40">
          <div class="w-16 h-16 border-4 border-[color:var(--primary-color)] border-t-transparent rounded-full animate-spin mb-6"></div>
          <p class="text-xl text-gray-400">Loading talent profile...</p>
        </div>

        <!-- Error State -->
        <div v-else-if="usersData.error" class="flex flex-col items-center justify-center py-40">
          <svg class="w-20 h-20 text-red-500 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <p class="text-xl text-red-400 font-medium">Failed to load profile</p>
          <p class="text-gray-500 mt-3">{{ usersData.error.message }}</p>
        </div>

        <!-- Profile View -->
        <div v-else-if="user" class="space-y-12">
          <!-- Header with Avatar -->
          <div class="text-center">
            <div class="inline-block relative">
              <div class="w-48 h-48 rounded-full overflow-hidden ring-8 ring-[color:var(--primary-color)/0.2] shadow-2xl">
                <img
                  src="~/assets/painter.png"
                  alt="Talent profile"
                  class="w-full h-full object-cover"
                />
              </div>
            </div>
            <h1 class="text-4xl lg:text-5xl font-bold mt-8">{{ user.first_name }} {{ user.last_name }}</h1>
            <p class="text-xl text-gray-400 mt-3">{{ user.email }}</p>
          </div>

          <!-- Personal Information Card -->
          <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-10 lg:p-14">
            <h2 class="text-2xl font-bold mb-10 flex items-center gap-4">
              <svg class="w-8 h-8 text-[color:var(--primary-color)]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z" />
              </svg>
              Personal Information
            </h2>

            <div class="grid grid-cols-1 md:grid-cols-2 gap-10">
              <div class="space-y-8">
                <div>
                  <p class="text-sm font-medium text-gray-400 uppercase tracking-wider">First Name</p>
                  <p class="text-2xl font-semibold text-white mt-2">{{ user.first_name }}</p>
                </div>
                <div>
                  <p class="text-sm font-medium text-gray-400 uppercase tracking-wider">Last Name</p>
                  <p class="text-2xl font-semibold text-white mt-2">{{ user.last_name }}</p>
                </div>
                <div>
                  <p class="text-sm font-medium text-gray-400 uppercase tracking-wider">Address</p>
                  <p class="text-xl text-gray-300 mt-2">{{ user.address || 'Not provided' }}</p>
                </div>
              </div>

              <div class="space-y-8">
                <div>
                  <p class="text-sm font-medium text-gray-400 uppercase tracking-wider">Email</p>
                  <p class="text-xl text-gray-300 mt-2 break-all">{{ user.email }}</p>
                </div>
                <div>
                  <p class="text-sm font-medium text-gray-400 uppercase tracking-wider">Phone Number</p>
                  <p class="text-xl text-gray-300 mt-2">{{ user.phone_number || 'Not provided' }}</p>
                </div>
              </div>
            </div>
          </div>

          <!-- Professional Summary -->
          <div class="bg-gray-800 rounded-3xl shadow-2xl border border-gray-700 p-10 lg:p-14">
            <h2 class="text-2xl font-bold mb-8 flex items-center gap-4">
              <svg class="w-8 h-8 text-[color:var(--primary-color)]" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
              Professional Summary
            </h2>
            <p class="text-lg text-gray-300 leading-relaxed">
              Experienced and reliable professional with a strong track record in delivering high-quality work. Skilled in various domains and committed to excellence. Open to new opportunities and collaborations.
            </p>
          </div>

          <!-- Action Buttons -->
          <div class="flex flex-col sm:flex-row gap-6 justify-center max-w-2xl mx-auto">
            <NuxtLink to="/admin/chats" class="flex-1">
              <button class="w-full py-5 bg-[color:var(--primary-color)] hover:bg-[color:var(--primary-color)/0.9] text-white rounded-2xl shadow-xl font-semibold text-lg transition flex items-center justify-center gap-4">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                </svg>
                Start Chat
              </button>
            </NuxtLink>
            <button class="flex-1 py-5 bg-green-600 hover:bg-green-700 text-white rounded-2xl shadow-xl font-semibold text-lg transition flex items-center justify-center gap-4">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a2 2 0 012-2h3.28a1 1 0 01.948.684l1.498 4.493a1 1 0 01-.502 1.21l-2.257 1.13a11.042 11.042 0 005.516 5.516l1.13-2.257a1 1 0 011.21-.502l4.493 1.498a1 1 0 01.684.949V19a2 2 0 01-2 2h-1C9.716 21 3 14.284 3 6V5z" />
              </svg>
              Call Talent
            </button>
          </div>
        </div>

        <!-- Not Found -->
        <div v-else class="flex flex-col items-center justify-center py-40">
          <svg class="w-24 h-24 text-gray-600 mb-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <p class="text-2xl font-medium text-gray-400">Talent not found</p>
        </div>
      </div>
    </main>
  </div>
</template>

<script lang="ts">
import { ArrowLeftCircleIcon } from "@heroicons/vue/24/outline";
import { defineComponent, onMounted, reactive, ref, computed } from "vue";
import { gql } from "graphql-tag";
import { CombinedError } from "@urql/core";

const USER_QUERY = gql`
  query {
    users {
      first_name
      last_name
      email
      phone_number
      address
    }
  }
`;

export default defineComponent({
  setup() {
    const searchTerm = ref("");
    const client = useUrqlClient();
    const route = useRoute();

    const usersData = reactive({
      users: [] as Array<any>,
      loading: false,
      error: null as CombinedError | null,
    });

    const user = computed(() => {
      return (
        usersData.users.find(
          (item) => item.first_name === route.params.name
        ) || null
      );
    });

    async function loadUsers() {
      usersData.loading = true;
      const response = await client.query(USER_QUERY, {}).toPromise();

      if (response.error) {
        usersData.error = response.error;
      } else {
        usersData.users = response.data.users;
      }

      usersData.loading = false;
    }

    onMounted(() => {
      loadUsers();
    });

    return {
      usersData,
      user,
      searchTerm,
    };
  },
});
</script>
