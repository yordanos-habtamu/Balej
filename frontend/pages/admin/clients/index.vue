<template>
  <div class="min-h-screen w-full bg-gray-100">
    <div class="container mx-auto px-4 py-6">
      <!-- Search Bar -->
      <div class="mb-6">
        <input 
          type="text" 
          v-model="searchTerm"
          placeholder="Search for something..." 
          class="w-full max-w-full sm:max-w-2xl mx-auto block border-2 border-black rounded-lg p-3 text-lg focus:outline-none focus:ring-2 focus:ring-black"
        />
      </div>

      <!-- Main Layout -->
      <div class="flex flex-col lg:flex-row gap-6 lg:h-[calc(100vh-7rem)]">
        <!-- Sidebar -->
        <div>
          <admin-sidebar></admin-sidebar>
        </div>
                      
        <!-- Main Content (Scrollable and Scrollbar Hidden) -->
        <div class="flex-1 overflow-y-auto pr-0 lg:pr-2 h-full hide-scrollbar">
               <!-- Modal -->
    <div v-if="isOpen" class="modal" @click.self="isOpen = false">
      <div class="modal-content">
        <span class="close-btn" @click="isOpen = false">&times;</span>
        <h2>You have successfuly added the user</h2>
        <p>to the archive</p>
      </div>
    </div>
          <div v-if="clientData.loading" class="flex items-center justify-center h-64">
            <p class="text-lg text-black">Loading...</p>
          </div>

          <div v-if="filteredUsers.length">
            <div
              v-for="user in filteredUsers"
              :key="user.email"
              class="w-full max-w-[800px] shadow-xl shadow-black/20 rounded-lg overflow-hidden flex flex-col justify-between bg-white mb-4"
            >
              <!-- User Info -->
              <div class="flex justify-around items-center w-full p-4">
                <div class="w-12 h-12 rounded-full overflow-hidden">
                  <img src="@/assets/painter.png" alt="Profile Image" class="object-cover w-full h-full" />
                </div>
                <div class="flex flex-col items-start">
                  <p class="font-bold">First Name</p>
                  <p class="text-sm">{{ user.first_name }}</p>
                </div>
                <div class="flex flex-col items-start">
                  <p class="font-bold">Last Name</p>
                  <p class="text-sm">{{ user.last_name }}</p>
                </div>
                <div class="flex flex-col items-start">
                  <p class="font-bold">Email</p>
                  <p class="text-sm">{{ user.email }}</p>
                </div>
              </div>

              <!-- Details -->
              <div class="px-4 pb-4">
                <p class="font-bold">Details</p>
                <p class="text-sm text-gray-600">
                  Lorem ipsum dolor sit amet consectetur adipisicing elit. Reiciendis sit quos aspernatur velit, illo tenetur.
                </p>
              </div>

              <!-- Action Buttons -->
              <div class="flex justify-between items-center p-4 mt-auto card-actions">
              <nuxt-link :to="`/admin/talents/${user.first_name}`">
                 <button
                  class="py-1 px-5 text-black bg-white border-2 border-black rounded-lg hover:scale-105 transition-all delay-75"
                >
                  Read More
                </button>
              </nuxt-link>
                <button
                  class="py-1 px-5 bg-black text-white rounded-lg hover:bg-black/25 transition-all delay-75"
                  @click="isOpen = true"
                >
                  Add to Catalogue
                </button>

              </div>
            </div>
          </div>
        </div>
           <aside 
          class="lg:block sticky top-6 self-start w-full lg:w-80 h-fit bg-white border-2 border-black rounded-lg p-6 overflow-y-auto max-h-[calc(100vh-7rem)]"
        >
          <div class="flex flex-col gap-4">
            <!-- Job Type Filters -->
            <div class=" pt-4">
              <h3 class="font-semibold mb-2">Job Type</h3>
              <div class="grid grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="part-time" class="h-4 w-4" />
                  <span>Part-time</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="contract" class="h-4 w-4" />
                  <span>Contract</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="full-time" class="h-4 w-4" />
                  <span>Full-time</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="internship" class="h-4 w-4" />
                  <span>Internship</span>
                </label>
              </div>
            </div>

            <!-- Category Filters -->
            <div class="border-t-2 border-black pt-4">
              <h3 class="font-semibold mb-2">Category</h3>
              <div class="grid grid-cols-2 gap-2">
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="construction" class="h-4 w-4" />
                  <span>Construction</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="house_keeping" class="h-4 w-4" />
                  <span>House Keeping</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="repairing" class="h-4 w-4" />
                  <span>Repairing</span>
                </label>
                <label class="flex items-center gap-2">
                  <input type="checkbox" value="management" class="h-4 w-4" />
                  <span>Management</span>
                </label>
              </div>
            </div>

            <!-- Price Range -->
            <div class="border-t-2 border-black pt-4">
              <h3 class="font-semibold mb-2">Price Range</h3>
              <div class="flex gap-2">
                <input type="number" placeholder="Min" class="w-1/2 border-2 border-black rounded-lg p-2" />
                <input type="number" placeholder="Max" class="w-1/2 border-2 border-black rounded-lg p-2" />
              </div>
            </div>
          </div>
        </aside>
      </div>
    </div>
  </div>
</template>

<style scoped>
/* ✅ Always hide scrollbars */
.hide-scrollbar {
  scrollbar-width: none; /* Firefox */
  -ms-overflow-style: none; /* IE 10+ */
}
.hide-scrollbar::-webkit-scrollbar {
  display: none; /* Chrome, Safari, Opera */
}

/* ✅ Mobile responsive tweaks */
@media (max-width: 1024px) {
  aside {
    position: static;
    width: 100%;
    max-height: none !important;
    margin-bottom: 1rem;
  }

  input[type="text"],
  input[type="number"] {
    padding: 0.5rem !important;
    font-size: 1rem !important;
  }

  button {
    padding: 0.75rem 1.25rem !important;
    font-size: 1rem !important;
  }

  .container {
    padding-left: 1rem;
    padding-right: 1rem;
  }
}

/* Stack buttons vertically on very small screens */
@media (max-width: 640px) {
  .card-actions {
    flex-direction: column;
    gap: 0.5rem;
  }
  .card-actions button {
    width: 100%;
  }
  /* Make profile images smaller */
  img {
    max-width: 48px;
    max-height: 48px;
  }
}


.modal {
  position: fixed;
  z-index: 1000;
  inset: 0;
  background-color: rgba(0,0,0,0.6);
  display: flex;
  justify-content: center;
  align-items: center;
}

.modal-content {
  background: #fff;
  padding: 20px;
  border-radius: 10px;
  width: 400px;
  max-width: 90%;
  text-align: center;
  box-shadow: 0 5px 15px rgba(0,0,0,0.3);
  animation: fadeIn 0.3s ease-in-out;
}

@keyframes fadeIn {
  from { opacity: 0; transform: scale(0.9); }
  to { opacity: 1; transform: scale(1); }
}

.close-btn {
  float: right;
  font-size: 20px;
  font-weight: bold;
  cursor: pointer;
  color: #333;
}
.close-btn:hover {
  color: red;
}
</style>


<script lang="ts">
import { defineComponent,onMounted, reactive, ref } from 'vue'
import { gql } from 'graphql-tag'
import { CombinedError } from '@urql/core'
import { defaultTypeResolver } from 'graphql'


const USER_QUERY = gql`
  query {
    users {
      first_name
      last_name
      email
      phone_number
      address
      role
    }
  }
`


export default defineComponent({
  setup(){
    const searchTerm = ref('')
    const isOpen = ref(false)
    const client= useUrqlClient();
 
const clientData = reactive({
  users: [] as Array<any>,
  loading: false,
  error: null as CombinedError | null
})
  

   const filteredUsers = computed(() => {
      if (!searchTerm.value) return clientData.users;
      const lowerSearch = searchTerm.value.toLowerCase();
      return clientData.users.filter(user =>
        user.first_name?.toLowerCase().includes(lowerSearch) ||
        user.email?.toLowerCase().includes(lowerSearch)
      );
    });
// Fetch users from backend
async function loadUsers() {
  clientData.loading = true
  const response = await client.query(USER_QUERY, {}).toPromise()

  if (response.error) {
    clientData.error = response.error
  } else {
    clientData.users = response.data.users
  }

  clientData.loading = false
}

// Load users on component mount
onMounted(() => {
  loadUsers()
})
 
return {
    clientData,
    searchTerm,
    filteredUsers,
    isOpen
}
  },
});
// Reactive object to store users and state


</script>
