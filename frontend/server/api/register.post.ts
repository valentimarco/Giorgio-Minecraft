export default defineEventHandler(async (event) => {
    const { backendUrl } = useRuntimeConfig(event)
    const body = await readValidatedBody(event, body => registerSchema.parse(body))

    console.log(body)

    const { token } = await $fetch<{ token: string }>(`${backendUrl}/api/v1/register`, {
        method: 'post',
        body
    })


    await setUserSession(event, {
        token,
        loggedInAt: Date.now()
    })




    return setResponseStatus(event, 201)
})