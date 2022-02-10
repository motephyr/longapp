<template>
  <main>
    <div class="grid-container w-full">
      Idstring管理
      <a
        :href="`/manage/users/${user.id}/idstrings/new`"
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
        新增Idstring
      </a>
      <div class="grid-item">
        <table class="border-separate border-green-800 w-full sm:text-center">
          <thead>
            <tr style="background-color: #1c4b73; color: white">
              <th class="border-green-600">Idstring</th>
              <th class="border-green-600"></th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(idstring, index) in userIdstrings"
              :style="{
                backgroundColor: index % 2 == 0 ? '#f89b3e' : '#fdf1e0',
              }"
              :key="index"
            >
              <td class="border-green-600">
                {{ idstring.idstring }}
              </td>
              <td class="border-green-600">
                <form
                  @submit.prevent="form.post(`/manage/users/${user.id}/idstrings/${idstring.id}/delete`)"
                >
                  <button type="submit" :disabled="form.processing">
                    刪除資料
                  </button>
                </form>
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
    user: Object,
    userIdstrings: Array,
  },
  data() {
    return {
      form: this.$inertia.form(),
    };
  }
};
</script>