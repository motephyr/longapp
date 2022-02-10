<template>
  <main>
    <div class="grid-container w-full">
      管理頁面
      <a
        href="/manage/users/new"
        class="
          h-12
          px-6
          m-2
          text-lg text-indigo-100
          transition-colors
          duration-150
          bg-indigo-700
          rounded-lg
          focus:shadow-outline
          hover:bg-indigo-800
          bg-blue-800
        "
      >
        新增帳號
      </a>
      <div class="grid-item">
        <table class="border-separate border-green-800 w-full sm:text-center">
          <thead>
            <tr style="background-color: #1c4b73; color: white">
              <th class="border-green-600">Username</th>
              <th class="border-green-600"></th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(user, index) in users"
              :style="{
                backgroundColor: index % 2 == 0 ? '#f89b3e' : '#fdf1e0',
              }"
              :key="index"
            >
              <td class="border-green-600">
                <a :href="`/manage/users/${user.id}`">
                  {{ user.username }}
                </a>
              </td>
              <td class="border-green-600">
                <a :href="`/manage/users/${user.id}/edit`" class="user-meta-small">
                  編輯資料
                </a>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>

<script>
import Layout from "@/js/Layout/Master";

export default {
  layout: Layout,
  props: {
    users: Array,
  },
};
</script>