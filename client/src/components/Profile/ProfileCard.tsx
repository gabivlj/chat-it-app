import React from 'react';
import { GetUser_user } from '../../queries/types/GetUser';
import Icon from '../Utils/Line';
import IconPlace from '../Utils/IconPlace';
import {
  CHECK_LOGED_LOCAL,
  UPLOAD_PROFILE_IMAGE
} from '../../queries/user_queries';
import { isLogged } from '../../queries/types/isLogged';
import { useQuery, useMutation } from '@apollo/react-hooks';
import { newImage, newImageVariables } from '../../queries/types/newImage';
import { TODO } from '../../utils/todo';
import PostsProfile from '../Posts/PostsProfile';
import { smoothScroll } from '../../utils/smoothScroll';
import UploadFile from '../Inputs/UploadFile';

type Props = {
  user: GetUser_user;
};

export default function ProfileCard({ user }: Props) {
  const { data } = useQuery<isLogged>(CHECK_LOGED_LOCAL);
  const [mutate, { error, loading }] = useMutation<newImage, newImageVariables>(
    UPLOAD_PROFILE_IMAGE,
    {
      onError: err => {
        TODO(`Handle err ${err}`);
      }
    }
  );

  function onClickSeePosts() {
    const posts = document.getElementById('posts');
    if (!posts) return;
    smoothScroll(posts);
  }

  function onChangeFile({
    target: {
      validity,
      files: [file]
    }
  }: any) {
    if (validity.valid) mutate({ variables: { image: file } });
  }

  return (
    <>
      <div className="max-w-4xl flex items-center h-auto lg:h-screen flex-wrap mx-auto my-32 lg:my-0">
        <div
          id="profile"
          className="w-full lg:w-3/5 rounded-lg lg:rounded-l-lg lg:rounded-r-none shadow-2xl bg-white opacity-75 mx-6 lg:mx-0"
        >
          <div className="p-4 md:p-12 text-center lg:text-left">
            {/* <!-- Image for mobile view--> */}
            <div
              className="block lg:hidden rounded-full shadow-xl mx-auto -mt-16 h-48 w-48 bg-cover bg-center"
              style={{
                backgroundImage: `url('${
                  user.profileImage && user.profileImage.urlXL.trim() !== ''
                    ? user.profileImage.urlXL
                    : 'https://www.netclipart.com/pp/m/232-2329525_person-svg-shadow-default-profile-picture-png.png'
                }')`
              }}
            ></div>

            <h1 className="text-3xl font-bold pt-8 lg:pt-0">{user.username}</h1>
            <div className="mx-auto lg:mx-0 w-4/5 pt-3 border-b-2 border-teal-500 opacity-25"></div>
            <p className="pt-4 text-base font-bold flex items-center justify-center lg:justify-start">
              <Icon /> todo: Put how many posts
            </p>
            <p className="pt-2 text-gray-600 text-xs lg:text-sm flex items-center justify-center lg:justify-start">
              <IconPlace /> Todo: Put how many comments
            </p>
            <p className="pt-8 text-sm">
              Totally optional short description about yourself, what you do and
              so on.
            </p>

            <div className="pt-12 pb-8">
              <button
                className="bg-teal-700 hover:bg-teal-900 text-white font-bold py-2 px-4 rounded-full mt-3"
                onClick={onClickSeePosts}
              >
                Show last posts
              </button>
              <button className="bg-teal-700 hover:bg-teal-900 text-white font-bold py-2 px-4 rounded-full ml-3 mt-3">
                Show last comments
              </button>
              {data &&
              data.loged &&
              data.loged.user &&
              data.loged.user.username === user.username ? (
                <UploadFile onChange={onChangeFile} />
              ) : (
                <></>
              )}
            </div>
            <div className="pt-12 pb-8"></div>
            {/* <SocialNetworks /> */}
          </div>
        </div>

        {/* <!--Img Col--> */}
        <div className="w-full lg:w-2/5">
          {/* <!-- Big profile image for side bar (desktop) --> */}
          <img
            src={
              user.profileImage && user.profileImage.urlXL.trim() !== ''
                ? user.profileImage.urlXL
                : 'https://www.netclipart.com/pp/m/232-2329525_person-svg-shadow-default-profile-picture-png.png'
            }
            className="rounded-none lg:rounded-lg shadow-2xl hidden lg:block"
            alt={``}
          />
          {loading && `Loading upload...`}
        </div>
      </div>
      <div id="posts">
        <h1 className="text-3xl font-bold pt-8 lg:pt-0">
          Posts from {user.username}
        </h1>
        <PostsProfile posts={user.posts} />
      </div>
    </>
  );
}
