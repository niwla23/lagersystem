import axios from "axios"
import {
  PartModel,
  UpdatePartData,
  CreatePartData,
  StoreByScannerResponse as StoreBoxByScannerResponse,
  BoxModel,
  ClearScannerResponse,
  TagModel,
  OperatorPositionsResponse,
  IOState,
} from "../types"

export function getImageUrl(imageId: string) {
  if (imageId === "00000000-0000-0000-0000-000000000000") {
    return undefined
  }
  return `/api/static/${imageId}`
}

export async function getAllParts(): Promise<PartModel[]> {
  let resp = await fetch("/api/parts")
  return await resp.json()
}

export async function getPartById(partId: number): Promise<PartModel> {
  let resp = await fetch(`/api/parts/${partId}`)
  return await resp.json()
}

export async function searchParts(query: string, filter: string): Promise<PartModel[]> {
  let resp = await fetch(`/api/parts/search?q=${query}`)
  return await resp.json()
}

export async function createPart(data: CreatePartData, image: File | undefined): Promise<PartModel> {
  let resp = await axios.post("/api/parts", data)
  let partId: number = resp.data["id"]

  if (!image) {
    return resp.data
  }
  const formData = new FormData()

  formData.append("image", image)
  let respImage = await axios.put(`/api/parts/${partId}/image`, formData)
  return resp.data
}

export async function updatePart(partId: number, data: UpdatePartData, image: File | undefined) {
  let resp = await axios.put(`/api/parts/${partId}`, data)

  if (!image) {
    return resp
  }
  const formData = new FormData()

  formData.append("image", image)
  let respImage = await axios.put(`/api/parts/${partId}/image`, formData)
  return respImage
}

export async function deletePart(partId: number) {
  let resp = await axios.delete(`/api/parts/${partId}`)
  return resp
}

// todo: correct type
export async function deliverPart(partId: number): Promise<object> {
  let resp = await axios.post(`/api/parts/${partId}/deliver`)
  return await resp.data
}

// todo: correct type
export async function deliverBox(boxId: string): Promise<object> {
  let resp = await axios.post(`/api/boxes/${boxId}/deliver`)
  return resp.data
}

export async function clearIo(): Promise<StoreBoxByScannerResponse> {
  let resp = await axios.post(`/api/warehouses/1/clearIo`)
  return resp.data
}

export async function clearScanner(): Promise<ClearScannerResponse> {
  let resp = await axios.post(`/api/warehouses/1/clearScanner`)
  return resp.data
}

export async function getBoxFromScanner(): Promise<ClearScannerResponse> {
  let resp = await axios.get(`/api/boxes/getFromScanner`)
  return resp.data
}

export async function getEmptyBox(): Promise<BoxModel> {
  const resp = await axios.get("/api/boxes/get-free-box")
  return resp.data
}

export async function getAllTags(): Promise<TagModel[]> {
  let resp = await fetch("/api/tags")
  return await resp.json()
}

export async function getPositions(): Promise<OperatorPositionsResponse> {
  let resp = await fetch("/api/warehouses/1/getPositions")
  return await resp.json()
}

export async function getIoState(): Promise<IOState> {
  let resp = await fetch("/api/warehouses/1/getIOState")
  return await resp.json()
}

export async function getFreeIoSlots(): Promise<number[]> {
  let ioState = await getIoState()
  return Object.entries(ioState)
    .filter((v) => v[1] === "free")
    .map((v) => parseInt(v[0]))
}

export async function bulkLinkParts(parts: number[], storeAfterwards: boolean): Promise<{ success: boolean }> {
  let resp = await axios.post("/api/parts/bulkLink", {
    partIds: parts,
    storeAfterwards: storeAfterwards,
  })
  return resp.data
}
