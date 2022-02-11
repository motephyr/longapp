<template>
  <main>
    <div class="grid-container w-full">
      管理頁面
      <div class="grid-item">
        <table class="border-separate border-green-800 w-full sm:text-center">
          <thead>
            <tr style="background-color: #1c4b73; color: white">
              <th class="border-solid">日期</th>
              <th class="border-solid">半夜活動事件數量</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(datestring, index) in Object.keys(datestrings).reverse()"
              :style="{
                backgroundColor: index % 2 == 0 ? '#f89b3e' : '#fdf1e0',
              }"
              :key="index"
            >
              <td class="border-solid">
                <a :href="`/manage/users/${user_id}/groups/${datestring}`">{{ datestring }}</a>
              </td>
              <td class="border-solid">{{ datestrings[datestring] }}</td>
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
    datestrings: Object,
    user_id: String
  },
};
</script>
