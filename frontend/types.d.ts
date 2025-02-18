declare module '@nuxt/ui' {
    interface SelectMenuItem<T = unknown> {
      value?: T
    }
}

declare module '#auth-utils' {
  // interface User {
    
  // }

  interface UserSession {
    readonly token: string
    readonly loggedInAt: number
  }

  // interface SecureSessionData {
  //   // Add your own fields
  // }
}

export {}