import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import GuideView from "@/views/GuideView.vue";
import SettingView from "@/views/SettingView.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/guide",
      name: "guide",
      component: GuideView,
    },
    {
      path: "/setting",
      name: "setting",
      component: SettingView,
    },
  ],
});

export default router;
