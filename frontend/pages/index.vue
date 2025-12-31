<template>
  <div class="bg-[#020205] text-white font-sans overflow-x-hidden selection:bg-cyan-500/30 animate-bg-pulse">
    
    <div id="cursor-glow" ref="glowRef" class="fixed w-[600px] h-[600px] bg-cyan-500/10 rounded-full blur-[120px] pointer-events-none z-0 transform -translate-x-1/2 -translate-y-1/2 transition-transform duration-300 ease-out hidden md:block"></div>
    
    <Transition name="scale-fade">
      <div v-if="activeNode" class="fixed inset-0 z-[100] flex items-center justify-center p-8 bg-black/60 backdrop-blur-md">
        <div class="bg-[#050505] border border-cyan-500 p-8 max-w-sm w-full shadow-[0_0_70px_rgba(0,243,255,0.2)] relative">
          <div class="flex justify-between items-center mb-6">
            <span class="font-mono text-cyan-500 text-[10px] tracking-[0.3em] uppercase">Status: Connected</span>
            <button @click="closeNode" class="text-white hover:text-cyan-500 transition-colors">✕</button>
          </div>
          <h4 class="text-4xl font-black italic uppercase mb-2 tracking-tighter">NODE_{{ activeNode.id }}</h4>
          <div class="space-y-2 mb-8 font-mono text-[10px] text-gray-400">
            <div class="flex justify-between"><span>LATENCY</span> <span class="text-cyan-400">{{ activeNode.latency }}ms</span></div>
            <div class="flex justify-between"><span>UPTIME</span> <span class="text-cyan-400">99.99%</span></div>
            <div class="flex justify-between"><span>CLUSTER</span> <span class="text-cyan-400">DECENTRALIZED</span></div>
          </div>
          <button @click="closeNode" class="w-full py-4 bg-cyan-500 text-black font-black uppercase text-[10px] tracking-[0.3em] hover:bg-white transition-all">Interface with Node</button>
        </div>
      </div>
    </Transition>

    <canvas ref="canvasRef" @click="handleCanvasClick" class="fixed inset-0 z-0 cursor-crosshair"></canvas>

    <div class="relative z-10">
      <nav class="fixed top-0 w-full z-50 px-6 md:px-12 py-8 flex justify-between items-center border-b border-white/5 backdrop-blur-md bg-black/40">
        <div class="flex items-center gap-4">
          <div class="w-3 h-3 bg-cyan-500 animate-pulse shadow-[0_0_15px_#00f3ff]"></div>
          <span class="font-black tracking-tighter text-2xl uppercase italic">Neural<span class="text-cyan-500">Talent</span></span>
        </div>
        <div class="hidden lg:flex gap-12 text-[10px] uppercase tracking-[0.5em] font-bold text-gray-400">
          <a href="#workflow" class="hover:text-cyan-400 transition">01_Process</a>
          <a href="#tech" class="hover:text-cyan-400 transition">02_Tech_Stack</a>
          <a href="#pricing" class="hover:text-cyan-400 transition">03_Pricing</a>
          <a href="#newsletter" class="hover:text-cyan-400 transition">04_Stream</a>
        </div>
        <a href="#signup" class="text-[10px] font-black uppercase tracking-[0.2em] px-8 py-3 border border-cyan-500 text-cyan-500 hover:bg-cyan-500 hover:text-black transition-all">Initialize_Uplink</a>
      </nav>

      <section class="h-screen flex flex-col justify-center items-center px-6 text-center pt-24">
        <div class="px-6 py-2 border border-cyan-500/20 rounded-full mb-12 bg-cyan-500/5">
          <span class="text-cyan-400 font-mono text-[10px] tracking-[0.6em] uppercase animate-pulse italic">Network_Expansion_Active</span>
        </div>
        <h1 class="text-6xl md:text-8xl lg:text-[10rem] font-black tracking-tighter leading-[0.8] mb-8 uppercase">
          Build at <br /><span class="text-transparent bg-clip-text bg-gradient-to-r from-cyan-400 to-blue-600 font-outline italic">Light Speed.</span>
        </h1>
        <p class="max-w-2xl text-gray-500 text-sm md:text-xl font-light leading-relaxed font-mono uppercase tracking-tight">
          Interfacing with global talent clusters through decentralized compute. <br/>
          <span class="text-cyan-500">Click the glowing pins to establish a direct link.</span>
        </p>
      </section>

      <section class="py-32 border-y border-white/5 bg-black/40 backdrop-blur-xl scroll-mt-24">
        <div class="max-w-7xl mx-auto px-6 grid grid-cols-2 md:grid-cols-4 gap-16">
          <div v-for="stat in stats" :key="stat.label" class="flex flex-col items-center">
            <span class="text-cyan-500 font-mono text-[10px] tracking-[0.4em] mb-4 uppercase text-center w-full">{{ stat.label }}</span>
            <span class="text-5xl md:text-7xl font-black italic tracking-tighter">{{ stat.value }}</span>
          </div>
        </div>
      </section>

      <section id="workflow" class="py-40 px-6 scroll-mt-24">
        <div class="max-w-6xl mx-auto">
          <div class="flex items-center gap-6 mb-20">
            <h2 class="text-xs font-bold tracking-[0.8em] text-cyan-500 uppercase">Step_01 // THE_PIPELINE</h2>
            <div class="h-px flex-1 bg-gradient-to-r from-cyan-500/50 to-transparent"></div>
          </div>
          <div class="grid lg:grid-cols-3 gap-1px bg-white/10 border border-white/10">
            <div v-for="(step, i) in workflowSteps" :key="i" class="bg-black p-16 hover:bg-cyan-950/10 transition-colors group">
              <span class="block text-cyan-500 font-mono mb-8 text-xl">0{{ i+1 }}_</span>
              <h3 class="text-3xl font-black uppercase mb-6 italic group-hover:translate-x-2 transition-transform">{{ step.title }}</h3>
              <p class="text-gray-500 leading-relaxed font-mono text-xs uppercase tracking-wider">{{ step.desc }}</p>
            </div>
          </div>
        </div>
      </section>

      <section id="tech" class="py-40 px-6 bg-cyan-500/[0.02] scroll-mt-24">
        <div class="max-w-6xl mx-auto grid lg:grid-cols-2 gap-20 items-center">
          <div>
            <h2 class="text-6xl font-black uppercase italic mb-8 tracking-tighter leading-none">Powered by <br/><span class="text-cyan-500">Decentralized AI.</span></h2>
            <p class="text-gray-400 font-mono text-sm leading-loose uppercase tracking-widest mb-12">We leverage zero-knowledge proofs and neural vector matching to find the perfect candidate without exposing proprietary data.</p>
            <div class="grid grid-cols-2 gap-6 font-mono text-[10px] text-cyan-500">
              <div class="p-4 border border-cyan-500/20 bg-cyan-500/5 uppercase tracking-widest">>> Llama-3 Optimization</div>
              <div class="p-4 border border-cyan-500/20 bg-cyan-500/5 uppercase tracking-widest">>> ZK-Proof Security</div>
              <div class="p-4 border border-cyan-500/20 bg-cyan-500/5 uppercase tracking-widest">>> Vector Indexing</div>
              <div class="p-4 border border-cyan-500/20 bg-cyan-500/5 uppercase tracking-widest">>> Realtime Sync</div>
            </div>
          </div>
          <div class="relative aspect-square border border-white/5 flex items-center justify-center bg-black/40">
            <div class="absolute inset-0 bg-gradient-to-br from-cyan-500/20 to-transparent animate-pulse"></div>
            <span class="text-white/20 font-black text-[12rem] italic select-none uppercase">CORE</span>
          </div>
        </div>
      </section>

      <section id="pricing" class="py-40 px-6 scroll-mt-24">
        <div class="max-w-6xl mx-auto grid lg:grid-cols-3 gap-0">
          <div v-for="plan in plans" :key="plan.name" class="p-16 border border-white/10 bg-black hover:border-cyan-500 transition-all z-10 hover:z-20">
            <h4 class="text-cyan-500 font-mono text-[10px] tracking-[0.5em] mb-12 uppercase">{{ plan.name }}</h4>
            <div class="text-7xl font-black mb-12 tracking-tighter italic">{{ plan.price }}<span class="text-xs text-gray-600">.usd</span></div>
            <ul class="space-y-6 mb-16 text-[10px] font-mono text-gray-400 uppercase tracking-[0.2em]">
              <li v-for="f in plan.features" :key="f" class="flex gap-4">
                <span class="text-cyan-500">>></span> {{ f }}
              </li>
            </ul>
            <button class="w-full py-6 bg-white text-black font-black uppercase text-[10px] tracking-[0.4em] hover:bg-cyan-500 transition-all">Provision_Node</button>
          </div>
        </div>
      </section>

      <section id="newsletter" class="py-40 px-6 border-t border-white/5 scroll-mt-24">
        <div class="max-w-4xl mx-auto bg-white/[0.02] p-20 border border-white/5 relative overflow-hidden">
          <div class="absolute top-0 right-0 w-64 h-64 bg-cyan-500/5 blur-[100px] rounded-full"></div>
          <div class="relative z-10 text-center">
            <h2 class="text-6xl font-black uppercase tracking-tighter mb-8 italic">Join the <span class="text-cyan-500 font-outline">Data Stream.</span></h2>
            <p class="text-gray-500 font-mono text-xs uppercase tracking-[0.3em] mb-12 max-w-md mx-auto leading-loose">Subscribe to receive real-time updates on global talent migration and compute availability.</p>
            <form class="flex flex-col md:flex-row gap-4" @submit.prevent>
              <input type="email" placeholder="NEURAL_IDENTIFIER@AI.IO" class="flex-1 bg-black border border-white/10 px-8 py-5 font-mono text-xs focus:border-cyan-500 outline-none uppercase tracking-widest">
              <button class="px-12 py-5 bg-cyan-500 text-black font-black uppercase text-[10px] tracking-[0.4em] hover:bg-white transition-all">Subscribe</button>
            </form>
          </div>
        </div>
      </section>

      <footer id="contact" class="min-h-screen grid lg:grid-cols-2 border-t border-white/5 scroll-mt-24">
        <div class="p-12 md:p-32 border-r border-white/5 flex flex-col justify-center bg-black/40 relative overflow-hidden">
          <canvas ref="contactCanvasRef" class="absolute inset-0 w-full h-full"></canvas>
          <div class="relative z-10">
            <h2 class="text-[12rem] font-black tracking-tighter text-outline text-transparent opacity-10 italic mb-0">UPLINK</h2>
            <div class="space-y-12">
              <div class="font-mono text-xs text-gray-500 space-y-2 uppercase tracking-widest">
                <p>Location: Zurich_Hub_01</p>
                <p>Security: AES_512_Quantum_Safe</p>
                <p>Status: Synchronized</p>
              </div>
            </div>
          </div>
        </div>
        <div id="signup" class="p-12 md:p-32 flex flex-col justify-center bg-cyan-500/[0.01]">
          <form class="space-y-16 max-w-sm" @submit.prevent>
            <div class="relative group">
              <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">01 // Identify_Self</label>
              <input type="text" placeholder="ENTITY_NAME" class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none">
            </div>
            <div class="relative group">
              <label class="text-[9px] text-cyan-500 font-mono uppercase tracking-[0.4em] mb-4 block italic">02 // Digital_Address</label>
              <input type="email" placeholder="CONTACT@DOMAIN.XYZ" class="w-full bg-transparent border-b border-white/10 pb-4 text-xs font-mono focus:border-cyan-500 transition-colors outline-none">
            </div>
            <button class="w-full py-8 bg-cyan-500 text-black font-black uppercase text-[12px] tracking-[0.5em] hover:bg-white transition-all shadow-[0_20px_40px_rgba(0,243,255,0.1)]">Establish Neural Link</button>
          </form>
        </div>
      </footer>
    </div>

    <div class="fixed bottom-0 w-full bg-black/90 backdrop-blur-2xl border-t border-white/10 py-4 z-50">
      <div class="flex items-center gap-12 whitespace-nowrap animate-ticker">
        <div v-for="(news, n) in newsFeed" :key="n" class="flex items-center gap-4">
          <span class="w-2 h-2 bg-cyan-500 rounded-full animate-pulse"></span>
          <span class="font-mono text-[10px] tracking-[0.3em] text-gray-400 uppercase italic font-bold">{{ news }}</span>
          <span class="text-cyan-900 px-12">>>></span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import * as THREE from 'three'
import { onMounted, ref, onUnmounted } from 'vue'

const canvasRef = ref(null)
const glowRef = ref(null)
const activeNode = ref(null)
const contactCanvasRef = ref(null)

const stats = [
  { label: "Active_Clusters", value: "9,102" },
  { label: "Sync_Latency", value: "0.02ms" },
  { label: "Throughput", value: "14.2TB/s" },
  { label: "Uptime", value: "100%" }
]

const workflowSteps = [
  { title: "Discovery", desc: "Autonomous agents crawl decentralized protocols to source peak-efficiency talent nodes." },
  { title: "Verification", desc: "Every node undergoes a zero-knowledge technical validation process via distributed consensus." },
  { title: "Sync", desc: "Instant uplink established via encrypted neural tunnels for project initialization." }
]

const plans = [
  { name: "Personal", price: "$0", features: ["1 Cluster Access", "Public Matching", "API Tier 1"] },
  { name: "Studio", price: "$499", features: ["10 Cluster Sync", "Verified Talent", "Priority Pipeline"] },
  { name: "Enterprise", price: "$1999", features: ["Unlimited Uplinks", "Dedicated Compute", "SLA Support"] }
]

const newsFeed = [
  "Compute Migration: Zurich -> Singapore", "ZK-Proof Verification System Updated", 
  "Neural Match v5.0 Active", "Node Connectivity +22% in LATAM"
]

let scene, camera, renderer, globeGroup, pins = [], raycaster, pointer = new THREE.Vector2()
let targetScrollY = 0, currentScrollY = 0
let isZoomed = false
const originalCamZ = 6

// IMPROVED RAYCASTING
const handleCanvasClick = (event) => {
  // Use correct client coordinate offset for responsive canvas
  const rect = canvasRef.value.getBoundingClientRect()
  pointer.x = ((event.clientX - rect.left) / rect.width) * 2 - 1
  pointer.y = -((event.clientY - rect.top) / rect.height) * 2 + 1
  
  raycaster.setFromCamera(pointer, camera)
  // Check for pins only
  const intersects = raycaster.intersectObjects(pins, false)
  
  if (intersects.length > 0) {
    const target = intersects[0].object
    activeNode.value = {
      id: Math.random().toString(36).substr(2, 6).toUpperCase(),
      latency: (Math.random() * 5 + 0.1).toFixed(2),
      position: target.position.clone()
    }
    isZoomed = true
  }
}

const closeNode = () => {
  activeNode.value = null
  isZoomed = false
}

onMounted(() => {
  scene = new THREE.Scene()
  camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000)
  camera.position.z = originalCamZ

  renderer = new THREE.WebGLRenderer({ canvas: canvasRef.value, antialias: true, alpha: true })
  renderer.setSize(window.innerWidth, window.innerHeight)
  renderer.setPixelRatio(Math.min(window.devicePixelRatio, 2))
  raycaster = new THREE.Raycaster()
  // Make raycaster more forgiving for small pins
  raycaster.params.Points.threshold = 0.1

  globeGroup = new THREE.Group()
  scene.add(globeGroup)

  const textureLoader = new THREE.TextureLoader()
  
  // 1. EARTH
  const earth = new THREE.Mesh(
    new THREE.SphereGeometry(2, 64, 64),
    new THREE.MeshPhongMaterial({
      map: textureLoader.load('https://threejs.org/examples/textures/planets/earth_atmos_2048.jpg'),
      emissiveMap: textureLoader.load('https://raw.githubusercontent.com/mrdoob/three.js/master/examples/textures/planets/earth_lights_2048.png'),
      emissive: new THREE.Color(0x00f3ff), emissiveIntensity: 0.9,
      transparent: true, opacity: 0.9
    })
  )
  globeGroup.add(earth)

  // 2. ATMOSPHERIC SHIELD (Wired sphere)
  const shield = new THREE.Mesh(
    new THREE.IcosahedronGeometry(2.2, 3),
    new THREE.MeshBasicMaterial({ color: 0x00f3ff, wireframe: true, transparent: true, opacity: 0.05 })
  )
  globeGroup.add(shield)

  // 3. ORBIT RING
  const ringGeom = new THREE.TorusGeometry(3.5, 0.005, 16, 100)
  const ringMat = new THREE.MeshBasicMaterial({ color: 0x00f3ff, transparent: true, opacity: 0.2 })
  const ring = new THREE.Mesh(ringGeom, ringMat)
  ring.rotation.x = Math.PI / 2.2
  globeGroup.add(ring)

  // 4. CLICKABLE PINS (Improved scale)
  const pinGeom = new THREE.SphereGeometry(0.06, 16, 16)
  const pinMat = new THREE.MeshBasicMaterial({ color: 0x00f3ff })
  for(let i=0; i<40; i++){
    const pin = new THREE.Mesh(pinGeom, pinMat.clone())
    const phi = Math.acos(-1 + (2 * i) / 40)
    const theta = Math.sqrt(40 * Math.PI) * phi
    pin.position.setFromSphericalCoords(2.05, phi, theta)
    // Adding a hidden larger hitbox for better interaction
    const hitbox = new THREE.Mesh(
      new THREE.SphereGeometry(0.15, 8, 8),
      new THREE.MeshBasicMaterial({ visible: false })
    )
    pin.add(hitbox)
    globeGroup.add(pin)
    pins.push(pin)
  }

  // 5. BACKGROUND STARS
  const starGeom = new THREE.BufferGeometry()
  const starsArr = new Float32Array(12000)
  for(let i=0; i<12000; i++) starsArr[i] = (Math.random() - 0.5) * 50
  starGeom.setAttribute('position', new THREE.BufferAttribute(starsArr, 3))
  scene.add(new THREE.Points(starGeom, new THREE.PointsMaterial({ color: 0x00f3ff, size: 0.015, transparent: true, opacity: 0.6 })))

  // LIGHTING
  scene.add(new THREE.AmbientLight(0xffffff, 0.4))
  const spotlight = new THREE.SpotLight(0x00f3ff, 2)
  spotlight.position.set(10, 10, 10)
  scene.add(spotlight)

  const animate = () => {
    currentScrollY += (targetScrollY - currentScrollY) * 0.08
    
    if (isZoomed && activeNode.value) {
      // Smooth Zoom
      const targetPos = activeNode.value.position.clone().applyMatrix4(globeGroup.matrixWorld)
      camera.position.lerp(targetPos.multiplyScalar(1.4), 0.08)
      camera.lookAt(targetPos)
    } else {
      // Normal View
      camera.position.lerp(new THREE.Vector3(0, 0, originalCamZ), 0.05)
      camera.lookAt(0, 0, 0)
      globeGroup.rotation.y += 0.001
      globeGroup.rotation.x = currentScrollY * 0.0005
    }

    globeGroup.position.y = -currentScrollY * 0.001
    shield.rotation.y -= 0.002
    ring.rotation.z += 0.005

    // Pulse Pins
    pins.forEach((p, i) => {
      const s = 1 + Math.sin(Date.now() * 0.004 + i) * 0.4
      p.scale.set(s, s, s)
      p.material.opacity = 0.5 + Math.sin(Date.now() * 0.004 + i) * 0.5
    })

    renderer.render(scene, camera)
    requestAnimationFrame(animate)
  }
  
  animate()
  window.addEventListener('resize', () => {
    camera.aspect = window.innerWidth / window.innerHeight
    camera.updateProjectionMatrix()
    renderer.setSize(window.innerWidth, window.innerHeight)
  })
  window.addEventListener('scroll', () => targetScrollY = window.scrollY)
  window.addEventListener('mousemove', (e) => {
    if(glowRef.value){
      glowRef.value.style.left = `${e.clientX}px`
      glowRef.value.style.top = `${e.clientY}px`
    }
  })

  // CONTACT SECTION THREE.JS SCENE
  if (contactCanvasRef.value) {
    const contactScene = new THREE.Scene()
    const contactCamera = new THREE.PerspectiveCamera(75, 1, 0.1, 1000)
    contactCamera.position.z = 3
    
    const contactRenderer = new THREE.WebGLRenderer({ 
      canvas: contactCanvasRef.value, 
      antialias: true, 
      alpha: true 
    })
    
    const updateContactSize = () => {
      const container = contactCanvasRef.value.parentElement
      if (container) {
        const width = container.clientWidth
        const height = container.clientHeight
        contactRenderer.setSize(width, height)
        contactCamera.aspect = width / height
        contactCamera.updateProjectionMatrix()
      }
    }
    updateContactSize()
    
    // Wireframe Torus
    const torusGeometry = new THREE.TorusGeometry(1, 0.4, 16, 100)
    const torusMaterial = new THREE.MeshBasicMaterial({ 
      color: 0x00f3ff, 
      wireframe: true,
      transparent: true,
      opacity: 0.3
    })
    const torus = new THREE.Mesh(torusGeometry, torusMaterial)
    contactScene.add(torus)
    
    // Inner rotating ring
    const innerRing = new THREE.Mesh(
      new THREE.TorusGeometry(0.7, 0.05, 16, 100),
      new THREE.MeshBasicMaterial({ color: 0x00f3ff, transparent: true, opacity: 0.6 })
    )
    contactScene.add(innerRing)
    
    // Orbiting particles
    const particleCount = 200
    const particlesGeometry = new THREE.BufferGeometry()
    const particlePositions = new Float32Array(particleCount * 3)
    
    for (let i = 0; i < particleCount * 3; i++) {
      particlePositions[i] = (Math.random() - 0.5) * 5
    }
    
    particlesGeometry.setAttribute('position', new THREE.BufferAttribute(particlePositions, 3))
    const particlesMaterial = new THREE.PointsMaterial({ 
      color: 0x00f3ff, 
      size: 0.02,
      transparent: true,
      opacity: 0.8
    })
    const particleSystem = new THREE.Points(particlesGeometry, particlesMaterial)
    contactScene.add(particleSystem)
    
    // Add point lights
    const pointLight1 = new THREE.PointLight(0x00f3ff, 1, 10)
    pointLight1.position.set(2, 2, 2)
    contactScene.add(pointLight1)
    
    const pointLight2 = new THREE.PointLight(0x0088ff, 0.5, 10)
    pointLight2.position.set(-2, -2, -2)
    contactScene.add(pointLight2)
    
    // Animation loop for contact scene
    const animateContact = () => {
      const time = Date.now() * 0.001
      
      torus.rotation.x = time * 0.3
      torus.rotation.y = time * 0.2
      
      innerRing.rotation.x = -time * 0.5
      innerRing.rotation.z = time * 0.3
      
      particleSystem.rotation.y = time * 0.1
      
      // Animate point lights
      pointLight1.position.x = Math.sin(time) * 2
      pointLight1.position.z = Math.cos(time) * 2
      
      contactRenderer.render(contactScene, contactCamera)
      requestAnimationFrame(animateContact)
    }
    animateContact()
    
    window.addEventListener('resize', updateContactSize)
  }
})
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Space+Grotesk:wght@300;400;700;900&display=swap');
body { font-family: 'Space Grotesk', sans-serif; background: #020205; cursor: crosshair; }
.font-outline { -webkit-text-stroke: 1px rgba(34, 211, 238, 0.5); color: transparent; }
@keyframes ticker { 0% { transform: translateX(0); } 100% { transform: translateX(-50%); } }
.animate-ticker { animation: ticker 30s linear infinite; }
.scale-fade-enter-active, .scale-fade-leave-active { transition: all 0.5s cubic-bezier(0.19, 1, 0.22, 1); }
.scale-fade-enter-from, .scale-fade-leave-to { opacity: 0; transform: scale(0.8) translateY(20px); }
::-webkit-scrollbar { width: 4px; }
::-webkit-scrollbar-thumb { background: #06b6d4; }
html { scroll-behavior: smooth; }
</style>