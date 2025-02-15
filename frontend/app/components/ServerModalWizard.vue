<script lang="ts" setup>




</script>

<template>
    <UModal title="Modal with close button" :close="{ color: 'primary', variant: 'outline', class: 'rounded-full' }">
        <UButton color="primary" label="Create Server" icon="i-lucide-plus" :ui="{ label: 'text-center w-full' }" />
        <template #body>
            <ServerWizard />
        </template>
    </UModal>
</template>