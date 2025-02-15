<script lang="ts" setup>
import type { NavigationMenuItem, SelectMenuItem } from '@nuxt/ui'

const { servers, curServer } = storeToRefs(useDataStore())

const itemsSelectMenu = computed<SelectMenuItem[][]>(() => [
    servers.value.map(
        (s) => ({
            label: s.name,
            value: s.id,
            chip: { color: s.online ? 'success' : 'error' },
            onSelect: () => navigateTo({ params: { id: s.id } })
        })
    )
])

const itemsNavMenu = computed<NavigationMenuItem[]>(() => [
    {
        label: 'Dashboard',
        icon: 'i-lucide-layout-dashboard',
        to: '/:id()/dashboard'
    },
    {
        label: 'Config',
        icon: 'i-lucide-book-open',
        to: '/:id()/config'
    },
    {
        label: 'Terminal',
        icon: 'i-lucide-terminal',
        to: '/:id()/terminal'
    },
])
</script>

<template>
    <div class="flex flex-col bg-base-100 rounded-lg m-2 p-2 gap-2 border-2 text-2xl">
        <ServerWizard />
        <USeparator size="xl" />
        <USelectMenu v-model="curServer" value-key="value" :items="itemsSelectMenu" class="w-full" size="xl" />
        <USeparator />
        <UNavigationMenu orientation="vertical" :items="itemsNavMenu" class="data-[orientation=vertical]:w-full grow"
            :ui="{ label: 'w-full flex items-center' }" />
        <USeparator />
        <UButton color="primary" label="User 1" icon="i-lucide-user" :ui="{ label: 'text-center w-full' }" />
    </div>
</template>