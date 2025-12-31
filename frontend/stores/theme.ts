import { defineStore } from 'pinia'

export const useThemeStore = defineStore('theme', {
    state: () => ({
        // THEME AESTHETICS
        accentColor: '#00f3ff', // The signature Cyan glow
        backgroundBase: '#020205', // Deep Obsidian
        
        // SYSTEM STATES
        isUplinkActive: false,    // Tracks if the globe is zoomed/logged in
        glitchMode: false,        // Global toggle for the glitch animation
        hudVisibility: true,      // Toggles UI overlays (scanners, news feed)
        
        // INTERFACE CONFIG
        terminalCollapsed: false, // Formerly sidebar
        encryptionLevel: 'AES_512',
        language: 'EN_US_NEURAL',  // Stylized language indicator
    }),
    
    actions: {
        // Updates the primary neon glow color across the app
        setAccentColor(hex: string) {
            this.accentColor = hex
            // Optional: Update CSS variable dynamically
            document.documentElement.style.setProperty('--accent-glow', hex)
        },
        
        // Triggers the "Neural Link" sequence we built
        initializeUplink() {
            this.isUplinkActive = true
            this.glitchMode = true
            setTimeout(() => { this.glitchMode = false }, 1500)
        },
        
        toggleHUD() {
            this.hudVisibility = !this.hudVisibility
        },
        
        toggleTerminal() {
            this.terminalCollapsed = !this.terminalCollapsed
        },

        setNeuralLanguage(lang: string) {
            this.language = lang
        }
    },
    persist: true 
})