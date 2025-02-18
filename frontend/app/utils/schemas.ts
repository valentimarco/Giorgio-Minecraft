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