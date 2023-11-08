export interface CommonNode {
  label?: string
  id?: number
  children: CommonNode[]
  selected?: boolean
}

export const hasSelectedChildren = (node: CommonNode | undefined): boolean => {
  if (!node || node.children.length == 0) {
    return node?.selected ?? false
  } else {
    let selected: boolean = false
    for (const child of node.children) {
      selected = hasSelectedChildren(child) || selected
    }
    return selected
  }
}

export const allChildrenSelected = (node: CommonNode | undefined): boolean => {
  if (!node || node.children.length == 0) {
    return node?.selected ?? false
  } else {
    let selected: boolean = false
    for (const child of node.children) {
      selected = allChildrenSelected(child) && selected
    }
    return selected
  }
}
