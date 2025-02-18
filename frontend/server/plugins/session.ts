export default defineNitroPlugin(() => {
    // Called when the session is fetched during SSR for the Vue composable (/api/_auth/session)
    // Or when we call useUserSession().fetch()
    sessionHooks.hook('fetch', async (session, event) => {
      // extend User Session by calling your database
      // or
      // throw createError({ ... }) if session is invalid for example
        const body = await readBody(event)
        console.log(body)
    })
  
    // Called when we call useUserSession().clear() or clearUserSession(event)
    // sessionHooks.hook('clear', async (session, event) => {
    //   // Log that user logged out
    // })
})