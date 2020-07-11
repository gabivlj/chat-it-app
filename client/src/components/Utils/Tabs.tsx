import React, { useState } from 'react';

type Props = {
  components: Array<[string, () => JSX.Element]>;
  onChange?: (currentElement: () => JSX.Element) => void;
  startWith?: number;
};

const classActive = `bg-white inline-block border-l border-t border-r rounded-t py-2 px-4 text-blue-700 font-semibold outline-none`;
const classInactive = `bg-white inline-block py-2 px-4 text-blue-500 hover:text-blue-800 font-semibold outline-none cursor-pointer`;

export default function Tabs({
  components,
  onChange = (_) => {},
  startWith = 0,
}: Props) {
  const [currentElement, setCurrentElement] = useState(startWith);
  const [currentName, Component] = components[currentElement];
  const getActive = (name: string) => {
    return name === currentName ? classActive : classInactive;
  };
  return (
    <div>
      <ul className="flex border-b">
        {components.map(([name], idx) => (
          <li
            className="-mb-px mr-1 cursor-pointer"
            key={name}
            onClick={() => {
              onChange(components[idx][1]);
              setCurrentElement(idx);
            }}
          >
            <span className={getActive(name)}>{name}</span>
          </li>
        ))}
      </ul>
      <Component />
    </div>
  );
}
