import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null as null | { exp: string;role: string ;user_id:string}
  }),
  actions: {
    setUser(user: { exp: string; role: string; user_id: string } | null) {
      this.user = user
    },
    logout() {
      this.user = null
    }
  }
})
