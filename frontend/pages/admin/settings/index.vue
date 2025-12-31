<template>
  <div :class="{'dark': themeStore.isDarkMode}" class="min-h-screen transition-colors duration-300" :style="pageStyle">
    <div class="min-h-screen bg-gray-50 dark:bg-gray-900 py-12 px-4 sm:px-6 lg:px-8 transition-colors duration-300">
      <div class="max-w-4xl mx-auto">
        <!-- Header -->
        <div class="mb-8 flex items-center justify-between">
          <div>
            <h1 class="text-3xl font-bold text-gray-900 dark:text-white">Settings</h1>
            <p class="mt-2 text-gray-600 dark:text-gray-400">Manage your dashboard preferences and system configurations.</p>
          </div>
          <nuxt-link to="/admin/profile" class="text-sm font-medium text-gray-500 hover:text-gray-700 dark:text-gray-400 dark:hover:text-gray-200 flex items-center">
            <svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-1" viewBox="0 0 20 20" fill="currentColor">
              <path fill-rule="evenodd" d="M9.707 16.707a1 1 0 01-1.414 0l-6-6a1 1 0 010-1.414l6-6a1 1 0 011.414 1.414L5.414 9H17a1 1 0 110 2H5.414l4.293 4.293a1 1 0 010 1.414z" clip-rule="evenodd" />
            </svg>
            Back to Profile
          </nuxt-link>
        </div>

        <div class="grid grid-cols-1 gap-8">
          <!-- Appearance Settings -->
          <div class="bg-white dark:bg-gray-800 shadow rounded-xl overflow-hidden">
            <div class="p-6 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-purple-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a4 4 0 01-4-4V5a2 2 0 012-2h4a2 2 0 012 2v12a4 4 0 01-4 4zm0 0h12a2 2 0 002-2v-4a2 2 0 00-2-2h-2.343M11 7.343l1.657-1.657a2 2 0 012.828 0l2.829 2.829a2 2 0 010 2.828l-8.486 8.485M7 17h.01" />
                </svg>
                Appearance
              </h2>
              <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Customize the look and feel of your dashboard.</p>
            </div>
            <div class="p-6 space-y-6">
              <!-- Dark Mode Toggle -->
              <div class="flex items-center justify-between">
                <div>
                  <h3 class="text-base font-medium text-gray-900 dark:text-white">Dark Mode</h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">Switch between light and dark themes.</p>
                </div>
                <button 
                  @click="themeStore.toggleDarkMode"
                  :class="themeStore.isDarkMode ? 'bg-purple-600' : 'bg-gray-200'"
                  class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-purple-500"
                >
                  <span 
                    :class="themeStore.isDarkMode ? 'translate-x-5' : 'translate-x-0'"
                    class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200"
                  ></span>
                </button>
              </div>

              <!-- Color Theme -->
              <div>
                <h3 class="text-base font-medium text-gray-900 dark:text-white mb-3">Accent Color</h3>
                <div class="flex space-x-4">
                  <button 
                    v-for="color in ['blue', 'green', 'purple', 'red', 'orange', 'teal']" 
                    :key="color"
                    @click="themeStore.setPrimaryColor(color)"
                    class="w-10 h-10 rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 dark:focus:ring-offset-gray-800 transition-transform hover:scale-110 flex items-center justify-center"
                    :class="[`bg-${color}-500`, themeStore.primaryColor === color ? 'ring-2 ring-offset-2 ring-gray-400 dark:ring-white scale-110' : '']"
                  >
                    <svg v-if="themeStore.primaryColor === color" xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 text-white" viewBox="0 0 20 20" fill="currentColor">
                      <path fill-rule="evenodd" d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z" clip-rule="evenodd" />
                    </svg>
                  </button>
                </div>
              </div>
            </div>
          </div>

          <!-- Notifications & Alerts -->
          <div class="bg-white dark:bg-gray-800 shadow rounded-xl overflow-hidden">
            <div class="p-6 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-yellow-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9" />
                </svg>
                Notifications
              </h2>
              <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Manage how you receive alerts.</p>
            </div>
            <div class="p-6 space-y-4">
               <div class="flex items-center justify-between">
                <div>
                  <h3 class="text-base font-medium text-gray-900 dark:text-white">Enable Notifications</h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">Receive system alerts and updates.</p>
                </div>
                <button 
                  @click="themeStore.toggleNotifications"
                  :class="themeStore.notificationsEnabled ? 'bg-green-500' : 'bg-gray-200'"
                  class="relative inline-flex flex-shrink-0 h-6 w-11 border-2 border-transparent rounded-full cursor-pointer transition-colors ease-in-out duration-200 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-green-500"
                >
                  <span 
                    :class="themeStore.notificationsEnabled ? 'translate-x-5' : 'translate-x-0'"
                    class="pointer-events-none inline-block h-5 w-5 rounded-full bg-white shadow transform ring-0 transition ease-in-out duration-200"
                  ></span>
                </button>
              </div>
              <div class="flex items-center justify-between opacity-75">
                <div>
                  <h3 class="text-base font-medium text-gray-900 dark:text-white">Email Digest</h3>
                  <p class="text-sm text-gray-500 dark:text-gray-400">Receive a daily summary of activities.</p>
                </div>
                <input type="checkbox" class="h-5 w-5 text-blue-600 rounded focus:ring-blue-500 border-gray-300" checked disabled>
              </div>
            </div>
          </div>

          <!-- Security Settings -->
          <div class="bg-white dark:bg-gray-800 shadow rounded-xl overflow-hidden">
            <div class="p-6 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-red-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
                Security
              </h2>
              <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Update password and secure your account.</p>
            </div>
            <div class="p-6 space-y-6">
              <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">New Password</label>
                  <input type="password" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white sm:text-sm p-2 border" placeholder="••••••••">
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Confirm Password</label>
                  <input type="password" class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white sm:text-sm p-2 border" placeholder="••••••••">
                </div>
              </div>
              <div class="flex justify-end">
                <button class="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700 transition-colors text-sm font-medium">
                  Update Password
                </button>
              </div>
              
              <div class="border-t border-gray-200 dark:border-gray-700 pt-6">
                 <div class="flex items-center justify-between">
                  <div>
                    <h3 class="text-base font-medium text-gray-900 dark:text-white">Two-Factor Authentication</h3>
                    <p class="text-sm text-gray-500 dark:text-gray-400">Add an extra layer of security.</p>
                  </div>
                  <button class="text-blue-600 hover:text-blue-800 font-medium text-sm">Enable 2FA</button>
                </div>
              </div>
            </div>
          </div>
          
           <!-- System Preferences -->
          <div class="bg-white dark:bg-gray-800 shadow rounded-xl overflow-hidden">
            <div class="p-6 border-b border-gray-200 dark:border-gray-700">
              <h2 class="text-xl font-semibold text-gray-900 dark:text-white flex items-center">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6 mr-2 text-gray-500" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z" />
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 11-6 0 3 3 0 016 0z" />
                </svg>
                System Preferences
              </h2>
              <p class="mt-1 text-sm text-gray-500 dark:text-gray-400">Global system settings.</p>
            </div>
            <div class="p-6 space-y-6">
               <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Language</label>
                  <select 
                    v-model="themeStore.language"
                    class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white sm:text-sm p-2 border"
                  >
                    <option value="en">English</option>
                    <option value="es">Spanish</option>
                    <option value="fr">French</option>
                    <option value="de">German</option>
                  </select>
                </div>
                <div>
                  <label class="block text-sm font-medium text-gray-700 dark:text-gray-300">Timezone</label>
                  <select class="mt-1 block w-full rounded-md border-gray-300 shadow-sm focus:border-blue-500 focus:ring-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:text-white sm:text-sm p-2 border">
                    <option>UTC (GMT+00:00)</option>
                    <option>EST (GMT-05:00)</option>
                    <option>PST (GMT-08:00)</option>
                    <option>CET (GMT+01:00)</option>
                  </select>
                </div>
              </div>
              
              <div class="flex items-center justify-between pt-4 border-t border-gray-200 dark:border-gray-700">
                  <button class="text-red-600 hover:text-red-800 text-sm font-medium flex items-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16" />
                      </svg>
                      Clear System Cache
                  </button>
                  <button class="text-gray-600 hover:text-gray-800 dark:text-gray-400 dark:hover:text-gray-200 text-sm font-medium flex items-center">
                      <svg xmlns="http://www.w3.org/2000/svg" class="h-4 w-4 mr-1" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                          <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-4l-4 4m0 0l-4-4m4 4V4" />
                      </svg>
                      Export Data
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
import { useThemeStore } from '~/stores/theme'
import { computed } from 'vue'

const themeStore = useThemeStore()

// Computed property to dynamically generate styles based on primary color
const pageStyle = computed(() => {
    // This is a simplified way to handle dynamic colors. 
    // In a real app with Tailwind, you might map 'blue' to specific hex codes or use CSS variables.
    // For now, we rely on the Tailwind classes like bg-blue-500 which are dynamically applied in the template.
    return {}
})
</script>
