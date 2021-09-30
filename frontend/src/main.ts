import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory } from "vue-router";
import Search from "./components/pages/search.vue";
import Watch from "./components/pages/torrents.vue";
import MassRename from "./components/pages/mass-rename.vue";
import moment from "moment";
import { store, key } from "./store";
moment.locale("ru");

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: "/", name: "root", redirect: "/search" },
    { path: "/search", component: Search },
    { path: "/watch", component: Watch },
    { path: "/mass-rename", component: MassRename },
  ],
});

const app = createApp(App);
app.use(store, key);
app.use(router).mount("#app");
app.config.performance = true;
