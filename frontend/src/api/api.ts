import client from "@/api/client";
import type { SystemInfo, UserLoginInfo } from "@/api/model";

export async function getSystemConfig(): Promise<SystemInfo> {
  return client.get("/api/v1/config/system");
}

export async function initUser(
  username: string,
  password: string
): Promise<UserLoginInfo> {
  return client.post("/api/v1/config/init_user", { username, password });
}
