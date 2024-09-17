export interface ApiResult<T> {
  code: number;
  message: string | null;
  data: T;
}

export interface SystemInfo {
  version: string;
  is_init_user: boolean;
  is_login: boolean;
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
  code: number;

  constructor(code: number, message: string) {
    super(message);
    this.code = code;
  }
}
