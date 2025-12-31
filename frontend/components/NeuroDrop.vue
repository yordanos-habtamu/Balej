<script setup lang="ts">
import { ref } from 'vue'
import { useThemeStore } from '@/stores/theme'

const themeStore = useThemeStore()

const props = defineProps<{
  modelValue: string | number
  options: readonly string[]
  placeholder: string
  label: string
}>()

const emit = defineEmits(['update:modelValue'])
const isOpen = ref(false)

function selectOption(option: string) {
  emit('update:modelValue', option)
  isOpen.value = false
}
</script>

<template>
  <div class="space-y-3">
    <label class="neural-label">>> {{ label }}</label>
    <div class="relative">
      <div
        @click="isOpen = !isOpen"
        class="neuro-drop-trigger cursor-pointer flex justify-between items-center pr-10"
        :class="{ 'open': isOpen }"
      >
        <span :class="modelValue ? 'text-white' : 'text-gray-600'">
          {{ modelValue || placeholder }}
        </span>
        <div class="absolute right-4 transition-transform" :class="{ 'rotate-180': isOpen }">▼</div>
      </div>

      <div v-if="isOpen" class="absolute z-[100] mt-1 w-full bg-[#0a0a0c] border border-white/20">
        <div
          v-for="opt in options"
          :key="opt"
          @click="selectOption(opt)"
          class="px-4 py-3 text-xs font-mono uppercase cursor-pointer hover:bg-white/5"
          :style="modelValue === opt ? { color: themeStore.accentColor } : {}"
        >
          {{ opt }}
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
@reference 'tailwindcss';
.neuro-drop-trigger { @apply w-full bg-[rgba(255,255,255,0.03)] border border-white/10 p-4 text-xs font-mono text-white transition-all uppercase outline-none; }
.open { border-color: v-bind('themeStore.accentColor'); }
</style>