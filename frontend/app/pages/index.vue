<script lang="ts" setup>

//const { loggedIn, session, user, clear, fetch } = useUserSession()
const firstTimeUser = true //TODO: get this from backend


const stepper = ref(false);



</script>

<template>
  <div class="grid grid-rows-2 justify-center p-2 m-2 text-2xl">
    <div class="flex items-center">
      <h1 class="text-center">Welcome to Giorgio Minecraft Administration</h1>
    </div>
    <div v-if="firstTimeUser" class="w-full h-full">
          <AuthForm v-if="!stepper" type="register" class="flex flex-col items-center space-y-4" @submit="(e) => {console.log(e);  stepper = !stepper}" />
          <ServerWizard v-else />
    </div>
    <div v-else>
      <AuthForm type="login" class="flex flex-col items-center space-y-4" @submit="(e) => console.log(e)" />
    </div>
  </div>
</template>
