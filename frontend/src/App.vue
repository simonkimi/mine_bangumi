<template>
  <n-config-provider :theme="darkTheme">
    <n-dialog-provider>
      <n-layout class="min-h-svh">
        <n-layout-header class="fixed top-0 left-0 z-50">
          <AppBar />
        </n-layout-header>
        <n-layout-content native-scrollbar>
          <router-view />
        </n-layout-content>
      </n-layout>
    </n-dialog-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { darkTheme } from "naive-ui";
import AppBar from "@/components/AppBar.vue";
</script>
