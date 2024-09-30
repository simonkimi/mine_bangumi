import axios, { type AxiosResponse } from "axios";
import {
  ApiError,
  type ApiResult,
  type SystemInfo,
  type UserLoginInfo,
} from "@/api/model";
import { ApiStatusEnum } from "@/gql/graphql";
import { useUserStore } from "@/stores/userStore";

const client = axios.create({
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});

client.interceptors.request.use((config) => {
  const userStore = useUserStore();
  const apiToken = userStore.apiToken;
  if (apiToken) {
    config.headers["Authorization"] = `Token ${apiToken}`;
  }
  return config;
});

client.interceptors.response.use(
  (response: AxiosResponse<ApiResult<any>>) => {
    const apiResult = response.data;

    if (apiResult.code !== ApiStatusEnum.Success) {
      return Promise.reject(
        new ApiError(apiResult.code, apiResult.message ?? "Unknown error")
      );
    }

    return apiResult.data;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export function useApi() {
  async function getSystemStatus(): Promise<SystemInfo> {
    return client.get("/api/v1/system/status");
  }

  async function initUser(
    username: string,
    password: string
  ): Promise<UserLoginInfo> {
    return client.post("/api/v1/user/init", { username, password });
  }

  async function login(
    username: string,
    password: string
  ): Promise<UserLoginInfo> {
    return client.post("/api/v1/user/login", { username, password });
  }

  return {
    getSystemStatus,
    initUser,
    login,
  };
}
