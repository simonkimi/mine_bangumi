import { getSystemStatus } from "@/api/api";

export const useSystemStore = defineStore("global", () => {
  const isSystemInit = ref(true);
  const version = ref("0.0.0");
  const isLogin = ref(false);

  async function loadSystemData() {
    const system = await getSystemStatus();
    version.value = system.version;
    isLogin.value = system.is_login;
    isSystemInit.value = system.is_system_init;
  }

  return {
    isSystemInit,
    version,
    isLogin,
    loadSystemData,
  };
});
