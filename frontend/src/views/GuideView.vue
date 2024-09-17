<template>
  <div class="flex flex-col items-center justify-center h-screen">
    <h1 class="text-4xl font-bold">初始配置</h1>
    <p class="mt-4 text-lg">欢迎使用Mine Bangumi</p>

    <n-card
      title="用户配置"
      class="mt-10 max-w-xl"
      v-if="currentStep === 'user'"
    >
      <UserForm
        :loading="loading"
        :username="username"
        @onNextStep="onUserNextStep"
        @onSkip="onUserSkip"
      />
    </n-card>

    <n-card
      title="下载器设置"
      class="mt-10 max-w-xl"
      v-if="currentStep === 'downloader'"
    >
      <DownloaderForm
        :loading="loading"
        :type="dlType"
        :api="dlApi"
        :username="dlUsername"
        :token="dlToken"
      />
    </n-card>
  </div>
</template>

<script setup lang="ts">
import { initUser } from "@/api/api";
import { ApiError } from "@/api/model";

const dialog = useDialog();

type Step = "user" | "downloader";
const loading = ref(false);
const currentStep = ref<Step>("user");
const username = ref("admin");

const dlType = ref("qbittorrent");
const dlApi = ref("http://localhost:8080");
const dlUsername = ref("admin");
const dlToken = ref("");

async function onUserNextStep(username: string, password: string) {
  try {
    loading.value = true;
    const token = await initUser(username, password);
    console.log(token);
    currentStep.value = "downloader";
  } catch (e) {
    if (e instanceof ApiError) {
      dialog.error({
        title: "错误",
        content: `初始化用户错误, 错误码: ${e.code}, 错误信息: ${e.message}`,
        positiveText: "确定",
      });
    } else {
      console.error(e);
    }
  } finally {
    loading.value = false;
  }
}

function onUserSkip() {
  currentStep.value = "downloader";
}
</script>

<style scoped lang="scss"></style>
