/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { User } from '../models/User';
import type { UserLogin } from '../models/UserLogin';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class AuthService {
    /**
     * User login
     * Log in a user and return a JWT token.
     * @param requestBody
     * @returns any Login successful
     * @throws ApiError
     */
    public static postLogin(
        requestBody: UserLogin,
    ): CancelablePromise<{
        message?: string;
        user?: User;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/login',
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                401: `Unauthorized`,
            },
        });
    }
    /**
     * Refresh JWT token
     * Refresh the JWT access token using the refresh token.
     * @returns any Token refreshed
     * @throws ApiError
     */
    public static postRefresh(): CancelablePromise<{
        message?: string;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/refresh',
            errors: {
                401: `Unauthorized`,
            },
        });
    }
    /**
     * User logout
     * Log out a user by clearing JWT cookies.
     * @returns any Logout successful
     * @throws ApiError
     */
    public static postLogout(): CancelablePromise<{
        message?: string;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/logout',
        });
    }
    /**
     * Get current user
     * Get the currently logged in user's information.
     * @returns User Successful response
     * @throws ApiError
     */
    public static getMe(): CancelablePromise<User> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/me',
            errors: {
                404: `User not found`,
            },
        });
    }
}
