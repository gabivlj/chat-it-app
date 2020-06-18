import React from 'react';
import { GetUser_user_posts } from '../../queries/types/GetUser';

type Props = {
  post: GetUser_user_posts;
};

export default function Post({ post }: Props) {
  return (
    <div className="w-full md:w-1/3 xl:w-1/4 p-6 flex flex-col">
      <a href="#">
        <img
          className="hover:grow hover:shadow-lg"
          style={{ width: '20rem', height: '10rem' }}
          src={`${post.image ? post.image.urlXL : ''}`}
        />
        <div className="pt-3 flex items-center justify-between">
          <p className="">{post.title}</p>
        </div>
        <p className="pt-1 text-gray-900">{post.text.slice(0, 10)}...</p>
      </a>
    </div>
  );
}
