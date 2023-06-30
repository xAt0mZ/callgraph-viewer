// import ReactFlow, {
//   ReactFlowProvider,
//   useNodesState,
//   useEdgesState,
//   Controls,
//   MiniMap,
//   Background,
//   BackgroundVariant,
//   Node,
// } from 'reactflow';

// import { Buttons } from './Buttons';

// import 'reactflow/dist/style.css';

// const initialNodes: Node[] = [
//   {
//     id: '1',
//     type: 'input',
//     data: { label: 'Node 1' },
//     position: { x: 250, y: 5 },
//   },
//   { id: '2', data: { label: 'Node 2' }, position: { x: 100, y: 100 } },
//   { id: '3', data: { label: 'Node 3' }, position: { x: 400, y: 100 } },
//   { id: '4', data: { label: 'Node 4' }, position: { x: 400, y: 200 } },
// ];

// const initialEdges = [
//   {
//     id: 'e1-2',
//     source: '1',
//     target: '2',
//   },
//   { id: 'e1-3', source: '1', target: '3' },
// ];

// const ProviderFlow = () => {
//   const [nodes, setNodes, onNodesChange] = useNodesState(initialNodes);
//   const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges);

//   return (
//     <ReactFlowProvider>
//       <ReactFlow
//         nodes={nodes}
//         edges={edges}
//         onNodesChange={onNodesChange}
//         onEdgesChange={onEdgesChange}
//         fitView
//       >
//         <Buttons />
//         <Controls />
//         <MiniMap />
//         <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
//       </ReactFlow>
//     </ReactFlowProvider>
//   );
// };

// export default ProviderFlow;
