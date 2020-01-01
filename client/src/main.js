import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router'
import router from './routes/Router'
import Vuex from 'vuex'
import Axios from 'axios'
import store from './helpers/store'
import BootstrapVue from 'bootstrap-vue'

import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'

Vue.use(VueRouter)
Vue.use(Vuex)
Vue.use(BootstrapVue)

Vue.prototype.$http = Axios;

const token = localStorage.getItem('token')
if (token) {
  Vue.prototype.$http.defaults.headers.common['Authorization'] = token
}

Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App),
}).$mount('#app')
