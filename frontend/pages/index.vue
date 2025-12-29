<template>
  <div class="h-[100vh] relative font-poppins">
    
    <!-- ✅ Fixed Top Navbar -->
    <div class="flex justify-evenly fixed top-0 z-50 w-full bg-white shadow ">
      <Logo />
      <Menu />
      <Action />
    </div>
    <Hero/>
  <div id="about" class="mt-12 h-full bg-gray-50 py-8">
    <h1 class="text-3xl font-bold text-center text-gray-800 mb-8">
      Lorem ipsum dolor sit amet consectetur adipisicing elit. Recusandae nihil inventore corrupti qui doloremque, neque iusto enim voluptatem culpa sunt tempore fugit quam aliquam dolor id, omnis, vel quas! Blanditiis!
    </h1>
    <div class="flex flex-wrap justify-center gap-6 px-4">
      <div class="w-full sm:w-[400px] bg-white border border-gray-200 rounded-xl shadow-lg shadow-black p-4 m-2 hover:shadow-xl hover:-translate-y-2 hover:scale-105 transition-all duration-300 ease-in-out">
        <div class="h-[300px] w-full max-w-[300px] mx-auto rounded-lg overflow-hidden">
          <img 
            src="https://images.pexels.com/photos/3186949/pexels-photo-3186949.jpeg" 
            alt="Service Image" 
            class="w-full h-full object-cover"
          >
        </div>
        <p class="mt-4 text-gray-600 text-center">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Incidunt quaerat deserunt autem quidem aliquam laudantium voluptatem alias quae. Expedita consequuntur voluptates voluptas ut. Ea cumque asperiores esse earum officiis doloremque iste facere illo laboriosam iusto!
        </p>
      </div>
      <div class="w-full sm:w-[400px] bg-white border border-gray-200 rounded-xl shadow-lg shadow-black p-4 m-2 hover:shadow-xl hover:-translate-y-2 hover:scale-105 transition-all duration-300 ease-in-out">
        <div class="h-[300px] w-full max-w-[300px] mx-auto rounded-lg overflow-hidden">
          <img 
            src="https://images.pexels.com/photos/3186949/pexels-photo-3186949.jpeg" 
            alt="Service Image" 
            class="w-full h-full object-cover"
          >
        </div>
        <p class="mt-4 text-gray-600 text-center">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Incidunt quaerat deserunt autem quidem aliquam laudantium voluptatem alias quae. Expedita consequuntur voluptates voluptas ut. Ea cumque asperiores esse earum officiis doloremque iste facere illo laboriosam iusto!
        </p>
      </div>
      <div class="w-full sm:w-[400px] bg-white border border-gray-200 rounded-xl shadow-lg shadow-black p-4 m-2 hover:shadow-2xl hover:-translate-y-2 hover:scale-105 transition-all duration-300 ease-in-out">
        <div class="h-[300px] w-full max-w-[300px] mx-auto rounded-lg overflow-hidden">
          <img 
            src="https://images.pexels.com/photos/3186949/pexels-photo-3186949.jpeg" 
            alt="Service Image" 
            class="w-full h-full object-cover"
          >
        </div>
        <p class="mt-4 text-gray-600 text-center">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Incidunt quaerat deserunt autem quidem aliquam laudantium voluptatem alias quae. Expedita consequuntur voluptates voluptas ut. Ea cumque asperiores esse earum officiis doloremque iste facere illo laboriosam iusto!
        </p>
      </div>
    </div>
  </div>
     <div id="contact" class="flex flex-col lg:flex-row justify-around items-center mt-8 min-h-[650px] bg-gradient-to-b from-gray-50 to-gray-100 py-12 px-4">
    <!-- Contact Form -->
    <div class="w-full lg:w-1/2 max-w-lg p-4">
      <div class="bg-white border border-gray-200 rounded-2xl shadow-lg p-8">
        <h1 class="text-2xl font-bold text-gray-800 mb-6">Contact Us</h1>
        <div class="space-y-4">
          <input 
            type="text" 
            placeholder="Name" 
            class="p-3 border border-gray-300 rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-black transition"
          >
          <input 
            type="email" 
            placeholder="Email" 
            class="p-3 border border-gray-300 rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-black transition"
          >
          <textarea 
            placeholder="Message" 
           
            rows="5" 
            class="p-3 border border-gray-300 rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-black transition resize-none"
          ></textarea>
          <button 
            class="w-full bg-black text-white px-4 py-3 rounded-full hover:bg-white hover:border-2 hover:border-black hover:text-black hover:shadow-md transition duration-300"
          >
            Submit
          </button>
        </div>
      </div>
    </div>
    <!-- Text Content -->
    <div class="w-full lg:w-1/2 max-w-lg p-4">
      <div class="text-center lg:text-left">
        <h1 class="text-3xl font-bold text-gray-800 mb-4">
          Get in Touch
        </h1>
        <p class="text-gray-600 max-w-md mx-auto lg:mx-0">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Recusandae nihil inventore corrupti qui doloremque, neque iusto enim voluptatem culpa sunt tempore fugit quam aliquam dolor id, omnis, vel quas! Blanditiis!
        </p>
        <p class="w-full p-2 bg-white border-black  border-2 rounded-full text-center text-black mt-5 shadow-lg shadow-black">Register for our daily NewsLetter</p>
      </div>
    </div>
  </div>
  <Footer/>
  </div>

</template>
<style scoped>
.font-poppins{
  font-family: 'Poppins', sans-serif;
}
</style>

<script setup></script>
