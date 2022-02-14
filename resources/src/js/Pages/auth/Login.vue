<template>
  <div class="flex items-center justify-center p-6 min-h-screen bg-indigo-800">
    <div class="w-full max-w-md">
      <Flash v-model="$page.props.flash.message" v-if="$page.props.flash && $page.props.flash.message" />
      <form
        class="mt-8 bg-white rounded-lg shadow-xl overflow-hidden"
        @submit.prevent="login"
      >
        <div class="px-10 py-12">
          <h1 class="text-center text-3xl font-bold">CarePLUS</h1>
          <div class="mt-6 mx-auto w-24 border-b-2" />
          <text-input
            v-model="form.username"
            :error="form.errors.username"
            class="mt-10"
            label="Username"
            type="text"
            autofocus
            autocapitalize="off"
          />
          <text-input
            v-model="form.password"
            :error="form.errors.password"
            class="mt-6"
            label="Password"
            type="password"
          />
        </div>
        <div class="flex px-10 py-4 bg-gray-100 border-t border-gray-100">
          <button :disabled="form.processing" class="flex items-center">
            <div v-if="form.processing" class="btn-spinner mr-2" />
            <div class="btn-indigo ml-auto" type="submit">Login</div>
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import TextInput from "@/js/Shared/TextInput";
import Flash from "@/js/Shared/Flash";

export default {
  components: {
    TextInput,
    Flash
  },
  data() {
    return {
      form: this.$inertia.form({
        username: "aaa",
        password: "aaa",
      }),
    };
  },
  methods: {
    login() {
      this.form.post("/auth/login");
    },
  },
};
</script>