/* tslint:disable */
/* eslint-disable */
// @generated
// This file was automatically generated and should not be edited.

import { Params } from "./../../types/graphql-global-types";

// ====================================================
// GraphQL query operation: posts
// ====================================================

export interface posts_posts_image {
  __typename: "Image";
  urlXL: string;
}

export interface posts_posts_user_profileImage {
  __typename: "Image";
  urlXL: string;
}

export interface posts_posts_user {
  __typename: "User";
  username: string;
  profileImage: posts_posts_user_profileImage | null;
}

export interface posts_posts {
  __typename: "Post";
  numberOfComments: number;
  text: string;
  id: string;
  title: string;
  image: posts_posts_image | null;
  user: posts_posts_user;
}

export interface posts {
  posts: posts_posts[];
}

export interface postsVariables {
  params?: Params | null;
}
