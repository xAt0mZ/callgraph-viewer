import ReactFlow, {
  Controls,
  Background,
  useNodesState,
  useEdgesState,
  BackgroundVariant,
  Panel,
  Edge,
} from 'reactflow';

import { Upload } from './Upload';
import { useCallback, useState } from 'react';
import { GNode, buildNodes } from './node';
import { buildEdges } from './edge';

export function App() {
  const [rawNodes, setRawNodes] = useState('');
  const [rawEdges, setRawEdges] = useState('');
  const [filter, setFilter] = useState('');
  const [allNodes, setAllNodes] = useState<GNode[]>([]);
  const [allEdges, setAllEdges] = useState<Edge[]>([]);
  const [nodes, setNodes, onNodesChange] = useNodesState([]);
  const [edges, setEdges, onEdgesChange] = useEdgesState([]);

  const compute = useCallback(() => {
    const nodes = buildNodes(rawNodes);
    const edges = buildEdges(rawEdges, nodes);
    setAllNodes(nodes);
    setAllEdges(edges);
  }, [rawEdges, rawNodes]);

  const render = useCallback(() => {
    const filteredNodes = allNodes.filter((n) => n.id.includes(filter));
    const nodeNames = filteredNodes.map((n) => n.id);

    const filteredEdges = allEdges.filter(
      (e) => nodeNames.includes(e.source) || nodeNames.includes(e.target)
    );
    const sourceNodes = 

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
        <div className='flex flex-col'>
          <Upload
            title='Upload nodes'
            onChange={setRawNodes}
            id='upload-nodes'
          />
          <Upload
            title='Upload edges'
            onChange={setRawEdges}
            id='upload-edges'
          />
          <button disabled={!rawNodes || !rawEdges} onClick={compute}>
            Submit
          </button>
        </div>
      </Panel>
      <Panel position='top-center'>
        <input onChange={(e) => setFilter(e.target.value)} />
        <button
          disabled={!allNodes.length || !allEdges.length}
          onClick={render}
        >
          Submit
        </button>
      </Panel>
      <Controls />
      {/* <MiniMap /> */}
      <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
    </ReactFlow>
  );
}
