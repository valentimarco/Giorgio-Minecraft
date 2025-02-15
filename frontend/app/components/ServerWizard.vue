<script lang="ts" setup>
const items = [
    {
        slot: 'provider',
        title: 'Select Provider',
        icon: 'i-lucide-database'
    }, {
        slot: 'config',
        title: 'Configure your Server',
        icon: 'i-lucide-settings'
    }, {
        slot: 'confirm',
        title: 'Confirm your settings',
        icon: 'i-lucide-check'
    }
]

const wizardModel = reactive<Partial<ServerWizard>>({
    type: 'VANILLA',
    version: '1.20.1',
    maxPlayers: 20,
    whitelist: undefined,
    memory: 6
})

const selectedProvider = ref<string>("CF")
const stepper = useTemplateRef('stepper')
const dropZone = useTemplateRef('dropZone')

const onDrop = (files: File[] | null) => {
    console.log(files)
}

useDropZone(dropZone, {
    onDrop,
    dataTypes: ['image/png'],
    multiple: true,
    preventDefaultForUnhandled: false,
})


const createServer = () => {
    console.log(wizardModel)
}
</script>

<template>
    <UModal title="Modal with close button" :close="{ color: 'primary', variant: 'outline', class: 'rounded-full' }">
        <UButton color="primary" label="Create Server" icon="i-lucide-plus" :ui="{ label: 'text-center w-full' }" />
        <template #body>
            <div class="flex flex-col h-full w-full ">
                <UStepper ref="stepper" :items="items" class="grow">
                    <template #provider>
                        <div class="grid grid-cols-[10rem_1fr] border-1 rounded-lg">
                            <div class="flex flex-col p-2 gap-1 m-2">
                                <UButton label="CurseForge" icon="i-lucide-anvil" :ui="{ label: 'text-center w-full' }"
                                    @click="selectedProvider = 'CF'" />
                                <UButton label="Modrinth" icon="i-lucide-anvil" :ui="{ label: 'text-center w-full' }"
                                    @click="selectedProvider = 'MR'" />
                                <UButton label="Custom" icon="i-lucide-anvil" :ui="{ label: 'text-center w-full' }"
                                    @click="selectedProvider = 'CU'" />
                            </div>

                            <div class="flex">
                                <!-- TODO: how to create this list based on the modpack list of each site? -->
                                <div v-if="selectedProvider === 'CF'">
                                    <ul>
                                        <li>
                                            <p>wa</p>
                                        </li>
                                    </ul>
                                </div>
                                <div v-else-if="selectedProvider === 'MR'">
                                    <ul>
                                        <li>
                                            <p>wu</p>
                                        </li>
                                    </ul>
                                </div>
                                <div v-else>
                                    <div ref="dropZone" class="">
                                        Drop files here
                                    </div>
                                </div>
                            </div>

                        </div>
                    </template>
                    <template #config>
                        <UForm :schema="serverWizardSchema" :state="wizardModel" class="space-y-4">
                            <div class="grid">
                                <USeparator label="Server Options" />
                                <div class="grid grid-cols-3 justify-center gap-2 m-2">
                                    <UFormField label="Type" name="type">
                                        <USelectMenu v-model="wizardModel.type" :items="[...SERVERTYPES]"
                                            :search-input="false" class="w-full" />
                                    </UFormField>
                                    <UFormField label="Version" name="version">
                                        <UInput v-model="wizardModel.version" placeholder="Version" />
                                    </UFormField>
                                    <UFormField label="Max Players" name="maxPlayers">
                                        <UInput v-model="wizardModel.maxPlayers" type="number" placeholder="20" />
                                    </UFormField>
                                    <UFormField label="Whitelist" class="col-start-2" name="whitelist">
                                        <UTextarea v-model="wizardModel.whitelist"
                                            placeholder="Type username or UUID" />
                                    </UFormField>
                                </div>

                                <USeparator label="Container Options" />
                                <div class="grid grid-cols-2 gap-2 m-2">

                                    <UFormField label="Memory" name="memory">
                                        <UInput v-model="wizardModel.memory" type="number" placeholder="6" />
                                    </UFormField>
                                    <UFormField label="Use Aikar Flag" class="justify-center" name="useAikarFlag">
                                        <UCheckbox v-model="wizardModel.useAikarFlag" />
                                    </UFormField>
                                </div>
                                <USeparator label="Container Options" />
                                <div class="flex gap-2 m-2 justify-center">
                                    <UButton color="primary" label="Advance Settings" icon="i-lucide-book-open-text"
                                        :ui="{ label: 'text-center w-full' }" disabled />
                                </div>
                            </div>
                        </UForm>
                    </template>
                    <template #confirm>
                        <div class="flex">
                            <h1>Table that display the form data</h1>
                        </div>
                        <UButton trailing-icon="i-material-symbols-settings">
                            View More
                        </UButton>
                    </template>
                </UStepper>

                <div class="flex gap-2 justify-between mt-4">
                    <UButton leading-icon="i-lucide-arrow-left" :disabled="!stepper?.hasPrev" @click="stepper?.prev()">
                        Prev
                    </UButton>


                    <UButton v-if="stepper?.hasNext" trailing-icon="i-lucide-arrow-right" :disabled="!stepper?.hasNext"
                        @click="stepper?.next()">
                        Next
                    </UButton>
                    <UButton v-else-if="stepper?.hasNext === false" trailing-icon="i-lucide-plane-takeoff"
                        @click="createServer">
                        Create
                    </UButton>
                </div>
            </div>
        </template>
    </UModal>
</template>