import axios, { type AxiosResponse } from "axios";
import { ApiError, type ApiResult } from "@/api/model";
import { ApiStatusEnum } from "@/gql/graphql";

const instance = axios.create({
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});

instance.interceptors.response.use(
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

export default instance;
