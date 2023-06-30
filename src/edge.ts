import { Edge } from 'reactflow';
import { GNode } from './node';

export function buildEdges(raw: string, nodes: GNode[]): Edge[] {
  const edges: Edge[] = [];
  const rawEdges = raw.split('\n');
  rawEdges.forEach((edge) => {
    const [from, to] = edge.replace(/"/g, '').split(' ');
    if (!from || !to) {
      return;
    }
    const fromNode = nodes.find((n) => n.id === from);
    const toNode = nodes.find((n) => n.id === to);
    if (!fromNode || !toNode) {
      return;
    }
    edges.push({
      id: `${fromNode.id}-${toNode.id}`,
      source: fromNode.id,
      target: toNode.id,
    });
  });
  return edges;
}
