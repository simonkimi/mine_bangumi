<template>
  <n-config-provider :theme="darkTheme" class="h-full">
    <n-dialog-provider>
      <n-message-provider>
        <n-layout class="h-full">
          <n-layout-header v-if="showAppBar" class="fixed top-0 left-0 z-50">
            <AppBar />
          </n-layout-header>
          <n-layout-content class="h-full" native-scrollbar>
            <router-view />
          </n-layout-content>
        </n-layout>
      </n-message-provider>
    </n-dialog-provider>
  </n-config-provider>
</template>

<script lang="ts" setup>
import { darkTheme } from "naive-ui";

const route = useRoute();
const showAppBar = computed(() => route.meta.showAppBar ?? true);
</script>
