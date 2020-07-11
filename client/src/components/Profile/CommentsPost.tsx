import React from 'react';
import CommentPost from './CommentPost';
import {
  GetUser_user,
  GetUser_user_commentsUser,
} from '../../queries/types/GetUser';

type Props = {
  user: GetUser_user;
  comments: GetUser_user_commentsUser[];
};

export default function CommentsPost({ user, comments }: Props) {
  return (
    <div>
      {comments.map((el) => (
        <CommentPost user={user} comment={el} key={el.id} />
      ))}
    </div>
  );
}
