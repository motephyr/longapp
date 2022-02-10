<template>
  <main>
    <div class="grid-container w-full">
      年長者頁面
      <a
        href="/olders/new"
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
        新增年長者
      </a>
      <div class="grid-item">
        <table class="border-separate border-green-800 w-full sm:text-center">
          <thead>
            <tr style="background-color: #1c4b73; color: white">
              <th class="border-green-600">姓名</th>
              <th class="border-green-600">房間</th>
              <th class="border-green-600"></th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(older, index) in olders"
              :style="{
                backgroundColor: index % 2 == 0 ? '#f89b3e' : '#fdf1e0',
              }"
              :key="index"
            >
              <td class="border-green-600">
                <a :href="`/olders/${older.id}`">
                  {{ older.name }}
                </a>
              </td>
              <td class="border-green-600">
                {{ older.room }}
              </td>
              <td class="border-green-600">
                <a :href="`/olders/${older.id}/edit`" class="older-meta-small">
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
    olders: Array,
  },
};
</script>