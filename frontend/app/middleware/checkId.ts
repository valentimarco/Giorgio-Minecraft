export default defineNuxtRouteMiddleware((to) => {
    const {servers} = useDataStore()

    if (!servers.some((s) => s.id === to.params.id)) {
        throw createError({ statusCode: 400, statusMessage: 'Page Not Found' })
    }
  })
  