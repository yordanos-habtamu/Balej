import * as HeroIcons from '@heroicons/vue/24/outline'; // Or /solid, depending on your needs
import { defineNuxtPlugin } from '#app';

export default defineNuxtPlugin((nuxtApp) => {
  Object.entries(HeroIcons).forEach(([name, component]) => {
    nuxtApp.vueApp.component(name, component);
  });
});