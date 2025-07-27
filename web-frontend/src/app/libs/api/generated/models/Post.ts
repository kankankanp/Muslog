/* generated using openapi-typescript-codegen -- do not edit */
/* istanbul ignore file */
/* tslint:disable */
/* eslint-disable */
import type { Tag } from './Tag';
import type { Track } from './Track';
export type Post = {
    id: number;
    title: string;
    description: string;
    userId: string;
    createdAt: Date;
    updatedAt: Date;
    tracks: Array<Track>;
    tags: Array<Tag>;
    likesCount: number;
};

