/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { UserQuery, Params } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL query operation: GetUser
// ====================================================

export interface GetUser_user_profileImage {
  __typename: "Image";
  urlXL: string;
}

export interface GetUser_user_postsUser_image {
  __typename: "Image";
  urlXL: string;
}

export interface GetUser_user_postsUser {
  __typename: "Post";
  text: string;
  id: string;
  title: string;
  image: GetUser_user_postsUser_image | null;
}

export interface GetUser_user_commentsUser_post_image {
  __typename: "Image";
  urlXL: string;
}

export interface GetUser_user_commentsUser_post {
  __typename: "Post";
  text: string;
  id: string;
  title: string;
  image: GetUser_user_commentsUser_post_image | null;
}

export interface GetUser_user_commentsUser {
  __typename: "Message";
  id: string;
  text: string;
  post: GetUser_user_commentsUser_post;
}

export interface GetUser_user {
  __typename: "User";
  username: string;
  id: string;
  numberOfComments: number;
  numberOfPosts: number;
  profileImage: GetUser_user_profileImage | null;
  postsUser: GetUser_user_postsUser[];
  commentsUser: GetUser_user_commentsUser[];
}

export interface GetUser {
  user: GetUser_user;
}

export interface GetUserVariables {
  query: UserQuery;
  paramsPost?: Params | null;
  paramsComments?: Params | null;
}
