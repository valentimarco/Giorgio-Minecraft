<script lang="ts" setup>
const { servers, curServer } = storeToRefs(useDataStore())
const route = useRoute();
const pathId = `/${route.params.id}/`

const itemsSelectMenu = computed(() => [
    servers.value.map(
        (s) => ({
             label: s.name, 
             value: s.id, 
             chip: { color: s.online ? 'success' : 'error' }
        })
    )
])

// TODO: Make so that only change the id part!
const onSelectValue = async (id: string) =>{
    await navigateTo({ params: { id: id } })
}

const itemsNavMenu = ref([
    {
        label: 'Dashboard',
        icon: 'i-lucide-layout-dashboard',
        to: pathId + 'dashboard'
    },
    {
        label: 'Config',
        icon: 'i-lucide-book-open',
        to: pathId + 'config'
    },
    {
        label: 'Terminal',
        icon: 'i-lucide-terminal',
        to: pathId + 'terminal'
    },
])
</script>

<template>
    <div class="flex flex-col bg-base-100 rounded-lg m-2 p-2 gap-2 border-2 text-2xl">

        <UButton color="primary" label="Create Server" icon="i-lucide-plus" :ui="{ label: 'text-center w-full' }" />
        <USeparator size="xl" />
        <USelectMenu v-model="curServer" value-key="value" :items="itemsSelectMenu" @update:model-value="onSelectValue" class="w-full" size="xl" />
        <USeparator />
        <UNavigationMenu 
            orientation="vertical" :items="itemsNavMenu" class="data-[orientation=vertical]:w-full grow"
            :ui="{ label: 'w-full flex items-center' }" 
        />
        <USeparator />
        <UButton color="primary" label="User 1" icon="i-lucide-user" :ui="{ label: 'text-center w-full' }" />

    </div>
</template>