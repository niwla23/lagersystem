import { PartModel, CreatePartData } from "../types"

export function getImageUrl(imageId: string) {
  return `/api/static/${imageId}`
}

export async function getAllParts(): Promise<PartModel[]> {
  let resp = await fetch("/api/parts")
  return await resp.json()
}

export async function createPart(data: CreatePartData, image: File | undefined) {
  console.log("creating:", data)
  let resp = await fetch("/api/parts", { method: "POST", body: JSON.stringify(data), headers: { "content-type": "application/json" } })
  let partId: number = (await resp.json())["id"]

  if (!image) {
    console.log("no image", image)
    return
  }
  const formData = new FormData()

  console.log(image)
  formData.append("image", image)
  let respImage = await fetch(`/api/parts/${partId}/image`, {
    method: "PUT",
    body: formData,
    // headers: { "content-type": `multipart/form-data; boundary=${formData.get}` },
  })
}
