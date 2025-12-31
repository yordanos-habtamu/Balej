
<template>
  <div class="min-h-screen flex">
    <!-- Sidebar - lowered with pt-8 -->
    <aside class="flex-shrink-0 pt-8">
      <admin-sidebar />
    </aside>

    <!-- Main Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="max-w-7xl mx-auto px-6 py-10 lg:px-12 lg:py-14">
        <!-- Header with Search -->
        <div class="mb-12">
          <h1 class="text-3xl lg:text-4xl font-bold mb-8">Talent Directory</h1>
          <div class="relative max-w-3xl mx-auto">
            <input
              type="text"
              v-model="searchTerm"
              placeholder="Search by name, email, skills..."
              class="w-full px-6 py-4 pl-14 text-lg bg-gray-800 border border-gray-700 rounded-2xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.6] transition-shadow placeholder-gray-500"
            />
            <svg class="absolute left-5 top-1/2 -translate-y-1/2 w-7 h-7 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
        </div>

        <!-- Loading -->
        <div v-if="usersData.loading" class="flex flex-col items-center justify-center py-40">
          <div class="w-16 h-16 border-4 border-[color:var(--primary-color)] border-t-transparent rounded-full animate-spin mb-6"></div>
          <p class="text-xl text-gray-400">Loading talents...</p>
        </div>

        <!-- Error -->
        <div v-else-if="usersData.error" class="flex flex-col items-center justify-center py-40">
          <svg class="w-20 h-20 text-red-500 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <p class="text-xl text-red-400 font-medium">Failed to load talents</p>
        </div>

        <!-- Content Grid -->
        <div v-else class="grid grid-cols-1 lg:grid-cols-4 gap-10">
          <!-- Talent List -->
          <div class="lg:col-span-3 space-y-6">
            <div class="flex items-center justify-between mb-6">
              <p class="text-gray-400">{{ filteredUsers.length }} talents found</p>
            </div>

            <div v-if="filteredUsers.length === 0" class="text-center py-32 bg-gray-800/50 rounded-3xl">
              <p class="text-2xl text-gray-500">No talents match your criteria.</p>
              <p class="text-gray-600 mt-4">Try adjusting search or filters.</p>
            </div>

            <!-- Talent Card -->
            <div
              v-for="user in filteredUsers"
              :key="user.email"
              class="bg-gray-800 rounded-2xl border border-gray-700 overflow-hidden hover:border-gray-600 transition-all duration-300"
            >
              <div class="p-8">
                <div class="flex items-start gap-6">
                  <!-- Avatar -->
                  <div class="w-24 h-24 rounded-2xl overflow-hidden ring-4 ring-gray-700 flex-shrink-0">
                    <img src="~/assets/painter.png" alt="Talent avatar" class="w-full h-full object-cover" />
                  </div>

                  <!-- Info -->
                  <div class="flex-1 min-w-0">
                    <h3 class="text-2xl font-bold text-white mb-1">
                      {{ user.first_name }} {{ user.last_name }}
                    </h3>
                    <p class="text-gray-400 mb-4">{{ user.email }}</p>
                    <p class="text-gray-300 line-clamp-3">
                      Dedicated professional with proven expertise. Open to full-time, part-time, and contract opportunities.
                    </p>
                  </div>
                </div>
              </div>

              <!-- Actions -->
              <div class="px-8 pb-8 flex gap-4">
                <NuxtLink :to="`/admin/talents/${user.first_name}`" class="flex-1">
                  <button class="w-full py-3.5 bg-gray-700 hover:bg-gray-600 text-white rounded-xl font-medium transition flex items-center justify-center gap-3">
                    <svg class="w-5 h-5" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                    </svg>
                    View Profile
                  </button>
                </NuxtLink>
                <button
                  @click="isOpen = true"
                  class="flex-1 py-3.5 bg-[color:var(--primary-color)] hover:bg-[color:var(--primary-color)/0.9] text-white rounded-xl font-medium transition"
                >
                  Add to Catalogue
                </button>
              </div>
            </div>
          </div>

          <!-- Filters Sidebar - Sticky (not fixed overlay) -->
          <aside class="lg:sticky lg:top-10 self-start lg:w-full bg-gray-800 rounded-3xl border border-gray-700 p-8 space-y-10">
            <h2 class="text-2xl font-bold mb-2">Refine Results</h2>

            <!-- Job Type -->
            <div>
              <h3 class="text-lg font-semibold mb-5">Employment Type</h3>
              <div class="space-y-4">
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Full-time</span>
                </label>
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Part-time</span>
                </label>
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Contract</span>
                </label>
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Internship</span>
                </label>
              </div>
            </div>

            <!-- Category -->
            <div>
              <h3 class="text-lg font-semibold mb-5">Category</h3>
              <div class="space-y-4">
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Construction</span>
                </label>
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">House Keeping</span>
                </label>
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Repairing</span>
                </label>
                <label class="flex items-center gap-4 cursor-pointer">
                  <input type="checkbox" class="w-5 h-5 rounded border-gray-600 focus:ring-[color:var(--primary-color)] text-[color:var(--primary-color)]" />
                  <span class="text-gray-300">Management</span>
                </label>
              </div>
            </div>

            <!-- Price Range -->
            <div>
              <h3 class="text-lg font-semibold mb-5">Hourly Rate ($)</h3>
              <div class="grid grid-cols-2 gap-5">
                <input type="number" placeholder="Min" class="px-5 py-3 bg-gray-700 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)/0.6]" />
                <input type="number" placeholder="Max" class="px-5 py-3 bg-gray-700 rounded-xl focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)/0.6]" />
              </div>
            </div>

            <!-- Action Buttons -->
            <div class="pt-6 border-t border-gray-700">
              <button class="w-full py-4 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-xl font-semibold transition mb-4">
                Apply Filters
              </button>
              <button class="w-full py-4 bg-gray-700 hover:bg-gray-600 text-white rounded-xl font-semibold transition">
                Clear All
              </button>
            </div>
          </aside>
        </div>
      </div>
    </main>

    <!-- Success Modal -->
    <div v-if="isOpen" class="fixed inset-0 bg-black/70 flex items-center justify-center z-50 p-4" @click.self="isOpen = false">
      <div class="bg-gray-800 rounded-3xl shadow-2xl p-12 max-w-md w-full text-center">
        <div class="w-24 h-24 mx-auto mb-8 rounded-full bg-green-600/20 flex items-center justify-center">
          <svg class="w-14 h-14 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h2 class="text-4xl font-bold text-white mb-6">Success!</h2>
        <p class="text-xl text-gray-300">Talent added to catalogue.</p>
        <button @click="isOpen = false" class="mt-10 px-10 py-5 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-2xl font-semibold text-lg transition">
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent,onMounted, reactive, ref } from 'vue'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { defaultTypeResolver } from 'graphql'


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
`


export default defineComponent({
  setup(){
    const searchTerm = ref('')
    const isOpen = ref(false)
    const client= useUrqlClient();
    const usersData = reactive({
  users: [] as Array<any>,
  loading: false,
  error: null as CombinedError | null
});
  

   const filteredUsers = computed(() => {
      if (!searchTerm.value) return usersData.users;
      const lowerSearch = searchTerm.value.toLowerCase();
      return usersData.users.filter(user =>
        user.first_name?.toLowerCase().includes(lowerSearch) ||
        user.email?.toLowerCase().includes(lowerSearch)
      );
    });
// Fetch users from backend
async function loadUsers() {
  usersData.loading = true
  const response = await client.query(USER_QUERY, {}).toPromise()

  if (response.error) {
    usersData.error = response.error
  } else {
    usersData.users = response.data.users
  }

  usersData.loading = false
}

// Load users on component mount
onMounted(() => {
  loadUsers()
})
 
return {
    usersData,
    searchTerm,
    filteredUsers,
    isOpen
}
  },
});
// Reactive object to store users and state


</script>
