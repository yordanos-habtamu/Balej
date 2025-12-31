<template>
  <aside
    :class="[
      'fixed left-0 top-0 h-screen transition-all duration-500 ease-[cubic-bezier(0.2,1,0.3,1)] z-[100]',
      'border-r border-white/10 bg-black/80 backdrop-blur-xl flex flex-col overflow-hidden',
      theme.terminalCollapsed ? 'w-20' : 'w-72'
    ]"
    @mouseenter="expandSidebar"
    @mouseleave="startCollapseDelay"
  >
    <div class="p-6 mb-8 flex items-center gap-4 min-w-[288px]">
      <div 
        class="w-8 h-8 flex-shrink-0 border-2 border-cyan-500 flex items-center justify-center animate-pulse"
        :style="{ borderColor: theme.accentColor, boxShadow: `0 0 15px ${theme.accentColor}` }"
      >
        <span class="text-[10px] font-black italic">NT</span>
      </div>
      <span 
        class="text-xs font-black italic tracking-[0.3em] uppercase transition-all duration-500 whitespace-nowrap"
        :class="theme.terminalCollapsed ? 'opacity-0 -translate-x-4 pointer-events-none' : 'opacity-100 translate-x-0'"
      >
        Neural<span :style="{ color: theme.accentColor }">Control</span>
      </span>
    </div>

    <nav class="flex-1 px-3">
      <ul class="space-y-4">
        <li v-for="(item, index) in navItems" :key="item.label">
          <NuxtLink
            :to="item.link"
            class="group relative flex items-center p-4 transition-all duration-300 overflow-hidden rounded-lg"
            :active-class="'is-active'"
          >
            <div class="absolute inset-0 bg-white/[0.03] translate-x-[-101%] group-hover:translate-x-0 transition-transform duration-300"></div>
            
            <div 
              class="absolute left-0 w-[2px] h-0 group-hover:h-full transition-all duration-300"
              :style="{ backgroundColor: theme.accentColor, boxShadow: `0 0 10px ${theme.accentColor}` }"
            ></div>

            <div class="w-6 h-6 flex items-center justify-center flex-shrink-0 mr-4">
              <component
                :is="item.icon"
                class="w-5 h-5 transition-all duration-300 text-gray-500 group-hover:scale-110"
                :style="[isRouteActive(item.link) ? { color: theme.accentColor } : {}]"
              />
            </div>

            <div 
              class="flex flex-col transition-all duration-500 whitespace-nowrap"
              :class="theme.terminalCollapsed ? 'opacity-0 translate-x-4' : 'opacity-100 translate-x-0'"
            >
              <span class="text-[10px] font-bold uppercase tracking-[0.2em] text-gray-400 group-hover:text-white transition-colors">
                {{ item.label }}
              </span>
              <span class="text-[7px] font-mono text-gray-700 uppercase group-hover:text-cyan-500 transition-colors">
                Access_Link_{{ index }}
              </span>
            </div>
          </NuxtLink>
        </li>
      </ul>
    </nav>

    <div class="p-6 border-t border-white/5 font-mono min-w-[288px]">
      <div 
        class="flex items-center gap-2 mb-2 transition-opacity duration-500"
        :class="theme.terminalCollapsed ? 'opacity-0' : 'opacity-100'"
      >
        <div class="w-1.5 h-1.5 rounded-full bg-green-500 animate-pulse"></div>
        <span class="text-[8px] text-gray-500">SYSTEM_STABLE</span>
      </div>
      <div 
        class="h-1 w-full bg-white/5 rounded-full overflow-hidden transition-all duration-500"
        :class="theme.terminalCollapsed ? 'opacity-0' : 'opacity-100'"
      >
        <div class="h-full bg-cyan-500 w-2/3" :style="{ backgroundColor: theme.accentColor }"></div>
      </div>
    </div>
  </aside>
</template>
 <script setup lang="ts">

import { useThemeStore } from '~/stores/theme'

import { useRoute } from 'vue-router'

import {

HomeIcon,

ChatBubbleLeftRightIcon,

DocumentChartBarIcon,

UsersIcon,

PlusIcon,

BriefcaseIcon

} from '@heroicons/vue/24/outline'


const theme = useThemeStore()

const route = useRoute()


let collapseTimeout: ReturnType<typeof setTimeout> | null = null


const expandSidebar = () => {

if (collapseTimeout) clearTimeout(collapseTimeout)

theme.terminalCollapsed = false

}


const startCollapseDelay = () => {

collapseTimeout = setTimeout(() => {

theme.terminalCollapsed = true

}, 400)

}


const isRouteActive = (link: string) => route.path === link


const navItems = [

{ label: 'Dashboard', icon: HomeIcon, link: '/admin/dashboard' },

{ label: 'Neural_Chat', icon: ChatBubbleLeftRightIcon, link: '/communication/chat' },

{ label: 'Analytics', icon: DocumentChartBarIcon, link: '/admin/reports' },

{ label: 'Talent_Nodes', icon: UsersIcon, link: '/talents' },

{ label: 'Add_Jobs', icon: PlusIcon, link: '/admin/postJob' },

{ label: 'Entities', icon: BriefcaseIcon, link: '/admin/clients' }

]

</script>


<style scoped>

/* Active State Styles */

.is-active {

background: rgba(255, 255, 255, 0.05);

}

.is-active div:first-child {

height: 100% !important;

}

.is-active span {

color: white !important;

}

</style> 