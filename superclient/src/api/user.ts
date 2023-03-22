import { http } from "@/utils/http";
import { Role } from "@/model/role";

export type UserResult = {
    success: boolean;
    trace_id: number;
    message: string;
    data: {
        user_id: number;
        name: string,
        username: string;
        accessToken: string;
        refreshToken: string;
        expires: number;
    };
};

export type RefreshTokenResult = {
    success: boolean;
    trace_id: number;
    message: string;
    data: {
        /** `token` */
        accessToken: string;
        /** 用于调用刷新`accessToken`的接口时所需的`token` */
        refreshToken: string;
        /** `accessToken`的过期时间（格式'xxxx/xx/xx xx:xx:xx'） */
        expires: number;
    };
};

/** 登录 */
export const getLogin = (data?: object) => {
    return http.request<UserResult>("post", "/api/user/login", { data });
};

/** 刷新token */
export const refreshTokenApi = (data?: object) => {
    return http.request<RefreshTokenResult>("post", "/api/user/refresh", { data });
};
