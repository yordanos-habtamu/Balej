<template>
 <!-- <div class="grid grid-cols-1 md:grid-cols-2  gap-4 w-full">
    <div class="w-200 h-200 hover:shadow-2xl hover:shadow-dark rounded-md">
       <p class="text-center">I am a Worker Looking for opportunities</p>
    </div>
    <div class="w-200 h-200 hover:shadow-2xl hover:shadow-dark rounded-md">
    <p class="text-center">I am client. I want workers with affordable price</p>
    </div>
 </div> -->

<div class="min-h-screen flex items-center justify-center bg-gray-100">
  <div class="flex flex-wrap justify-center items-center gap-6 p-6">
    <!-- Client Card -->
   <nuxt-link to="/registerClient">
    <div class="w-80 h-80 bg-white border-2 border-black rounded-2xl shadow-md hover:shadow-2xl hover:scale-105 transition-all duration-300 p-6 flex flex-col justify-center">
      <p class="text-center text-lg font-semibold text-gray-800">
        I am a <span class="text-blue-600 font-bold">Client</span>, I want professionals with affordable price
      </p>
    </div>
   </nuxt-link> 

    <!-- Professional Card -->
   <nuxt-link to="/registerTalent">
    <div class="w-80 h-80 bg-white border-2 border-black rounded-2xl shadow-md hover:shadow-2xl hover:scale-105 transition-all duration-300 p-6 flex flex-col justify-center">
      <p class="text-center text-lg font-semibold text-gray-800">
        I am a <span class="text-green-600 font-bold">Professional</span>, I want Clients to work with
      </p>
    </div>
   </nuxt-link>
    
  </div>
</div>

</template>
<script setup lang="ts">
</script>