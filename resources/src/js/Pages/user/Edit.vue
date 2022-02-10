<template>
<div>
  <form :action="`/manage/users/${user.id}`" method="POST" class="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4">

    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="name">
        Username
      </label>
      <input
        class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
        name="username" type="text" :value="user.name" placehuser="Username">
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="room">
        Password
      </label>
      <input
        class="shadow appearance-none border border-red-500 rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline"
        name="password" type="text" :value="user.room">
    </div>
    <div class="flex items-center justify-between">
      <button
        class="hover:bg-blue-700 font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        type="submit">
        送出
      </button>
    </div>
  </form>
</div>
</template>

<script>
import Layout from "@/js/Layout/Master";

export default {
  layout: Layout,
  props: {
    user: Object,
  },
};
</script>