import { jwtDecode } from 'jwt-decode'
import { useAuthStore } from '~/stores/auth'

export default defineNuxtRouteMiddleware((to) => {
  const token = useCookie('auth_token')?.value

  if (!token) return navigateTo('/login')

  try {
    const decoded: any = jwtDecode(token)

    const path = to.path

    if (path.startsWith('/admin') && decoded.role !== 'admin') {
      return navigateTo('/unauthorized')
    }

    if (path.startsWith('/user') && decoded.role !== 'user') {
      return navigateTo('/unauthorized')
    }

    const auth = useAuthStore()
    auth.setUser(decoded)
  } catch (err) {
    return navigateTo('/login')
  }
})
