import client from "@/api/client";
import type { SystemInfo, UserLoginInfo } from "@/api/model";

export async function getSystemStatus(): Promise<SystemInfo> {
  return client.get("/api/v1/system/status");
}

export async function initUser(
  username: string,
  password: string
): Promise<UserLoginInfo> {
  return client.post("/api/v1/user/init", { username, password });
}

export async function login(
  username: string,
  password: string
): Promise<UserLoginInfo> {
  return client.post("/api/v1/user/login", { username, password });
}
