// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate : '2025-08-27',
  modules: ['@nuxtjs/tailwindcss', '@pinia/nuxt', '@nuxtjs/google-fonts'],
  googleFonts: {
    families: {
      Roboto: [100, 300, 400, 500, 700, 900],
      'Open+Sans': [100, 300, 400, 500, 700], // Use '+' for spaces in font names
      Lato: [100, 300, 400, 500, 700, 900],
      Montserrat: [100, 300, 400, 500, 700, 900],
      Poppins: [100, 300, 400, 500, 700, 900],
    },
    // Optional: Additional settings for better control
    download: true, // Download fonts for self-hosting (optional for offline support)
    inject: true, // Automatically inject CSS into your app
  },
  plugins: ['~/plugins/heroicons.ts'],
  devtools: { enabled: true },
});