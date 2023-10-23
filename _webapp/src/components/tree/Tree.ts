import type { Sensor, Station } from '@/http-client'

export interface TreeNode {}

export interface Tree {
  label: string
  children: Tree[]
  sensor: Sensor | null
  station: Station | null
}

export const traverseChildren = (node: Tree | undefined): Sensor[] => {
  if (!node || node.children.length == 0) {
    if (node?.sensor) {
      return [node.sensor]
    } else {
      return []
    }
  } else {
    let children: Sensor[] = []
    for (const child of node.children) {
      children = children.concat(traverseChildren(child))
    }
    if (node.sensor) {
      children.push(node.sensor)
    }
    return children
  }
}
