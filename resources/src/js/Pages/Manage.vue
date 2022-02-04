<template>
  <main>
    <div class="grid-container w-full">
      管理頁面
      <div class="grid-item">
        <table class="border-separate border-green-800 w-full sm:text-center">
          <thead>
            <tr style="background-color: #1c4b73; color: white">
              <th class="border-green-600">日期</th>
              <th class="border-green-600">半夜活動事件數量</th>
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
              <td class="border-green-600">
                <a :href="`/manage/${datestring}`">{{ datestring }}</a>
              </td>
              <td class="border-green-600">{{ datestrings[datestring] }}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </main>
</template>
<script>
import Layout from "../Layout/Master";
import helper from "../../js/helpers";

export default {
  layout: Layout,
  props: {
    datestrings: Object,
  },
  methods: {
    toHHMMSS(time) {
      return helper.toHHMMSS(time);
    },
  },
};
</script>
