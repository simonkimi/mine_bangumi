import client from "@/api/client";
import type { ApiResult, SystemInfo } from "@/api/model";

export async function getSystemConfig(): Promise<ApiResult<SystemInfo>> {
  return client.get("/api/v1/system/config");
}
