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
              <th class="border-solid">姓名</th>
              <th class="border-solid">房間</th>
              <th class="border-solid"></th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="(older, index) in theOlders"
              :style="{
                backgroundColor: index % 2 == 0 ? '#f89b3e' : '#fdf1e0',
              }"
              :key="index"
            >
              <td class="border-solid">
                <a :href="`/olders/${older.id}`">
                  {{ older.name }}
                </a>
              </td>
              <td class="border-solid">
                {{ older.room }}
              </td>
              <td class="border-solid">
                <a :href="`/olders/${older.id}/edit`" class="older-meta-small">
                  編輯資料
                </a>
                <form @submit.prevent="form.post(`/olders/${older.id}/delete`)">
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
import modelQuery from "@/js/Plugin/modelQuery.js";
import axios from "axios";

export default {
  mixins: [modelQuery],
  layout: Layout,
  props: {
    olders: Array,
  },
  data() {
    return {
      theOlders: this.olders,
      form: this.$inertia.form(),
    };
  },
  queryModels: {
    model: "theOlders",
    watch: async function (newObj, oldObj) {
      let result = await axios.get(
        `/olders/query?query=${encodeURIComponent(JSON.stringify(newObj))}`
      );

      this.handleResults(result);
    },
    variables: function () {
      return JSON.parse(
        JSON.stringify({
          sort: { id: -1 },
        })
      );
    },
  },
  sockets: {
    connect: function () {
      this.$socket.emit("join", {
        roomName: `/user/${this.$page.props.user.id.toString()}/older`,
      });
    },
    insert: async function (id) {
      let result = await axios.get(
        `/olders/query/${id}?query=${encodeURIComponent(
          JSON.stringify(this.variables)
        )}`
      );
      await this.handleResult(result, id);
    },
    update: async function (id) {
      let result = await axios.get(
        `/olders/query/${id}?query=${encodeURIComponent(
          JSON.stringify(this.variables)
        )}`
      );

      await this.handleResult(result, id);
    },
    delete: function (id) {
      this.theOlders = this.theOlders.filter((x) => {
        return x.id !== id;
      });
    },
  },
  destroyed() {
    this.$socket.emit("unjoin", {
      roomName: `/user/${this.$page.props.user.id.toString()}/older`,
    });
  },
};
</script>