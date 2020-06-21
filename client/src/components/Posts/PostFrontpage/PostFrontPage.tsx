import React from 'react';
import { posts_posts } from '../../../queries/types/posts';
import { Link } from 'react-router-dom';

type Props = {
  post: posts_posts;
};

export default function PostFrontPage({ post }: Props) {
  return (
    <div className="h-auto flex-wrap flex w-full rounded overflow-hidden shadow-lg mt-3">
      <Link to={`/post/${post.id}`} className="">
        {post.image && post.image.urlXL !== '' && (
          <img
            className="w-full"
            src={!post.image || post.image.urlXL === '' ? '' : post.image.urlXL}
            alt={post.title}
          />
        )}

        <div className="px-6 py-4 w-full">
          <div className="font-bold text-xl mb-2">{post.title}</div>
          <p className="text-gray-700 text-base">
            {post.text.slice(0, 100)} {post.text.length > 100 && '...'}
          </p>
        </div>
      </Link>
      <div className="px-6 py-6 w-full">
        <Link to={`/user/${post.user.username}`}>
          <span className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 ">
            <img
              className="w-4 h-4 rounded-full mr-3 float-left mt-1"
              src={
                post.user.profileImage &&
                post.user.profileImage.urlXL.trim() != ''
                  ? post.user.profileImage.urlXL
                  : 'https://www.netclipart.com/pp/m/232-2329525_person-svg-shadow-default-profile-picture-png.png'
              }
            ></img>
            <p className="float-right">{post.user.username}</p>
          </span>
        </Link>
        <span className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 float-left">
          Comments: 1221
        </span>
      </div>
    </div>
  );
}
