<script lang="ts" setup>
  definePageMeta({
    middleware: ["check-id"]
  })  
</script>

<template>
    <div class="flex min-h-dvh scroll-smooth antialiased transition-colors">
    <div class="grid grid-cols-[16rem_1fr] grow">
      <NavBar />
      <div class="grid grid-rows-[3rem_1fr] rounded-lg m-2 p-2 gap-2 border">
        <ServerHeader />
        <div class="flex rounded-lg p-2 border">
          <NuxtPage />
        </div>
      </div>
    </div>
    </div>  
  </template>