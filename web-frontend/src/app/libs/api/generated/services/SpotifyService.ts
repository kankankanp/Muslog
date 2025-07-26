/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Track } from '../models/Track';
import type { CancelablePromise } from '../core/CancelablePromise';
import { OpenAPI } from '../core/OpenAPI';
import { request as __request } from '../core/request';
export class SpotifyService {
    /**
     * Search Spotify tracks
     * Search for tracks on Spotify by a query string.
     * @param q The search query string for tracks.
     * @returns any Successful response with a list of tracks.
     * @throws ApiError
     */
    public static getSpotifySearch(
        q: string,
    ): CancelablePromise<{
        message?: string;
        tracks?: Array<Track>;
    }> {
        return __request(OpenAPI, {
            method: 'GET',
            url: '/spotify/search',
            query: {
                'q': q,
            },
            errors: {
                400: `Bad request, missing search term.`,
                500: `Internal server error.`,
            },
        });
    }
}
