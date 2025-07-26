/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Post } from '../models/Post';
import type { Track } from '../models/Track';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class BlogsService {
    /**
     * Get all blogs
     * Get all post posts
     * @returns any Successful response
     * @throws ApiError
     */
    public static getBlogs(): CancelablePromise<{
        message?: string;
        posts?: Array<Post>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/blogs',
        });
    }
    /**
     * Create a new post
     * Create a new post post
     * @param requestBody
     * @returns any Post created successfully
     * @throws ApiError
     */
    public static postBlogs(
        requestBody: {
            title?: string;
            description?: string;
            userId?: string;
            tracks?: Array<Track>;
        },
    ): CancelablePromise<{
        message?: string;
        post?: Post;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/blogs',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get a post by ID
     * Get a single post post by its ID
     * @param id Post ID
     * @returns any Successful response
     * @throws ApiError
     */
    public static getBlogs1(
        id: number,
    ): CancelablePromise<{
        message?: string;
        post?: Post;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/blogs/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }
    /**
     * Update a post
     * Update an existing post post
     * @param id Post ID
     * @param requestBody
     * @returns any Post updated successfully
     * @throws ApiError
     */
    public static putBlogs(
        id: number,
        requestBody: {
            title?: string;
            description?: string;
        },
    ): CancelablePromise<{
        message?: string;
        post?: Post;
    }> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/blogs/{id}',
            path: {
                'id': id,
            },
            body: requestBody,
            mediaType: 'application/json',
            errors: {
                404: `Not Found`,
            },
        });
    }
    /**
     * Delete a post
     * Delete a post post by its ID
     * @param id Post ID
     * @returns any Post deleted successfully
     * @throws ApiError
     */
    public static deleteBlogs(
        id: number,
    ): CancelablePromise<{
        message?: string;
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/blogs/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Get blogs by page
     * Get post posts paginated
     * @param page Page number
     * @returns any Successful response
     * @throws ApiError
     */
    public static getBlogsPage(
        page: number,
    ): CancelablePromise<{
        message?: string;
        posts?: Array<Post>;
        totalCount?: number;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/blogs/page/{page}',
            path: {
                'page': page,
            },
        });
    }
}
