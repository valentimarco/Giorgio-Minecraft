<script lang="ts" setup>
const items = [
    {
        slot: 'provider',
        title: 'Select Provider',
        icon: 'i-lucide-database-zap'
    }, {
        slot: 'config',
        title: 'Configure your Server',
        icon: 'i-lucide-cog'
    }, {
        slot: 'confirm',
        title: 'Confirm your settings',
        icon: 'i-lucide-view'
    }
]

const stepper = useTemplateRef('stepper')
const selectedProvider = ref<string>("CF")
</script>

<template>
    <UModal title="Create Server" description="Wizard to create dedicated servers" :close="{ color: 'primary', variant: 'outline', class: 'rounded-full' }">
        <template #body>
            <div class="w-full">
                <UStepper ref="stepper" :items="items">
                    <template #provider>
                        <div class="grid grid-cols-[10rem_1fr] border-1 rounded-lg">
                            <div class="flex flex-col p-2 gap-1">
                                <UButton label="CurseForge" icon="i-lucide-anvil" @click="(e) => {selectedProvider = 'CF'}" :ui="{ label: 'text-center w-full' }"/>
                                <UButton label="Modrinth" icon="i-lucide-anvil" @click="(e) => {selectedProvider = 'MR'}" :ui="{ label: 'text-center w-full' }"/>
                                <UButton label="Custom" icon="i-lucide-anvil" @click="(e) => {selectedProvider = 'CU'}" :ui="{ label: 'text-center w-full' }"/>                           
                            </div>
                            <div class="flex">
                                <!-- TODO: how to create this list based on the modpack list of each site? -->
                                <div v-show="selectedProvider === 'CF'">
                                    <ul>
                                        <li>
                                            <p>wa</p>
                                        </li>
                                    </ul>
                                </div>
                                <div v-show="selectedProvider === 'MR'">
                                    <ul>
                                        <li>
                                            <p>wu</p>
                                        </li>
                                    </ul>
                                </div>
                                <div v-show="selectedProvider === 'CU'">
                                    <ul>
                                        <li>
                                            <!-- TODO: add drag and drop -->
                                            <p>wu</p>
                                        </li>
                                    </ul>
                                </div>
                            </div>

                        </div>
                    </template>
                </UStepper>

                <div class="flex gap-2 justify-between mt-4">
                    <UButton leading-icon="i-lucide-arrow-left" :disabled="!stepper?.hasPrev" @click="stepper?.prev()">
                        Prev
                    </UButton>

                    <UButton trailing-icon="i-lucide-arrow-right" :disabled="!stepper?.hasNext" @click="stepper?.next()">
                        Next
                    </UButton>
                </div>
            </div>
        </template>
    </UModal>
</template>