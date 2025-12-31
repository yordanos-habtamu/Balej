import { defineStore } from 'pinia'

export const useThemeStore = defineStore('theme', {
    state: () => ({
        primaryColor: 'blue', // Default Tailwind color name
        isDarkMode: false,
        sidebarCollapsed: false,
        notificationsEnabled: true,
        language: 'en',
    }),
    actions: {
        setPrimaryColor(color: string) {
            this.primaryColor = color
        },
        toggleDarkMode() {
            this.isDarkMode = !this.isDarkMode
        },
        toggleSidebar() {
            this.sidebarCollapsed = !this.sidebarCollapsed
        },
        toggleNotifications() {
            this.notificationsEnabled = !this.notificationsEnabled
        },
        setLanguage(lang: string) {
            this.language = lang
        }
    },
    persist: true // Requires @pinia-plugin-persistedstate if installed, otherwise manual handling might be needed but for now we assume basic store
})
