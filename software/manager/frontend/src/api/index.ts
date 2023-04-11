import axios from "axios"
import { PartModel, CreatePartData } from "../types"

export function getImageUrl(imageId: string) {
  if (imageId === "00000000-0000-0000-0000-000000000000") {
    return null
  }
  return `/api/static/${imageId}`
}

export async function getAllParts(): Promise<PartModel[]> {
  let resp = await fetch("/api/parts")
  return await resp.json()
}

export async function searchParts(query: string, filter: string): Promise<PartModel[]> {
  let resp = await fetch(`/api/parts/search?q=${query}`)
  return await resp.json()
}

export async function createPart(data: CreatePartData, image: File | undefined) {
  // let resp = await fetch("/api/parts", { method: "POST", body: JSON.stringify(data), headers: { "content-type": "application/json" } })
  let resp = await axios.post("/api/parts", data)
  let partId: number = (resp.data)["id"]

  if (!image) {
    return resp
  }
  const formData = new FormData()

  formData.append("image", image)
  let respImage = await axios.put(`/api/parts/${partId}/image`, formData)
  return respImage
}

// todo: correct type
export async function deliverPart(partId: number): Promise<object> {
  let resp = await fetch(`/api/parts/${partId}/deliver`)
  return await resp.json()
}

// todo: correct type
export async function deliverBox(partId: number): Promise<object> {
  let resp = await fetch(`/api/boxes/${partId}/deliver`)
  return await resp.json()
}

// todo: correct type
export async function storeBoxByScanner(): Promise<object> {
  let resp = await fetch(`/api/store/by-scanner`)
  return await resp.json()
}