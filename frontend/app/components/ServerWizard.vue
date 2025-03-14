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

const [DefineTemplate, ReuseTemplate] = createReusableTemplate<{
    isResume: boolean
}>()

const onDrop = (files: File[] | null) => {
    console.log(files)
}

useDropZone(dropZone, {
    onDrop,
    dataTypes: ['image/png'],
    multiple: false,
    preventDefaultForUnhandled: false,
})

const { open, onChange } = useFileDialog({
    accept: 'image/*', // Set to accept only image files
    directory: false, // Select directories instead of files if set true
})

onChange((files) => {
    const file = files?.item(0)
    // TODO: idk if needs
    if (!file) {
        console.log("Something goes wrong")
    }
    onDrop([file!])
})


const createServer = () => {
    console.log(wizardModel)
}
</script>

<template>
    <DefineTemplate v-slot="{ isResume }">
        <div v-if="!isResume" class="grid border-1 rounded-lg p-2 m-2 h-full">
            <USeparator label="Server Options" :ui="{ label: 'text-xl' }" />
            <div class="grid justify-evenly m-2 p-2 gap-y-5 h-full">
                <UFormField label="Type" name="type">
                    <USelectMenu v-model="wizardModel.type" :items="[...SERVERTYPES]" :search-input="false" />
                </UFormField>
                <UFormField label="Version" name="version">
                    <UInput v-model="wizardModel.version" placeholder="Version" />
                </UFormField>
                <UFormField label="Max Players" name="maxPlayers">
                    <UInput v-model="wizardModel.maxPlayers" type="number" placeholder="20" />
                </UFormField>
                <UFormField label="Whitelist" class="col-start-2" name="whitelist">
                    <UTextarea v-model="wizardModel.whitelist" placeholder="Type username or UUID" />
                </UFormField>
            </div>

            <USeparator label="Container Options" :ui="{ label: 'text-xl' }" />
            <div class="grid justify-evenly m-2 p-2 gap-y-5 h-full">
                <UFormField label="Memory" name="memory">
                    <UInput v-model="wizardModel.memory" type="number" placeholder="6" />
                </UFormField>
                <UFormField label="Use Aikar Flag" class="justify-center" name="useAikarFlag">
                    <UCheckbox v-model="wizardModel.useAikarFlag" />
                </UFormField>
            </div>
            <div class="flex p-2 mt-4 justify-center">
                <UButton class="w-fit" color="primary" label="Advance Settings" icon="i-lucide-book-open-text"
                    :ui="{ label: 'text-center w-full' }" disabled />
            </div>
        </div>
        <div v-else class="grid p-2 m-2 h-full space-y-4">
            <USeparator label="Server Options" :ui="{ label: 'text-xl' }" />
            <div class="grid justify-evenly m-2 p-2 gap-y-5 h-full">
                <UFormField label="Type" name="type">
                    <USelectMenu v-model="wizardModel.type" :items="[...SERVERTYPES]" :search-input="false" disabled />
                </UFormField>
                <UFormField label="Version" name="version">
                    <UInput v-model="wizardModel.version" placeholder="Version" disabled />
                </UFormField>
                <UFormField label="Max Players" name="maxPlayers">
                    <UInput v-model="wizardModel.maxPlayers" type="number" placeholder="20" disabled />
                </UFormField>
                <UFormField label="Whitelist" class="col-start-2" name="whitelist">
                    <UTextarea v-model="wizardModel.whitelist" placeholder="Type username or UUID" disabled />
                </UFormField>
            </div>

            <USeparator label="Container Options" :ui="{ label: 'text-xl' }" />
            <div class="grid justify-evenly m-2 p-2 gap-y-5 h-full">
                <UFormField label="Memory" name="memory">
                    <UInput v-model="wizardModel.memory" type="number" placeholder="6" disabled />
                </UFormField>
                <UFormField label="Use Aikar Flag" class="justify-center" name="useAikarFlag">
                    <UCheckbox v-model="wizardModel.useAikarFlag" disabled />
                </UFormField>
            </div>
        </div>
    </DefineTemplate>

    <div class="grid grid-rows-[1fr_3rem] overflow-hidden">
        <UStepper ref="stepper" :items="items" size="xl" class="">
            <template #provider>
                <div class="grid grid-cols-[10rem_1fr] border-1 rounded-lg p-2 m-2 h-full">
                    <div class="flex flex-col p-2 gap-1 m-2">
                        <UButton label="CurseForge" icon="i-lucide-anvil" :ui="{ label: 'text-center w-full' }"
                            @click="selectedProvider = 'CF'" />
                        <UButton label="Modrinth" icon="i-lucide-anvil" :ui="{ label: 'text-center w-full' }"
                            @click="selectedProvider = 'MR'" />
                        <UButton label="Custom" icon="i-lucide-anvil" :ui="{ label: 'text-center w-full' }"
                            @click="selectedProvider = 'CU'" />
                    </div>


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
                    <!-- Use this https://vueuse.org/core/useFileDialog/ -->
                    <div v-else class="flex justify-center">
                        <div ref="dropZone" class="flex justify-center items-center border border-dashed w-full"
                            @click="open">
                            <h2 class="p-2">Drop files here or click this area!</h2>
                        </div>
                    </div>

                </div>
            </template>
            <template #config>
                <UForm :schema="serverWizardSchema" :state="wizardModel" class="space-y-4 h-full" >
                    <ReuseTemplate :is-resume="false"/>
                </UForm >
            </template>
            <template #confirm>
                <div class="grid grid-rows-[1fr_5rem] border-1 rounded-lg p-2 m-2 h-full">
                    <ReuseTemplate :is-resume="true"/>
                    <div class="flex flex-col p-2 mt-4 items-center">
                        <UButton class="w-[10rem]" icon="i-material-symbols-settings" label="View More"
                            :ui="{ label: 'text-center w-full text-xl' }" size="sm" />
                    </div>
                </div>
            </template>
        </UStepper>

        <div class="flex gap-2 justify-around mt-4">
            <UButton leading-icon="i-lucide-arrow-left" :disabled="!stepper?.hasPrev" @click="stepper?.prev()">
                Prev
            </UButton>


            <UButton v-if="stepper?.hasNext" trailing-icon="i-lucide-arrow-right" :disabled="!stepper?.hasNext"
                @click="stepper?.next()">
                Next
            </UButton>
            <UButton v-else trailing-icon="i-lucide-plane-takeoff" @click="createServer">
                Create
            </UButton>
        </div>
    </div>
</template>