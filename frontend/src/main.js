import Vue from 'vue'
import VueRouter from 'vue-router'
import App from './App.vue'
import Search from './components/search'
import Watch from './components/torrents'
import MassRename from './components/mass-rename'
import moment from 'moment'
// import 'spectre.css/dist/spectre.min.css'
// import 'spectre.css/dist/spectre-exp.min.css'
// import 'spectre.css/dist/spectre-icons.min.css'

moment.locale('ru')

Vue.config.productionTip = false

Vue.use(VueRouter)

const router = new VueRouter({
  routes: [
    { path: '/', redirect: '/search' },
    { path: '/search', component: Search },
    { path: '/watch', component: Watch },
    { path: '/mass-rename', component: MassRename }
  ]
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')

window.api = require('./js/api').default
