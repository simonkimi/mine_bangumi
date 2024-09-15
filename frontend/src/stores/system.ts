export const useSystemStore = defineStore("global", () => {
  const forceGuide = ref(false);
  const version = ref("0.0.0");

  async function loadSystemData() {}

  return {
    forceGuide,
    version,
    loadSystemData,
  };
});
