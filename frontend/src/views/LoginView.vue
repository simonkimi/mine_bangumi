<template>
  <div class="flex flex-col items-center justify-center h-screen">
    <p class="mt-4 text-2xl">登录</p>
    <n-card class="mt-10 max-w-md">
      <n-form ref="formRef" :model="form" :rules="rules">
        <n-form-item label="用户名" path="username">
          <n-input v-model:value="form.username" placeholder="" />
        </n-form-item>
        <n-form-item label="密码" path="password">
          <n-input
            v-model:value="form.password"
            placeholder=""
            type="password"
          />
        </n-form-item>
        <div class="flex mt-10">
          <n-button class="flex-1" type="primary" @click="login">登录</n-button>
        </div>
      </n-form>
    </n-card>
  </div>
</template>

<script lang="ts" setup>
import type { FormInst, FormRules } from "naive-ui";
import { useApi } from "@/api/api";
import { validateForm } from "@/utils/form";
import { ApiError } from "@/api/model";
import { errorMessage } from "@/api/errno";
import { useSystemStore } from "@/stores/systemStore";
import { useUserStore } from "@/stores/userStore";

const api = useApi();

const isLoading = ref(false);
const formRef = ref<FormInst | null>(null);
const dialog = useDialog();
const userStore = useUserStore();
const systemStore = useSystemStore();
const router = useRouter();
const route = useRoute();
const message = useMessage();

const form = ref({
  username: userStore.username,
  password: "",
});

const rules: FormRules = {
  username: [{ required: true, message: "请输入用户名" }],
  password: [{ required: true, message: "请输入密码" }],
};

async function login() {
  isLoading.value = true;
  try {
    if (!(await validateForm(formRef.value))) {
      return;
    }

    const result = await api.login(form.value.username, form.value.password);
    await userStore.setApiToken(result.token);
    await systemStore.setLogin(true);
    message.success("登录成功");
    const redirect = route.query.redirect as string | undefined;
    if (redirect) {
      await router.push(redirect);
    } else {
      await router.push("/");
    }
  } catch (e) {
    if (e instanceof ApiError) {
      dialog.error({
        title: "错误",
        content: errorMessage(e.code),
        positiveText: "确定",
      });
    }
  } finally {
    isLoading.value = false;
  }
}
</script>

<style lang="scss" scoped></style>
