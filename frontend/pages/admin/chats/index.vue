<template>
  <div class="h-screen flex p-8">
    <!-- Sidebar -->
    <aside class="m-2">
      <admin-sidebar />
    </aside>

    <!-- Main content -->
    <main class="flex-1 h-full bg-gray-200 flex p-6">
      <div v-if="chats.length" class="w-[100%]">
        <div v-for="(chat,index) in chats" :key="index">
         <div class="h-[70px]  bg-white rounded-lg p-6 m-2 flex items-center justify-start hover:border-2 hover:border-blue-300 transition-all duration-75 ease-in" @click="showMessenger(chat.user.first_name)">
        <div class="h-10 w-10 rounded-full align-left">
          <img
            :src="chat.user.profile"
            alt=""
            class="w-10 h-10 rounded-full object-cover"
          />
        </div>
        <div class="ml-5">
          <div class="flex flex-col ">
            <div class="flex">
              <h1 class="pr-2">{{ chat.user.first_name }}</h1>
              <h1>{{ chat.user.last_name }}</h1>
            </div>
            <div class="text-sm text-center w-[300px]">
              <h1 class="truncate">
               {{ chat.last_message }}
              </h1>
            </div>
          </div>
        </div>
      </div>
        </div>
       
      </div>
      <div v-else>
       <div class="flex items-center justify-center text-center">
        <h1>Start Messaging by selecting users</h1>
       </div>
      </div>
          <div v-show="show" class="h-full w-full">
            <div v-if="currentUser" class=" h-full w-full bg-white rounded-lg p-4 m-2 flex flex-col justify-between">
            <div class="flex items-center bg-gray-200 m-0 mb-4">
          <div class="h-10 w-10">
            <img
              src="~/assets/painter.png"
              alt=""
              class="w-10 h-10 rounded-full object-cover"
            />
          </div>
          <div class="flex pb-2 pl-2">
            <h1 class="pr-2">{{ currentUser.user.first_name }}</h1>
            <h2> {{ currentUser.user.last_name }}</h2>
          </div>
        </div>

        <!-- MESSAGE CONTAINER - ONLY ADDITION -->
        <div
          v-if="messages.length || currentUser.last_message"
          class="space-y-2 overflow-y-auto max-h-[calc(100vh-250px)]"
          style="word-break: break-word;"
        >
               <div
              class="bg-gray-200 p-4 rounded-2xl max-w-[70vw] inline-block break-words whitespace-normal"
              style="word-break: break-word;"
            >
              {{currentUser.last_message }}
            </div>
          <div v-for="(message, index) in messages" :key="index">
          
            <div
              class="bg-gray-200 p-4 rounded-2xl max-w-[70vw] inline-block break-words whitespace-normal"
              style="word-break: break-word;"
            >
              {{ message }}
            </div>
          </div>
        </div>
        <div v-else>
          <h1 class="text-lg text-center">Start messaging now</h1>
        </div>

        <div class="relative w-full mt-4">
          <input
            type="text"
            v-model="content"
            @keyup.enter="sendMessage"
            placeholder="Type your message here..."
            class="w-full h-10 rounded-lg p-2 pr-10 bg-gray-200 caret-blue-400 focus:outline-none focus:ring-2 focus:ring-blue-400"
          />
          <paper-airplane-icon
            @click="sendMessage"
            class="absolute right-3 top-1/2 -translate-y-1/2 h-6 w-6 text-white p-1 bg-blue-500 rounded-full cursor-pointer hover:bg-blue-600 transition"
          />
        </div>
      </div>
      
      </div>
        
   
  
    </main>
  </div>
</template>

<script setup lang="ts">
import { ref,reactive } from "vue";

const show=ref(false)
const chats = reactive(
  [
    {
     "user":{
      "first_name":"cabscas",
      "last_name":"aakfjkasjfsa",
      "profile":"/images/painter.png"
     },
     "last_message":"Lorem ipsum dolor sit amet consectetur adipisicing elit. Soluta autem quod suscipit illo possimus quaerat!",
     "messages":[]
    },
   {
     "user":{
      "first_name":"abc",
      "last_name":"aakfjkasjfsa",
      "profile":"/images/painter.png"
     },
     "last_message":"Lorem ipsum dolor sit amet consectetur adipisicing elit. Soluta autem quod suscipit illo possimus quaerat!",
     "messages":[]
    },
    {
     "user":{
      "first_name":"abcas",
      "last_name":"aakfjkasjfsa",
      "profile":"/images/painter.png"
     },
     "last_message":"Lorem ipsum dolor sit amet consectetur adipisicing elit. Soluta autem quod suscipit illo possimus quaerat!",
     "messages":[]
    },
  ]
)
const currentUser = ref()

const loadMessage =(first_name: string)=>{
 const user = chats.find(item => item.user.first_name === first_name)
 currentUser.value = user

   
}
const showMessenger = (first_name : string)=>{
   show.value ? show.value=false :show.value=true;
   console.log(show.value)
   loadMessage(first_name)
  
  }
const content = ref("");
const messages = ref<string[]>([]);

const sendMessage = () => {
  if (!content.value.trim()) return;
  messages.value.push(content.value.trim());

  content.value = "";

};
</script>
