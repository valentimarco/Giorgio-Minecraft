import { z } from "zod"

export const loginSchema = z.object({
    username: z.string(),
    password: z.string().min(8, 'Must be at least 8 characters')
})

export type UserLogin = z.output<typeof loginSchema>

export const registerSchema = loginSchema.extend({
    confirmPassword: z.string() 
}).refine((data) => data.password === data.confirmPassword, {
    message: "The passwords don't match!",
    path: ['confirmPassword']
})

export type UserRegister = z.output<typeof registerSchema>