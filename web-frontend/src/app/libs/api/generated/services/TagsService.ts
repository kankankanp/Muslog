/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Tag } from '../models/Tag';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class TagsService {
    /**
     * Create a new tag
     * Create a new tag
     * @param requestBody
     * @returns any Tag created successfully
     * @throws ApiError
     */
    public static postTags(
        requestBody: {
            name?: string;
        },
    ): CancelablePromise<{
        message?: string;
        tag?: Tag;
    }> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/tags',
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get all tags
     * Get all tags
     * @returns any Successful response
     * @throws ApiError
     */
    public static getTags(): CancelablePromise<{
        message?: string;
        tags?: Array<Tag>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/tags',
        });
    }
    /**
     * Get a tag by ID
     * Get a single tag by its ID
     * @param id Tag ID
     * @returns any Successful response
     * @throws ApiError
     */
    public static getTags1(
        id: number,
    ): CancelablePromise<{
        message?: string;
        tag?: Tag;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/tags/{id}',
            path: {
                'id': id,
            },
            errors: {
                404: `Not Found`,
            },
        });
    }
    /**
     * Update a tag
     * Update an existing tag
     * @param id Tag ID
     * @param requestBody
     * @returns any Tag updated successfully
     * @throws ApiError
     */
    public static putTags(
        id: number,
        requestBody: {
            name?: string;
        },
    ): CancelablePromise<{
        message?: string;
        tag?: Tag;
    }> {
        return __request(OpenAPI, {
            method: 'PUT',
            url: '/tags/{id}',
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
     * Delete a tag
     * Delete a tag by its ID
     * @param id Tag ID
     * @returns any Tag deleted successfully
     * @throws ApiError
     */
    public static deleteTags(
        id: number,
    ): CancelablePromise<{
        message?: string;
    }> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/tags/{id}',
            path: {
                'id': id,
            },
        });
    }
    /**
     * Add tags to a post
     * Add tags to a specific post
     * @param postId Post ID
     * @param requestBody
     * @returns void
     * @throws ApiError
     */
    public static postTagsPosts(
        postId: number,
        requestBody: {
            tag_names?: Array<string>;
        },
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'POST',
            url: '/tags/posts/{postID}',
            path: {
                'postID': postId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Remove tags from a post
     * Remove tags from a specific post
     * @param postId Post ID
     * @param requestBody
     * @returns void
     * @throws ApiError
     */
    public static deleteTagsPosts(
        postId: number,
        requestBody: {
            tag_names?: Array<string>;
        },
    ): CancelablePromise<void> {
        return __request(OpenAPI, {
            method: 'DELETE',
            url: '/tags/posts/{postID}',
            path: {
                'postID': postId,
            },
            body: requestBody,
            mediaType: 'application/json',
        });
    }
    /**
     * Get tags by post ID
     * Get all tags associated with a specific post
     * @param postId Post ID
     * @returns any Successful response
     * @throws ApiError
     */
    public static getTagsPosts(
        postId: number,
    ): CancelablePromise<{
        message?: string;
        tags?: Array<Tag>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/tags/posts/{postID}',
            path: {
                'postID': postId,
            },
        });
    }
}
