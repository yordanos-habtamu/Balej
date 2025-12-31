<template>
  <div class="bg-[#020205] text-white font-sans min-h-screen overflow-hidden selection:bg-cyan-500/30 relative">
    <!-- Background Canvas -->
    <canvas ref="bgCanvasRef" class="fixed inset-0 z-0"></canvas>
    
    <!-- Cursor Glow -->
    <div ref="glowRef" class="fixed w-[600px] h-[600px] bg-cyan-500/10 rounded-full blur-[120px] pointer-events-none z-0 transform -translate-x-1/2 -translate-y-1/2 transition-transform duration-300 ease-out hidden md:block"></div>

    <div class="relative z-10 min-h-screen flex items-center justify-center p-6">
      <div id="onboarding-container" class="max-w-5xl w-full grid lg:grid-cols-2 bg-black/40 backdrop-blur-xl border border-white/5 shadow-[0_0_100px_rgba(0,243,255,0.05)] relative overflow-hidden">
        <!-- Scanning Line -->
        <div class="absolute top-0 left-0 w-full h-[1px] bg-cyan-500/30 animate-[scan_4s_linear_infinite] pointer-events-none"></div>

        <!-- Left Side: Visual/Info -->
        <div class="hidden lg:flex flex-col justify-center p-16 border-r border-white/5 relative overflow-hidden">
          <div class="relative z-10">
            <h2 class="text-[8rem] font-black tracking-tighter text-outline text-transparent opacity-10 italic mb-0 leading-none">CLIENT</h2>
            <div class="space-y-8 mt-8">
              <div class="flex items-center gap-4">
                <div class="w-2 h-2 bg-cyan-500 animate-pulse shadow-[0_0_100px_#00f3ff]"></div>
                <span class="font-mono text-[10px] tracking-[0.4em] text-gray-400 uppercase">Organization_Setup</span>
              </div>
              <p class="text-gray-500 font-mono text-xs leading-relaxed uppercase tracking-widest">
                Configuring your recruitment parameters for the global mesh. <br/>
                Step: 03 // Role_Specific_Onboarding <br/>
                Status: Awaiting_Organization_Data
              </p>
            </div>
          </div>
          <div class="absolute -bottom-20 -left-20 w-64 h-64 bg-cyan-500/5 blur-[100px] rounded-full"></div>
        </div>

        <!-- Right Side: Form -->
        <div class="p-8 md:p-16 flex flex-col justify-center bg-cyan-500/[0.01]">
          <div class="mb-12">
            <h2 class="text-5xl font-black uppercase italic tracking-tighter text-cyan-500">Sync_<span class="text-white">Org</span></h2>
            <p class="text-[9px] text-gray-500 tracking-[0.5em] uppercase mt-2 font-mono">Client_Onboarding_Protocol</p>
          </div>

          <form class="space-y-8" @submit.prevent="handleOnboarding">
            <!-- Company Name -->
            <div class="relative group">
              <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">Entity_Name // Company</label>
              <input 
                type="text" 
                v-model="formData.companyName" 
                placeholder="Enter company or individual name" 
                class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none uppercase"
                required
              >
            </div>

            <!-- Company Size -->
            <div class="relative group">
              <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">Organization_Scale</label>
              <select v-model="formData.companySize" class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none uppercase appearance-none cursor-pointer">
                <option value="" disabled class="bg-[#020205]">Select Size</option>
                <option value="1-10" class="bg-[#020205]">1-10 // Startup</option>
                <option value="11-50" class="bg-[#020205]">11-50 // Growth</option>
                <option value="51-200" class="bg-[#020205]">51-200 // Mid-Market</option>
                <option value="201+" class="bg-[#020205]">201+ // Enterprise</option>
              </select>
            </div>

            <!-- Hiring Intent -->
            <div class="relative group">
              <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">Recruitment_Intent</label>
              <div class="grid grid-cols-2 gap-4">
                <button 
                  v-for="intent in ['jobs', 'projects', 'both']" 
                  :key="intent"
                  type="button"
                  @click="formData.hiringIntent = intent"
                  :class="formData.hiringIntent === intent ? 'border-cyan-500 bg-cyan-500/10 text-cyan-500' : 'border-white/10 bg-white/5 text-gray-500'"
                  class="p-3 border text-[9px] font-mono uppercase tracking-widest transition-all"
                >
                  {{ intent }}
                </button>
              </div>
            </div>

            <!-- Budget Range -->
            <div class="relative group">
              <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">Annual_Budget_Range // Optional</label>
              <select v-model="formData.budget" class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none uppercase appearance-none cursor-pointer">
                <option value="" class="bg-[#020205]">Not Specified</option>
                <option value="<50k" class="bg-[#020205]">Under $50k</option>
                <option value="50k-200k" class="bg-[#020205]">$50k - $200k</option>
                <option value="200k-1M" class="bg-[#020205]">$200k - $1M</option>
                <option value="1M+" class="bg-[#020205]">$1M+</option>
              </select>
            </div>
            
            <div class="space-y-4 pt-4">
              <button 
                type="submit"
                :disabled="loading"
                class="w-full py-6 bg-cyan-500 text-black font-black uppercase text-[12px] tracking-[0.5em] hover:bg-white transition-all shadow-[0_20px_40px_rgba(0,243,255,0.1)] disabled:opacity-50"
              >
                {{ loading ? 'Synchronizing...' : 'Complete_Onboarding' }}
              </button>
            </div>
          </form>
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
import { onMounted, ref, reactive } from 'vue'

const router = useRouter()
const loading = ref(false)

const formData = reactive({
  companyName: '',
  companySize: '',
  hiringIntent: 'jobs',
  budget: '',
})

const handleOnboarding = async () => {
  loading.value = true
  // Simulate API call
  setTimeout(() => {
    loading.value = false
    router.push('/client/dashboard')
  }, 2000)
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

input:focus, select:focus {
  box-shadow: 0 0 15px rgba(6, 182, 212, 0.1);
}

::-webkit-scrollbar { width: 4px; }
::-webkit-scrollbar-thumb { background: #06b6d4; }
</style>
