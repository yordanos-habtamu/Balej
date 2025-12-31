<template>
  <div :class="{'dark': themeStore.isDarkMode}" class="min-h-screen transition-colors duration-300">
    <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8 transition-colors duration-300">
    <!-- Loading State -->
    <div v-if="fetching" class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2" :class="`border-${themeStore.primaryColor}-600`"></div>
    </div>
    
    <!-- Error State -->
    <div v-else-if="error" class="text-center text-red-600 bg-red-50 p-4 rounded-lg max-w-md mx-auto">
      <p class="font-semibold">Error loading profile</p>
      <p class="text-sm mt-1">{{ error.message }}</p>
    </div>

    <!-- Profile Content -->
    <div v-else-if="user" class="max-w-6xl mx-auto">
      <!-- Admin Header -->
      <div class="bg-white dark:bg-gray-800 shadow-2xl rounded-xl overflow-hidden mb-8 border-t-4" :class="`border-${themeStore.primaryColor}-600`">
        <div class="h-32 relative overflow-hidden" :class="`bg-${themeStore.primaryColor}-800`">
            <div class="absolute inset-0 opacity-20 bg-[url('https://www.transparenttextures.com/patterns/carbon-fibre.png')]"></div>
            <div class="absolute top-4 right-4 text-white opacity-50">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-16 w-16" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1" d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z" />
                </svg>
            </div>
        </div>
        <div class="px-8 pb-6 relative">
          <div class="relative -mt-16 mb-4 flex flex-col sm:flex-row items-end">
            <div class="w-32 h-32 rounded-lg border-4 border-white dark:border-gray-700 shadow-lg overflow-hidden bg-gray-200 z-10">
              <img 
                :src="`https://ui-avatars.com/api/?name=${user.first_name}+${user.last_name}&background=1f2937&color=fff&size=256&bold=true`" 
                alt="Profile" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="mt-4 sm:mt-0 sm:ml-6 mb-1">
              <h1 class="text-3xl font-bold text-gray-900 dark:text-white">{{ user.first_name }} {{ user.last_name }}</h1>
              <div class="flex items-center mt-1 text-gray-600 dark:text-gray-300">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M2.166 4.999A11.954 11.954 0 0010 1.944 11.954 11.954 0 0017.834 5c.11.65.166 1.32.166 2.001 0 5.225-3.34 9.67-8 11.317C5.34 16.67 2 12.225 2 7c0-.682.057-1.35.166-2.001zm11.541 3.708a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clip-rule="evenodd" />
                  </svg>
                  <span class="font-semibold uppercase tracking-wider text-sm">System Administrator</span>
              </div>
            </div>
            <div class="flex-grow"></div>
            <div class="mt-4 sm:mt-0 flex gap-3">
                 <nuxt-link to="/admin/settings" class="text-white px-5 py-2 rounded-md transition-all shadow flex items-center gap-2 text-sm font-medium" :class="`bg-${themeStore.primaryColor}-900 hover:bg-${themeStore.primaryColor}-800`">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                    Settings
                </nuxt-link>
            </div>
          </div>
        </div>
      </div>

      <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
        <!-- Left Column: Personal Info -->
        <div class="lg:col-span-2 space-y-8">
            <div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                <h3 class="text-lg font-bold text-gray-800 dark:text-white mb-4 border-b dark:border-gray-700 pb-2">Administrator Details</h3>
                <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                     <div>
                        <label class="block text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Email</label>
                        <p class="text-gray-900 dark:text-gray-200 font-medium mt-1">{{ user.email }}</p>
                    </div>
                    <div>
                        <label class="block text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Phone</label>
                        <p class="text-gray-900 dark:text-gray-200 font-medium mt-1">{{ user.phone_number }}</p>
                    </div>
                     <div>
                        <label class="block text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Role ID</label>
                        <p class="text-gray-900 dark:text-gray-200 font-medium mt-1 font-mono bg-gray-100 dark:bg-gray-700 inline-block px-2 py-0.5 rounded text-sm">ADMIN-{{ user.id }}</p>
                    </div>
                    <div>
                        <label class="block text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Status</label>
                        <span class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-green-100 text-green-800 mt-1">
                            Active
                        </span>
                    </div>
                     <div class="md:col-span-2">
                        <label class="block text-xs font-semibold text-gray-500 dark:text-gray-400 uppercase">Office Address</label>
                        <p class="text-gray-900 dark:text-gray-200 font-medium mt-1">{{ user.address }}</p>
                    </div>
                </div>
            </div>
        </div>

        <!-- Right Column: Quick Actions / Stats -->
        <div class="space-y-6">
            <!-- Quick Stats -->
             <div class="text-white rounded-xl shadow-lg p-6" :class="`bg-${themeStore.primaryColor}-900`">
                <h3 class="text-lg font-bold mb-4">System Overview</h3>
                <div class="space-y-4">
                    <div class="flex justify-between items-center border-b border-gray-700 pb-2">
                        <span class="text-gray-300">Total Users</span>
                        <span class="font-bold text-xl">1,234</span>
                    </div>
                    <div class="flex justify-between items-center border-b border-gray-700 pb-2">
                        <span class="text-gray-300">Active Jobs</span>
                        <span class="font-bold text-xl text-green-400">56</span>
                    </div>
                    <div class="flex justify-between items-center">
                        <span class="text-gray-300">Pending Reviews</span>
                        <span class="font-bold text-xl text-yellow-400">12</span>
                    </div>
                </div>
            </div>

            <!-- Quick Actions -->
            <div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm border border-gray-200 dark:border-gray-700 p-6">
                 <h3 class="text-lg font-bold text-gray-800 dark:text-white mb-4">Management Console</h3>
                 <div class="space-y-3">
                     <button class="w-full text-left px-4 py-3 rounded-lg bg-gray-50 dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors flex items-center group">
                        <div class="bg-blue-100 dark:bg-blue-900 text-blue-600 dark:text-blue-300 p-2 rounded-md mr-3 group-hover:bg-blue-200 dark:group-hover:bg-blue-800">
                             <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z" />
                            </svg>
                        </div>
                        <span class="font-medium text-gray-700 dark:text-gray-200">Manage Users</span>
                     </button>
                     <button class="w-full text-left px-4 py-3 rounded-lg bg-gray-50 dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors flex items-center group">
                        <div class="bg-purple-100 dark:bg-purple-900 text-purple-600 dark:text-purple-300 p-2 rounded-md mr-3 group-hover:bg-purple-200 dark:group-hover:bg-purple-800">
                             <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2m-3 7h3m-3 4h3m-6-4h.01M9 16h.01" />
                            </svg>
                        </div>
                        <span class="font-medium text-gray-700 dark:text-gray-200">View Reports</span>
                     </button>
                     <button class="w-full text-left px-4 py-3 rounded-lg bg-gray-50 dark:bg-gray-700 hover:bg-gray-100 dark:hover:bg-gray-600 transition-colors flex items-center group">
                        <div class="bg-red-100 dark:bg-red-900 text-red-600 dark:text-red-300 p-2 rounded-md mr-3 group-hover:bg-red-200 dark:group-hover:bg-red-800">
                             <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-3L13.732 4c-.77-1.333-2.694-1.333-3.464 0L3.34 16c-.77 1.333.192 3 1.732 3z" />
                            </svg>
                        </div>
                        <span class="font-medium text-gray-700 dark:text-gray-200">System Alerts</span>
                     </button>
                 </div>
            </div>
        </div>
      </div>
    </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useCookie } from '#app'
import { jwtDecode } from 'jwt-decode'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { useThemeStore } from '~/stores/theme'

const themeStore = useThemeStore()
const tokenCookie = useCookie('token')
const client = useUrqlClient()
const fetching = ref(true)
const error = ref<CombinedError | null>(null)
const user = ref<any>(null)
const userId = ref<string | null>(null)

// Decode token to get user ID
if (tokenCookie.value) {
  try {
    const decoded: any = jwtDecode(tokenCookie.value)
    // Check for common ID fields in JWT
    userId.value = decoded.user_id || decoded.id || decoded.sub 
    console.log("Decoded User ID:", userId.value)
  } catch (e) {
    console.error('Failed to decode token', e)
    error.value = { message: "Invalid token" } as any
    fetching.value = false
  }
} else {
    fetching.value = false
    error.value = { message: "No token found" } as any
}

const UserQuery = gql`
  query GetUser($id: ID!) {
    user(id: $id) {
      id
      first_name
      last_name
      email
      phone_number
      address
      role
      is_active
    }
  }
`

const fetchUser = async () => {
    if (!userId.value) return

    fetching.value = true
    error.value = null
    
    try {
        const result = await client
        .query(UserQuery, { id: userId.value })
        .toPromise()
    
        if (result.error) {
            error.value = result.error
        } else {
            user.value = result.data?.user
        }
    } catch (err: any) {
        error.value = err
    } finally {
        fetching.value = false
    }
}

onMounted(() => {
    fetchUser()
})

</script>