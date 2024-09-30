export const useUserStore = defineStore(
  "user",
  () => {
    const username = ref("");
    const apiToken = ref("");

    async function setUsername(v: string) {
      username.value = v;
    }

    async function setApiToken(v: string) {
      apiToken.value = v;
    }

    return {
      username,
      setUsername,
      apiToken,
      setApiToken,
    };
  },
  {
    persist: true,
  }
);
