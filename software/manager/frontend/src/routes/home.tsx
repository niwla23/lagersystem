import React, { useEffect } from "react"
import PartCard from "../components/partCard"
import * as api from "../api"
import { PartModel } from "../types"

export default function Home() {
  let [parts, setParts] = React.useState<PartModel[]>([])

  const loadAllParts = async () => {
    let x = await api.getAllParts()
    setParts(x)
  }

  useEffect(() => {
    loadAllParts()
  }, [])

  let renderedParts = parts.map((part) => {
    let tags = part.tags ? part.tags.map((tag) => tag.name) : []

    return (
      <PartCard
        id={part.id}
        onActionClick={() => {}}
        actionText="deliver"
        name={part.name}
        description={part.description || "No description"}
        tags={tags}
        properties={part.properties || []}
        imageUrl={api.getImageUrl(part.imageId)}
      />
    )
  })

  return (
    <div className="h-full p-4">
      <div className="input-group w-full pb-4">
        <input type="text" placeholder="Search…" className="input input-bordered w-full" />
        <button className="btn btn-primary btn-square">
          <svg xmlns="http://www.w3.org/2000/svg" className="h-6 w-6" fill="none" viewBox="0 0 24 24" stroke="currentColor">
            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" />
          </svg>
        </button>
      </div>
      <main className="grid gap-2">{renderedParts}</main>
    </div>
  )
}
