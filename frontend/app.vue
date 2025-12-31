<template>
  <div :class="{'dark': themeStore.isDarkMode}">
    <div class="bg-gray-100 dark:bg-gray-900 min-h-screen font-poppins text-gray-900 dark:text-gray-100 transition-colors duration-300">
      <NuxtPage/>
    </div>
  </div>
</template>

<script setup>
import { useThemeStore } from '~/stores/theme'
const themeStore = useThemeStore()
</script>