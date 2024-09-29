<template>
  <n-config-provider :theme="darkTheme">
    <n-dialog-provider>
      <n-layout class="min-h-svh">
        <n-layout-header class="fixed top-0 left-0 z-50" v-if="showAppBar">
          <AppBar />
        </n-layout-header>
        <n-layout-content native-scrollbar>
          <router-view />
        </n-layout-content>
      </n-layout>
    </n-dialog-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import { darkTheme } from "naive-ui";
import { useSystemStore } from "@/stores/system";

const systemStore = useSystemStore();
const routerStore = useRouter();

const route = useRoute();
const showAppBar = computed(() => route.meta.showAppBar ?? true);

onMounted(async () => {
  await systemStore.loadSystemData();
  if (!systemStore.isSystemInit) {
    console.info("未初始化用户, 跳转到引导页");
    await routerStore.push("/guide");
  }
});
</script>
