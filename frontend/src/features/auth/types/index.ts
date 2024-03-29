import { z } from 'zod'

export const UserSchema = z.object({
  id: z.number(),
  email: z.string(),
  name: z.string(),
})

export type User = z.infer<typeof UserSchema>
