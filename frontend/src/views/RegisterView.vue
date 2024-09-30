<template>
  <div class="flex flex-col items-center justify-center h-screen px-5">
    <h1 class="text-4xl font-bold">初始配置</h1>
    <p class="mt-4 text-lg">欢迎使用Mine Bangumi</p>
    <n-card class="mt-10 max-w-xl" title="用户配置">
      <n-form ref="userFormRef" :model="userForm" :rules="userFormRules">
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
        <div class="flex space-x-4 mt-2">
          <n-button
            :loading="isLoading"
            class="flex-grow-[3]"
            type="primary"
            @click="userNextStep"
          >
            注册
          </n-button>
          <n-button class="flex-grow-[1]" @click="skipUser">跳过</n-button>
        </div>
      </n-form>
    </n-card>
  </div>
</template>

<script lang="tsx" setup>
import type { FormInst, FormRules } from "naive-ui";
import { useApi } from "@/api/api";

const { initUser } = useApi();
const dialog = useDialog();
const userFormRef = ref<FormInst | null>(null);
const isLoading = ref(false);

interface UserModel {
  username: string;
  password: string;
  repeatPassword: string;
}

const userForm = ref<UserModel>({
  username: "",
  password: "",
  repeatPassword: "",
});

async function userNextStep() {
  isLoading.value = true;
  try {
    if (!(await userFormRef.value?.validate())) return;
    await initUser(userForm.value.username, userForm.value.password);
  } catch (error) {
    dialog.error({
      title: "注册失败",
    });
  } finally {
    isLoading.value = false;
  }
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
      positiveText: "跳过",
      negativeText: "取消",
      onPositiveClick: () => resolve(true),
      onNegativeClick: () => resolve(false),
      onMaskClick: () => resolve(false),
    });
  });

  if (!isSkip) return;
}

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
</script>

<style lang="scss" scoped></style>
