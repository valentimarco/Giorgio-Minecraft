<script lang="ts" setup>
const { servers, curServer } = storeToRefs(useDataStore())

const itemsSelectMenu = computed(() => [
    servers.value.map(
        (s) => ({ label: s.name, value: s.id })
    )
])
</script>

<template>
    <div class="grid grid-cols-[16rem_1fr] grow">
        <div class="flex flex-col bg-base-100 rounded-lg m-2 p-2 gap-2 border-2">

            <USelectMenu v-model="curServer" value-key="value" :items="itemsSelectMenu" class="w-full" />
            <USeparator />

            <UNavigationMenu orientation="vertical" :items="itemsSelectMenu" class="data-[orientation=vertical]:w-full grow" />

            <USeparator />
            <UButton color="warning" label="User 1" icon="i-lucide-user" :ui="{label: 'text-center w-full'}"/>

        </div>
    </div>
</template>