import React from 'react';
import IconPosts from '../Utils/Line';
import { getPostAndMessages_post } from '../../queries/types/getPostAndMessages';
import { Link } from 'react-router-dom';

type Props = {
  post: getPostAndMessages_post;
};

export default function PostCard({ post }: Props) {
  return (
    <>
      <div className="max-w-4xl flex items-center h-auto flex-wrap mx-auto my-32 lg:my-0">
        {/* <!--Img Col--> */}
        <div className="w-full lg:w-4/5 justify-center flex">
          {/* <!-- Big profile image for side bar (desktop) --> */}
          <div>
            {post.image && (
              <img
                src={post.image.urlXL}
                className="rounded-none lg:rounded-lg text-center shadow-2xl hidden lg:block"
                alt={``}
              />
            )}
          </div>
        </div>
        <div
          id="profile"
          className="w-full rounded-lg lg:rounded-l-lg lg:rounded-r-none shadow-2xl bg-white opacity-75 mx-6 lg:mx-0"
        >
          <div className="p-4 md:p-12 text-center lg:text-left">
            {/* <!-- Image for mobile view--> */}
            {post.image && (
              <div
                className="block lg:hidden shadow-xl mx-auto -mt-16 h-48 w-48 bg-cover bg-center"
                style={{
                  backgroundImage: `url('${post.image.urlXL}')`,
                }}
              ></div>
            )}

            <h1 className="text-3xl font-bold pt-8 lg:pt-0">{post.title}</h1>
            <div className="mx-auto lg:mx-0 w-4/5 pt-3 border-b-2 border-teal-500 opacity-25"></div>
            <Link to={`/user/${post.user.username}`}>
              <p className="pt-4 text-base font-bold flex items-center justify-center lg:justify-start underline">
                <img
                  className="w-4 rounded-full mr-3"
                  src={
                    post.user.profileImage ? post.user.profileImage.urlXL : ''
                  }
                ></img>
                {post.user.username}
              </p>
            </Link>
            <div className="mt-12 h-64 overflow-y-auto">
              <p className="text-md break-words">{post.text}</p>
            </div>

            <div className="pt-12 pb-8"></div>
            {/* <SocialNetworks /> */}
          </div>
        </div>
      </div>
    </>
  );
}
