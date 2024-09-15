export interface ApiResult<T> {
  code: number;
  message: string | null;
  data: T;
}

export interface SystemInfo {
  version: string;
  is_first_run: boolean;
  is_login: boolean;
}
