import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import GuideView from "@/views/GuideView.vue";
import SettingView from "@/views/SettingView.vue";
import { useSystemStore } from "@/stores/system";

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
      meta: {
        showAppBar: false,
      },
    },
    {
      path: "/setting",
      name: "setting",
      component: SettingView,
    },
  ],
});

router.beforeEach((to, from, next) => {
  const systemStore = useSystemStore();
  if (!systemStore.isInitUser && to.name !== "guide") {
    console.log("未初始化用户信息，跳转到引导页");
    next({ name: "guide" });
  } else {
    next();
  }
});

export default router;
