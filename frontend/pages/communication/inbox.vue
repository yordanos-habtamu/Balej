<template>
  <div class="min-h-screen flex p-4 sm:p-6 md:p-8 gap-4">
    <!-- Sidebar (desktop only) -->
    <aside class="hidden md:block m-2 flex-shrink-0">
      <admin-sidebar />
    </aside>

    <!-- Main Chat Area -->
    <div class="flex-1 flex flex-col md:flex-row gap-4 max-w-7xl mx-auto w-full">
      <!-- Chat List Panel -->
      <div
        :class="[
          'bg-gray-800 rounded-xl overflow-hidden transition-all duration-300 relative',
          selectedChat ? 'w-full md:w-96 md:flex-shrink-0' : 'w-full flex-grow'
        ]"
      >
        <div class="h-full flex flex-col">
          <!-- Header -->
          <div class="p-4 md:p-6 border-b border-gray-700 flex items-center justify-between">
            <h2 class="text-xl font-semibold text-gray-100">Chats</h2>
          </div>

          <!-- Active Chats -->
          <div v-if="activeChats.length" class="flex-1 overflow-y-auto p-4 space-y-3">
            <div
              v-for="chat in activeChats"
              :key="chat.id"
              @click="selectChat(chat)"
              class="group flex items-center gap-4 p-4 h-[70px] bg-gray-700/40 rounded-xl cursor-pointer 
                     transition-all hover:bg-gray-700/70"
              :class="{ 'ring-2 ring-[color:var(--primary-color)] bg-gray-700/70': selectedChat?.id === chat.id }"
            >
              <img
                :src="chat.profile || '/images/default-avatar.png'"
                alt="profile"
                class="w-12 h-12 rounded-full object-cover ring-2 ring-gray-600 group-hover:ring-[color:var(--primary-color)] transition"
              />
              <div class="flex-1 min-w-0">
                <div class="flex items-baseline gap-2">
                  <h3 class="font-medium text-gray-100 truncate">{{ chat.first_name }}</h3>
                  <h3 class="text-gray-400 truncate">{{ chat.last_name }}</h3>
                </div>
                <p class="text-sm text-gray-400 truncate mt-1">
                  {{ chat.last_message || 'No messages yet' }}
                </p>
              </div>
            </div>
          </div>

          <!-- Empty State -->
          <div v-else class="flex-1 flex flex-col items-center justify-center text-gray-500 px-8">
            <svg class="w-28 h-28 mb-8 text-gray-700" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"
                d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z" />
            </svg>
            <p class="text-xl font-medium mb-2">No chats yet</p>
            <p class="text-sm text-center">Tap the + button below to start a new conversation</p>
          </div>
        </div>

        <!-- Floating + Button -->
        <button
          @click="showAddModal = true"
          class="absolute bottom-6 right-6 w-14 h-14 bg-[color:var(--primary-color)] text-white rounded-full shadow-2xl 
                 flex items-center justify-center hover:scale-110 active:scale-95 transition-all duration-200 z-10"
        >
          <svg class="w-8 h-8" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M12 4v16m8-8H4" />
          </svg>
        </button>
      </div>

      <!-- Messenger Panel -->
      <div
        v-if="selectedChat"
        class="flex-1 bg-gray-800 rounded-xl overflow-hidden flex flex-col"
        :class="{ 'fixed inset-0 z-50 md:relative md:inset-auto': true }"
      >
        <!-- Header with Three Dots Menu -->
        <div class="relative flex items-center justify-between p-4 md:p-6 border-b border-gray-700 bg-gray-800/80">
          <!-- Mobile Back Button + User Info -->
          <div class="flex items-center gap-4 flex-1">
            <button @click="selectedChat = null" class="md:hidden p-2 rounded-lg hover:bg-gray-700">
              <svg class="w-6 h-6 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7" />
              </svg>
            </button>
            <img
              :src="selectedChat.profile || '/images/default-avatar.png'"
              alt="profile"
              class="w-10 h-10 md:w-12 md:h-12 rounded-full ring-2 ring-[color:var(--primary-color)/0.6]"
            />
            <div>
              <h2 class="font-semibold text-gray-100 text-base md:text-xl">
                {{ selectedChat.first_name }} {{ selectedChat.last_name }}
              </h2>
            </div>
          </div>

          <!-- Three Dots Button -->
          <button
            ref="menuButton"
            @click.stop="toggleMenu"
            class="p-2 rounded-lg hover:bg-gray-700 transition"
          >
            <svg class="w-6 h-6 text-gray-400" fill="currentColor" viewBox="0 0 20 20">
              <path d="M10 6a2 2 0 110-4 2 2 0 010 4zM10 12a2 2 0 110-4 2 2 0 010 4zM10 18a2 2 0 110-4 2 2 0 010 4z" />
            </svg>
          </button>

          <!-- Dropdown Menu -->
          <div
            v-if="menuOpen"
            class="absolute right-4 top-full mt-2 bg-gray-700 rounded-lg shadow-2xl py-2 w-48 z-50"
          >
            <button @click="archiveChat" class="w-full px-4 py-2 text-left text-gray-200 hover:bg-gray-600 transition">
              Archive
            </button>
            <button @click="prepareDelete" class="w-full px-4 py-2 text-left text-red-400 hover:bg-gray-600 transition">
              Delete Chat
            </button>
          </div>
        </div>

        <!-- Messages -->
        <div
          ref="messagesContainer"
          class="flex-1 overflow-y-auto p-4 md:p-6 space-y-4 bg-gray-900"
          style="max-height: calc(100vh - 280px); word-break: break-word;"
        >
          <div v-if="messages.length === 0" class="text-center text-gray-500 mt-20">
            <p class="text-lg">Start the conversation!</p>
          </div>
          <div v-for="(msg, i) in messages" :key="i" class="flex justify-end">
            <div class="bg-[color:var(--primary-color)] text-white p-4 rounded-2xl rounded-br-none max-w-[70%] shadow-lg">
              {{ msg }}
            </div>
          </div>
        </div>

        <!-- Input -->
        <div class="p-4 md:p-6 border-t border-gray-700 bg-gray-800">
          <div class="relative">
            <input
              v-model="newMessage"
              @keyup.enter="sendMessage"
              placeholder="Type a message..."
              class="w-full px-5 py-3 pr-14 rounded-xl bg-gray-700 text-white placeholder-gray-400
                     focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)] transition"
            />
            <PaperAirplaneIcon
              @click="sendMessage"
              class="absolute right-3 top-1/2 -translate-y-1/2 w-8 h-8 p-2 bg-[color:var(--primary-color)] text-white rounded-full cursor-pointer hover:scale-105 transition"
            />
          </div>
        </div>
      </div>
    </div>

    <!-- Add New Chat Modal -->
    <div v-if="showAddModal" class="fixed inset-0 bg-black/70 flex items-center justify-center z-50 p-4" @click.self="showAddModal = false">
      <div class="bg-gray-800 rounded-2xl p-8 max-w-md w-full shadow-2xl">
        <h3 class="text-2xl font-bold text-gray-100 mb-6">Start New Chat</h3>
        <form @submit.prevent="addNewChat">
          <div class="space-y-4">
            <input v-model="newPerson.first_name" type="text" placeholder="First Name" required class="w-full px-4 py-3 rounded-lg bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]" />
            <input v-model="newPerson.last_name" type="text" placeholder="Last Name" required class="w-full px-4 py-3 rounded-lg bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]" />
            <input v-model="newPerson.profile" type="url" placeholder="Profile Image URL (optional)" class="w-full px-4 py-3 rounded-lg bg-gray-700 text-white placeholder-gray-400 focus:outline-none focus:ring-2 focus:ring-[color:var(--primary-color)]" />
          </div>
          <div class="flex gap-4 mt-8">
            <button type="button" @click="showAddModal = false" class="flex-1 py-3 rounded-lg bg-gray-700 text-gray-300 hover:bg-gray-600 transition">Cancel</button>
            <button type="submit" class="flex-1 py-3 rounded-lg bg-[color:var(--primary-color)] text-white hover:opacity-90 transition">Start Chat</button>
          </div>
        </form>
      </div>
    </div>

    <!-- Delete Confirmation Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black/70 flex items-center justify-center z-50 p-4" @click.self="showDeleteModal = false">
      <div class="bg-gray-800 rounded-2xl p-8 max-w-sm w-full shadow-2xl text-center">
        <h3 class="text-xl font-bold text-gray-100 mb-4">Delete chat?</h3>
        <p class="text-gray-400 mb-8">This will permanently remove the conversation with {{ selectedChat?.first_name }} {{ selectedChat?.last_name }}.</p>
        <div class="flex gap-4">
          <button @click="showDeleteModal = false" class="flex-1 py-3 rounded-lg bg-gray-700 text-gray-300 hover:bg-gray-600 transition">Cancel</button>
          <button @click="confirmDelete" class="flex-1 py-3 rounded-lg bg-red-600 text-white hover:bg-red-700 transition">Delete</button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, nextTick, onMounted, watch, computed, onUnmounted } from 'vue'
import { useThemeStore } from '~/stores/theme'
import { PaperAirplaneIcon } from '@heroicons/vue/24/outline'

const theme = useThemeStore()

// Data
const allChats = ref<Array<any>>([])
const selectedChat = ref<any>(null)
const messages = ref<string[]>([])
const newMessage = ref('')

// Modals
const showAddModal = ref(false)
const showDeleteModal = ref(false)

// Menu
const menuOpen = ref(false)
const menuButton = ref<HTMLElement>()

const newPerson = ref({ first_name: '', last_name: '', profile: '' })

let nextId = 1

const activeChats = computed(() => allChats.value.filter(c => !c.archived))

// Actions
const addNewChat = () => {
  if (!newPerson.value.first_name.trim() || !newPerson.value.last_name.trim()) return

  const newChat = {
    id: nextId++,
    first_name: newPerson.value.first_name.trim(),
    last_name: newPerson.value.last_name.trim(),
    profile: newPerson.value.profile || null,
    last_message: null,
    archived: false
  }

  allChats.value.unshift(newChat)
  showAddModal.value = false
  newPerson.value = { first_name: '', last_name: '', profile: '' }
}

const selectChat = (chat: any) => {
  selectedChat.value = chat
  messages.value = []
  menuOpen.value = false
}

const toggleMenu = () => {
  menuOpen.value = !menuOpen.value
}

// Close menu when clicking outside
const handleClickOutside = (event: MouseEvent) => {
  if (menuButton.value && !menuButton.value.contains(event.target as Node)) {
    menuOpen.value = false
  }
}

onMounted(() => {
  document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
  document.removeEventListener('click', handleClickOutside)
})

const archiveChat = () => {
  if (!selectedChat.value) return
  selectedChat.value.archived = true
  selectedChat.value = null
  messages.value = []
  menuOpen.value = false
}

const prepareDelete = () => {
  showDeleteModal.value = true
  menuOpen.value = false
}

const confirmDelete = () => {
  if (!selectedChat.value) return

  allChats.value = allChats.value.filter(c => c.id !== selectedChat.value.id)
  selectedChat.value = null
  messages.value = []
  showDeleteModal.value = false
}

const sendMessage = () => {
  if (!newMessage.value.trim() || !selectedChat.value) return

  messages.value.push(newMessage.value.trim())
  selectedChat.value.last_message = newMessage.value.trim()
  newMessage.value = ''

  nextTick(() => {
    const container = messagesContainer.value
    if (container) container.scrollTop = container.scrollHeight
  })
}

const messagesContainer = ref<HTMLElement>()

// Theme primary color
onMounted(() => {
  const updateColor = () => {
    const map: Record<string, string> = {
      blue: '#3b82f6',
      indigo: '#6366f1',
      purple: '#a78bfa',
      pink: '#ec4899',
      emerald: '#10b981',
      red: '#ef4444'
    }
    document.documentElement.style.setProperty('--primary-color', map[theme.primaryColor] || '#3b82f6')
  }
  updateColor()
  watch(() => theme.primaryColor, updateColor)
})
</script>