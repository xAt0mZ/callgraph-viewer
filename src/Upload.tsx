import { ChangeEvent } from 'react';

type Props = {
  title: string;
  onChange: (input: string) => void;
};

export function Upload({ title, onChange }: Props) {
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
  return <input type='file' onChange={(e) => handleChange(e)} title={title} />;
}
