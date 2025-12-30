<template>
  <div class="min-h-screen bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
    <!-- Loading State -->
    <div v-if="fetching" class="flex justify-center items-center h-64">
      <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
    </div>
    
    <!-- Error State -->
    <div v-else-if="error" class="text-center text-red-600 bg-red-50 p-4 rounded-lg max-w-md mx-auto">
      <p class="font-semibold">Error loading profile</p>
      <p class="text-sm mt-1">{{ error.message }}</p>
    </div>

    <!-- Profile Content -->
    <div v-else-if="user" class="max-w-5xl mx-auto">
      <!-- Profile Header -->
      <div class="bg-white shadow-xl rounded-2xl overflow-hidden mb-8 transform transition-all hover:shadow-2xl duration-300">
        <div class="h-40 bg-gradient-to-r from-blue-600 via-purple-600 to-indigo-600 relative overflow-hidden">
            <div class="absolute inset-0 bg-black opacity-10"></div>
            <div class="absolute -bottom-10 -right-10 w-40 h-40 bg-white opacity-10 rounded-full blur-2xl"></div>
            <div class="absolute top-10 left-10 w-20 h-20 bg-white opacity-10 rounded-full blur-xl"></div>
        </div>
        <div class="px-8 pb-8 relative">
          <div class="relative -mt-20 mb-6 flex flex-col sm:flex-row items-center sm:items-end">
            <div class="w-40 h-40 rounded-full border-4 border-white shadow-lg overflow-hidden bg-white z-10">
              <img 
                :src="`https://ui-avatars.com/api/?name=${user.first_name}+${user.last_name}&background=random&size=256`" 
                alt="Profile" 
                class="w-full h-full object-cover"
              />
            </div>
            <div class="mt-4 sm:mt-0 sm:ml-6 text-center sm:text-left mb-2">
              <h1 class="text-3xl font-bold text-gray-900">{{ user.first_name }} {{ user.last_name }}</h1>
              <p class="text-blue-600 font-medium text-lg uppercase tracking-wide">{{ user.role }}</p>
            </div>
            <div class="flex-grow"></div>
            <div class="mt-4 sm:mt-0">
                <button class="bg-blue-600 text-white px-6 py-2.5 rounded-full hover:bg-blue-700 transition-all shadow-md hover:shadow-lg flex items-center gap-2 font-medium">
                    <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                        <path d="M13.586 3.586a2 2 0 112.828 2.828l-.793.793-2.828-2.828.793-.793zM11.379 5.793L3 14.172V17h2.828l8.38-8.379-2.83-2.828z" />
                    </svg>
                    Edit Profile
                </button>
            </div>
          </div>
          
        </div>
      </div>

      <!-- Info Grid -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8">
        <!-- Contact Info -->
        <div class="bg-white rounded-2xl shadow-md p-8 hover:shadow-lg transition-shadow border border-gray-100">
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="bg-blue-100 p-2.5 rounded-xl mr-3 text-blue-600">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 5.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
              </svg>
            </span>
            Contact Information
          </h2>
          <div class="space-y-5">
            <div class="group flex items-center p-3 rounded-lg hover:bg-gray-50 transition-colors">
              <div class="w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center mr-4 text-gray-500 group-hover:bg-blue-100 group-hover:text-blue-600 transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                      <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
                      <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
                  </svg>
              </div>
              <div>
                  <p class="text-sm text-gray-500 font-medium">Email Address</p>
                  <p class="text-gray-900 font-medium">{{ user.email }}</p>
              </div>
            </div>

            <div class="group flex items-center p-3 rounded-lg hover:bg-gray-50 transition-colors">
               <div class="w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center mr-4 text-gray-500 group-hover:bg-blue-100 group-hover:text-blue-600 transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                      <path d="M2 3a1 1 0 011-1h2.153a1 1 0 01.986.836l.74 4.435a1 1 0 01-.54 1.06l-1.548.773a11.037 11.037 0 006.105 6.105l.774-1.548a1 1 0 011.059-.54l4.435.74a1 1 0 01.836.986V17a1 1 0 01-1 1h-2C7.82 18 2 12.18 2 5V3z" />
                  </svg>
              </div>
              <div>
                  <p class="text-sm text-gray-500 font-medium">Phone Number</p>
                  <p class="text-gray-900 font-medium">{{ user.phone_number }}</p>
              </div>
            </div>

            <div class="group flex items-center p-3 rounded-lg hover:bg-gray-50 transition-colors">
               <div class="w-10 h-10 rounded-full bg-gray-100 flex items-center justify-center mr-4 text-gray-500 group-hover:bg-blue-100 group-hover:text-blue-600 transition-colors">
                  <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M5.05 4.05a7 7 0 119.9 9.9L10 18.9l-4.95-4.95a7 7 0 010-9.9zM10 11a2 2 0 100-4 2 2 0 000 4z" clip-rule="evenodd" />
                  </svg>
              </div>
              <div>
                  <p class="text-sm text-gray-500 font-medium">Address</p>
                  <p class="text-gray-900 font-medium">{{ user.address }}</p>
              </div>
            </div>
          </div>
        </div>

        <!-- Account Status -->
        <div class="bg-white rounded-2xl shadow-md p-8 hover:shadow-lg transition-shadow border border-gray-100">
          <h2 class="text-xl font-bold text-gray-800 mb-6 flex items-center">
            <span class="bg-green-100 p-2.5 rounded-xl mr-3 text-green-600">
              <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
              </svg>
            </span>
            Account Status
          </h2>
          <div class="space-y-4">
            <div class="flex items-center justify-between p-4 bg-gray-50 rounded-xl border border-gray-100">
              <span class="text-gray-600 font-medium">Current Status</span>
              <span :class="user.is_active ? 'bg-green-100 text-green-700 border border-green-200' : 'bg-red-100 text-red-700 border border-red-200'" class="px-4 py-1.5 rounded-full text-sm font-bold flex items-center gap-2">
                <span class="w-2 h-2 rounded-full" :class="user.is_active ? 'bg-green-500' : 'bg-red-500'"></span>
                {{ user.is_active ? 'Active' : 'Inactive' }}
              </span>
            </div>
            <div class="flex items-center justify-between p-4 bg-gray-50 rounded-xl border border-gray-100">
              <span class="text-gray-600 font-medium">Member ID</span>
              <span class="text-gray-900 font-mono text-sm bg-white px-3 py-1 rounded border border-gray-200 shadow-sm">#{{ user.id }}</span>
            </div>
             <div class="mt-6 p-4 bg-blue-50 rounded-xl border border-blue-100">
                <h3 class="text-blue-800 font-semibold mb-2">Profile Completion</h3>
                <div class="w-full bg-blue-200 rounded-full h-2.5">
                    <div class="bg-blue-600 h-2.5 rounded-full" style="width: 85%"></div>
                </div>
                <p class="text-xs text-blue-600 mt-2 text-right font-medium">85% Complete</p>
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