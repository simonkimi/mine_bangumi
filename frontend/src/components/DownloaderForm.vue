<template>
  <n-form
    ref="downloaderFormRef"
    :model="downloaderForm"
    :rules="downloaderFormRules"
  >
    <n-form-item label="下载器" required path="downloader">
      <n-select
        v-model:value="downloaderForm.downloader"
        :options="[
          {
            label: 'qBittorrent',
            value: 'qbittorrent',
          },
          {
            label: 'Aria2',
            value: 'aria2',
          },
          {
            label: '内置',
            value: 'builtin',
          },
        ]"
        placeholder="请选择下载器"
      >
      </n-select>
    </n-form-item>
    <n-form-item label="Api" path="api">
      <n-input
        v-model:value="downloaderForm.api"
        :placeholder="apiPlaceholder"
      />
    </n-form-item>
    <n-form-item label="用户名" path="username" v-if="displayUsername">
      <n-input v-model:value="downloaderForm.username" placeholder="" />
    </n-form-item>
    <n-form-item
      :label="apiPasswordTitle"
      path="password"
      v-if="displayPassword"
    >
      <n-input
        v-model:value="downloaderForm.password"
        type="password"
        placeholder=""
      />
    </n-form-item>
    <div class="flex space-x-4 mt-2">
      <n-button
        class="flex-grow-[3]"
        type="primary"
        @click="downloaderNextStep"
        :loading="props.loading"
      >
        下一步
      </n-button>
      <n-button class="flex-grow-[1]" @click="skipDownloader">跳过</n-button>
    </div>
  </n-form>
</template>

<script setup lang="ts">
import type { FormInst, FormRules } from "naive-ui";
import type { DownloaderInfo, DownloaderType } from "@/api/model";

const dialog = useDialog();

const props = defineProps<{
  type: DownloaderType;
  api: string;
  username: string;
  token: string;
  loading: boolean;
}>();

const emit = defineEmits<{
  onSkip: [];
  onNextStep: [model: DownloaderInfo];
}>();

const downloaderFormRef = ref<FormInst | null>(null);

const downloaderForm = ref<DownloaderInfo>({
  type: props.type,
  api: props.api,
  username: props.username,
  token: props.token,
});

const apiPlaceholder = computed(() => {
  switch (downloaderForm.value.type) {
    case "qbittorrent":
      return "host:port";
    case "aria2":
      return "scheme://host:port/jsonrpc";
    default:
      return "";
  }
});

function skipDownloader() {
  dialog.info({
    title: "跳过下载器配置",
    content: "您后续可以在设置中进行配置, 但是在未配置下载器前无法使用下载功能",
    positiveText: "跳过",
    negativeText: "取消",
    onNegativeClick: () => {},
    onPositiveClick: () => {
      emit("onSkip");
    },
  });
}

async function downloaderNextStep() {
  const isValid = await downloaderFormRef.value?.validate();
  if (!isValid) return;
  emit("onNextStep", downloaderForm.value);
}

const apiPasswordTitle = computed(() => {
  switch (downloaderForm.value.type) {
    case "qbittorrent":
      return "密码";
    case "aria2":
      return "Token";
    default:
      return "None";
  }
});

const displayUsername = computed(() => {
  return downloaderForm.value.type === "qbittorrent";
});

const displayPassword = computed(() =>
  ["qbittorrent", "aria2"].includes(downloaderForm.value.type)
);

const downloaderFormRules: FormRules = {
  downloader: [
    {
      required: true,
      message: "请选择下载器",
      trigger: "change",
    },
  ],
  api: [
    {
      required: true,
      message: "请输入Api地址",
      trigger: "blur",
    },
    {
      required: true,
      message: "请输入正确的Api地址",
      validator: (rule, value) => {
        switch (downloaderForm.value.type) {
          case "qbittorrent":
            return /^([a-zA-Z0-9.-]+):(\d{1,5})$/.test(value);
          case "aria2":
            return /^(https?|wss?):\/\/[^\s$.?#].\S*\/[a-zA-Z0-9_-]+$/.test(
              value
            );
          default:
            return true;
        }
      },
    },
  ],
};
</script>

<style scoped lang="scss"></style>
