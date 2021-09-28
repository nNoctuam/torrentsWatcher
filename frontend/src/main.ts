import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHistory } from "vue-router";
import Search from "./components/search.vue";
import Watch from "./components/torrents.vue";
import MassRename from "./components/mass-rename.vue";
import moment from "moment";
moment.locale("ru");

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", name: "root", redirect: "/search" },
    { path: "/search", component: Search },
    { path: "/watch", component: Watch },
    { path: "/mass-rename", component: MassRename },
  ],
});

createApp(App).use(router).mount("#app");

// window.api = require('./ts/api').default
