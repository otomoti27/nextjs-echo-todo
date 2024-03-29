import { cookies } from 'next/headers'
import { getCsrf } from '@/features/auth'

export const dynamic = 'force-dynamic'
export const fetchCache = 'force-no-store'

export async function GET() {
  const token = await getCsrf()

  cookies().set({
    name: '_csrf',
    value: token,
    httpOnly: true,
    sameSite: 'none',
    secure: true,
  })

  const response = Response.json(
    { csrf_token: token },
    {
      status: 200,
    },
  )

  return response
}
