import Vue from 'vue'

import { createInertiaApp } from '@inertiajs/inertia-vue'
import { InertiaProgress } from '@inertiajs/progress'
import helper from "@/js/helpers";

Vue.prototype.$helper = helper


createInertiaApp({
  resolve: name => require(`../js/Pages/${name}`),
  setup({ el, App, props, plugin }) {
    Vue.use(plugin)

    new Vue({
      render: h => h(App, props),
    }).$mount(el)
  },
})

InertiaProgress.init()
