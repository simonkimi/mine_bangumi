<template>
  <div class="pt-14 flex justify-center h-full">
    <div class="flex max-w-7xl w-full">
      <div class="min-w-60">
        <n-menu :options="menuOption" :value="currentSelect"></n-menu>
      </div>
      <div class="ml-5 w-full">
        <router-view></router-view>
      </div>
    </div>
  </div>
</template>

<script lang="tsx" setup>
import type { MenuOption } from "naive-ui";
import { NIcon } from "naive-ui";
import { RouterLink } from "vue-router";
import {
  UserSettingsRoute,
  RssSettingsRoute,
  DownloaderSettingsRoute,
} from "@/router";

const route = useRoute();

const currentSelect = computed(() => {
  switch (route.name) {
    case UserSettingsRoute:
      return "user";
    case RssSettingsRoute:
      return "rss";
    case DownloaderSettingsRoute:
      return "downloader";
    default:
      return "user";
  }
});

const menuOption: MenuOption[] = [
  {
    key: "user",
    label: () => (
      <RouterLink to={{ name: UserSettingsRoute }}>用户设置</RouterLink>
    ),
    icon: () => (
      <NIcon>
        <i-tabler-user />
      </NIcon>
    ),
  },
  {
    key: "rss",
    label: () => (
      <RouterLink to={{ name: RssSettingsRoute }}>RSS 设置</RouterLink>
    ),
    icon: () => (
      <NIcon>
        <i-tabler-rss />
      </NIcon>
    ),
  },
  {
    key: "downloader",
    label: () => (
      <RouterLink to={{ name: DownloaderSettingsRoute }}>下载器设置</RouterLink>
    ),
    icon: () => (
      <NIcon>
        <i-tabler-download />
      </NIcon>
    ),
  },
];
</script>

<style lang="scss" scoped></style>
