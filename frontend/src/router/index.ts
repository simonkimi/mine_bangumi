import { createRouter, createWebHistory } from "vue-router";
import HomeView from "@/views/HomeView.vue";
import { useSystemStore } from "@/stores/systemStore";
import RegisterView from "@/views/RegisterView.vue";
import { useUserStore } from "@/stores/userStore";

export const HomeRoute = "Home";
export const RegisterRoute = "Register";
export const SettingsRoute = "Settings";
export const UserSettingsRoute = "UserSettings";
export const RssSettingsRoute = "RssSettings";
export const DownloaderSettingsRoute = "DownloaderSettings";
export const LoginRoute = "Login";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: "/",
      name: HomeRoute,
      component: HomeView,
    },
    {
      path: "/register",
      name: RegisterRoute,
      component: RegisterView,
      meta: {
        showAppBar: false,
      },
    },
    {
      path: "/login",
      name: LoginRoute,
      component: () => import("@/views/LoginView.vue"),
      meta: {
        showAppBar: false,
      },
    },
    {
      path: "/settings",
      name: SettingsRoute,
      component: () => import("@/views/SettingsView.vue"),
      children: [
        {
          path: "",
          redirect: { name: UserSettingsRoute },
        },
        {
          path: "user",
          name: UserSettingsRoute,
          component: () => import("@/components/UserSettings.vue"),
        },
        {
          path: "rss",
          name: RssSettingsRoute,
          component: () => import("@/components/RssSettings.vue"),
        },
        {
          path: "downloader",
          name: DownloaderSettingsRoute,
          component: () => import("@/components/DownloaderSettings.vue"),
        },
      ],
    },
  ],
});

router.beforeEach(async (to, from, next) => {
  const systemStore = useSystemStore();
  const userStore = useUserStore();
  if (!systemStore.isStoreInit) {
    await systemStore.loadSystemData();
    await userStore.setUsername(systemStore.username);
  }

  if (!systemStore.isSystemInit && to.name !== RegisterRoute) {
    console.log("未初始化用户信息，跳转到引导页");
    next({ name: RegisterRoute });
    return;
  }

  if (!systemStore.isLogin && to.name !== LoginRoute) {
    console.log("未登录，跳转到登录页");
    next({ name: LoginRoute, query: { redirect: to.fullPath } });
    return;
  }

  next();
});

export default router;
