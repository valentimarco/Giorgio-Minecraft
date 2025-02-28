export default defineNuxtRouteMiddleware(() => {
    const { loggedIn, user, session } = useUserSession()

    console.log(user.value)

    console.log(session.value)
  
    // redirect the user to the login screen if they're not authenticated
    if (!loggedIn.value) {
      return navigateTo('/')
    }
})