<script lang="ts" setup generic="T extends 'login' | 'register', S = T extends 'login' ? UserLogin : UserRegister">
const { type } = defineProps<{
  type: T
}>()

defineEmits<{
  submit: [data: S]
}>()

const credentials = reactive<Partial<UserLogin | UserRegister>>({})

const schema = computed(() => type === 'login' ? loginSchema : registerSchema)

const submitButtonText = computed(() => type === 'login' ? 'Log In' : 'Register')
</script>

<template>
  <UForm :schema="schema" :state="credentials" @submit="$emit('submit', $event.data)">
    <UFormField label="Username" name="username" size="xl">
      <UInput v-model="credentials.username" />
    </UFormField>

    <UFormField label="Password" name="password" size="xl">
      <UInput v-model="credentials.password" type="password" />
    </UFormField>

    <UFormField v-if="type === 'register'" label="Confirm Password" name="confirmPassword" size="xl">
      <UInput v-model="(credentials as UserRegister).confirmPassword" type="password" />
    </UFormField>

    <UButton type="submit" size="xl">
      {{ submitButtonText }}
    </UButton>
  </UForm>
</template>
