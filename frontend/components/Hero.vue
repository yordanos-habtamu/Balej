<template>
     <div class="relative flex flex-col justify-center items-center h-[650px] z-10 pt-20">
      <!-- Background image -->
      <div
        class="absolute inset-0 w-full flex flex-col items-center py-12 bg-cover bg-center z-[-1]"
        :style="{ backgroundImage: `url(${painterImage})` }"
      ></div>

      <h1 class="text-4xl font-bold text-center text-white mt-5 p-1.5 mb-2">
        Connect With a Top Talent or Find your Job today!
      </h1>
      <!-- Search Input -->
      <div class="w-[800px] max-w-full px-4">
        <input
          type="text"
          placeholder="Search..."
          class="p-4 w-full rounded-lg bg-white border border-gray-300 focus:outline-none focus:ring-2 focus:ring-black"
        />
      </div>

      <!-- Buttons -->
      <div class="flex justify-between mt-7 gap-4">
        <button
          class="px-5 py-2 bg-black text-white rounded-full hover:bg-white hover:text-black hover:border-2 hover:border-black transition"
        >
          Hire a Handy Man
        </button>
        <button
          class="px-5 py-2 hover:border-2 text-white hover:rounded-full hover:bg-black transition"
        >
          Find a Job
        </button>
      </div>
    </div>
    
  </template>

  <script setup>
import painterImage from '@/assets/painter.png'
</script>
