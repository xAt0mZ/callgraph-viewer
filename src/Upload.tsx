import { ChangeEvent } from 'react';

type Props = {
  title: string;
  id: string;
  onChange: (input: string) => void;
};

export function Upload({ title, onChange, id }: Props) {
  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      const fileReader = new FileReader();
      fileReader.readAsText(e.target.files[0], 'UTF-8');
      fileReader.onload = (e) => {
        if (e.target?.result) {
          onChange(e.target.result as string);
        }
      };
    }
  };
  return (
    <div className='relative'>
      <input
        type='file'
        id={id}
        onChange={(e) => handleChange(e)}
        className='block px-2.5 pb-1.5 pt-3 w-full text-sm text-gray-900 bg-transparent rounded-lg border-1 border-gray-300 appearance-none dark:text-white dark:border-gray-600 dark:focus:border-blue-500 focus:outline-none focus:ring-0 focus:border-blue-600 peer'
      />
      <label
        htmlFor={id}
        className='absolute text-sm text-gray-500 dark:text-gray-400 duration-300 transform -translate-y-3 scale-75 top-1 z-10 origin-[0] bg-white dark:bg-gray-900 px-2 peer-focus:px-2 peer-focus:text-blue-600 peer-focus:dark:text-blue-500 peer-placeholder-shown:scale-100 peer-placeholder-shown:-translate-y-1/2 peer-placeholder-shown:top-1/2 peer-focus:top-1 peer-focus:scale-75 peer-focus:-translate-y-3 left-1'
      >
        {title}
      </label>
    </div>
  );
}
