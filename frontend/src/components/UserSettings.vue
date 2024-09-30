<template>
  <div class="mt-5">
    <h1 class="text-2xl font-semibold">密码设置</h1>
    <n-divider />
    <n-form
      ref="userRef"
      :model="userForm"
      :rules="userRules"
      class="max-w-md"
      size="small"
    >
      <n-form-item label="用户名" path="username" required>
        <n-input v-model:value="userForm.username" placeholder="" />
      </n-form-item>
      <n-form-item label="密码" path="password">
        <n-input
          v-model:value="userForm.password"
          placeholder=""
          type="password"
          @keydown.enter.prevent
        />
      </n-form-item>
      <n-form-item label="重复密码" path="repeatPassword">
        <n-input
          v-model:value="userForm.repeatPassword"
          :disabled="!userForm.password"
          placeholder=""
          type="password"
          @keydown.enter.prevent
        />
      </n-form-item>
      <div class="flex flex-row items-center">
        <n-button size="small" type="primary" @click="updateUser"
          >修改密码
        </n-button>
        <p class="ml-5 text-info">更新密码同时会更新api key</p>
      </div>
    </n-form>

    <div class="h-10" />

    <h1 class="text-2xl font-semibold">ApiKey</h1>
    <n-divider />
    <n-form class="max-w-md" size="small">
      <n-form-item label="Apikey" path="username">
        <n-input :value="apiKey" placeholder="" readonly />
        <n-tooltip>
          <template #trigger>
            <n-button class="ml-3" size="small" type="primary">
              <n-icon>
                <i-ph-copy-fill />
              </n-icon>
            </n-button>
          </template>
          <span>点击复制</span>
        </n-tooltip>
      </n-form-item>
      <div class="flex">
        <n-button size="small" type="primary">更新ApiKey</n-button>
      </div>
    </n-form>
  </div>
</template>

<script lang="ts" setup>
import type { FormInst, FormItemRule, FormRules } from "naive-ui";
import { useUserStore } from "@/stores/userStore";
import { validateForm } from "@/utils/form";

const userStore = useUserStore();

const apiKey = ref(userStore.apiToken);

const userRef = ref<FormInst | null>();
const userForm = ref({
  username: userStore.username,
  password: "",
  repeatPassword: "",
});

async function updateUser() {
  console.log(userForm.value);
  if (!(await validateForm(userRef.value))) {
    return;
  }
}

const userRules: FormRules = {
  username: [
    {
      required: true,
      message: "请输入用户名",
    },
    {
      pattern: /^[a-zA-Z0-9_]{4,16}$/,
      message: "用户名必须为4-16位字母、数字或下划线",
    },
  ],
  password: {
    validator: (rule: FormItemRule, value: string) => {
      if (!value || value.length < 4 || value.length > 40) {
        return new Error("密码可以为空, 否则密码长度必须为4-40位");
      }
      return true;
    },
  },
  repeatPassword: {
    validator: (rule: FormItemRule, value: string) => {
      if (value !== userForm.value.password) {
        return new Error("两次输入的密码不一致");
      }
      return true;
    },
  },
};
</script>

<style lang="scss" scoped></style>
