import { getSystemConfig } from "@/api/api";

export const useSystemStore = defineStore("global", () => {
  const displayGuide = ref(false);
  const version = ref("0.0.0");
  const isLogin = ref(false);

  async function loadSystemData() {
    const system = await getSystemConfig();
    version.value = system.data.version;
    isLogin.value = system.data.is_login;
    displayGuide.value = system.data.is_first_run;
  }

  return {
    displayGuide,
    version,
    isLogin,
    loadSystemData,
  };
});
