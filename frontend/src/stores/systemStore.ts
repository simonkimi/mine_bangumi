import { useApi } from "@/api/api";

export const useSystemStore = defineStore("global", () => {
  const { getSystemStatus } = useApi();
  const isSystemInit = ref(true);
  const version = ref("0.0.0");
  const isLogin = ref(false);
  const username = ref("");
  const isStoreInit = ref(false);

  async function loadSystemData() {
    console.log("Load system data");
    const system = await getSystemStatus();
    version.value = system.version;
    isLogin.value = system.is_login;
    isSystemInit.value = system.is_system_init;
    username.value = system.username;
    isStoreInit.value = true;
    console.log("System data loaded");
  }

  async function setLogin(v: boolean) {
    isLogin.value = v;
  }

  return {
    isStoreInit,
    isSystemInit,
    version,
    isLogin,
    setLogin,
    username,
    loadSystemData,
  };
});
