<script lang="ts" setup>
import AuthForm from '~/components/AuthForm.vue';

//const { loggedIn, session, user, clear, fetch } = useUserSession()
const firstTimeUser = true //TODO: get this from backend

</script>

<template>
  <div class="grid grid-rows-2 justify-center p-2 m-2 gap-2 text-2xl">
    <div class="flex items-center">
      <h1 class="text-center">Welcome to Giorgio Minecraft Administration</h1>
    </div>
    <div v-if="firstTimeUser" class="flex flex-col items-center space-y-4">
      <h1 class="text-center">
        To start create a user and follow the wizard
      </h1>
      <AuthForm type="register" class="flex flex-col items-center space-y-4" @submit="(e) => console.log(e)" />
    </div>
    <div v-else>
      <AuthForm type="login" class="flex flex-col items-center space-y-4" @submit="(e) => console.log(e)" />
    </div>
  </div>
</template>
