import React, { useEffect, useState } from 'react';

export default function Loading() {
  const [n, setN] = useState(0);
  useEffect(() => {
    let interval = setInterval(() => {
      setN(prev => (prev + 0.5) % 100);
    }, 10);
    return () => clearInterval(interval);
  }, []);

  return (
    <div className="container mt-3">
      <div className="relative pt-1">
        <div className="flex mb-2 items-center ml-3">
          <div>
            <span className="text-xs font-semibold inline-block py-1 px-2 uppercase rounded-full text-teal-600 bg-teal-200 text-center">
              Just hold on one sec...
            </span>
          </div>
        </div>
        <div className="overflow-hidden h-2 mb-4 text-xs flex rounded bg-teal-200">
          <div
            style={{ width: `${n}%` }}
            className="shadow-none flex flex-col text-center whitespace-nowrap text-white justify-center bg-teal-500"
          ></div>
        </div>
      </div>
    </div>
  );
}
