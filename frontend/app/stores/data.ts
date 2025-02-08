export const useDataStore = defineStore('data', () => {
  const servers = reactive([
    {
      name: "hello",
      id: "ops"
    },
    {
      name: "world",
      id: "spo"
    }
  ])

  const curServer = ref<string>(servers[0]!.id)

  return {servers, curServer}
})