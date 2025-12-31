<template>
  <aside
    :class="[
      'rounded-lg p-4 transition-all duration-300 flex flex-col',
      theme.sidebarCollapsed ? 'w-20' : 'w-64',
      'bg-gray-800'  // Solid dark gray as base (works beautifully in both modes)
    ]"
    @mouseenter="expandSidebar"
    @mouseleave="startCollapseDelay"
  >
    <!-- Nav -->
    <nav class="flex-1">
      <ul class="flex flex-col gap-2">
        <li v-for="item in navItems" :key="item.label">
          <NuxtLink
            :to="item.link"
            class="group flex items-center gap-3 p-3 rounded-lg transition duration-200 ease-in-out"
            :class="[
              'hover:bg-[color:var(--primary-color)/0.2]',  // Light tint of accent
              'active:bg-[color:var(--primary-color)/0.3]',
              'hover:text-[color:var(--primary-color)]'
            ]"
          >
            <!-- Icon -->
            <component
              :is="item.icon"
              class="w-5 h-5 flex-shrink-0 transition duration-200 
                     text-gray-400 group-hover:text-[color:var(--primary-color)]"
            />
            <!-- Label -->
            <span
              v-if="!theme.sidebarCollapsed"
              class="text-sm font-medium whitespace-nowrap text-gray-300 
                     group-hover:text-[color:var(--primary-color)]"
            >
              {{ item.label }}
            </span>
          </NuxtLink>
        </li>
      </ul>
    </nav>
  </aside>
</template>


<script setup lang="ts">
import { useThemeStore } from '~/stores/theme'
import {
  HomeIcon,
  ChatBubbleLeftRightIcon,
  DocumentChartBarIcon,
  UsersIcon,
  PlusIcon,
  BriefcaseIcon
} from '@heroicons/vue/24/outline'

const theme = useThemeStore()

let collapseTimeout: ReturnType<typeof setTimeout> | null = null

const expandSidebar = () => {
  if (collapseTimeout) {
    clearTimeout(collapseTimeout)
    collapseTimeout = null
  }
  theme.sidebarCollapsed = false
}

const startCollapseDelay = () => {
  collapseTimeout = setTimeout(() => {
    theme.sidebarCollapsed = true
  }, 400) // Adjust delay as needed (300–500ms feels natural)
}

const navItems = [
  { label: 'Overview', icon: HomeIcon, link: '/admin' },
  { label: 'Chats', icon: ChatBubbleLeftRightIcon, link: '/admin/chats' },
  { label: 'Reports', icon: DocumentChartBarIcon, link: '/admin/reports' },
  { label: 'Manage Talents', icon: UsersIcon, link: '/admin/talents' },
  { label: 'Add Jobs', icon: PlusIcon, link: '/admin/postJob' },
  { label: 'Clients', icon: BriefcaseIcon, link: '/admin/clients' }
]

// Primary color CSS variable (same as before)
onMounted(() => {
  const updatePrimaryColor = () => {
    const colorMap: Record<string, string> = {
      blue: '#2563eb',    // blue-600
      indigo: '#4f46e5',
      purple: '#9333ea',
      pink: '#ec4899',
      red: '#ef4444',
      green: '#10b981',
      // add more as needed
    }
    const color = colorMap[theme.primaryColor] || theme.primaryColor
    document.documentElement.style.setProperty('--primary-color', color)
  }

  updatePrimaryColor()
  watch(() => theme.primaryColor, updatePrimaryColor)
})
</script>