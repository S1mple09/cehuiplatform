import { defineStore } from 'pinia'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    isAuthenticated: false,
    user: null
  }),
  actions: {
    login(userData) {
      this.isAuthenticated = true
      this.user = userData
    },
    logout() {
      this.isAuthenticated = false
      this.user = null
    },
    updateUser(userData) {
      this.user = { ...this.user, ...userData }
    }
  },
  persist: true
})