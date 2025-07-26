/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Post } from '../models/Post';
import type { User } from '../models/User';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class UsersService {
    /**
     * Get all users
     * Get all users
     * @returns any Successful response
     * @throws ApiError
     */
    public static getUsers(): CancelablePromise<{
        message?: string;
        users?: Array<User>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/users',
        });
    }
    /**
     * Get a user by ID
     * Get a single user by their ID
     * @param id User ID
     * @returns any Successful response
     * @throws ApiError
     */
    public static getUsers1(
        id: string,
    ): CancelablePromise<{
        message?: string;
        user?: User;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/users/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }
    /**
     * Get user posts
     * Get all posts by a user
     * @param id User ID
     * @returns any Successful response
     * @throws ApiError
     */
    public static getUsersPosts(
        id: string,
    ): CancelablePromise<{
        message?: string;
        posts?: Array<Post>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/users/{id}/posts',
            path: {
                'id': id,
            },
        });
    }
}
