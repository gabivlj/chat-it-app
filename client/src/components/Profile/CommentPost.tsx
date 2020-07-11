import React from 'react';
import { Link } from 'react-router-dom';
import {
  GetUser_user,
  GetUser_user_commentsUser,
} from '../../queries/types/GetUser';

type Props = {
  user: GetUser_user;
  comment: GetUser_user_commentsUser;
};

export default function CommentPost({ user, comment }: Props) {
  return (
    <div className="h-auto flex-wrap flex w-full rounded overflow-hidden shadow-lg mt-3">
      <div className="px-6 py-4 w-full items-center">
        <div className="ml-40">
          {comment.post.image && (
            <Link to={`/post/${comment.post.id}`}>
              <img
                className="w-2/4 h-2/4 rounded-sm "
                src={comment.post.image.urlXL}
              ></img>
            </Link>
          )}
        </div>
        <div className="mb-2">
          <Link to={`/user/${user.username}`}>
            <img
              className="w-8 h-8 rounded-full mr-3 float-left"
              src={
                user.profileImage && user.profileImage.urlXL.trim() != ''
                  ? user.profileImage.urlXL
                  : 'https://www.netclipart.com/pp/m/232-2329525_person-svg-shadow-default-profile-picture-png.png'
              }
            ></img>
          </Link>
          <div className="mt-4"></div>
          <span className="font-bold text-xl mr-2">{user.username}</span> said
          on{' '}
          <Link
            to={`/post/${comment.post.id}`}
            className="text text-xl font-medium text-blue-500 hover:text-blue-800 "
          >
            {comment.post.title}
          </Link>
        </div>

        <span
          className="inline-block bg-gray-200 rounded-semifull px-3 py-1 text-sm font-semibold text-gray-700 mr-2 float-left w-3/4 "
          style={{ wordWrap: 'break-word' }}
        >
          {comment.text}
        </span>
      </div>

      <div className="px-6 py-6 w-full"></div>
    </div>
  );
}
