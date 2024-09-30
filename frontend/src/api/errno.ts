import { ApiStatusEnum } from "@/generate/graphql";

export function errorMessage(code: ApiStatusEnum): string {
  switch (code) {
    case ApiStatusEnum.BadRequest:
      return "请求错误";
    case ApiStatusEnum.Cancel:
      return "请求取消";
    case ApiStatusEnum.DatabaseMigrationError:
      return "数据库需要迁移";
    case ApiStatusEnum.Forbidden:
      return "拒绝访问";
    case ApiStatusEnum.FormValidationError:
      return "表单验证错误";
    case ApiStatusEnum.InternalServerError:
      return "内部服务器错误";
    case ApiStatusEnum.NotFound:
      return "数据不存在";
    case ApiStatusEnum.Success:
      return "成功";
    case ApiStatusEnum.ThirdPartyApiError:
      return "三方API请求错误";
    case ApiStatusEnum.Timeout:
      return "请求超时";
    case ApiStatusEnum.Unauthorized:
      return "未登录";
    case ApiStatusEnum.UserCredentialsError:
      return "用户名或者密码错误";
    default:
      return "Unknown Error";
  }
}
