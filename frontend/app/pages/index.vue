<script lang="ts" setup>
const { backendUrl } = useRuntimeConfig().public
const { data: firstTimeUser } = await useFetch(`${backendUrl}/api/v1/first-time`, {
  transform: (d: { firstTime: boolean }) => d.firstTime,
  default: () => false
})
const { fetch } = useUserSession()


const stepper = ref(false);

const registerUser = async (e: UserRegister) => {
  await $fetch('/api/register', {
    method: 'post',
    body: {
      username: e.username,
      password: e.password,
      confirmPassword: e.confirmPassword
    },
  }).then(() => fetch())

  stepper.value = !stepper.value
}

</script>

<template>
  <div class="grid grid-rows-[4rem_1fr] text-2xl h-dvh overflow-y-hidden">
    <div class="flex justify-center">
      <h1 class="text-center p-2 m-2">Welcome to Giorgio Minecraft Administration</h1>
    </div>
    <!-- <div v-if="firstTimeUser" class="w-full h-full">
      <AuthForm v-if="!stepper" type="register" class="flex flex-col items-center space-y-4" @submit="registerUser" />
      <ServerWizard v-else />
    </div>
    <div v-else>
      <AuthForm type="login" class="flex flex-col items-center space-y-4" @submit="(e) => console.log(e)" />
    </div> -->
    <ServerWizard />
  </div>
</template>
