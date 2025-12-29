<template>
  <div class="min-h-screen w-full bg-gray-50">
    <!-- Back button -->
    <div class="p-4">
      <nuxt-link to="/admin/talents">
        <ArrowLeftCircleIcon
          class="h-8 w-8 text-gray-700 hover:text-gray-900 transition-colors"
        />
      </nuxt-link>
    </div>

    <!-- Loading state -->
    <div v-if="usersData.loading" class="flex items-center justify-center h-[70vh]">
      <p class="text-gray-600 text-lg animate-pulse">Loading user data...</p>
    </div>

    <!-- Error state -->
    <div v-else-if="usersData.error" class="flex items-center justify-center h-[70vh]">
      <p class="text-red-500 font-medium">Error: {{ usersData.error.message }}</p>
    </div>

    <!-- User profile -->
    <div v-else-if="user" class="flex flex-col items-center px-4 md:px-12">
      <!-- Profile image -->
      <div class="w-40 h-40 md:w-60 md:h-60 rounded-full overflow-hidden shadow-md">
        <img
          src="~/assets/painter.png"
          alt="profile pic"
          class="w-full h-full object-cover"
        />
      </div>

      <!-- User info card -->
      <div
        class="bg-white rounded-2xl shadow-md w-full max-w-4xl mt-6 p-6 md:p-10 grid grid-cols-1 md:grid-cols-2 gap-6"
      >
        <!-- Left column -->
        <div class="space-y-4">
          <div>
            <h2 class="text-sm font-semibold text-gray-500">First Name</h2>
            <p class="text-lg font-medium text-gray-900">{{ user.first_name }}</p>
          </div>
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Last Name</h2>
            <p class="text-lg font-medium text-gray-900">{{ user.last_name }}</p>
          </div>
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Address</h2>
            <p class="text-lg font-medium text-gray-900">{{ user.address }}</p>
          </div>
        </div>

        <!-- Right column -->
        <div class="space-y-4">
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Email</h2>
            <p class="text-lg font-medium text-gray-900">{{ user.email }}</p>
          </div>
          <div>
            <h2 class="text-sm font-semibold text-gray-500">Phone Number</h2>
            <p class="text-lg font-medium text-gray-900">{{ user.phone_number }}</p>
          </div>
        </div>
      </div>

      <!-- Details Section -->
      <div class="bg-white rounded-2xl shadow-md w-full max-w-4xl mt-6 p-6 md:p-10">
        <h2 class="text-xl font-semibold mb-3 text-gray-800">Details</h2>
        <p class="text-gray-600 leading-relaxed">
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Reiciendis illo
          a nulla eligendi vel aperiam minus dolores, ratione id ex nemo itaque
          provident veniam in voluptas deleniti quam laudantium magni quasi dolore
          quos necessitatibus adipisci qui esse.
        </p>

        <!-- Actions -->
        <div class="grid grid-cols-2 sm:flex-row gap-4 mt-6">
         <nuxt-link to="/admin/chats">
             <button
            class="px-6 py-3 bg-blue-500 hover:bg-blue-600 text-white rounded-lg shadow transition"
          >
            Start Chat
          </button>
         </nuxt-link>
        <nuxt-link to="/admin">
            <button
            class=" px-6 py-3 bg-green-500 hover:bg-green-600 text-white rounded-lg shadow transition"
          >
            Call
          </button>
        </nuxt-link>
        
        </div>
      </div>
    </div>

    <!-- User not found -->
    <div v-else class="flex items-center justify-center h-[70vh]">
      <p class="text-gray-600 text-lg">User not found</p>
    </div>
  </div>
</template>

<script lang="ts">
import { ArrowLeftCircleIcon } from "@heroicons/vue/24/outline";
import { defineComponent, onMounted, reactive, ref, computed } from "vue";
import { gql } from "graphql-tag";
import { CombinedError } from "@urql/core";

const USER_QUERY = gql`
  query {
    users {
      first_name
      last_name
      email
      phone_number
      address
    }
  }
`;

export default defineComponent({
  setup() {
    const searchTerm = ref("");
    const client = useUrqlClient();
    const route = useRoute();

    const usersData = reactive({
      users: [] as Array<any>,
      loading: false,
      error: null as CombinedError | null,
    });

    const user = computed(() => {
      return (
        usersData.users.find(
          (item) => item.first_name === route.params.name
        ) || null
      );
    });

    async function loadUsers() {
      usersData.loading = true;
      const response = await client.query(USER_QUERY, {}).toPromise();

      if (response.error) {
        usersData.error = response.error;
      } else {
        usersData.users = response.data.users;
      }

      usersData.loading = false;
    }

    onMounted(() => {
      loadUsers();
    });

    return {
      usersData,
      user,
      searchTerm,
    };
  },
});
</script>
