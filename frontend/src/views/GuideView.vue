<template>
  <div class="flex flex-col items-center justify-center h-screen">
    <h1 class="text-4xl font-bold">初始配置</h1>
    <p class="mt-4 text-lg">欢迎使用Mine Bangumi</p>

    <n-card title="用户配置" class="mt-10 max-w-xl">
      <n-form ref="userFormRef" :model="userForm" :rules="userFormRules">
        <n-form-item label="用户名" path="username" required>
          <n-input v-model:value="userForm.username" placeholder="" />
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input
            v-model:value="userForm.password"
            type="password"
            placeholder=""
            @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item label="重复密码" path="repeatPassword">
          <n-input
            v-model:value="userForm.repeatPassword"
            :disabled="!userForm.password"
            type="password"
            placeholder=""
            @keydown.enter.prevent
          />
        </n-form-item>
        <div class="flex space-x-4 mt-2">
          <n-button class="flex-grow-[3]" type="primary" @click="userNextStep">
            下一步
          </n-button>
          <n-button class="flex-grow-[1]" @click="skipUser">跳过</n-button>
        </div>
      </n-form>
    </n-card>

    <n-card title="下载器设置" class="mt-10 max-w-xl">
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
          >
            下一步
          </n-button>
          <n-button class="flex-grow-[1]" @click="skipDownloader"
            >跳过
          </n-button>
        </div>
      </n-form>
    </n-card>
  </div>
</template>

<script setup lang="tsx">
import type { FormInst, FormRules } from "naive-ui";

const userFormRef = ref<FormInst | null>(null);
const downloaderFormRef = ref<FormInst | null>(null);
const dialog = useDialog();

type downloaderType = "qbittorrent" | "aria2" | "builtin";

interface UserModel {
  username: string;
  password: string;
  repeatPassword: string;
}

interface DownloaderModel {
  downloader: downloaderType;
  api: string;
  username: string;
  password: string;
}

const userForm = ref<UserModel>({
  username: "",
  password: "",
  repeatPassword: "",
});

const downloaderForm = ref<DownloaderModel>({
  downloader: "qbittorrent",
  api: "",
  username: "",
  password: "",
});

const apiPlaceholder = computed(() => {
  switch (downloaderForm.value.downloader) {
    case "qbittorrent":
      return "host:port";
    case "aria2":
      return "scheme://host:port/jsonrpc";
    default:
      return "";
  }
});

const apiPasswordTitle = computed(() => {
  switch (downloaderForm.value.downloader) {
    case "qbittorrent":
      return "密码";
    case "aria2":
      return "Token";
    default:
      return "None";
  }
});

const displayUsername = computed(() => {
  return downloaderForm.value.downloader === "qbittorrent";
});

const displayPassword = computed(() =>
  ["qbittorrent", "aria2"].includes(downloaderForm.value.downloader)
);

const userFormRules: FormRules = {
  username: [
    {
      required: true,
      validator: (rule, value) => {
        if (!value) return new Error("请输入用户名");
        if (!/^[a-zA-Z0-9_]{4,16}$/.test(value))
          return new Error("用户名长度为4-16位，只能包含字母、数字和下划线");
        if (!/^[a-zA-Z]/.test(value)) return new Error("用户名必须以字母开头");
        return true;
      },
      trigger: "blur",
    },
  ],
  password: [
    {
      required: true,
      validator: (rule, value) => {
        if (!value) return new Error("请输入密码");
        if (!/^[\x20-\x7E]{4,50}$/.test(value))
          return new Error("密码长度为4-50位，只能包含英文、数字和特殊字符");
        return true;
      },
      trigger: "blur",
    },
  ],
  repeatPassword: [
    {
      required: true,
      message: "请重复密码",
      trigger: "blur",
    },
    {
      required: true,
      message: "两次输入的密码不一致",
      validator: (rule, value) => {
        return value === userForm.value.password;
      },
    },
  ],
};

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
        switch (downloaderForm.value.downloader) {
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

async function userNextStep() {
  const isValid = await userFormRef.value?.validate();
  if (!isValid) return;
}

async function downloaderNextStep() {
  const isValid = await downloaderFormRef.value?.validate();
  if (!isValid) return;
}

async function skipUser() {
  const isSkip = await new Promise((resolve) => {
    dialog.warning({
      title: "是否跳过配置",
      content: () => (
        <div class="my-2">
          <p>跳过用户配置将进行无密码登录</p>
          <p class="text-red-500 text-sm">
            <strong>请确保在安全的环境下使用</strong>
          </p>
        </div>
      ),
      positiveText: "继续",
      negativeText: "取消",
      onPositiveClick: () => resolve(true),
      onNegativeClick: () => resolve(false),
      onMaskClick: () => resolve(false),
    });
  });

  if (!isSkip) return;
}

function skipDownloader() {
  dialog.info({
    title: "跳过下载器配置",
    content: "您后续可以在设置中进行配置, 但是在未配置下载器前无法使用下载功能",
    positiveText: "继续",
    onPositiveClick: () => {},
  });
}
</script>

<style scoped lang="scss"></style>
