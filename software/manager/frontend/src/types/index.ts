export interface TagModel {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
}

export interface PropertyModel {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
  value: string
  type: "string" | "number" | "boolean"
}

export interface PartModel {
  id: number
  name: string
  description?: string
  tags?: TagModel[]
  properties?: PropertyModel[]
  amount: number
  imageId: string
}

export interface PropertyAddData {
  value: string
  type: string
}

export interface CreatePartData {
  name: string
  description: string
  amount: number
  tags: string[]
  properties: { [key: string]: PropertyAddData }
}