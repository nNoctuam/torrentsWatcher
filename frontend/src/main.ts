import { createApp } from "vue";
import App from "./App.vue";
import { createRouter, createWebHashHistory } from "vue-router";
import Search from "./components/search.vue";
import Watch from "./components/torrents.vue";
import MassRename from "./components/mass-rename.vue";
import moment from "moment";
import api from "@/ts/api";
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

(async () => {
  const config = await fetch("/config.json");
  const configData = await config.json();

  console.log("config:", configData);
  api.setRpcDomain(
    `${window.location.protocol}//${window.location.hostname}:${configData.PORT_GRPC}`
  );

  const app = createApp(App);
  app.use(router).mount("#app");
  app.config.performance = true;
})();
