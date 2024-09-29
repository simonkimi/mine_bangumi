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

<script lang="tsx" setup>
import type { MenuOption } from "naive-ui";
import { NIcon } from "naive-ui";
import { Home20Regular } from "@vicons/fluent";
import { Settings20Regular } from "@vicons/fluent";
import { RouterLink } from "vue-router";

const menuOptions: MenuOption[] = [
  {
    label: () => (
      <RouterLink to={{ name: "home" }}>
        <p>首页</p>
      </RouterLink>
    ),
    key: "home",
    icon: () => (
      <NIcon>
        <Home20Regular />
      </NIcon>
    ),
  },
  {
    label: () => (
      <RouterLink to={{ name: "setting" }}>
        <p>设置</p>
      </RouterLink>
    ),
    key: "setting",
    icon: () => (
      <NIcon>
        <Settings20Regular />
      </NIcon>
    ),
  },
];
</script>

<style lang="scss" scoped></style>
