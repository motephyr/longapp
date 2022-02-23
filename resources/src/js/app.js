import Vue from 'vue'

import { createInertiaApp } from '@inertiajs/inertia-vue'
import { InertiaProgress } from '@inertiajs/progress'
import helper from "@/js/helpers";

Vue.prototype.$helper = helper

import VueSocketIO from 'vue-socket.io'
import SocketIO from 'socket.io-client'

Vue.use(new VueSocketIO({
  debug: true,
  connection: SocketIO('http://localhost:15001'),
  vuex: {} //Optional options
}))


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
