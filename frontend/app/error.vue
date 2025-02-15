<script setup lang="ts">
import type { NuxtError } from '#app'

defineProps<{
  error: NuxtError
}>()
</script>

<template>
  <div class="flex flex-col items-center p-20 m-20 gap-2">
    <div class="flex-row items-center">
      <h1 class="text-xl text-maroon-50">{{ error.statusCode }} {{ error.statusMessage }}</h1> 
    </div>
    
    <NuxtLink to="/" class="text-xl text-primary-800">Go back home</NuxtLink>
  </div>
</template>
