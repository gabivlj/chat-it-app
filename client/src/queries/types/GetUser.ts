/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { UserQuery } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL query operation: GetUser
// ====================================================

export interface GetUser_user_profileImage {
  __typename: "Image";
  urlXL: string;
}

export interface GetUser_user_posts_image {
  __typename: "Image";
  urlXL: string;
}

export interface GetUser_user_posts {
  __typename: "Post";
  text: string;
  image: GetUser_user_posts_image | null;
  title: string;
}

export interface GetUser_user {
  __typename: "User";
  username: string;
  id: string;
  profileImage: GetUser_user_profileImage | null;
  posts: GetUser_user_posts[];
}

export interface GetUser {
  user: GetUser_user;
}

export interface GetUserVariables {
  query: UserQuery;
}
