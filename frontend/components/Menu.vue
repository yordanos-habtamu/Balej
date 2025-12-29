<template>
     <div class="ml-3">
      <nav class="top-0 w-[600px] "  >
        <ul class="flex justify-evenly p-3 bg-black text-white rounded-b-2xl ">
          <nuxt-link to="/"><li class=" hover:bg-white hover:px-6 hover:py-2 hover:rounded-full hover:text-black p-2 transition duration-300 rounded-full">Home</li></nuxt-link>
          <nuxt-link to="#about"><li class=" hover:bg-white hover:px-6 hover:py-2 hover:rounded-full hover:text-black p-2 transition duration-300 rounded-full">About</li></nuxt-link>
         <nuxt-link to="#contact"><li class=" hover:bg-white hover:px-6 hover:py-2 hover:rounded-full hover:text-black p-2 transition duration-300 rounded-full">Contact</li></nuxt-link>
          <nuxt-link to="/blogs"><li class=" hover:bg-white hover:px-6 hover:py-2 hover:rounded-full hover:text-black p-2 transition duration-300 rounded-full">Blogs</li></nuxt-link>
          <nuxt-link to="/">  <li class=" hover:bg-white hover:px-6 hover:py-2 hover:rounded-full hover:text-black p-2 transition duration-300 rounded-full">Services</li></nuxt-link>
        </ul>
      </nav>
    </div>
</template>