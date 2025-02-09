export const useDataStore = defineStore('data', () => {
  const servers = reactive([
    {
      name: "hello",
      id: "ops",
      online: true
    },
    {
      name: "world",
      id: "spo",
      online: false
    }
  ])

  const curServer = ref<string>(servers[0]!.id)

  return {servers, curServer}
})