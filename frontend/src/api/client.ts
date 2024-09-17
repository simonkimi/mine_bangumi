import axios, { type AxiosResponse } from "axios";
import { isWailsRuntime } from "@/utils/runtime";
import { GetIpv4Port } from "../../wailsjs/go/gui/App";
import memoize from "lodash/memoize";
import { ApiError, type ApiResult } from "@/api/model";

const getBaseUrl = memoize(async () => {
  if (isWailsRuntime()) {
    const post = await GetIpv4Port();
    return `http://localhost:${post}`;
  }
  return void 0;
});

const instance = axios.create({
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});

instance.interceptors.request.use(
  async (config) => {
    if (!config.baseURL) {
      config.baseURL = await getBaseUrl();
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

instance.interceptors.response.use(
  (response: AxiosResponse<ApiResult<any>>) => {
    const apiResult = response.data;

    if (apiResult.code !== 200) {
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

export default instance;
