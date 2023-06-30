import { Node } from 'reactflow';

export type GNode = Node & {
  idx: number;
  pkg: string;
  fn: string;
};

function splitLastOccurrence(str: string, substring: string) {
  const lastIndex = str.lastIndexOf(substring);
  const before = str.slice(0, lastIndex);
  const after = str.slice(lastIndex + 1);
  return [before, after];
}

function buildNode(id: string, idx: number): GNode {
  const [pkg, fn] = splitLastOccurrence(id, '.');
  return {
    id,
    idx,
    position: { x: 0, y: 0 },
    data: {
      label: (
        <>
          <p className='bg-gray-200'>{pkg}</p>
          <p>{fn}</p>
        </>
      ),
    },
    pkg,
    fn,
  };
}

export function buildNodes(raw: string): GNode[] {
  return raw.split('\n').map(buildNode);
}
