import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Search from './components/search'
import Watch from './components/torrents'
import moment from 'moment'
// import 'spectre.css/dist/spectre.min.css'
// import 'spectre.css/dist/spectre-exp.min.css'
// import 'spectre.css/dist/spectre-icons.min.css'

moment.locale('ru')

Vue.config.productionTip = false

Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    { path: '/', component: Search },
    { path: '/search', component: Search },
    { path: '/watch', component: Watch }
  ]
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')

window.api = require('./js/api').default
