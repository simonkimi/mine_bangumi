<template>
  <div class="flex flex-row h-14">
    <div class="flex flex-row items-center justify-between w-full px-4">
      <div>
        <n-menu :options="menuOptions" mode="horizontal" responsive></n-menu>
      </div>
      <div>
        <n-button> 添加</n-button>
      </div>
    </div>
  </div>
</template>

<script setup lang="tsx">
import type { MenuOption } from "naive-ui";
import { Home20Regular } from "@vicons/fluent";
import { Settings20Regular } from "@vicons/fluent";

const menuOptions: MenuOption[] = [
  {
    label: "首页",
    key: "home",
    icon: () => (
      <n-icon>
        <Home20Regular />
      </n-icon>
    ),
  },
  {
    label: "设置",
    key: "setting",
    icon: () => (
      <n-icon>
        <Settings20Regular />
      </n-icon>
    ),
  },
];
</script>

<style scoped lang="scss"></style>
