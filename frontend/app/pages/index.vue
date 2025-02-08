<script lang="ts" setup>

</script>


<template>
  <div class="flex min-h-dvh scroll-smooth antialiased transition-colors">
    <NavBar />
    <div class="grid grid-rows-[4rem_1fr] m-2 border">
      <!-- <ServerHeader /> -->
      {children}
    </div>
  </div>
</template>
