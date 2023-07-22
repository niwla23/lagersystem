type UUID = string

export interface TagModel {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
  description: string
}

export interface PropertyModel {
  id: number
  createdAt: Date
  updatedAt: Date
  name: string
  value: string
  type: "string" | "number" | "boolean"
}

export interface PositionModel {
  id: number
  createdAt: Date
  updatedAt: Date
}

export interface BoxModel {
  id: UUID
  createdAt: Date
  updatedAt: Date
  parts?: PartModel[]
  position?: PositionModel
}

export interface PartModel {
  createdAt: string
  updatedAt: string
  id: number
  name: string
  description?: string
  tags?: TagModel[]
  properties?: PropertyModel[]
  amount: number
  imageId: string
  box?: BoxModel
}

// todo: outdated
export interface StoreByScannerResponse {
  status: string
  boxId: string
  positionID: number
  duration: number
}

export interface ClearScannerResponse {
  status: "success"
  boxId: UUID
  duration: number
}

export interface OperatorPositionsResponse {
  positions: { [posId: string]: { x: number; y: number; box?: BoxModel } }
}

export interface IOState {
  [posId: string]: "occupied" | "free"
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
  boxId?: string
}

export interface UpdatePartData {
  name: string
  description: string
  amount: number
  tags: string[]
  properties: { [key: string]: PropertyAddData }
  boxId?: string
}
