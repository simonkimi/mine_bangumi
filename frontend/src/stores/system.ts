import { getSystemConfig } from "@/api/api";

export const useSystemStore = defineStore("global", () => {
  const isInitUser = ref(true);
  const version = ref("0.0.0");
  const isLogin = ref(false);

  async function loadSystemData() {
    const system = await getSystemConfig();
    version.value = system.version;
    isLogin.value = system.is_login;
    isInitUser.value = system.is_init_user;
  }

  return {
    isInitUser,
    version,
    isLogin,
    loadSystemData,
  };
});
