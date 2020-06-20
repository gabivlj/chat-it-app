import React from 'react';
import { Link } from 'react-router-dom';

export default function NotFound() {
  return (
    <div>
      <p className="text-sm container">
        Sorry but there has been an error... :( Maybe we did not found what you
        were searching for{' '}
        <Link to="/" className="underline text-green-500">
          Click here to take you to the safe zone!
        </Link>
      </p>
    </div>
  );
}
