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
import { RouterLink } from "vue-router";
import { HomeRoute, SettingsRoute } from "@/router";

const menuOptions: MenuOption[] = [
  {
    label: () => (
      <RouterLink to={{ name: HomeRoute }}>
        <p>首页</p>
      </RouterLink>
    ),
    key: "home",
    icon: () => (
      <NIcon>
        <i-tabler-home />
      </NIcon>
    ),
  },
  {
    label: () => (
      <RouterLink to={{ name: SettingsRoute }}>
        <p>设置</p>
      </RouterLink>
    ),
    key: "settings",
    icon: () => (
      <NIcon>
        <i-ant-design-setting-outlined />
      </NIcon>
    ),
  },
];
</script>

<style lang="scss" scoped></style>
