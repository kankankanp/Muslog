/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class LikesService {
    /**
     * Like a post
     * Like a post by its ID
     * @param postId Post ID
     * @returns any Post liked successfully
     * @throws ApiError
     */
    public static postPostsLike(
        postId: number,
    ): CancelablePromise<{
        message?: string;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/posts/{postID}/like',
            path: {
                'postID': postId,
            },
            errors: {
                400: `Bad request`,
                401: `Unauthorized`,
                500: `Internal server error`,
            },
        });
    }
    /**
     * Unlike a post
     * Unlike a post by its ID
     * @param postId Post ID
     * @returns any Post unliked successfully
     * @throws ApiError
     */
    public static deletePostsUnlike(
        postId: number,
    ): CancelablePromise<{
        message?: string;
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/posts/{postID}/unlike',
            path: {
                'postID': postId,
            },
            errors: {
                400: `Bad request`,
                401: `Unauthorized`,
                500: `Internal server error`,
            },
        });
    }
    /**
     * Check if post is liked by user
     * Check if the current user has liked a specific post
     * @param postId Post ID
     * @returns any Successful response
     * @throws ApiError
     */
    public static getPostsLiked(
        postId: number,
    ): CancelablePromise<{
        isLiked?: boolean;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/posts/{postID}/liked',
            path: {
                'postID': postId,
            },
            errors: {
                401: `Unauthorized`,
                500: `Internal server error`,
            },
        });
    }
}
