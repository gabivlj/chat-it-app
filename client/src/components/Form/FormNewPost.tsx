import React, { useState } from 'react';
import UploadFile from '../Inputs/UploadFile';
import { useMutation } from '@apollo/react-hooks';
import { CREATE_POST, POST_FRONTPAGE } from '../../queries/post_queries';
import {
  createPost,
  createPostVariables
} from '../../queries/types/createPost';
import Loading from '../Utils/Loading';

export default function FormNewPost() {
  const [state, setState] = useState({ text: '', title: '', file: null });
  const [loading, setLoading] = useState(false);
  const [mutate] = useMutation<createPost, createPostVariables>(CREATE_POST);

  function onChange(e: any) {
    setState({ ...state, [e.target.name]: e.target.value });
  }

  function onChangeFile({
    target: {
      validity,
      files: [file]
    }
  }: any) {
    if (validity.valid) setState({ ...state, file });
  }

  function onSubmit(e: any) {
    e.preventDefault();
    const { text, file, title } = state;
    setLoading(true);
    mutate({
      variables: {
        form: {
          text,
          image: file,
          title
        }
      },
      update: (proxy, { data: { newPost } }: any) => {
        // Read the data from our cache for this query.
        const data = proxy.readQuery({
          query: POST_FRONTPAGE,
          variables: {
            params: {
              limit: 5
            }
          }
        }) as any;
        // Write our data back to the cache with the new comment in it
        proxy.writeQuery({
          query: POST_FRONTPAGE,
          variables: {
            params: {
              limit: 5
            }
          },
          data: {
            ...data,
            posts: [newPost, ...data.posts]
          }
        });
        setLoading(false);
        setState({ text: '', title: '', file: null });
      }
    });
  }

  return (
    <div className="h-auto flex-wrap flex w-full rounded overflow-hidden shadow-lg mt-3">
      <div className="container mx-auto flex items-center flex-wrap pt-4 pb-12">
        <div className="w-full md:w-1/3 px-3 mb-6 md:mb-0">
          <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
            <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
              Text
            </label>
          </div>
          <textarea
            name="text"
            value={state.text}
            onChange={onChange}
            className="resize-y md:w-2/3 border rounded focus:outline-none focus:shadow-outline appearance-none block w-full bg-gray-200 text-gray-700 border rounded py-3 px-4 mb-3 leading-tight focus:outline-none focus:bg-white"
          ></textarea>
        </div>
        <div className="w-full md:w-1/3 px-3 mb-6 md:mb-0">
          <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
            <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
              Title
            </label>
          </div>
          <input
            name="title"
            value={state.title}
            onChange={onChange}
            className="appearance-none block w-full bg-gray-200 text-gray-700 border border-gray-200 rounded py-3 px-4 leading-tight focus:outline-none focus:bg-white focus:border-gray-500"
            type="text"
          />
        </div>
        <div className="w-full md:w-1/3 px-3 mb-6 md:mb-0">
          <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
            <label className="block uppercase tracking-wide text-gray-700 text-xs font-bold mb-2">
              Upload Image
            </label>
          </div>
          <UploadFile onChange={onChangeFile} uploadImageText="" />
        </div>
        <div className="w-full md:w-1/3 px-3 mb-6 md:mb-0">
          <div className="w-full md:w-1/2 px-3 mb-6 md:mb-0">
            <div className="md:w-2/3">
              <form onSubmit={onSubmit}>
                <input
                  className="shadow bg-gray-200 hover:bg-white focus:shadow-outline focus:outline-none text-gray-700 font-bold py-2 px-4 rounded hover:pointer"
                  type="submit"
                  value="Send post"
                ></input>
              </form>
            </div>
          </div>
        </div>
      </div>
      {loading && <Loading />}
    </div>
  );
}
