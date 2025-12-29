<template>
  <div>
    <label :for="keyName" class="block text-sm font-medium text-gray-700">
      {{ label }}
    </label>
    <input
      :id="keyName"
      :type="keyName === 'password' ? 'password' : 'text'"
      class="mt-1 p-3 w-full border rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
      :class="{ 'border-red-500': error }"
      :value="modelValue"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      :placeholder="`Enter ${label}`"
    />
    <p v-if="error" class="text-red-500 text-sm mt-1">{{ error }}</p>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
  modelValue: string
  keyName: string
  error?: string
}>()

const label = computed(() =>
  props.keyName.replace(/_/g, ' ').replace(/\b\w/g, (l) => l.toUpperCase())
)
</script>
