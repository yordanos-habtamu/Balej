<template>
  <div class="min-h-screen flex">
    <!-- Sidebar -->
    <aside class="flex-shrink-0 pt-8">
      <admin-sidebar />
    </aside>

    <!-- Main Content -->
    <main class="flex-1 bg-gray-900 text-gray-100 overflow-y-auto">
      <div class="p-6 md:p-8 lg:p-10 max-w-7xl mx-auto">
        <!-- Header -->
        <div class="mb-12">
          <h1 class="text-4xl font-bold mb-2">Clients Directory</h1>
          <p class="text-xl text-gray-400">Manage and connect with your valued clients</p>
        </div>

        <!-- Search Bar -->
        <div class="mb-10">
          <div class="relative max-w-3xl mx-auto">
            <input
              type="text"
              v-model="searchTerm"
              placeholder="Search clients by name, email, or company..."
              class="w-full px-8 py-5 pl-16 text-xl bg-gray-800 border border-gray-700 rounded-3xl focus:outline-none focus:ring-4 focus:ring-[color:var(--primary-color)/0.5] transition-shadow placeholder-gray-500"
            />
            <svg class="absolute left-6 top-1/2 -translate-y-1/2 w-8 h-8 text-gray-500" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
            </svg>
          </div>
        </div>

        <!-- Loading -->
        <div v-if="clientData.loading" class="flex flex-col items-center justify-center py-40">
          <div class="w-16 h-16 border-4 border-[color:var(--primary-color)] border-t-transparent rounded-full animate-spin mb-6"></div>
          <p class="text-xl text-gray-400">Loading clients...</p>
        </div>

        <!-- Error -->
        <div v-else-if="clientData.error" class="flex flex-col items-center justify-center py-40">
          <svg class="w-20 h-20 text-red-500 mb-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
          </svg>
          <p class="text-xl text-red-400 font-medium">Failed to load clients</p>
        </div>

        <!-- Clients List -->
        <div v-else>
          <div class="flex items-center justify-between mb-6">
            <p class="text-lg text-gray-400">{{ filteredUsers.length }} client{{ filteredUsers.length !== 1 ? 's' : '' }} found</p>
          </div>

          <div v-if="filteredUsers.length === 0" class="text-center py-32 bg-gray-800/50 rounded-3xl">
            <p class="text-2xl text-gray-500">No clients match your search.</p>
          </div>

          <!-- Compact Client Cards -->
          <div class="space-y-4">
            <div
              v-for="user in filteredUsers"
              :key="user.email"
              class="bg-gray-800 rounded-xl shadow-lg border border-gray-700 overflow-hidden transition-all duration-300 hover:shadow-[0_0_20px_rgba(var(--primary-color-rgb),0.15)] hover:border-[color:var(--primary-color)/0.3] flex items-center gap-6 p-5"
            >
              <!-- Avatar -->
              <div class="w-16 h-16 rounded-xl overflow-hidden ring-4 ring-gray-700 flex-shrink-0">
                <img src="~/assets/painter.png" alt="Client" class="w-full h-full object-cover" />
              </div>

              <!-- Info -->
              <div class="flex-1 min-w-0">
                <div class="flex items-center gap-3">
                  <h3 class="text-lg font-bold text-white truncate">
                    {{ user.first_name }} {{ user.last_name }}
                  </h3>
                  <span class="px-3 py-1 bg-[color:var(--primary-color)/0.2] text-[color:var(--primary-color)] rounded-full text-xs font-medium">
                    Premium
                  </span>
                </div>
                <p class="text-sm text-gray-400 truncate mt-1">{{ user.email }}</p>
                <p class="text-sm text-gray-500 mt-1">TechCorp Solutions • Active since 2024</p>
              </div>

              <!-- Actions -->
              <div class="flex gap-3">
                <NuxtLink to="/admin/chats" class="flex-shrink-0">
                  <button class="px-5 py-2.5 bg-[color:var(--primary-color)] hover:bg-[color:var(--primary-color)/0.9] text-white rounded-lg font-medium transition flex items-center gap-2 text-sm">
                    <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
                    </svg>
                    Chat
                  </button>
                </NuxtLink>
               <NuxtLink :to="`/admin/clients/${user.first_name}`"><p
                 
                  class="px-5 py-2.5 bg-gray-700 hover:bg-gray-600 text-white rounded-lg font-medium transition text-sm"
                >
                  Projects
              </p></NuxtLink>
              </div>
            </div>
          </div>
        </div>
      </div>
    </main>

    <!-- Success Modal -->
    <div v-if="isOpen" class="fixed inset-0 bg-black/70 flex items-center justify-center z-50 p-4" @click.self="isOpen = false">
      <div class="bg-gray-800 rounded-3xl shadow-2xl p-10 max-w-md w-full text-center">
        <div class="w-20 h-20 mx-auto mb-6 rounded-full bg-green-600/20 flex items-center justify-center">
          <svg class="w-12 h-12 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M5 13l4 4L19 7" />
          </svg>
        </div>
        <h2 class="text-3xl font-bold text-white mb-4">Client Added!</h2>
        <p class="text-lg text-gray-300">Successfully added to your catalogue.</p>
        <button @click="isOpen = false" class="mt-8 px-8 py-4 bg-[color:var(--primary-color)] hover:opacity-90 text-white rounded-2xl font-semibold transition">
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* ✅ Always hide scrollbars */
.hide-scrollbar {
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE 10+ */
}
.hide-scrollbar::-webkit-scrollbar {
  display: none; /* Chrome, Safari, Opera */
}

/* ✅ Mobile responsive tweaks */
@media (max-width: 1024px) {
  aside {
    position: static;
    width: 100%;
    max-height: none !important;
    margin-bottom: 1rem;
  }

  input[type="text"],
  input[type="number"] {
    padding: 0.5rem !important;
    font-size: 1rem !important;
  }

  button {
    padding: 0.75rem 1.25rem !important;
    font-size: 1rem !important;
  }

  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}

/* Stack buttons vertically on very small screens */
@media (max-width: 640px) {
  .card-actions {
    flex-direction: column;
    gap: 0.5rem;
  }
  .card-actions button {
    width: 100%;
  }
  /* Make profile images smaller */
  img {
    max-width: 48px;
    max-height: 48px;
  }
}


.modal {
  position: fixed;
  z-index: 1000;
  inset: 0;
  background-color: rgba(0,0,0,0.6);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  width: 400px;
  max-width: 90%;
  text-align: center;
  box-shadow: 0 5px 15px rgba(0,0,0,0.3);
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; transform: scale(1); }
}

.close-btn {
  float: right;
  font-size: 20px;
  font-weight: bold;
  cursor: pointer;
  color: #333;
}
.close-btn:hover {
  color: red;
}
</style>


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
      role
    }
  }
`


export default defineComponent({
  setup(){
    const searchTerm = ref('')
    const isOpen = ref(false)
    const client= useUrqlClient();
 
const clientData = reactive({
  users: [] as Array<any>,
  loading: false,
  error: null as CombinedError | null
})
  

   const filteredUsers = computed(() => {
      if (!searchTerm.value) return clientData.users;
      const lowerSearch = searchTerm.value.toLowerCase();
      return clientData.users.filter(user =>
        user.first_name?.toLowerCase().includes(lowerSearch) ||
        user.email?.toLowerCase().includes(lowerSearch)
      );
    });
// Fetch users from backend
async function loadUsers() {
  clientData.loading = true
  const response = await client.query(USER_QUERY, {}).toPromise()

  if (response.error) {
    clientData.error = response.error
  } else {
    clientData.users = response.data.users
  }

  clientData.loading = false
}

// Load users on component mount
onMounted(() => {
  loadUsers()
})
 
return {
    clientData,
    searchTerm,
    filteredUsers,
    isOpen
}
  },
});
// Reactive object to store users and state


</script>
