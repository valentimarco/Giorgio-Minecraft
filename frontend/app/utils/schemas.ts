import { z } from "zod";

export const SERVERTYPES = ['VANILLA', 'CUSTOM', 'CURSEFORGE', 'FTBA', 'MODRINTH'] as const

export const serverWizardSchema = z.object({
    type: z.enum(SERVERTYPES),
    version: z.string(),
    maxPlayers: z.number().positive().min(1).int().safe(),
    whitelist: z.string().optional(),
    memory: z.number().positive().min(1).int().safe(),
    useAikarFlag: z.boolean()
})

export type ServerWizard = z.output<typeof serverWizardSchema>

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