import { ApiStatusEnum } from "@/gql/graphql";

export interface ApiResult<T> {
  code: ApiStatusEnum;
  message: string | null;
  data: T;
}

export interface SystemInfo {
  version: string;
  app_database_version: number;
  current_database_version: number;
  is_system_init: boolean;
  is_login: boolean;
  username: string;
}

export type DownloaderType = "qbittorrent" | "aria2" | "builtin";

export interface DownloaderInfo {
  type: DownloaderType;
  api: string;
  username: string | null;
  token: string | null;
}

export interface UserLoginInfo {
  token: string;
}

export class ApiError extends Error {
  code: ApiStatusEnum;

  constructor(code: ApiStatusEnum, message: string) {
    super(message);
    this.code = code;
  }
}
