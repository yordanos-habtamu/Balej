<template>
  <div class="bg-[#020205] text-white font-sans min-h-screen overflow-hidden selection:bg-cyan-500/30 relative">
    <!-- Background Canvas -->
    <canvas ref="bgCanvasRef" class="fixed inset-0 z-0"></canvas>
    
    <!-- Cursor Glow -->
    <div ref="glowRef" class="fixed w-[600px] h-[600px] bg-cyan-500/10 rounded-full blur-[120px] pointer-events-none z-0 transform -translate-x-1/2 -translate-y-1/2 transition-transform duration-300 ease-out hidden md:block"></div>

    <div class="relative z-10 min-h-screen flex items-center justify-center p-6">
      <div id="auth-container" class="max-w-4xl w-full grid lg:grid-cols-2 bg-black/40 backdrop-blur-xl border border-white/5 shadow-[0_0_100px_rgba(0,243,255,0.05)] relative overflow-hidden">
        <!-- Scanning Line -->
        <div class="absolute top-0 left-0 w-full h-[1px] bg-cyan-500/30 animate-[scan_4s_linear_infinite] pointer-events-none"></div>

        <!-- Left Side: Visual/Info -->
        <div class="hidden lg:flex flex-col justify-center p-16 border-r border-white/5 relative overflow-hidden">
          <div class="relative z-10">
            <h2 class="text-[8rem] font-black tracking-tighter text-outline text-transparent opacity-10 italic mb-0 leading-none">AUTH</h2>
            <div class="space-y-8 mt-8">
              <div class="flex items-center gap-4">
                <div class="w-2 h-2 bg-cyan-500 animate-pulse shadow-[0_0_10px_#00f3ff]"></div>
                <span class="font-mono text-[10px] tracking-[0.4em] text-gray-400 uppercase">Neural_Link_Protocol</span>
              </div>
              <p class="text-gray-500 font-mono text-xs leading-relaxed uppercase tracking-widest">
                Establishing secure uplink to global talent mesh. <br/>
                Encryption: AES_512_Quantum_Safe <br/>
                Status: Awaiting_Credentials
              </p>
            </div>
          </div>
          <!-- Decorative Elements -->
          <div class="absolute -bottom-20 -left-20 w-64 h-64 bg-cyan-500/5 blur-[100px] rounded-full"></div>
        </div>

        <!-- Right Side: Form -->
        <div class="p-8 md:p-16 flex flex-col justify-center bg-cyan-500/[0.01]">
          <div class="w-full">
            <div class="mb-12">
              <h2 class="text-5xl font-black uppercase italic tracking-tighter">System_<span class="text-cyan-500">Login</span></h2>
              <p class="text-[9px] text-gray-500 tracking-[0.5em] uppercase mt-2 font-mono">Authorization_Required</p>
            </div>

            <form class="space-y-10" @submit.prevent="handleLogin">
              <div class="relative group">
                <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">Neural_ID (Email)</label>
                <input 
                  type="email" 
                  v-model="loginData.email"
                  placeholder="OPERATOR@NEURAL.IO" 
                  class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none uppercase"
                  required
                >
              </div>
              <div class="relative group">
                <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">Access_Hash (Password)</label>
                <input 
                  type="password" 
                  v-model="loginData.password"
                  placeholder="••••••••" 
                  class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none"
                  required
                >
              </div>
              
              <div class="space-y-4">
                <button 
                  type="submit"
                  :disabled="mutationLoading"
                  class="w-full py-6 bg-cyan-500 text-black font-black uppercase text-[12px] tracking-[0.5em] hover:bg-white transition-all shadow-[0_20px_40px_rgba(0,243,255,0.1)] disabled:opacity-50"
                >
                  {{ mutationLoading ? 'Syncing...' : 'Establish_Session' }}
                </button>
                <nuxt-link to="/auth/register" class="block w-full py-4 border border-white/10 text-center text-gray-500 font-black uppercase text-[10px] tracking-[0.3em] hover:text-cyan-500 transition-all">
                  Initialize_New_Node
                </nuxt-link>
              </div>
              <p v-if="mutationError" class="text-[8px] text-red-500 text-center font-mono uppercase mt-4">{{ errorMessage }}</p>
            </form>
          </div>
        </div>
      </div>
    </div>

    <!-- Ticker Footer -->
    <div class="fixed bottom-0 w-full bg-black/90 backdrop-blur-2xl border-t border-white/10 py-4 z-50">
      <div class="flex items-center gap-12 whitespace-nowrap animate-ticker">
        <div v-for="n in 10" :key="n" class="flex items-center gap-4">
          <span class="w-2 h-2 bg-cyan-500 rounded-full animate-pulse"></span>
          <span class="font-mono text-[10px] tracking-[0.3em] text-gray-400 uppercase italic font-bold">System_Status: Secure</span>
          <span class="text-cyan-900 px-12">>>></span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import * as THREE from 'three'
import { onMounted, ref, reactive, computed } from 'vue'
import { jwtDecode } from 'jwt-decode'
import { useAuthStore } from '~/stores/auth'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'

const authStore = useAuthStore()
const tokenCookie = useCookie('token')
const client = useUrqlClient()

const mutationLoading = ref(false)
const mutationError = ref<CombinedError | null>(null)
const errorMessage = computed(() => mutationError.value?.message || '')

const loginData = reactive({
  email: '',
  password: '',
})

const LoginMutation = gql`
  mutation LoginUser($input: LoginInput!) {
    LoginUser(input: $input) {
      token
      user {
        id
        first_name
        last_name
        email
        role
      }
    }
  }
`

const handleLogin = async () => {
  mutationLoading.value = true
  mutationError.value = null
  
  try {
    const result = await client
      .mutation(LoginMutation, { input: { email: loginData.email, password: loginData.password } })
      .toPromise()

    if (result.data?.LoginUser?.token) {
      const token = result.data.LoginUser.token
      tokenCookie.value = token
      const decoded: any = jwtDecode(token)
      authStore.setUser(decoded)

      if (decoded.role === 'admin') await navigateTo('/admin/dashboard')
      else if (decoded.role === 'talent') await navigateTo('/talent')
      else if (decoded.role === 'client') await navigateTo('/client')
    } else if (result.error) {
      mutationError.value = result.error
    }
  } catch (err: any) {
    mutationError.value = err
  } finally {
    mutationLoading.value = false
  }
}

// --- THREE.JS LOGIC ---
const bgCanvasRef = ref(null)
const glowRef = ref(null)
let scene, camera, renderer, torus, particles

onMounted(() => {
  if (!bgCanvasRef.value) return

  scene = new THREE.Scene()
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
  camera.position.z = 5

  renderer = new THREE.WebGLRenderer({ canvas: bgCanvasRef.value, antialias: true, alpha: true })
  renderer.setSize(window.innerWidth, window.innerHeight)
  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))

  const torusGeom = new THREE.TorusGeometry(3, 0.8, 16, 100)
  const torusMat = new THREE.MeshBasicMaterial({ color: 0x00f3ff, wireframe: true, transparent: true, opacity: 0.1 })
  torus = new THREE.Mesh(torusGeom, torusMat)
  scene.add(torus)

  const particleCount = 5000
  const posArr = new Float32Array(particleCount * 3)
  for(let i=0; i<particleCount*3; i++) posArr[i] = (Math.random() - 0.5) * 20
  const particleGeom = new THREE.BufferGeometry()
  particleGeom.setAttribute('position', new THREE.BufferAttribute(posArr, 3))
  const particleMat = new THREE.PointsMaterial({ color: 0x00f3ff, size: 0.01, transparent: true, opacity: 0.5 })
  particles = new THREE.Points(particleGeom, particleMat)
  scene.add(particles)

  const animate = () => {
    const time = Date.now() * 0.0005
    torus.rotation.x = time * 0.2
    torus.rotation.y = time * 0.3
    particles.rotation.y = time * 0.1
    renderer.render(scene, camera)
    requestAnimationFrame(animate)
  }
  animate()

  window.addEventListener('resize', () => {
    camera.aspect = window.innerWidth / window.innerHeight
    camera.updateProjectionMatrix()
    renderer.setSize(window.innerWidth, window.innerHeight)
  })

  window.addEventListener('mousemove', (e) => {
    if(glowRef.value){
      glowRef.value.style.left = `${e.clientX}px`
      glowRef.value.style.top = `${e.clientY}px`
    }
  })
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@300;400;700;900&display=swap');

.font-sans { font-family: 'Space Grotesk', sans-serif; }
.font-outline { -webkit-text-stroke: 1px rgba(34, 211, 238, 0.5); color: transparent; }

@keyframes scan {
  0% { transform: translateY(-100px); opacity: 0; }
  50% { opacity: 1; }
  100% { transform: translateY(800px); opacity: 0; }
}

@keyframes ticker { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
.animate-ticker { animation: ticker 30s linear infinite; }

input:focus {
  box-shadow: 0 0 15px rgba(6, 182, 212, 0.1);
}

::-webkit-scrollbar { width: 4px; }
::-webkit-scrollbar-thumb { background: #06b6d4; }
</style>
