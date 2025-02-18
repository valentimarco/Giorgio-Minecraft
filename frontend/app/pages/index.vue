<script lang="ts" setup>
const { backendUrl } = useRuntimeConfig().public
// const { user, fetch } = useUserSession()
const { data: firstTimeUser } = await useLazyFetch(`${backendUrl}/api/v1/first-time`, {
  transform: (d: { firstTime: boolean }) => d.firstTime,
  default: () => false
})


const stepper = ref(false);

const registerUser = async (e: UserRegister) => {
  await $fetch('/api/register', {
    method: 'post',
    body: {
      username: e.username,
      password: e.password
    },
  })
}



</script>

<template>
  <div class="grid grid-rows-2 justify-center p-2 m-2 text-2xl">
    <div class="flex items-center">
      <h1 class="text-center">Welcome to Giorgio Minecraft Administration</h1>
    </div>
    <div v-if="firstTimeUser" class="w-full h-full">
      <AuthForm v-if="!stepper" type="register" class="flex flex-col items-center space-y-4" @submit="registerUser" />
      <ServerWizard v-else />
    </div>
    <div v-else>
      <AuthForm type="login" class="flex flex-col items-center space-y-4" @submit="(e) => console.log(e)" />
    </div>
  </div>
</template>
