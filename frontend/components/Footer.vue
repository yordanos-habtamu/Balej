<template>
    <div class="container mx-auto px-4">
      <div class="flex flex-col md:flex-row justify-between items-start md:items-center gap-8">
        <!-- About Section -->
        <div class="w-full md:w-1/3">
          <h2 class="text-2xl font-bold mb-4">Our Company</h2>
          <p class="text-gray-400 max-w-xs">
            Connect,find and grow. We connect skilled professionals with top opportunities.
          </p>
        </div>

        <!-- Links Section -->
        <div class="w-full md:w-1/3">
          <h2 class="text-xl font-semibold mb-4">Quick Links</h2>
          <ul class="space-y-2">
            <li><a href="#" class="text-black hover:text-gray-400 transition duration-300">About Us</a></li>
            <li><a href="#" class="text-black hover:text-gray-400 transition duration-300">Blog</a></li>
            <li><a href="#" class="text-black hover:text-gray-400 transition duration-300">Careers</a></li>
            <li><a href="#" class="text-black hover:text-gray-400 transition duration-300">Contact</a></li>
          </ul>
        </div>

        <!-- Newsletter Section -->
        <div class="w-full md:w-1/3">
          <h2 class="text-xl font-semibold mb-4">Subscribe to Our Newsletter</h2>
          <div class="flex flex-col sm:flex-row gap-2">
            <input 
              type="email" 
              placeholder="Enter your email" 
              class="p-3 bg-white border border-black rounded-lg w-full focus:outline-none focus:ring-2 focus:ring-black text-black transition-all duration-300"
            >
            <button 
              class="px-4 py-3 bg-black text-white rounded-lg hover:bg-white hover:text-black hover:border-2 hover:border-black transition duration-300"
            >
              Subscribe
            </button>
          </div>
          <!-- Social Media Icons -->
          <div class="flex gap-4 mt-4">
            <a href="#" class="text-black hover:text-gray-400 transition duration-300">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 2.04c-5.5 0-9.96 4.46-9.96 9.96 0 4.41 3.58 8.04 8.04 8.88v-6.28h-2.43v-2.6h2.43v-1.98c0-2.4 1.46-3.71 3.59-3.71.72 0 1.44.04 2.16.12v2.5h-1.48c-1.16 0-1.39.55-1.39 1.36v1.71h2.78l-.36 2.6h-2.42v6.28c4.46-.84 8.04-4.47 8.04-8.88 0-5.5-4.46-9.96-9.96-9.96z"/>
              </svg>
            </a>
            <a href="#" class="text-black hover:text-gray-400 transition duration-300">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M22.46 6c-.77.34-1.6.57-2.47.67  .89-.53 1.57-1.37 1.89-2.37-.83.49-1.75.85-2.73 1.04-.78-.83-1.89-1.35-3.12-1.35-2.36 0-4.28 1.92-4.28 4.28 0 .34.04.67.11 1-3.56-.18-6.72-1.89-8.84-4.48-.37.63-.58 1.36-.58 2.14 0 1.48.75 2.78 1.89 3.54-.69-.02-1.34-.21-1.91-.52v.05c0 2.07 1.47 3.8 3.42 4.19-.36.1-.73.15-1.12.15-.27 0-.54-.03-.8-.08.54 1.69 2.11 2.92 3.97 2.96-1.46 1.15-3.29 1.83-5.28 1.83-.34 0-.68-.02-1.02-.06 1.89 1.21 4.13 1.92 6.54 1.92 7.85 0 12.14-6.51 12.14-12.14 0-.18 0-.36-.01-.54.83-.6 1.55-1.35 2.12-2.2z"/>
              </svg>
            </a>
            <a href="#" class="text-black hover:text-gray-400 transition duration-300">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M4.98 3.5c0 1.38-1.12 2.5-2.5 2.5S0 4.88 0 3.5 1.12 1 2.5 1s2.5 1.12 2.5 2.5zm.02 4.5h-5v16h5v-16zm7.98 0h-5v16h5v-16zm7.98 0h-5v16h5v-16zm-7.98-6c0 1.38-1.12 2.5-2.5 2.5s-2.5-1.12-2.5-2.5 1.12-2.5 2.5-2.5 2.5 1.12 2.5 2.5zm.02 4.5h-5v16h5v-16z"/>
              </svg>
            </a>
          </div>
        </div>
      </div>
      <!-- Bottom Bar -->
      <div class="mt-8 pt-8 border-t border-gray-800 text-center text-gray-400 text-sm">
        <p>&copy; 2025 Balej. All Rights Reserved</p>
      </div>
     
    </div>
    
</template>
<script setup></script>