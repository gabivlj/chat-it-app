import React from 'react';
import { posts_posts } from '../../../queries/types/posts';
import PostFrontPage from './PostFrontPage';

type Props = {
  posts: posts_posts[];
};

export default function PostsFrontPage({ posts }: Props) {
  return (
    <div className="container mx-auto flex items-center flex-wrap pt-4 pb-12">
      {posts.map(post => (
        <PostFrontPage key={post.id} post={post} />
      ))}
    </div>
  );
}
