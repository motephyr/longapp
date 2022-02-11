<template>
  <main>
    <div class="grid-container w-full">
      帳號頁面
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
              <th class="border-solid">Username</th>
              <th class="border-solid"></th>
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
              <td class="border-solid">
                  {{ user.username }}
              </td>
              <td class="border-solid">
                <a :href="`/manage/users/${user.id}/idstrings`">
                  idString管理
                </a>
                <a :href="`/manage/users/${user.id}/groups`">
                  影片管理
                </a>
                <a :href="`/manage/users/${user.id}/edit`" class="user-meta-small">
                  編輯資料
                </a>
                <form
                  @submit.prevent="formSubmit(user)"
                  class="inline-block"
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
    users: Array,
  },
  data() {
    return {
      form: this.$inertia.form(),
    };
  },
  methods:{
    formSubmit(user){
      if (confirm(`確定要刪除嗎？`) === true){
       this.form.post(`/manage/users/${user.id}/delete`)
      }
    }
  }
}
</script>