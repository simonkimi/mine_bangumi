import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import SettingView from "@/views/SettingView.vue";
import { useSystemStore } from "@/stores/system";
import RegisterView from "@/views/RegisterView.vue";

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/register",
      name: "register",
      component: RegisterView,
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
  if (!systemStore.isSystemInit && to.name !== "register") {
    console.log("未初始化用户信息，跳转到引导页");
    next({ name: "register" });
  } else {
    next();
  }
});

export default router;
