<template>
  <aside
    :class="[
      'bg-white shadow-sm shadow-black/20 rounded-lg p-4 transition-all duration-300 flex flex-col',
      isCollapsed ? 'w-20' : 'w-64'
    ]"
    @mouseenter="isCollapsed = false"
    @mouseleave="isCollapsed = true"
  >
   
    <!-- Nav -->
    <nav class="flex-1">
      <ul class="flex flex-col gap-2">
        <li v-for="item in navItems" :key="item.label">
          <NuxtLink
            :to="item.link"
            class="group flex items-center gap-3 p-3 rounded-lg transition duration-200 ease-in-out hover:bg-gray-100 hover:text-blue-600"
          >
            <!-- Icon -->
            <component
              :is="item.icon"
              class="w-5 h-5 text-gray-400 group-hover:text-blue-600 transition duration-200 flex-shrink-0"
            />

            <!-- Label (hidden when collapsed) -->
            <span
              v-if="!isCollapsed"
              class="text-sm font-medium whitespace-nowrap"
            >
              {{ item.label }}
            </span>
          </NuxtLink>
        </li>
      </ul>
    </nav>
  </aside>
</template>

<script setup>
import { ref } from 'vue'
import {
  HomeIcon,
  ChatBubbleLeftRightIcon,
  DocumentChartBarIcon,
  UsersIcon,
  PlusIcon,
  BriefcaseIcon
} from '@heroicons/vue/24/outline'

const isCollapsed = ref(false)

const navItems = [
  { label: 'Overview', icon: HomeIcon, link: '/admin' },
  { label: 'Chats', icon: ChatBubbleLeftRightIcon, link: '/admin/chats' },
  { label: 'Reports', icon: DocumentChartBarIcon, link: '/admin/reports' },
  { label: 'Manage Talents', icon: UsersIcon, link: '/admin/talents'},
  { label: 'Add Jobs', icon: PlusIcon, link: '/admin/postJob' },
  { label: 'Clients', icon: BriefcaseIcon, link: '/admin/clients' }]
</script>
