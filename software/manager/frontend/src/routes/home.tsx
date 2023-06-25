import React, { useEffect } from "react"
import PartCard from "../components/partCard"
import * as api from "../api"
import { PartModel } from "../types"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { faPen } from "@fortawesome/free-solid-svg-icons"
import { useNavigate } from "react-router-dom"

export default function Home() {
  const navigate = useNavigate()
  let [parts, setParts] = React.useState<PartModel[]>([])
  let [searchQuery, setSearchQuery] = React.useState("")

  const loadAllParts = async () => {
    let x = await api.searchParts(searchQuery, "")
    console.log(x)
    setParts(x)
  }

  useEffect(() => {
    loadAllParts()
  }, [searchQuery])

  let renderedParts = parts.map((part) => {
    let tags = part.tags ? part.tags.map((tag) => tag.name) : []

    return (
      <PartCard
        id={part.id}
        actions={[
          <button className="btn btn-primary flex-grow sm:flex-grow-0" onClick={() => {}}>
            Deliver
          </button>,
          <button className="btn btn-secondary" onClick={() => {navigate(`/parts/edit/${part.id}`)}}>
            <FontAwesomeIcon icon={faPen} />
          </button>,
        ]}
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
        <input
          type="text"
          placeholder="Search…"
          className="input input-bordered w-full"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
        />
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
