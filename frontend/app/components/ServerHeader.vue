<script lang="ts" setup>
const items = ref([
    {
        label: 'Home',
        icon: 'i-lucide-house'
    },
])
</script>

<template>
    <div class="flex p-2">
        <UButton color="neutral" icon="i-lucide-align-justify" size="sm" />
        <USeparator orientation="vertical" class="p-2" />
        <UBreadcrumb :items="items" class="p-2" />
    </div>
</template>