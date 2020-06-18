import React from 'react';
import Post from './Post';
import { GetUser_user_posts } from '../../queries/types/GetUser';

type Props = {
  posts: GetUser_user_posts[];
};
export default function PostsProfile({ posts }: Props) {
  return (
    <div className="container mx-auto flex items-center flex-wrap pt-4 pb-12">
      {posts.map(post => (
        <Post post={post} />
      ))}
    </div>
  );
}
