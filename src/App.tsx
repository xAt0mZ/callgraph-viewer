import splitSpacesExcludeQuotes from 'quoted-string-space-split';
import ReactFlow, {
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  BackgroundVariant,
  Panel,
  Node,
  Edge,
} from 'reactflow';

import 'reactflow/dist/style.css';
import { Upload } from './Upload';
import { useCallback, useState } from 'react';

const initialNodes = [
  { id: '1', position: { x: 0, y: 0 }, data: { label: '1' } },
  { id: '2', position: { x: 0, y: 100 }, data: { label: '2' } },
];
const initialEdges = [{ id: 'e1-2', source: '1', target: '2' }];

function build(raw: string): [Node[], Edge[]] {
  const nodes: {
    [k: string]: Node;
  } = {};
  const edges: Edge[] = [];

  const links = raw.split('\n');
  links.forEach((link, idx) => {
    const [from, to] = link.split(' ').map((s) => s.replace('"', ''));

    let fromNode: Node = nodes[from];
    if (!fromNode) {
      nodes[from] = fromNode = {
        id: from,
        data: { label: from },
        position: { x: idx * 100, y: idx * 100 },
      };
    }

    let toNode: Node = nodes[to];
    if (!toNode) {
      nodes[to] = toNode = {
        id: to,
        data: { label: to },
        position: { x: idx * 100 + 50, y: idx * 100 + 50 },
      };
    }

    edges.push({
      id: `${fromNode.id}-${toNode.id}`,
      source: fromNode.id,
      target: toNode.id,
    });
  });
  return [Object.values(nodes), edges];
}

export function App() {
  const [filter, setFilter] = useState('');
  const [allNodes, setAllNodes] = useState<Node[]>([]);
  const [allEdges, setAllEdges] = useState<Edge[]>([]);
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  const handleChange = useCallback(
    (raw: string) => {
      const [nodes, edges] = build(raw);
      setAllNodes(nodes);
      setAllEdges(edges);
    },
    [setAllNodes, setAllEdges]
  );

  const render = useCallback(() => {
    const filteredNodes = allNodes.filter((n) => n.id.includes(filter));
    const nodeNames = filteredNodes.map((n) => n.id);
    const filteredEdges = allEdges.filter(
      (e) => nodeNames.includes(e.source) || nodeNames.includes(e.target)
    );
    filteredNodes.map((n, idx) => {
      n.position = { x: idx * 100, y: idx * 100 };
    });
    setNodes(filteredNodes);
    setEdges(filteredEdges);
  }, [allEdges, allNodes, filter, setEdges, setNodes]);

  return (
    <ReactFlow
      nodes={nodes}
      edges={edges}
      onNodesChange={onNodesChange}
      onEdgesChange={onEdgesChange}
    >
      <Panel position='top-left'>
        <Upload title='Upload nodes' onChange={handleChange} />
      </Panel>
      <Panel position='top-center'>
        <input onChange={(e) => setFilter(e.target.value)} />
        <button onClick={() => render()}>Submit</button>
      </Panel>
      <Controls />
      {/* <MiniMap /> */}
      <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
    </ReactFlow>
  );
}
